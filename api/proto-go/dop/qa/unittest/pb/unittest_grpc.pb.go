// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: unittest.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// UnitTestServiceClient is the client API for UnitTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnitTestServiceClient interface {
	Callback(ctx context.Context, in *TestCallBackRequest, opts ...grpc.CallOption) (*TestCallBackResponse, error)
	GetTestTypes(ctx context.Context, in *TestTypeRequest, opts ...grpc.CallOption) (*TestTypeResponse, error)
	ListRecords(ctx context.Context, in *TestRecordPagingRequest, opts ...grpc.CallOption) (*TestRecordPagingResponse, error)
	GetRecord(ctx context.Context, in *TestRecordGetRequest, opts ...grpc.CallOption) (*TestRecordGetResponse, error)
}

type unitTestServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewUnitTestServiceClient(cc grpc1.ClientConnInterface) UnitTestServiceClient {
	return &unitTestServiceClient{cc}
}

func (c *unitTestServiceClient) Callback(ctx context.Context, in *TestCallBackRequest, opts ...grpc.CallOption) (*TestCallBackResponse, error) {
	out := new(TestCallBackResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.qa.unittest.UnitTestService/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitTestServiceClient) GetTestTypes(ctx context.Context, in *TestTypeRequest, opts ...grpc.CallOption) (*TestTypeResponse, error) {
	out := new(TestTypeResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.qa.unittest.UnitTestService/GetTestTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitTestServiceClient) ListRecords(ctx context.Context, in *TestRecordPagingRequest, opts ...grpc.CallOption) (*TestRecordPagingResponse, error) {
	out := new(TestRecordPagingResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.qa.unittest.UnitTestService/ListRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unitTestServiceClient) GetRecord(ctx context.Context, in *TestRecordGetRequest, opts ...grpc.CallOption) (*TestRecordGetResponse, error) {
	out := new(TestRecordGetResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.qa.unittest.UnitTestService/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnitTestServiceServer is the server API for UnitTestService service.
// All implementations should embed UnimplementedUnitTestServiceServer
// for forward compatibility
type UnitTestServiceServer interface {
	Callback(context.Context, *TestCallBackRequest) (*TestCallBackResponse, error)
	GetTestTypes(context.Context, *TestTypeRequest) (*TestTypeResponse, error)
	ListRecords(context.Context, *TestRecordPagingRequest) (*TestRecordPagingResponse, error)
	GetRecord(context.Context, *TestRecordGetRequest) (*TestRecordGetResponse, error)
}

// UnimplementedUnitTestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUnitTestServiceServer struct {
}

func (*UnimplementedUnitTestServiceServer) Callback(context.Context, *TestCallBackRequest) (*TestCallBackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (*UnimplementedUnitTestServiceServer) GetTestTypes(context.Context, *TestTypeRequest) (*TestTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestTypes not implemented")
}
func (*UnimplementedUnitTestServiceServer) ListRecords(context.Context, *TestRecordPagingRequest) (*TestRecordPagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecords not implemented")
}
func (*UnimplementedUnitTestServiceServer) GetRecord(context.Context, *TestRecordGetRequest) (*TestRecordGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}

func RegisterUnitTestServiceServer(s grpc1.ServiceRegistrar, srv UnitTestServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_UnitTestService_serviceDesc(srv, opts...), srv)
}

var _UnitTestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.dop.qa.unittest.UnitTestService",
	HandlerType: (*UnitTestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "unittest.proto",
}

func _get_UnitTestService_serviceDesc(srv UnitTestServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_UnitTestService_Callback_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Callback(ctx, req.(*TestCallBackRequest))
	}
	var _UnitTestService_Callback_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UnitTestService_Callback_info = transport.NewServiceInfo("erda.dop.qa.unittest.UnitTestService", "Callback", srv)
		_UnitTestService_Callback_Handler = h.Interceptor(_UnitTestService_Callback_Handler)
	}

	_UnitTestService_GetTestTypes_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTestTypes(ctx, req.(*TestTypeRequest))
	}
	var _UnitTestService_GetTestTypes_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UnitTestService_GetTestTypes_info = transport.NewServiceInfo("erda.dop.qa.unittest.UnitTestService", "GetTestTypes", srv)
		_UnitTestService_GetTestTypes_Handler = h.Interceptor(_UnitTestService_GetTestTypes_Handler)
	}

	_UnitTestService_ListRecords_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListRecords(ctx, req.(*TestRecordPagingRequest))
	}
	var _UnitTestService_ListRecords_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UnitTestService_ListRecords_info = transport.NewServiceInfo("erda.dop.qa.unittest.UnitTestService", "ListRecords", srv)
		_UnitTestService_ListRecords_Handler = h.Interceptor(_UnitTestService_ListRecords_Handler)
	}

	_UnitTestService_GetRecord_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetRecord(ctx, req.(*TestRecordGetRequest))
	}
	var _UnitTestService_GetRecord_info transport.ServiceInfo
	if h.Interceptor != nil {
		_UnitTestService_GetRecord_info = transport.NewServiceInfo("erda.dop.qa.unittest.UnitTestService", "GetRecord", srv)
		_UnitTestService_GetRecord_Handler = h.Interceptor(_UnitTestService_GetRecord_Handler)
	}

	var serviceDesc = _UnitTestService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Callback",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestCallBackRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UnitTestServiceServer).Callback(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UnitTestService_Callback_info)
				}
				if interceptor == nil {
					return _UnitTestService_Callback_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.qa.unittest.UnitTestService/Callback",
				}
				return interceptor(ctx, in, info, _UnitTestService_Callback_Handler)
			},
		},
		{
			MethodName: "GetTestTypes",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestTypeRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UnitTestServiceServer).GetTestTypes(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UnitTestService_GetTestTypes_info)
				}
				if interceptor == nil {
					return _UnitTestService_GetTestTypes_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.qa.unittest.UnitTestService/GetTestTypes",
				}
				return interceptor(ctx, in, info, _UnitTestService_GetTestTypes_Handler)
			},
		},
		{
			MethodName: "ListRecords",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestRecordPagingRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UnitTestServiceServer).ListRecords(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UnitTestService_ListRecords_info)
				}
				if interceptor == nil {
					return _UnitTestService_ListRecords_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.qa.unittest.UnitTestService/ListRecords",
				}
				return interceptor(ctx, in, info, _UnitTestService_ListRecords_Handler)
			},
		},
		{
			MethodName: "GetRecord",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TestRecordGetRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(UnitTestServiceServer).GetRecord(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _UnitTestService_GetRecord_info)
				}
				if interceptor == nil {
					return _UnitTestService_GetRecord_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.qa.unittest.UnitTestService/GetRecord",
				}
				return interceptor(ctx, in, info, _UnitTestService_GetRecord_Handler)
			},
		},
	}
	return &serviceDesc
}
