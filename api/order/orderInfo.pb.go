// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: order/orderInfo.proto

package order

import (
	base "github.com/jianglianggithub/proto/api/base"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             *base.BaseReq          `protobuf:"bytes,5,opt,name=a,proto3" json:"a,omitempty"`
	Passport      string                 `protobuf:"bytes,1,opt,name=Passport,proto3" json:"Passport,omitempty"` // v: required
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"` // v: required
	Nickname      string                 `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"` // v: required
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	mi := &file_order_orderInfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_orderInfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_order_orderInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetA() *base.BaseReq {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *CreateReq) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *CreateReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type CreateRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            string                 `protobuf:"bytes,1,opt,name=Ok,proto3" json:"Ok,omitempty"` // status
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	mi := &file_order_orderInfo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_order_orderInfo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_order_orderInfo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRes) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

var File_order_orderInfo_proto protoreflect.FileDescriptor

const file_order_orderInfo_proto_rawDesc = "" +
	"\n" +
	"\x15order/orderInfo.proto\x12\x05order\x1a\x11base/common.proto\"|\n" +
	"\tCreateReq\x12\x1b\n" +
	"\x01a\x18\x05 \x01(\v2\r.base.BaseReqR\x01a\x12\x1a\n" +
	"\bPassport\x18\x01 \x01(\tR\bPassport\x12\x1a\n" +
	"\bPassword\x18\x02 \x01(\tR\bPassword\x12\x1a\n" +
	"\bNickname\x18\x03 \x01(\tR\bNickname\"\x1b\n" +
	"\tCreateRes\x12\x0e\n" +
	"\x02Ok\x18\x01 \x01(\tR\x02Ok2<\n" +
	"\tOrderInfo\x12/\n" +
	"\aCreate1\x12\x10.order.CreateReq\x1a\x10.order.CreateRes\"\x00B-Z+github.com/jianglianggithub/proto/api/orderb\x06proto3"

var (
	file_order_orderInfo_proto_rawDescOnce sync.Once
	file_order_orderInfo_proto_rawDescData []byte
)

func file_order_orderInfo_proto_rawDescGZIP() []byte {
	file_order_orderInfo_proto_rawDescOnce.Do(func() {
		file_order_orderInfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_orderInfo_proto_rawDesc), len(file_order_orderInfo_proto_rawDesc)))
	})
	return file_order_orderInfo_proto_rawDescData
}

var file_order_orderInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_orderInfo_proto_goTypes = []any{
	(*CreateReq)(nil),    // 0: order.CreateReq
	(*CreateRes)(nil),    // 1: order.CreateRes
	(*base.BaseReq)(nil), // 2: base.BaseReq
}
var file_order_orderInfo_proto_depIdxs = []int32{
	2, // 0: order.CreateReq.a:type_name -> base.BaseReq
	0, // 1: order.OrderInfo.Create1:input_type -> order.CreateReq
	1, // 2: order.OrderInfo.Create1:output_type -> order.CreateRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_orderInfo_proto_init() }
func file_order_orderInfo_proto_init() {
	if File_order_orderInfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_orderInfo_proto_rawDesc), len(file_order_orderInfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_orderInfo_proto_goTypes,
		DependencyIndexes: file_order_orderInfo_proto_depIdxs,
		MessageInfos:      file_order_orderInfo_proto_msgTypes,
	}.Build()
	File_order_orderInfo_proto = out.File
	file_order_orderInfo_proto_goTypes = nil
	file_order_orderInfo_proto_depIdxs = nil
}
