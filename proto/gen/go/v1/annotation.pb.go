// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: v1/annotation.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthMethod int32

const (
	AuthMethod_AUTH_METHOD_UNSPECIFIED AuthMethod = 0
	// IAM uses the standard IAM authorization check on the organizational resources.
	AuthMethod_IAM AuthMethod = 1
	// Custom authorization method.
	AuthMethod_CUSTOM AuthMethod = 2
)

// Enum value maps for AuthMethod.
var (
	AuthMethod_name = map[int32]string{
		0: "AUTH_METHOD_UNSPECIFIED",
		1: "IAM",
		2: "CUSTOM",
	}
	AuthMethod_value = map[string]int32{
		"AUTH_METHOD_UNSPECIFIED": 0,
		"IAM":                     1,
		"CUSTOM":                  2,
	}
)

func (x AuthMethod) Enum() *AuthMethod {
	p := new(AuthMethod)
	*p = x
	return p
}

func (x AuthMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_annotation_proto_enumTypes[0].Descriptor()
}

func (AuthMethod) Type() protoreflect.EnumType {
	return &file_v1_annotation_proto_enumTypes[0]
}

func (x AuthMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMethod.Descriptor instead.
func (AuthMethod) EnumDescriptor() ([]byte, []int) {
	return file_v1_annotation_proto_rawDescGZIP(), []int{0}
}

var file_v1_annotation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         100000,
		Name:          "bytebase.v1.allow_without_credential",
		Tag:           "varint,100000,opt,name=allow_without_credential",
		Filename:      "v1/annotation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         100001,
		Name:          "bytebase.v1.permission",
		Tag:           "bytes,100001,opt,name=permission",
		Filename:      "v1/annotation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*AuthMethod)(nil),
		Field:         100002,
		Name:          "bytebase.v1.auth_method",
		Tag:           "varint,100002,opt,name=auth_method,enum=bytebase.v1.AuthMethod",
		Filename:      "v1/annotation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         100003,
		Name:          "bytebase.v1.audit",
		Tag:           "varint,100003,opt,name=audit",
		Filename:      "v1/annotation.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional bool allow_without_credential = 100000;
	E_AllowWithoutCredential = &file_v1_annotation_proto_extTypes[0]
	// optional string permission = 100001;
	E_Permission = &file_v1_annotation_proto_extTypes[1]
	// optional bytebase.v1.AuthMethod auth_method = 100002;
	E_AuthMethod = &file_v1_annotation_proto_extTypes[2]
	// optional bool audit = 100003;
	E_Audit = &file_v1_annotation_proto_extTypes[3]
)

var File_v1_annotation_proto protoreflect.FileDescriptor

var file_v1_annotation_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x3e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x41, 0x4d, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x10, 0x02, 0x3a, 0x5a, 0x0a, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xa0, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x57,
	0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x3a, 0x40, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa1,
	0x8d, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x3a, 0x5a, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa2, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x36,
	0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x77, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x06, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x42, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x42, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x42, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x42, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_annotation_proto_rawDescOnce sync.Once
	file_v1_annotation_proto_rawDescData = file_v1_annotation_proto_rawDesc
)

func file_v1_annotation_proto_rawDescGZIP() []byte {
	file_v1_annotation_proto_rawDescOnce.Do(func() {
		file_v1_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_annotation_proto_rawDescData)
	})
	return file_v1_annotation_proto_rawDescData
}

var file_v1_annotation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_annotation_proto_goTypes = []any{
	(AuthMethod)(0),                    // 0: bytebase.v1.AuthMethod
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_v1_annotation_proto_depIdxs = []int32{
	1, // 0: bytebase.v1.allow_without_credential:extendee -> google.protobuf.MethodOptions
	1, // 1: bytebase.v1.permission:extendee -> google.protobuf.MethodOptions
	1, // 2: bytebase.v1.auth_method:extendee -> google.protobuf.MethodOptions
	1, // 3: bytebase.v1.audit:extendee -> google.protobuf.MethodOptions
	0, // 4: bytebase.v1.auth_method:type_name -> bytebase.v1.AuthMethod
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_annotation_proto_init() }
func file_v1_annotation_proto_init() {
	if File_v1_annotation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_annotation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_v1_annotation_proto_goTypes,
		DependencyIndexes: file_v1_annotation_proto_depIdxs,
		EnumInfos:         file_v1_annotation_proto_enumTypes,
		ExtensionInfos:    file_v1_annotation_proto_extTypes,
	}.Build()
	File_v1_annotation_proto = out.File
	file_v1_annotation_proto_rawDesc = nil
	file_v1_annotation_proto_goTypes = nil
	file_v1_annotation_proto_depIdxs = nil
}
