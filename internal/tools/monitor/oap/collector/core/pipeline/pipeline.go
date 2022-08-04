// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipeline

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda/internal/apps/msp/apm/trace"
	"github.com/erda-project/erda/internal/tools/monitor/core/log"
	"github.com/erda-project/erda/internal/tools/monitor/core/metric"
	"github.com/erda-project/erda/internal/tools/monitor/oap/collector/core/config"
	"github.com/erda-project/erda/internal/tools/monitor/oap/collector/core/model"
	odata2 "github.com/erda-project/erda/internal/tools/monitor/oap/collector/core/model/odata"
	"github.com/erda-project/erda/internal/tools/monitor/oap/collector/lib"
)

type Pipeline struct {
	name       string
	Log        logs.Logger
	cfg        config.Pipeline
	dtype      odata2.DataType
	receivers  []*model.RuntimeReceiver
	processors []*model.RuntimeProcessor
	exporters  []*model.RuntimeExporter
	rp, pe     chan odata2.ObservableData
}

var (
	dataReceived                *prometheus.CounterVec
	dataProcessed, dataExported *prometheus.CounterVec
)

func init() {
	dataReceived = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "data_pipeline",
		Name:      "receiver_consumed",
		Help:      "event count for certain receiver consumed",
	}, []string{"pipeline", "dtype", "receiver"})

	dataProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "data_pipeline",
		Name:      "processor_consumed",
		Help:      "event count for certain processor consumed",
	}, []string{"pipeline", "dtype", "processor"})

	dataExported = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "data_pipeline",
		Name:      "exporter_consumed",
		Help:      "event count for certain exporter consumed",
	}, []string{"pipeline", "dtype", "exporter"})
}

func NewPipeline(name string, logger logs.Logger, cfg config.Pipeline, dtype odata2.DataType) *Pipeline {
	p := &Pipeline{
		name:  name,
		Log:   logger,
		cfg:   cfg,
		dtype: dtype,
	}
	p.rp = make(chan odata2.ObservableData, cfg.RPChannelCap)
	p.pe = make(chan odata2.ObservableData, cfg.PEChannelCap)
	p.initStats()
	return p
}

func (p *Pipeline) initStats() {
	promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "data_pipeline",
		Name:      "rp_channel_used",
		Help:      "the current channel used of receiver to processor",
		ConstLabels: prometheus.Labels{
			"pipeline": p.name,
		},
	}, func() float64 {
		return float64(len(p.pe))
	})

	promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "data_pipeline",
		Name:      "pe_channel_used",
		Help:      "the current channel used of processor to exporter",
		ConstLabels: prometheus.Labels{
			"pipeline": p.name,
		},
	}, func() float64 {
		return float64(len(p.pe))
	})
}

func (p *Pipeline) InitComponents(receivers, processors, exporters []model.ComponentUnit) error {
	rs, err := p.rsFromComponent(receivers)
	if err != nil {
		return err
	}
	prs, err := p.prsFromComponent(processors)
	if err != nil {
		return err
	}
	es, err := p.esFromComponent(exporters)
	if err != nil {
		return err
	}

	p.receivers = rs
	p.processors = prs
	p.exporters = es
	return nil
}

func (p *Pipeline) rsFromComponent(coms []model.ComponentUnit) ([]*model.RuntimeReceiver, error) {
	res := make([]*model.RuntimeReceiver, 0, len(coms))
	for _, com := range coms {
		c, ok := com.Component.(model.Receiver)
		if !ok {
			return nil, fmt.Errorf("invalid component<%s> type<%T>", com.Name, com.Component)
		}
		res = append(res, &model.RuntimeReceiver{Name: com.Name, Receiver: c, Filter: com.Filter})
	}
	return res, nil
}

func (p *Pipeline) prsFromComponent(coms []model.ComponentUnit) ([]*model.RuntimeProcessor, error) {
	res := make([]*model.RuntimeProcessor, 0, len(coms))
	for _, com := range coms {
		c, ok := com.Component.(model.Processor)
		if !ok {
			return nil, fmt.Errorf("invalid component<%s> type<%T>", com.Name, com.Component)
		}
		res = append(res, &model.RuntimeProcessor{Name: com.Name, Processor: c, Filter: com.Filter})
	}
	return res, nil
}
func (p *Pipeline) esFromComponent(coms []model.ComponentUnit) ([]*model.RuntimeExporter, error) {
	res := make([]*model.RuntimeExporter, 0, len(coms))
	for _, com := range coms {
		c, ok := com.Component.(model.Exporter)
		if !ok {
			return nil, fmt.Errorf("invalid component<%s> type<%T>", com.Name, com.Component)
		}
		res = append(res, &model.RuntimeExporter{
			Name:     com.Name,
			Logger:   p.Log.Sub("exporter-" + com.Name),
			Exporter: c,
			DType:    p.dtype,
			Filter:   com.Filter,
			Timer:    time.NewTimer(lib.RandomDuration(p.cfg.FlushInterval, p.cfg.FlushJitter)),
			Interval: p.cfg.FlushInterval,
			Jitter:   p.cfg.FlushJitter,
			Buffer:   odata2.NewBuffer(p.cfg.BatchSize),
		})
	}
	return res, nil
}

func (p *Pipeline) StartStream(ctx context.Context) {
	go p.StartExporters(ctx, p.pe)
	go p.startProcessors(ctx, p.rp, p.pe)
	p.startReceivers(ctx, p.rp)
}

func (p *Pipeline) StartExporters(ctx context.Context, out <-chan odata2.ObservableData) {
	for _, e := range p.exporters {
		go func(exp *model.RuntimeExporter) {
			go exp.Start(ctx)
		}(e)
	}

	for {
		select {
		case data, ok := <-out:
			if !ok {
				return
			}
			// TODO. fan-out connector
			var wg sync.WaitGroup
			wg.Add(len(p.exporters))
			for _, e := range p.exporters {
				go func(exp *model.RuntimeExporter, od odata2.ObservableData) {
					defer wg.Done()
					dataExported.WithLabelValues(p.name, string(p.dtype), exp.Name).Inc()
					exp.Add(od)
				}(e, data)
			}
			wg.Wait()
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pipeline) startProcessors(ctx context.Context, in <-chan odata2.ObservableData, out chan<- odata2.ObservableData) {
	for _, r := range p.processors {
		rp, ok := r.Processor.(model.RunningProcessor)
		if !ok {
			continue
		}
		// TODO. may not needed
		rp.StartProcessor(p.newConsumer(ctx, r.Name, out))
	}

	for {
	begin:
		select {
		case data, ok := <-in:
			if !ok {
				return
			}
			// TODO. Parallelism
			for _, pr := range p.processors {
				if !pr.Filter.Selected(data) {
					continue
				}
				dataProcessed.WithLabelValues(p.name, string(p.dtype), pr.Name).Inc()
				switch p.dtype {
				case odata2.MetricType:
					tmp, err := pr.Processor.ProcessMetric(data.(*metric.Metric))
					if err != nil {
						p.Log.Errorf("Processor<%s> process data error: %s", pr.Name, err)
						continue
					}
					if tmp == nil {
						goto begin
					}
					data = tmp
				case odata2.LogType:
					tmp, err := pr.Processor.ProcessLog(data.(*log.Log))
					if err != nil {
						p.Log.Errorf("Processor<%s> process data error: %s", pr.Name, err)
						continue
					}
					if tmp == nil {
						goto begin
					}
					data = tmp
				case odata2.SpanType:
					tmp, err := pr.Processor.ProcessSpan(data.(*trace.Span))
					if err != nil {
						p.Log.Errorf("Processor<%s> process data error: %s", pr.Name, err)
						continue
					}
					if tmp == nil {
						goto begin
					}
					data = tmp
				case odata2.RawType:
					tmp, err := pr.Processor.ProcessRaw(data.(*odata2.Raw))
					if err != nil {
						p.Log.Errorf("Processor<%s> process data error: %s", pr.Name, err)
						continue
					}
					if tmp == nil {
						goto begin
					}
					data = tmp
				default:
					continue
				}
			}

			// wait forever
			select {
			case out <- data:
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pipeline) startReceivers(ctx context.Context, out chan<- odata2.ObservableData) {
	for _, r := range p.receivers {
		r.Receiver.RegisterConsumer(p.newConsumer(ctx, r.Name, out))
	}
}

func (p *Pipeline) newConsumer(pctx context.Context, name string, out chan<- odata2.ObservableData) model.ObservableDataConsumerFunc {
	return func(od odata2.ObservableData) error {
		select {
		case out <- od:
			dataReceived.WithLabelValues(p.name, string(p.dtype), name).Inc()
		case <-pctx.Done():
			return nil
		}
		return nil
	}
}