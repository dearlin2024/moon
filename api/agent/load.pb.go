// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: agent/load.proto

package agent

import (
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

type StrategyGroupAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StrategyGroupAllRequest) Reset() {
	*x = StrategyGroupAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyGroupAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyGroupAllRequest) ProtoMessage() {}

func (x *StrategyGroupAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyGroupAllRequest.ProtoReflect.Descriptor instead.
func (*StrategyGroupAllRequest) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{0}
}

type StrategyGroupAllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupList []*GroupSimple `protobuf:"bytes,1,rep,name=groupList,proto3" json:"groupList,omitempty"`
}

func (x *StrategyGroupAllReply) Reset() {
	*x = StrategyGroupAllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyGroupAllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyGroupAllReply) ProtoMessage() {}

func (x *StrategyGroupAllReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyGroupAllReply.ProtoReflect.Descriptor instead.
func (*StrategyGroupAllReply) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{1}
}

func (x *StrategyGroupAllReply) GetGroupList() []*GroupSimple {
	if x != nil {
		return x.GroupList
	}
	return nil
}

type GroupSimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 策略组ID
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 策略组名称
	Name       string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Strategies []*StrategySimple `protobuf:"bytes,3,rep,name=strategies,proto3" json:"strategies,omitempty"`
}

func (x *GroupSimple) Reset() {
	*x = GroupSimple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSimple) ProtoMessage() {}

func (x *GroupSimple) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSimple.ProtoReflect.Descriptor instead.
func (*GroupSimple) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{2}
}

func (x *GroupSimple) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupSimple) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupSimple) GetStrategies() []*StrategySimple {
	if x != nil {
		return x.Strategies
	}
	return nil
}

type StrategySimple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 策略ID
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 告警名称
	Alert string `protobuf:"bytes,2,opt,name=alert,proto3" json:"alert,omitempty"`
	// 表达式
	Expr string `protobuf:"bytes,3,opt,name=expr,proto3" json:"expr,omitempty"`
	// 持续时间
	Duration *api.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// 标签
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 注解
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 策略组ID
	GroupId uint32 `protobuf:"varint,7,opt,name=groupId,proto3" json:"groupId,omitempty"`
	// 告警级别ID
	AlarmLevelId uint32 `protobuf:"varint,8,opt,name=alarmLevelId,proto3" json:"alarmLevelId,omitempty"`
	// 数据源
	Endpoint string `protobuf:"bytes,9,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *StrategySimple) Reset() {
	*x = StrategySimple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategySimple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategySimple) ProtoMessage() {}

func (x *StrategySimple) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategySimple.ProtoReflect.Descriptor instead.
func (*StrategySimple) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{3}
}

func (x *StrategySimple) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StrategySimple) GetAlert() string {
	if x != nil {
		return x.Alert
	}
	return ""
}

func (x *StrategySimple) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *StrategySimple) GetDuration() *api.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *StrategySimple) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *StrategySimple) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *StrategySimple) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *StrategySimple) GetAlarmLevelId() uint32 {
	if x != nil {
		return x.AlarmLevelId
	}
	return 0
}

func (x *StrategySimple) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type StrategyGroupDiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StrategyGroupDiffRequest) Reset() {
	*x = StrategyGroupDiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyGroupDiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyGroupDiffRequest) ProtoMessage() {}

func (x *StrategyGroupDiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyGroupDiffRequest.ProtoReflect.Descriptor instead.
func (*StrategyGroupDiffRequest) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{4}
}

type StrategyGroupDiffReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppendItems []*GroupSimple `protobuf:"bytes,1,rep,name=appendItems,proto3" json:"appendItems,omitempty"`
	DeleteItems []*GroupSimple `protobuf:"bytes,2,rep,name=deleteItems,proto3" json:"deleteItems,omitempty"`
	UpdateItems []*GroupSimple `protobuf:"bytes,3,rep,name=updateItems,proto3" json:"updateItems,omitempty"`
}

func (x *StrategyGroupDiffReply) Reset() {
	*x = StrategyGroupDiffReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyGroupDiffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyGroupDiffReply) ProtoMessage() {}

func (x *StrategyGroupDiffReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyGroupDiffReply.ProtoReflect.Descriptor instead.
func (*StrategyGroupDiffReply) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{5}
}

func (x *StrategyGroupDiffReply) GetAppendItems() []*GroupSimple {
	if x != nil {
		return x.AppendItems
	}
	return nil
}

func (x *StrategyGroupDiffReply) GetDeleteItems() []*GroupSimple {
	if x != nil {
		return x.DeleteItems
	}
	return nil
}

func (x *StrategyGroupDiffReply) GetUpdateItems() []*GroupSimple {
	if x != nil {
		return x.UpdateItems
	}
	return nil
}

type EvaluateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupList []*GroupSimple `protobuf:"bytes,1,rep,name=groupList,proto3" json:"groupList,omitempty"`
}

func (x *EvaluateRequest) Reset() {
	*x = EvaluateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateRequest) ProtoMessage() {}

func (x *EvaluateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateRequest.ProtoReflect.Descriptor instead.
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{6}
}

func (x *EvaluateRequest) GetGroupList() []*GroupSimple {
	if x != nil {
		return x.GroupList
	}
	return nil
}

type EvaluateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EvaluateReply) Reset() {
	*x = EvaluateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_load_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateReply) ProtoMessage() {}

func (x *EvaluateReply) ProtoReflect() protoreflect.Message {
	mi := &file_agent_load_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateReply.ProtoReflect.Descriptor instead.
func (*EvaluateReply) Descriptor() ([]byte, []int) {
	return file_agent_load_proto_rawDescGZIP(), []int{7}
}

var File_agent_load_proto protoreflect.FileDescriptor

var file_agent_load_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x0b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x22, 0xd7, 0x03, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65,
	0x78, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12,
	0x29, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1a, 0x0a, 0x18,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x0b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x38, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x47, 0x0a, 0x0f, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xff, 0x01, 0x0a, 0x04,
	0x4c, 0x6f, 0x61, 0x64, 0x12, 0x58, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b,
	0x0a, 0x11, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x69, 0x66, 0x66, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x31, 0x0a,
	0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x22, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_load_proto_rawDescOnce sync.Once
	file_agent_load_proto_rawDescData = file_agent_load_proto_rawDesc
)

func file_agent_load_proto_rawDescGZIP() []byte {
	file_agent_load_proto_rawDescOnce.Do(func() {
		file_agent_load_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_load_proto_rawDescData)
	})
	return file_agent_load_proto_rawDescData
}

var file_agent_load_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_agent_load_proto_goTypes = []interface{}{
	(*StrategyGroupAllRequest)(nil),  // 0: api.agent.StrategyGroupAllRequest
	(*StrategyGroupAllReply)(nil),    // 1: api.agent.StrategyGroupAllReply
	(*GroupSimple)(nil),              // 2: api.agent.GroupSimple
	(*StrategySimple)(nil),           // 3: api.agent.StrategySimple
	(*StrategyGroupDiffRequest)(nil), // 4: api.agent.StrategyGroupDiffRequest
	(*StrategyGroupDiffReply)(nil),   // 5: api.agent.StrategyGroupDiffReply
	(*EvaluateRequest)(nil),          // 6: api.agent.EvaluateRequest
	(*EvaluateReply)(nil),            // 7: api.agent.EvaluateReply
	nil,                              // 8: api.agent.StrategySimple.LabelsEntry
	nil,                              // 9: api.agent.StrategySimple.AnnotationsEntry
	(*api.Duration)(nil),             // 10: api.Duration
}
var file_agent_load_proto_depIdxs = []int32{
	2,  // 0: api.agent.StrategyGroupAllReply.groupList:type_name -> api.agent.GroupSimple
	3,  // 1: api.agent.GroupSimple.strategies:type_name -> api.agent.StrategySimple
	10, // 2: api.agent.StrategySimple.duration:type_name -> api.Duration
	8,  // 3: api.agent.StrategySimple.labels:type_name -> api.agent.StrategySimple.LabelsEntry
	9,  // 4: api.agent.StrategySimple.annotations:type_name -> api.agent.StrategySimple.AnnotationsEntry
	2,  // 5: api.agent.StrategyGroupDiffReply.appendItems:type_name -> api.agent.GroupSimple
	2,  // 6: api.agent.StrategyGroupDiffReply.deleteItems:type_name -> api.agent.GroupSimple
	2,  // 7: api.agent.StrategyGroupDiffReply.updateItems:type_name -> api.agent.GroupSimple
	2,  // 8: api.agent.EvaluateRequest.groupList:type_name -> api.agent.GroupSimple
	0,  // 9: api.agent.Load.StrategyGroupAll:input_type -> api.agent.StrategyGroupAllRequest
	4,  // 10: api.agent.Load.StrategyGroupDiff:input_type -> api.agent.StrategyGroupDiffRequest
	6,  // 11: api.agent.Load.Evaluate:input_type -> api.agent.EvaluateRequest
	1,  // 12: api.agent.Load.StrategyGroupAll:output_type -> api.agent.StrategyGroupAllReply
	5,  // 13: api.agent.Load.StrategyGroupDiff:output_type -> api.agent.StrategyGroupDiffReply
	7,  // 14: api.agent.Load.Evaluate:output_type -> api.agent.EvaluateReply
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_agent_load_proto_init() }
func file_agent_load_proto_init() {
	if File_agent_load_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_load_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyGroupAllRequest); i {
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
		file_agent_load_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyGroupAllReply); i {
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
		file_agent_load_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupSimple); i {
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
		file_agent_load_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategySimple); i {
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
		file_agent_load_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyGroupDiffRequest); i {
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
		file_agent_load_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyGroupDiffReply); i {
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
		file_agent_load_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateRequest); i {
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
		file_agent_load_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateReply); i {
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
			RawDescriptor: file_agent_load_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_load_proto_goTypes,
		DependencyIndexes: file_agent_load_proto_depIdxs,
		MessageInfos:      file_agent_load_proto_msgTypes,
	}.Build()
	File_agent_load_proto = out.File
	file_agent_load_proto_rawDesc = nil
	file_agent_load_proto_goTypes = nil
	file_agent_load_proto_depIdxs = nil
}
