// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: server/prom/notify/chat_group.proto

package notify

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

const (
	ChatGroup_CreateChatGroup_FullMethodName = "/api.server.prom.notify.ChatGroup/CreateChatGroup"
	ChatGroup_UpdateChatGroup_FullMethodName = "/api.server.prom.notify.ChatGroup/UpdateChatGroup"
	ChatGroup_DeleteChatGroup_FullMethodName = "/api.server.prom.notify.ChatGroup/DeleteChatGroup"
	ChatGroup_GetChatGroup_FullMethodName    = "/api.server.prom.notify.ChatGroup/GetChatGroup"
	ChatGroup_ListChatGroup_FullMethodName   = "/api.server.prom.notify.ChatGroup/ListChatGroup"
	ChatGroup_SelectChatGroup_FullMethodName = "/api.server.prom.notify.ChatGroup/SelectChatGroup"
)

// ChatGroupClient is the client API for ChatGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatGroupClient interface {
	// 创建通知群组
	CreateChatGroup(ctx context.Context, in *CreateChatGroupRequest, opts ...grpc.CallOption) (*CreateChatGroupReply, error)
	// 更新通知群组
	UpdateChatGroup(ctx context.Context, in *UpdateChatGroupRequest, opts ...grpc.CallOption) (*UpdateChatGroupReply, error)
	// 删除通知群组
	DeleteChatGroup(ctx context.Context, in *DeleteChatGroupRequest, opts ...grpc.CallOption) (*DeleteChatGroupReply, error)
	// 获取通知群组
	GetChatGroup(ctx context.Context, in *GetChatGroupRequest, opts ...grpc.CallOption) (*GetChatGroupReply, error)
	// 获取通知群组列表
	ListChatGroup(ctx context.Context, in *ListChatGroupRequest, opts ...grpc.CallOption) (*ListChatGroupReply, error)
	// 获取通知群组列表(下拉选择)
	SelectChatGroup(ctx context.Context, in *SelectChatGroupRequest, opts ...grpc.CallOption) (*SelectChatGroupReply, error)
}

type chatGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewChatGroupClient(cc grpc.ClientConnInterface) ChatGroupClient {
	return &chatGroupClient{cc}
}

func (c *chatGroupClient) CreateChatGroup(ctx context.Context, in *CreateChatGroupRequest, opts ...grpc.CallOption) (*CreateChatGroupReply, error) {
	out := new(CreateChatGroupReply)
	err := c.cc.Invoke(ctx, ChatGroup_CreateChatGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGroupClient) UpdateChatGroup(ctx context.Context, in *UpdateChatGroupRequest, opts ...grpc.CallOption) (*UpdateChatGroupReply, error) {
	out := new(UpdateChatGroupReply)
	err := c.cc.Invoke(ctx, ChatGroup_UpdateChatGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGroupClient) DeleteChatGroup(ctx context.Context, in *DeleteChatGroupRequest, opts ...grpc.CallOption) (*DeleteChatGroupReply, error) {
	out := new(DeleteChatGroupReply)
	err := c.cc.Invoke(ctx, ChatGroup_DeleteChatGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGroupClient) GetChatGroup(ctx context.Context, in *GetChatGroupRequest, opts ...grpc.CallOption) (*GetChatGroupReply, error) {
	out := new(GetChatGroupReply)
	err := c.cc.Invoke(ctx, ChatGroup_GetChatGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGroupClient) ListChatGroup(ctx context.Context, in *ListChatGroupRequest, opts ...grpc.CallOption) (*ListChatGroupReply, error) {
	out := new(ListChatGroupReply)
	err := c.cc.Invoke(ctx, ChatGroup_ListChatGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGroupClient) SelectChatGroup(ctx context.Context, in *SelectChatGroupRequest, opts ...grpc.CallOption) (*SelectChatGroupReply, error) {
	out := new(SelectChatGroupReply)
	err := c.cc.Invoke(ctx, ChatGroup_SelectChatGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatGroupServer is the server API for ChatGroup service.
// All implementations must embed UnimplementedChatGroupServer
// for forward compatibility
type ChatGroupServer interface {
	// 创建通知群组
	CreateChatGroup(context.Context, *CreateChatGroupRequest) (*CreateChatGroupReply, error)
	// 更新通知群组
	UpdateChatGroup(context.Context, *UpdateChatGroupRequest) (*UpdateChatGroupReply, error)
	// 删除通知群组
	DeleteChatGroup(context.Context, *DeleteChatGroupRequest) (*DeleteChatGroupReply, error)
	// 获取通知群组
	GetChatGroup(context.Context, *GetChatGroupRequest) (*GetChatGroupReply, error)
	// 获取通知群组列表
	ListChatGroup(context.Context, *ListChatGroupRequest) (*ListChatGroupReply, error)
	// 获取通知群组列表(下拉选择)
	SelectChatGroup(context.Context, *SelectChatGroupRequest) (*SelectChatGroupReply, error)
	mustEmbedUnimplementedChatGroupServer()
}

// UnimplementedChatGroupServer must be embedded to have forward compatible implementations.
type UnimplementedChatGroupServer struct {
}

func (UnimplementedChatGroupServer) CreateChatGroup(context.Context, *CreateChatGroupRequest) (*CreateChatGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatGroup not implemented")
}
func (UnimplementedChatGroupServer) UpdateChatGroup(context.Context, *UpdateChatGroupRequest) (*UpdateChatGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatGroup not implemented")
}
func (UnimplementedChatGroupServer) DeleteChatGroup(context.Context, *DeleteChatGroupRequest) (*DeleteChatGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatGroup not implemented")
}
func (UnimplementedChatGroupServer) GetChatGroup(context.Context, *GetChatGroupRequest) (*GetChatGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatGroup not implemented")
}
func (UnimplementedChatGroupServer) ListChatGroup(context.Context, *ListChatGroupRequest) (*ListChatGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChatGroup not implemented")
}
func (UnimplementedChatGroupServer) SelectChatGroup(context.Context, *SelectChatGroupRequest) (*SelectChatGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectChatGroup not implemented")
}
func (UnimplementedChatGroupServer) mustEmbedUnimplementedChatGroupServer() {}

// UnsafeChatGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatGroupServer will
// result in compilation errors.
type UnsafeChatGroupServer interface {
	mustEmbedUnimplementedChatGroupServer()
}

func RegisterChatGroupServer(s grpc.ServiceRegistrar, srv ChatGroupServer) {
	s.RegisterService(&ChatGroup_ServiceDesc, srv)
}

func _ChatGroup_CreateChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGroupServer).CreateChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatGroup_CreateChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGroupServer).CreateChatGroup(ctx, req.(*CreateChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGroup_UpdateChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGroupServer).UpdateChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatGroup_UpdateChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGroupServer).UpdateChatGroup(ctx, req.(*UpdateChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGroup_DeleteChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGroupServer).DeleteChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatGroup_DeleteChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGroupServer).DeleteChatGroup(ctx, req.(*DeleteChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGroup_GetChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGroupServer).GetChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatGroup_GetChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGroupServer).GetChatGroup(ctx, req.(*GetChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGroup_ListChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGroupServer).ListChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatGroup_ListChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGroupServer).ListChatGroup(ctx, req.(*ListChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGroup_SelectChatGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectChatGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGroupServer).SelectChatGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatGroup_SelectChatGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGroupServer).SelectChatGroup(ctx, req.(*SelectChatGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatGroup_ServiceDesc is the grpc.ServiceDesc for ChatGroup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatGroup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.prom.notify.ChatGroup",
	HandlerType: (*ChatGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChatGroup",
			Handler:    _ChatGroup_CreateChatGroup_Handler,
		},
		{
			MethodName: "UpdateChatGroup",
			Handler:    _ChatGroup_UpdateChatGroup_Handler,
		},
		{
			MethodName: "DeleteChatGroup",
			Handler:    _ChatGroup_DeleteChatGroup_Handler,
		},
		{
			MethodName: "GetChatGroup",
			Handler:    _ChatGroup_GetChatGroup_Handler,
		},
		{
			MethodName: "ListChatGroup",
			Handler:    _ChatGroup_ListChatGroup_Handler,
		},
		{
			MethodName: "SelectChatGroup",
			Handler:    _ChatGroup_SelectChatGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/prom/notify/chat_group.proto",
}