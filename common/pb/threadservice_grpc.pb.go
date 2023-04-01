// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: threadService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ThreadServiceClient is the client API for ThreadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThreadServiceClient interface {
	Get(ctx context.Context, in *GetThreadRequest, opts ...grpc.CallOption) (*GetThreadResponse, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error)
	Create(ctx context.Context, in *CreateThreadReq, opts ...grpc.CallOption) (*Empty, error)
	AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*Empty, error)
}

type threadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThreadServiceClient(cc grpc.ClientConnInterface) ThreadServiceClient {
	return &threadServiceClient{cc}
}

func (c *threadServiceClient) Get(ctx context.Context, in *GetThreadRequest, opts ...grpc.CallOption) (*GetThreadResponse, error) {
	out := new(GetThreadResponse)
	err := c.cc.Invoke(ctx, "/chitchat.ThreadService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/chitchat.ThreadService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadServiceClient) Create(ctx context.Context, in *CreateThreadReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chitchat.ThreadService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadServiceClient) AddPost(ctx context.Context, in *AddPostRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chitchat.ThreadService/AddPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThreadServiceServer is the server API for ThreadService service.
// All implementations must embed UnimplementedThreadServiceServer
// for forward compatibility
type ThreadServiceServer interface {
	Get(context.Context, *GetThreadRequest) (*GetThreadResponse, error)
	GetAll(context.Context, *Empty) (*GetAllResponse, error)
	Create(context.Context, *CreateThreadReq) (*Empty, error)
	AddPost(context.Context, *AddPostRequest) (*Empty, error)
	mustEmbedUnimplementedThreadServiceServer()
}

// UnimplementedThreadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThreadServiceServer struct {
}

func (UnimplementedThreadServiceServer) Get(context.Context, *GetThreadRequest) (*GetThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedThreadServiceServer) GetAll(context.Context, *Empty) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedThreadServiceServer) Create(context.Context, *CreateThreadReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedThreadServiceServer) AddPost(context.Context, *AddPostRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPost not implemented")
}
func (UnimplementedThreadServiceServer) mustEmbedUnimplementedThreadServiceServer() {}

// UnsafeThreadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThreadServiceServer will
// result in compilation errors.
type UnsafeThreadServiceServer interface {
	mustEmbedUnimplementedThreadServiceServer()
}

func RegisterThreadServiceServer(s grpc.ServiceRegistrar, srv ThreadServiceServer) {
	s.RegisterService(&ThreadService_ServiceDesc, srv)
}

func _ThreadService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chitchat.ThreadService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServiceServer).Get(ctx, req.(*GetThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThreadService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chitchat.ThreadService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThreadService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateThreadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chitchat.ThreadService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServiceServer).Create(ctx, req.(*CreateThreadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThreadService_AddPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServiceServer).AddPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chitchat.ThreadService/AddPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServiceServer).AddPost(ctx, req.(*AddPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThreadService_ServiceDesc is the grpc.ServiceDesc for ThreadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThreadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chitchat.ThreadService",
	HandlerType: (*ThreadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ThreadService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ThreadService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ThreadService_Create_Handler,
		},
		{
			MethodName: "AddPost",
			Handler:    _ThreadService_AddPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "threadService.proto",
}
