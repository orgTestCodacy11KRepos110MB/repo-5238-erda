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

package query

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/pkg/transport"
	"github.com/erda-project/erda-infra/providers/i18n"
	"github.com/erda-project/erda/internal/tools/monitor/core/metric/model"
	tsql "github.com/erda-project/erda/internal/tools/monitor/core/metric/query/es-tsql"
	"github.com/erda-project/erda/internal/tools/monitor/core/metric/query/es-tsql/formats"
	"github.com/erda-project/erda/internal/tools/monitor/core/metric/query/metricmeta"
	"github.com/erda-project/erda/internal/tools/monitor/core/metric/storage"
)

type queryer struct {
	storage   storage.Storage `autowired:"metric-storage"`
	ckStorage storage.Storage `autowired:"metric-storage-clickhouse"`
	meta      *metricmeta.Manager
	log       logs.Logger
}

// New .
func New(meta *metricmeta.Manager, esStorage storage.Storage, ckStorage storage.Storage, log logs.Logger) Queryer {
	return &queryer{
		meta:      meta,
		storage:   esStorage,
		ckStorage: ckStorage,
		log:       log,
	}
}

const hourms = int64(time.Hour) / int64(time.Millisecond)

func (q *queryer) buildTSQLParser(ctx context.Context, ql, statement string, params map[string]interface{}, filters []*model.Filter, options url.Values) (parser tsql.Parser, others map[string]interface{}, err error) {
	idx := strings.Index(ql, ":")
	if idx > 0 {
		if ql[idx+1:] == "ast" && ql[0:idx] == "influxql" {
			statement, err = ConvertAstToStatement(statement)
			if err != nil {
				return nil, nil, err
			}
		}
		ql = ql[0:idx]
	}
	if ql != "influxql" {
		return nil, nil, fmt.Errorf("not support tsql '%s'", ql)
	}
	start, end, err := ParseTimeRange(options.Get("start"), options.Get("end"), options.Get("timestamp"), options.Get("latest"))
	if err != nil {
		return nil, nil, err
	}
	if end < hourms {
		end = hourms
	}
	if start < 0 || start >= end {
		start = end - hourms
	}
	fs, others := ParseFilters(options)
	filters = append(fs, filters...)

	_, debug := options["debug"]

	parser = tsql.New(start*int64(time.Millisecond), end*int64(time.Millisecond), ql, statement, debug)

	if parser == nil {
		return nil, nil, fmt.Errorf("not support tsql '%s'", ql)
	}

	parser.SetMeta(q.meta)
	parser, err = parser.SetFilter(filters)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid filter on parse filter: %s", err)
	}
	if params == nil {
		params = others
	}
	parser = parser.SetParams(params)

	// process header
	if ctx == nil {
		ctx = context.Background()
	}
	if head := transport.ContextHeader(ctx); head != nil {
		if orgNames := head.Get("org"); len(orgNames) > 0 {
			parser = parser.SetOrgName(orgNames[0])
		}
		if tkeys := head.Get("terminus_key"); len(tkeys) > 0 {
			parser = parser.SetTerminusKey(tkeys[0])
		}

	} else {
		header := ctx.Value("header")
		if v, ok := header.(http.Header); ok && v != nil {
			parser = parser.SetOrgName(v.Get("org"))
			parser = parser.SetTerminusKey(v.Get("terminus_key"))
		}
	}

	unit := options.Get("epoch") // Keep the same parameters as the influxdb.
	if len(unit) > 0 {
		unit, err := tsql.ParseTimeUnit(unit)
		if err != nil {
			return nil, nil, err
		}
		parser.SetTargetTimeUnit(unit)
	}
	tf := options.Get("time_field")
	idx = strings.Index(tf, "::")
	if len(tf) > 4 && idx > 0 {
		tf = tf[idx+2:] + "s." + tf[0:idx]
	}
	if len(tf) > 0 {
		parser.SetTimeKey(tf)
		if tf == model.TimestampKey {
			parser.SetOriginalTimeUnit(tsql.Nanosecond)
		} else {
			tu := options.Get("time_unit")
			if len(tu) > 0 {
				unit, err := tsql.ParseTimeUnit(tu)
				if err != nil {
					return nil, nil, err
				}
				parser.SetOriginalTimeUnit(unit)
			}
		}
	}
	return parser, others, nil
}

func (q *queryer) doQuery(ctx context.Context, ql, statement string, params map[string]interface{}, filters []*model.Filter, options url.Values) (*model.ResultSet, tsql.Query, map[string]interface{}, error) {
	parser, others, err := q.buildTSQLParser(ctx, ql, statement, params, filters, options)
	if err != nil {
		return nil, nil, nil, err
	}

	err = parser.Build()
	if err != nil {
		return nil, nil, nil, err
	}

	result, query, err := q.GetResult(parser)
	return result, query, others, err
}

func (q *queryer) GetResult(parser tsql.Parser) (*model.ResultSet, tsql.Query, error) {
	var query tsql.Query
	if q.ckStorage != nil {
		metrics, err := parser.Metrics()
		if err == nil {
			if q.ckStorage.Select(metrics) {
				queries, err := parser.ParseQuery(model.ClickhouseKind)
				if err != nil {
					return nil, nil, err
				}
				query = queries[0]
				result, err := q.ckStorage.Query(context.Background(), query)
				return result, query, err
			}
		}
	}
	queries, err := parser.ParseQuery("")
	if err != nil {
		return nil, nil, err
	}
	if len(queries) != 1 {
		return nil, nil, fmt.Errorf("only support one statement")
	}
	query = queries[0]
	result, err := q.storage.Query(context.Background(), query)
	return result, query, err
}

// Query .
func (q *queryer) Query(ctx context.Context, tsql, statement string, params map[string]interface{}, options url.Values) (*model.ResultSet, error) {
	rs, _, _, err := q.doQuery(ctx, tsql, statement, params, nil, options)
	if err != nil {
		q.log.Error("query tsql is error:", err)
	}
	return rs, err
}

// QueryWithFormat .
func (q *queryer) QueryWithFormat(ctx context.Context, tsql, statement, format string, langCodes i18n.LanguageCodes, params map[string]interface{}, filters []*model.Filter, options url.Values) (*model.ResultSet, interface{}, error) {
	rs, query, opts, err := q.doQuery(ctx, tsql, statement, params, filters, options)
	if err != nil {
		q.log.Error("query tsql is error:", err)
		return nil, nil, err
	}
	if rs.Details != nil {
		return rs, nil, err
	}
	data, err := formats.Format(format, query, rs.Data, opts)
	if err != nil {
		q.log.Error("parse query result is error:", err)
	}
	return rs, data, err
}