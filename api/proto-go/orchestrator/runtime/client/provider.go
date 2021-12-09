// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: runtime.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/orchestrator/runtime/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.orchestrator.runtime",
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
	clientsType              = reflect.TypeOf((*Client)(nil)).Elem()
	runtimeServiceClientType = reflect.TypeOf((*pb.RuntimeServiceClient)(nil)).Elem()
	runtimeServiceServerType = reflect.TypeOf((*pb.RuntimeServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.orchestrator.runtime-client":
		return p.client
	case "erda.orchestrator.runtime.RuntimeService":
		return &runtimeServiceWrapper{client: p.client.RuntimeService(), opts: opts}
	case "erda.orchestrator.runtime.RuntimeService.client":
		return p.client.RuntimeService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case runtimeServiceClientType:
		return p.client.RuntimeService()
	case runtimeServiceServerType:
		return &runtimeServiceWrapper{client: p.client.RuntimeService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.orchestrator.runtime-client", &servicehub.Spec{
		Services: []string{
			"erda.orchestrator.runtime.RuntimeService",
			"erda.orchestrator.runtime.RuntimeService.client",
			"erda.orchestrator.runtime-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			runtimeServiceClientType,
			// server types
			runtimeServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
