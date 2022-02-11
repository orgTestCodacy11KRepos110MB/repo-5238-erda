// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: diagnotor.proto, diagnotor_agent.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/diagnotor/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.monitor.diagnotor",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                     = reflect.TypeOf((*Client)(nil)).Elem()
	diagnotorServiceClientType      = reflect.TypeOf((*pb.DiagnotorServiceClient)(nil)).Elem()
	diagnotorServiceServerType      = reflect.TypeOf((*pb.DiagnotorServiceServer)(nil)).Elem()
	diagnotorAgentServiceClientType = reflect.TypeOf((*pb.DiagnotorAgentServiceClient)(nil)).Elem()
	diagnotorAgentServiceServerType = reflect.TypeOf((*pb.DiagnotorAgentServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.monitor.diagnotor-client":
		return p.client
	case "erda.core.monitor.diagnotor.DiagnotorAgentService":
		return &diagnotorAgentServiceWrapper{client: p.client.DiagnotorAgentService(), opts: opts}
	case "erda.core.monitor.diagnotor.DiagnotorAgentService.client":
		return p.client.DiagnotorAgentService()
	case "erda.core.monitor.diagnotor.DiagnotorService":
		return &diagnotorServiceWrapper{client: p.client.DiagnotorService(), opts: opts}
	case "erda.core.monitor.diagnotor.DiagnotorService.client":
		return p.client.DiagnotorService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case diagnotorServiceClientType:
		return p.client.DiagnotorService()
	case diagnotorServiceServerType:
		return &diagnotorServiceWrapper{client: p.client.DiagnotorService(), opts: opts}
	case diagnotorAgentServiceClientType:
		return p.client.DiagnotorAgentService()
	case diagnotorAgentServiceServerType:
		return &diagnotorAgentServiceWrapper{client: p.client.DiagnotorAgentService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.diagnotor-client", &servicehub.Spec{
		Services: []string{
			"erda.core.monitor.diagnotor.DiagnotorAgentService",
			"erda.core.monitor.diagnotor.DiagnotorAgentService.client",
			"erda.core.monitor.diagnotor.DiagnotorService",
			"erda.core.monitor.diagnotor.DiagnotorService.client",
			"erda.core.monitor.diagnotor-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			diagnotorServiceClientType,
			diagnotorAgentServiceClientType,
			// server types
			diagnotorServiceServerType,
			diagnotorAgentServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
