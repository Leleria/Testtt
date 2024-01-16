// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: ServiceLoyalty.proto

package Grpc_Service_App_sv1

import (
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

type AddNewPromoCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypeDiscount     int32                  `protobuf:"varint,2,opt,name=typeDiscount,proto3" json:"typeDiscount,omitempty"`
	ValueDiscount    int32                  `protobuf:"varint,3,opt,name=valueDiscount,proto3" json:"valueDiscount,omitempty"`
	DateStartActive  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=dateStartActive,proto3" json:"dateStartActive,omitempty"`
	DateFinishActive *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=dateFinishActive,proto3" json:"dateFinishActive,omitempty"`
	CountUses        int32                  `protobuf:"varint,6,opt,name=countUses,proto3" json:"countUses,omitempty"`
}

func (x *AddNewPromoCodeRequest) Reset() {
	*x = AddNewPromoCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServiceLoyalty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewPromoCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewPromoCodeRequest) ProtoMessage() {}

func (x *AddNewPromoCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ServiceLoyalty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewPromoCodeRequest.ProtoReflect.Descriptor instead.
func (*AddNewPromoCodeRequest) Descriptor() ([]byte, []int) {
	return file_ServiceLoyalty_proto_rawDescGZIP(), []int{0}
}

func (x *AddNewPromoCodeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddNewPromoCodeRequest) GetTypeDiscount() int32 {
	if x != nil {
		return x.TypeDiscount
	}
	return 0
}

func (x *AddNewPromoCodeRequest) GetValueDiscount() int32 {
	if x != nil {
		return x.ValueDiscount
	}
	return 0
}

func (x *AddNewPromoCodeRequest) GetDateStartActive() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStartActive
	}
	return nil
}

func (x *AddNewPromoCodeRequest) GetDateFinishActive() *timestamppb.Timestamp {
	if x != nil {
		return x.DateFinishActive
	}
	return nil
}

func (x *AddNewPromoCodeRequest) GetCountUses() int32 {
	if x != nil {
		return x.CountUses
	}
	return 0
}

type AddNewPromoCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddNewPromoCodeResponse) Reset() {
	*x = AddNewPromoCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServiceLoyalty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewPromoCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewPromoCodeResponse) ProtoMessage() {}

func (x *AddNewPromoCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ServiceLoyalty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewPromoCodeResponse.ProtoReflect.Descriptor instead.
func (*AddNewPromoCodeResponse) Descriptor() ([]byte, []int) {
	return file_ServiceLoyalty_proto_rawDescGZIP(), []int{1}
}

func (x *AddNewPromoCodeResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_ServiceLoyalty_proto protoreflect.FileDescriptor

var file_ServiceLoyalty_proto_rawDesc = []byte{
	0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x53, 0x65, 0x72, 0x76, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02,
	0x0a, 0x16, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x46, 0x0a, 0x10,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x73, 0x22, 0x31, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x60, 0x0a, 0x0e, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x2e,
	0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x3a, 0x73, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServiceLoyalty_proto_rawDescOnce sync.Once
	file_ServiceLoyalty_proto_rawDescData = file_ServiceLoyalty_proto_rawDesc
)

func file_ServiceLoyalty_proto_rawDescGZIP() []byte {
	file_ServiceLoyalty_proto_rawDescOnce.Do(func() {
		file_ServiceLoyalty_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServiceLoyalty_proto_rawDescData)
	})
	return file_ServiceLoyalty_proto_rawDescData
}

var file_ServiceLoyalty_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ServiceLoyalty_proto_goTypes = []interface{}{
	(*AddNewPromoCodeRequest)(nil),  // 0: Serv.AddNewPromoCodeRequest
	(*AddNewPromoCodeResponse)(nil), // 1: Serv.AddNewPromoCodeResponse
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
}
var file_ServiceLoyalty_proto_depIdxs = []int32{
	2, // 0: Serv.AddNewPromoCodeRequest.dateStartActive:type_name -> google.protobuf.Timestamp
	2, // 1: Serv.AddNewPromoCodeRequest.dateFinishActive:type_name -> google.protobuf.Timestamp
	0, // 2: Serv.LoyaltyService.AddNewPromoCode:input_type -> Serv.AddNewPromoCodeRequest
	1, // 3: Serv.LoyaltyService.AddNewPromoCode:output_type -> Serv.AddNewPromoCodeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ServiceLoyalty_proto_init() }
func file_ServiceLoyalty_proto_init() {
	if File_ServiceLoyalty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ServiceLoyalty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewPromoCodeRequest); i {
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
		file_ServiceLoyalty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewPromoCodeResponse); i {
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
			RawDescriptor: file_ServiceLoyalty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ServiceLoyalty_proto_goTypes,
		DependencyIndexes: file_ServiceLoyalty_proto_depIdxs,
		MessageInfos:      file_ServiceLoyalty_proto_msgTypes,
	}.Build()
	File_ServiceLoyalty_proto = out.File
	file_ServiceLoyalty_proto_rawDesc = nil
	file_ServiceLoyalty_proto_goTypes = nil
	file_ServiceLoyalty_proto_depIdxs = nil
}
