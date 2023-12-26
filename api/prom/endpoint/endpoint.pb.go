// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: prom/endpoint/endpoint.proto

package endpoint

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "prometheus-manager/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 添加监控端点请求参数
type AppendEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 监控端点名称
	AgentName string `protobuf:"bytes,1,opt,name=agentName,proto3" json:"agentName,omitempty"`
	// 监控端点地址列表信息, 最少有一个, 不允许为空对象
	Endpoints []*api.PrometheusServerItem `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *AppendEndpointRequest) Reset() {
	*x = AppendEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_endpoint_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEndpointRequest) ProtoMessage() {}

func (x *AppendEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_endpoint_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEndpointRequest.ProtoReflect.Descriptor instead.
func (*AppendEndpointRequest) Descriptor() ([]byte, []int) {
	return file_prom_endpoint_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *AppendEndpointRequest) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *AppendEndpointRequest) GetEndpoints() []*api.PrometheusServerItem {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// 添加监控端点响应参数
type AppendEndpointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应信息
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AppendEndpointReply) Reset() {
	*x = AppendEndpointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_endpoint_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEndpointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEndpointReply) ProtoMessage() {}

func (x *AppendEndpointReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_endpoint_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEndpointReply.ProtoReflect.Descriptor instead.
func (*AppendEndpointReply) Descriptor() ([]byte, []int) {
	return file_prom_endpoint_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEndpointReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

// 删除监控端点请求参数
type DeleteEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 监控端点名称
	AgentName string `protobuf:"bytes,1,opt,name=agentName,proto3" json:"agentName,omitempty"`
	// 要删除的监控端点uuid列表, 最少有一个
	Uuids []string `protobuf:"bytes,2,rep,name=uuids,proto3" json:"uuids,omitempty"`
}

func (x *DeleteEndpointRequest) Reset() {
	*x = DeleteEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_endpoint_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointRequest) ProtoMessage() {}

func (x *DeleteEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_endpoint_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointRequest.ProtoReflect.Descriptor instead.
func (*DeleteEndpointRequest) Descriptor() ([]byte, []int) {
	return file_prom_endpoint_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteEndpointRequest) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *DeleteEndpointRequest) GetUuids() []string {
	if x != nil {
		return x.Uuids
	}
	return nil
}

// 删除监控端点响应参数
type DeleteEndpointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应信息
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteEndpointReply) Reset() {
	*x = DeleteEndpointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_endpoint_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEndpointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointReply) ProtoMessage() {}

func (x *DeleteEndpointReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_endpoint_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointReply.ProtoReflect.Descriptor instead.
func (*DeleteEndpointReply) Descriptor() ([]byte, []int) {
	return file_prom_endpoint_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteEndpointReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

// 获取监控端点列表请求参数, 这里每次拉取都是全部, 所以不做分页查询等
type ListEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEndpointRequest) Reset() {
	*x = ListEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_endpoint_endpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointRequest) ProtoMessage() {}

func (x *ListEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prom_endpoint_endpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointRequest.ProtoReflect.Descriptor instead.
func (*ListEndpointRequest) Descriptor() ([]byte, []int) {
	return file_prom_endpoint_endpoint_proto_rawDescGZIP(), []int{4}
}

type ListEndpointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应信息
	Response *api.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	// 监控端点列表
	Endpoints []*api.PrometheusServerItem `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *ListEndpointReply) Reset() {
	*x = ListEndpointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prom_endpoint_endpoint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEndpointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointReply) ProtoMessage() {}

func (x *ListEndpointReply) ProtoReflect() protoreflect.Message {
	mi := &file_prom_endpoint_endpoint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointReply.ProtoReflect.Descriptor instead.
func (*ListEndpointReply) Descriptor() ([]byte, []int) {
	return file_prom_endpoint_endpoint_proto_rawDescGZIP(), []int{5}
}

func (x *ListEndpointReply) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ListEndpointReply) GetEndpoints() []*api.PrometheusServerItem {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_prom_endpoint_endpoint_proto protoreflect.FileDescriptor

var file_prom_endpoint_endpoint_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x20, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01, 0x09, 0x08, 0x01, 0x22, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x40,
	0x0a, 0x13, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x67, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x20, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01, 0x09, 0x08, 0x01, 0x22, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x22, 0x40, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x77, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x32, 0x9c, 0x03, 0x0a, 0x08,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7e, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x44, 0x0a, 0x11, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50,
	0x01, 0x5a, 0x2d, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2d, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prom_endpoint_endpoint_proto_rawDescOnce sync.Once
	file_prom_endpoint_endpoint_proto_rawDescData = file_prom_endpoint_endpoint_proto_rawDesc
)

func file_prom_endpoint_endpoint_proto_rawDescGZIP() []byte {
	file_prom_endpoint_endpoint_proto_rawDescOnce.Do(func() {
		file_prom_endpoint_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_prom_endpoint_endpoint_proto_rawDescData)
	})
	return file_prom_endpoint_endpoint_proto_rawDescData
}

var file_prom_endpoint_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_prom_endpoint_endpoint_proto_goTypes = []interface{}{
	(*AppendEndpointRequest)(nil),    // 0: api.prom.endpoint.AppendEndpointRequest
	(*AppendEndpointReply)(nil),      // 1: api.prom.endpoint.AppendEndpointReply
	(*DeleteEndpointRequest)(nil),    // 2: api.prom.endpoint.DeleteEndpointRequest
	(*DeleteEndpointReply)(nil),      // 3: api.prom.endpoint.DeleteEndpointReply
	(*ListEndpointRequest)(nil),      // 4: api.prom.endpoint.ListEndpointRequest
	(*ListEndpointReply)(nil),        // 5: api.prom.endpoint.ListEndpointReply
	(*api.PrometheusServerItem)(nil), // 6: api.PrometheusServerItem
	(*api.Response)(nil),             // 7: api.Response
}
var file_prom_endpoint_endpoint_proto_depIdxs = []int32{
	6, // 0: api.prom.endpoint.AppendEndpointRequest.endpoints:type_name -> api.PrometheusServerItem
	7, // 1: api.prom.endpoint.AppendEndpointReply.response:type_name -> api.Response
	7, // 2: api.prom.endpoint.DeleteEndpointReply.response:type_name -> api.Response
	7, // 3: api.prom.endpoint.ListEndpointReply.response:type_name -> api.Response
	6, // 4: api.prom.endpoint.ListEndpointReply.endpoints:type_name -> api.PrometheusServerItem
	0, // 5: api.prom.endpoint.Endpoint.AppendEndpoint:input_type -> api.prom.endpoint.AppendEndpointRequest
	2, // 6: api.prom.endpoint.Endpoint.DeleteEndpoint:input_type -> api.prom.endpoint.DeleteEndpointRequest
	4, // 7: api.prom.endpoint.Endpoint.ListEndpoint:input_type -> api.prom.endpoint.ListEndpointRequest
	1, // 8: api.prom.endpoint.Endpoint.AppendEndpoint:output_type -> api.prom.endpoint.AppendEndpointReply
	3, // 9: api.prom.endpoint.Endpoint.DeleteEndpoint:output_type -> api.prom.endpoint.DeleteEndpointReply
	5, // 10: api.prom.endpoint.Endpoint.ListEndpoint:output_type -> api.prom.endpoint.ListEndpointReply
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_prom_endpoint_endpoint_proto_init() }
func file_prom_endpoint_endpoint_proto_init() {
	if File_prom_endpoint_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prom_endpoint_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEndpointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prom_endpoint_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEndpointReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prom_endpoint_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEndpointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prom_endpoint_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEndpointReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prom_endpoint_endpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prom_endpoint_endpoint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEndpointReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prom_endpoint_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prom_endpoint_endpoint_proto_goTypes,
		DependencyIndexes: file_prom_endpoint_endpoint_proto_depIdxs,
		MessageInfos:      file_prom_endpoint_endpoint_proto_msgTypes,
	}.Build()
	File_prom_endpoint_endpoint_proto = out.File
	file_prom_endpoint_endpoint_proto_rawDesc = nil
	file_prom_endpoint_endpoint_proto_goTypes = nil
	file_prom_endpoint_endpoint_proto_depIdxs = nil
}