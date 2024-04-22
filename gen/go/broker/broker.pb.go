// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: broker/broker.proto

package brokerv1

import (
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

type BuyingProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ProductId int64  `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *BuyingProductRequest) Reset() {
	*x = BuyingProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyingProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyingProductRequest) ProtoMessage() {}

func (x *BuyingProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broker_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyingProductRequest.ProtoReflect.Descriptor instead.
func (*BuyingProductRequest) Descriptor() ([]byte, []int) {
	return file_broker_broker_proto_rawDescGZIP(), []int{0}
}

func (x *BuyingProductRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BuyingProductRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type BuyingProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlreadyBuy bool `protobuf:"varint,1,opt,name=already_buy,json=alreadyBuy,proto3" json:"already_buy,omitempty"`
}

func (x *BuyingProductResponse) Reset() {
	*x = BuyingProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyingProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyingProductResponse) ProtoMessage() {}

func (x *BuyingProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_broker_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyingProductResponse.ProtoReflect.Descriptor instead.
func (*BuyingProductResponse) Descriptor() ([]byte, []int) {
	return file_broker_broker_proto_rawDescGZIP(), []int{1}
}

func (x *BuyingProductResponse) GetAlreadyBuy() bool {
	if x != nil {
		return x.AlreadyBuy
	}
	return false
}

var File_broker_broker_proto protoreflect.FileDescriptor

var file_broker_broker_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x4b, 0x0a,
	0x14, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x15, 0x42, 0x75,
	0x79, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x62,
	0x75, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x42, 0x75, 0x79, 0x32, 0x56, 0x0a, 0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x4c,
	0x0a, 0x0d, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1c, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b,
	0x73, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x6b, 0x6f, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x3b, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_broker_broker_proto_rawDescOnce sync.Once
	file_broker_broker_proto_rawDescData = file_broker_broker_proto_rawDesc
)

func file_broker_broker_proto_rawDescGZIP() []byte {
	file_broker_broker_proto_rawDescOnce.Do(func() {
		file_broker_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_broker_proto_rawDescData)
	})
	return file_broker_broker_proto_rawDescData
}

var file_broker_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_broker_broker_proto_goTypes = []interface{}{
	(*BuyingProductRequest)(nil),  // 0: broker.BuyingProductRequest
	(*BuyingProductResponse)(nil), // 1: broker.BuyingProductResponse
}
var file_broker_broker_proto_depIdxs = []int32{
	0, // 0: broker.Broker.BuyingProduct:input_type -> broker.BuyingProductRequest
	1, // 1: broker.Broker.BuyingProduct:output_type -> broker.BuyingProductResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_broker_broker_proto_init() }
func file_broker_broker_proto_init() {
	if File_broker_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyingProductRequest); i {
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
		file_broker_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyingProductResponse); i {
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
			RawDescriptor: file_broker_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_broker_broker_proto_goTypes,
		DependencyIndexes: file_broker_broker_proto_depIdxs,
		MessageInfos:      file_broker_broker_proto_msgTypes,
	}.Build()
	File_broker_broker_proto = out.File
	file_broker_broker_proto_rawDesc = nil
	file_broker_broker_proto_goTypes = nil
	file_broker_broker_proto_depIdxs = nil
}
