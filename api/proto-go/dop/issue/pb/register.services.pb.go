// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: stream.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterCommentIssueStreamServiceImp stream.proto
func RegisterCommentIssueStreamServiceImp(regester transport.Register, srv CommentIssueStreamServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterCommentIssueStreamServiceHandler(regester, CommentIssueStreamServiceHandler(srv), _ops.HTTP...)
	RegisterCommentIssueStreamServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.dop.issue.CommentIssueStreamService",
	)
}

var (
	commentIssueStreamServiceClientType  = reflect.TypeOf((*CommentIssueStreamServiceClient)(nil)).Elem()
	commentIssueStreamServiceServerType  = reflect.TypeOf((*CommentIssueStreamServiceServer)(nil)).Elem()
	commentIssueStreamServiceHandlerType = reflect.TypeOf((*CommentIssueStreamServiceHandler)(nil)).Elem()
)

// CommentIssueStreamServiceClientType .
func CommentIssueStreamServiceClientType() reflect.Type { return commentIssueStreamServiceClientType }

// CommentIssueStreamServiceServerType .
func CommentIssueStreamServiceServerType() reflect.Type { return commentIssueStreamServiceServerType }

// CommentIssueStreamServiceHandlerType .
func CommentIssueStreamServiceHandlerType() reflect.Type { return commentIssueStreamServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		commentIssueStreamServiceClientType,
		// server types
		commentIssueStreamServiceServerType,
		// handler types
		commentIssueStreamServiceHandlerType,
	}
}