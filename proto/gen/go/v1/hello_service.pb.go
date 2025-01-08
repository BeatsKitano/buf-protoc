// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: v1/hello_service.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求参数
type Req struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 年龄
	Age int32 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	// 购买日期
	PurchaseDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	// 交付日期
	DeliveryDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	// 物品名
	Name          string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Req) Reset() {
	*x = Req{}
	mi := &file_v1_hello_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hello_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_v1_hello_service_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Req) GetPurchaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PurchaseDate
	}
	return nil
}

func (x *Req) GetDeliveryDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveryDate
	}
	return nil
}

func (x *Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 用户信息
type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_v1_hello_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hello_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_v1_hello_service_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_v1_hello_service_proto protoreflect.FileDescriptor

var file_v1_hello_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x4c,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x3a, 0xba, 0x48, 0x37,
	0xba, 0x01, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x2e, 0x61, 0x67, 0x65, 0x1a, 0x29, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x3c, 0x20, 0x31, 0x38, 0x20, 0x3f, 0x20, 0x27, 0xe7, 0x94, 0xa8, 0xe6, 0x88,
	0xb7, 0xe8, 0x87, 0xb3, 0xe5, 0xb0, 0x91, 0xe5, 0xb9, 0xb4, 0xe6, 0xbb, 0xa1, 0x31, 0x38, 0xe5,
	0xb2, 0x81, 0x27, 0x3a, 0x20, 0x27, 0x27, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0d,
	0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a,
	0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x95, 0x01, 0xea, 0x41, 0x24, 0x0a, 0x0c, 0x78,
	0x2e, 0x78, 0x2e, 0x78, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x67,
	0x7d, 0xba, 0x48, 0x6b, 0x1a, 0x69, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x2e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x20, 0x64, 0x61, 0x74, 0x65, 0x20, 0xe5, 0xbf, 0x85, 0xe9, 0xa1, 0xbb, 0xe5,
	0x9c, 0xa8, 0x20, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x20, 0x64, 0x61, 0x74, 0x65,
	0xe4, 0xb9, 0x8b, 0xe5, 0x90, 0x8e, 0x1a, 0x27, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x20, 0x3e, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x3a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x32, 0xd4, 0x01, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x49, 0xea, 0x80, 0x1f, 0x28, 0x90, 0xea, 0x30, 0x01, 0x9a, 0xea, 0x30, 0x0f,
	0x68, 0x67, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x73, 0x65, 0x74, 0xa0,
	0xea, 0x30, 0x01, 0xa8, 0xea, 0x30, 0x01, 0xb0, 0xea, 0x30, 0x0a, 0xb8, 0xea, 0x30, 0xe8, 0x07,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x53, 0xca,
	0x41, 0x15, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x68, 0x75, 0x69, 0x67, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x1f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x68, 0x75, 0x69, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x8a, 0xd4, 0xdb, 0xd2, 0x0f, 0x13, 0x76,
	0x31, 0x5f, 0x32, 0x30, 0x32, 0x35, 0x30, 0x38, 0x32, 0x31, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x42, 0x60, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_hello_service_proto_rawDescOnce sync.Once
	file_v1_hello_service_proto_rawDescData = file_v1_hello_service_proto_rawDesc
)

func file_v1_hello_service_proto_rawDescGZIP() []byte {
	file_v1_hello_service_proto_rawDescOnce.Do(func() {
		file_v1_hello_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_hello_service_proto_rawDescData)
	})
	return file_v1_hello_service_proto_rawDescData
}

var file_v1_hello_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_hello_service_proto_goTypes = []any{
	(*Req)(nil),                   // 0: api.v1.Req
	(*User)(nil),                  // 1: api.v1.User
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_v1_hello_service_proto_depIdxs = []int32{
	2, // 0: api.v1.Req.purchase_date:type_name -> google.protobuf.Timestamp
	2, // 1: api.v1.Req.delivery_date:type_name -> google.protobuf.Timestamp
	0, // 2: api.v1.HelloService.GetUser:input_type -> api.v1.Req
	1, // 3: api.v1.HelloService.GetUser:output_type -> api.v1.User
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_hello_service_proto_init() }
func file_v1_hello_service_proto_init() {
	if File_v1_hello_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_hello_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_hello_service_proto_goTypes,
		DependencyIndexes: file_v1_hello_service_proto_depIdxs,
		MessageInfos:      file_v1_hello_service_proto_msgTypes,
	}.Build()
	File_v1_hello_service_proto = out.File
	file_v1_hello_service_proto_rawDesc = nil
	file_v1_hello_service_proto_goTypes = nil
	file_v1_hello_service_proto_depIdxs = nil
}
