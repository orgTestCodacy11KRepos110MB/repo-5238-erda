// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: template.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/dicehub/template/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.dicehub.template",
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
	clientsType               = reflect.TypeOf((*Client)(nil)).Elem()
	templateServiceClientType = reflect.TypeOf((*pb.TemplateServiceClient)(nil)).Elem()
	templateServiceServerType = reflect.TypeOf((*pb.TemplateServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.dicehub.template-client":
		return p.client
	case "erda.core.dicehub.template.TemplateService":
		return &templateServiceWrapper{client: p.client.TemplateService(), opts: opts}
	case "erda.core.dicehub.template.TemplateService.client":
		return p.client.TemplateService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case templateServiceClientType:
		return p.client.TemplateService()
	case templateServiceServerType:
		return &templateServiceWrapper{client: p.client.TemplateService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.dicehub.template-client", &servicehub.Spec{
		Services: []string{
			"erda.core.dicehub.template.TemplateService",
			"erda.core.dicehub.template.TemplateService.client",
			"erda.core.dicehub.template-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			templateServiceClientType,
			// server types
			templateServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}