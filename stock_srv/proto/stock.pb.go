// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/stock.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GoodsInvInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int32 `protobuf:"varint,1,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	Num     int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GoodsInvInfo) Reset() {
	*x = GoodsInvInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInvInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInvInfo) ProtoMessage() {}

func (x *GoodsInvInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInvInfo.ProtoReflect.Descriptor instead.
func (*GoodsInvInfo) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInvInfo) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GoodsInvInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type SellInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsInfo []*GoodsInvInfo `protobuf:"bytes,1,rep,name=goodsInfo,proto3" json:"goodsInfo,omitempty"`
	OrderSn   string          `protobuf:"bytes,2,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
}

func (x *SellInfo) Reset() {
	*x = SellInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellInfo) ProtoMessage() {}

func (x *SellInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellInfo.ProtoReflect.Descriptor instead.
func (*SellInfo) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{1}
}

func (x *SellInfo) GetGoodsInfo() []*GoodsInvInfo {
	if x != nil {
		return x.GoodsInfo
	}
	return nil
}

func (x *SellInfo) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

//
//+-------------------------------------------------------+
//|   Demo                                                |
//+-------------------------------------------------------+
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_proto_stock_proto_rawDescGZIP(), []int{3}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_stock_proto protoreflect.FileDescriptor

var file_proto_stock_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3a, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x51, 0x0a, 0x08,
	0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x22,
	0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe5, 0x01, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x12, 0x0d, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0d, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x04, 0x53,
	0x65, 0x6c, 0x6c, 0x12, 0x09, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x09, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stock_proto_rawDescOnce sync.Once
	file_proto_stock_proto_rawDescData = file_proto_stock_proto_rawDesc
)

func file_proto_stock_proto_rawDescGZIP() []byte {
	file_proto_stock_proto_rawDescOnce.Do(func() {
		file_proto_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stock_proto_rawDescData)
	})
	return file_proto_stock_proto_rawDescData
}

var file_proto_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_stock_proto_goTypes = []interface{}{
	(*GoodsInvInfo)(nil),  // 0: GoodsInvInfo
	(*SellInfo)(nil),      // 1: SellInfo
	(*HelloRequest)(nil),  // 2: HelloRequest
	(*HelloReply)(nil),    // 3: HelloReply
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_proto_stock_proto_depIdxs = []int32{
	0, // 0: SellInfo.goodsInfo:type_name -> GoodsInvInfo
	2, // 1: Stock.SayHello:input_type -> HelloRequest
	0, // 2: Stock.SetInv:input_type -> GoodsInvInfo
	0, // 3: Stock.InvDetail:input_type -> GoodsInvInfo
	1, // 4: Stock.Sell:input_type -> SellInfo
	1, // 5: Stock.Reback:input_type -> SellInfo
	3, // 6: Stock.SayHello:output_type -> HelloReply
	4, // 7: Stock.SetInv:output_type -> google.protobuf.Empty
	0, // 8: Stock.InvDetail:output_type -> GoodsInvInfo
	4, // 9: Stock.Sell:output_type -> google.protobuf.Empty
	4, // 10: Stock.Reback:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_stock_proto_init() }
func file_proto_stock_proto_init() {
	if File_proto_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInvInfo); i {
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
		file_proto_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellInfo); i {
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
		file_proto_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_proto_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stock_proto_goTypes,
		DependencyIndexes: file_proto_stock_proto_depIdxs,
		MessageInfos:      file_proto_stock_proto_msgTypes,
	}.Build()
	File_proto_stock_proto = out.File
	file_proto_stock_proto_rawDesc = nil
	file_proto_stock_proto_goTypes = nil
	file_proto_stock_proto_depIdxs = nil
}
