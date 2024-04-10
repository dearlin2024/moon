// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: perrors/error.proto

package perrors

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Err int32

const (
	// 未知错误
	Err_UNKNOWN Err = 0
	// 参数错误
	Err_INVALID_PARAMS Err = 1
	// 资源不存在
	Err_NOT_FOUND Err = 2
	// 数据已存在
	Err_ALREADY_EXISTS Err = 3
	// 权限不足
	Err_PERMISSION_DENIED Err = 4
	// 未登录
	Err_UNAUTHORIZED Err = 5
)

// Enum value maps for Err.
var (
	Err_name = map[int32]string{
		0: "UNKNOWN",
		1: "INVALID_PARAMS",
		2: "NOT_FOUND",
		3: "ALREADY_EXISTS",
		4: "PERMISSION_DENIED",
		5: "UNAUTHORIZED",
	}
	Err_value = map[string]int32{
		"UNKNOWN":           0,
		"INVALID_PARAMS":    1,
		"NOT_FOUND":         2,
		"ALREADY_EXISTS":    3,
		"PERMISSION_DENIED": 4,
		"UNAUTHORIZED":      5,
	}
)

func (x Err) Enum() *Err {
	p := new(Err)
	*p = x
	return p
}

func (x Err) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Err) Descriptor() protoreflect.EnumDescriptor {
	return file_perrors_error_proto_enumTypes[0].Descriptor()
}

func (Err) Type() protoreflect.EnumType {
	return &file_perrors_error_proto_enumTypes[0]
}

func (x Err) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Err.Descriptor instead.
func (Err) EnumDescriptor() ([]byte, []int) {
	return file_perrors_error_proto_rawDescGZIP(), []int{0}
}

var File_perrors_error_proto protoreflect.FileDescriptor

var file_perrors_error_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9c, 0x01, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x12,
	0x11, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45,
	0xf4, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x53, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x13, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94,
	0x03, 0x12, 0x18, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x53, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44,
	0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55,
	0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xb0, 0x04, 0x42, 0x40, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x64, 0x65, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2f,
	0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x3b, 0x70, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_perrors_error_proto_rawDescOnce sync.Once
	file_perrors_error_proto_rawDescData = file_perrors_error_proto_rawDesc
)

func file_perrors_error_proto_rawDescGZIP() []byte {
	file_perrors_error_proto_rawDescOnce.Do(func() {
		file_perrors_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_perrors_error_proto_rawDescData)
	})
	return file_perrors_error_proto_rawDescData
}

var file_perrors_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_perrors_error_proto_goTypes = []interface{}{
	(Err)(0), // 0: api.perrors.Err
}
var file_perrors_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_perrors_error_proto_init() }
func file_perrors_error_proto_init() {
	if File_perrors_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_perrors_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_perrors_error_proto_goTypes,
		DependencyIndexes: file_perrors_error_proto_depIdxs,
		EnumInfos:         file_perrors_error_proto_enumTypes,
	}.Build()
	File_perrors_error_proto = out.File
	file_perrors_error_proto_rawDesc = nil
	file_perrors_error_proto_goTypes = nil
	file_perrors_error_proto_depIdxs = nil
}
