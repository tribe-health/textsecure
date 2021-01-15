//*
// Copyright (C) 2019 Open Whisper Systems
//
// Licensed according to the LICENSE file in this repository.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: InternalSerialization.proto

package signalservice

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignalServiceContentProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalAddress *AddressProto  `protobuf:"bytes,1,opt,name=localAddress" json:"localAddress,omitempty"`
	Metadata     *MetadataProto `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// Types that are assignable to Data:
	//	*SignalServiceContentProto_LegacyDataMessage
	//	*SignalServiceContentProto_Content
	Data isSignalServiceContentProto_Data `protobuf_oneof:"data"`
}

func (x *SignalServiceContentProto) Reset() {
	*x = SignalServiceContentProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InternalSerialization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalServiceContentProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalServiceContentProto) ProtoMessage() {}

func (x *SignalServiceContentProto) ProtoReflect() protoreflect.Message {
	mi := &file_InternalSerialization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalServiceContentProto.ProtoReflect.Descriptor instead.
func (*SignalServiceContentProto) Descriptor() ([]byte, []int) {
	return file_InternalSerialization_proto_rawDescGZIP(), []int{0}
}

func (x *SignalServiceContentProto) GetLocalAddress() *AddressProto {
	if x != nil {
		return x.LocalAddress
	}
	return nil
}

func (x *SignalServiceContentProto) GetMetadata() *MetadataProto {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (m *SignalServiceContentProto) GetData() isSignalServiceContentProto_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SignalServiceContentProto) GetLegacyDataMessage() *DataMessage {
	if x, ok := x.GetData().(*SignalServiceContentProto_LegacyDataMessage); ok {
		return x.LegacyDataMessage
	}
	return nil
}

func (x *SignalServiceContentProto) GetContent() *Content {
	if x, ok := x.GetData().(*SignalServiceContentProto_Content); ok {
		return x.Content
	}
	return nil
}

type isSignalServiceContentProto_Data interface {
	isSignalServiceContentProto_Data()
}

type SignalServiceContentProto_LegacyDataMessage struct {
	LegacyDataMessage *DataMessage `protobuf:"bytes,3,opt,name=legacyDataMessage,oneof"`
}

type SignalServiceContentProto_Content struct {
	Content *Content `protobuf:"bytes,4,opt,name=content,oneof"`
}

func (*SignalServiceContentProto_LegacyDataMessage) isSignalServiceContentProto_Data() {}

func (*SignalServiceContentProto_Content) isSignalServiceContentProto_Data() {}

type MetadataProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                  *AddressProto `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	SenderDevice             *int32        `protobuf:"varint,2,opt,name=senderDevice" json:"senderDevice,omitempty"`
	Timestamp                *int64        `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	ServerReceivedTimestamp  *int64        `protobuf:"varint,5,opt,name=serverReceivedTimestamp" json:"serverReceivedTimestamp,omitempty"`
	ServerDeliveredTimestamp *int64        `protobuf:"varint,6,opt,name=serverDeliveredTimestamp" json:"serverDeliveredTimestamp,omitempty"`
	NeedsReceipt             *bool         `protobuf:"varint,4,opt,name=needsReceipt" json:"needsReceipt,omitempty"`
}

func (x *MetadataProto) Reset() {
	*x = MetadataProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InternalSerialization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataProto) ProtoMessage() {}

func (x *MetadataProto) ProtoReflect() protoreflect.Message {
	mi := &file_InternalSerialization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataProto.ProtoReflect.Descriptor instead.
func (*MetadataProto) Descriptor() ([]byte, []int) {
	return file_InternalSerialization_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataProto) GetAddress() *AddressProto {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *MetadataProto) GetSenderDevice() int32 {
	if x != nil && x.SenderDevice != nil {
		return *x.SenderDevice
	}
	return 0
}

func (x *MetadataProto) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *MetadataProto) GetServerReceivedTimestamp() int64 {
	if x != nil && x.ServerReceivedTimestamp != nil {
		return *x.ServerReceivedTimestamp
	}
	return 0
}

func (x *MetadataProto) GetServerDeliveredTimestamp() int64 {
	if x != nil && x.ServerDeliveredTimestamp != nil {
		return *x.ServerDeliveredTimestamp
	}
	return 0
}

func (x *MetadataProto) GetNeedsReceipt() bool {
	if x != nil && x.NeedsReceipt != nil {
		return *x.NeedsReceipt
	}
	return false
}

type AddressProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	E164  *string `protobuf:"bytes,2,opt,name=e164" json:"e164,omitempty"`
	Relay *string `protobuf:"bytes,3,opt,name=relay" json:"relay,omitempty"`
}

func (x *AddressProto) Reset() {
	*x = AddressProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InternalSerialization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressProto) ProtoMessage() {}

func (x *AddressProto) ProtoReflect() protoreflect.Message {
	mi := &file_InternalSerialization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressProto.ProtoReflect.Descriptor instead.
func (*AddressProto) Descriptor() ([]byte, []int) {
	return file_InternalSerialization_proto_rawDescGZIP(), []int{2}
}

func (x *AddressProto) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *AddressProto) GetE164() string {
	if x != nil && x.E164 != nil {
		return *x.E164
	}
	return ""
}

func (x *AddressProto) GetRelay() string {
	if x != nil && x.Relay != nil {
		return *x.Relay
	}
	return ""
}

var File_InternalSerialization_proto protoreflect.FileDescriptor

var file_InternalSerialization_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x1a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98,
	0x02, 0x0a, 0x19, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x0a, 0x0c,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0c, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x4a, 0x0a, 0x11, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x11, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x02, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3a, 0x0a, 0x18, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x73,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e,
	0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x4c, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x31, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x31, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x4f, 0x0a, 0x3a, 0x6f, 0x72, 0x67,
	0x2e, 0x77, 0x68, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
}

var (
	file_InternalSerialization_proto_rawDescOnce sync.Once
	file_InternalSerialization_proto_rawDescData = file_InternalSerialization_proto_rawDesc
)

func file_InternalSerialization_proto_rawDescGZIP() []byte {
	file_InternalSerialization_proto_rawDescOnce.Do(func() {
		file_InternalSerialization_proto_rawDescData = protoimpl.X.CompressGZIP(file_InternalSerialization_proto_rawDescData)
	})
	return file_InternalSerialization_proto_rawDescData
}

var file_InternalSerialization_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_InternalSerialization_proto_goTypes = []interface{}{
	(*SignalServiceContentProto)(nil), // 0: textsecure.SignalServiceContentProto
	(*MetadataProto)(nil),             // 1: textsecure.MetadataProto
	(*AddressProto)(nil),              // 2: textsecure.AddressProto
	(*DataMessage)(nil),               // 3: signalservice.DataMessage
	(*Content)(nil),                   // 4: signalservice.Content
}
var file_InternalSerialization_proto_depIdxs = []int32{
	2, // 0: textsecure.SignalServiceContentProto.localAddress:type_name -> textsecure.AddressProto
	1, // 1: textsecure.SignalServiceContentProto.metadata:type_name -> textsecure.MetadataProto
	3, // 2: textsecure.SignalServiceContentProto.legacyDataMessage:type_name -> signalservice.DataMessage
	4, // 3: textsecure.SignalServiceContentProto.content:type_name -> signalservice.Content
	2, // 4: textsecure.MetadataProto.address:type_name -> textsecure.AddressProto
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_InternalSerialization_proto_init() }
func file_InternalSerialization_proto_init() {
	if File_InternalSerialization_proto != nil {
		return
	}
	file_SignalService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InternalSerialization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalServiceContentProto); i {
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
		file_InternalSerialization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataProto); i {
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
		file_InternalSerialization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressProto); i {
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
	file_InternalSerialization_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SignalServiceContentProto_LegacyDataMessage)(nil),
		(*SignalServiceContentProto_Content)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_InternalSerialization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InternalSerialization_proto_goTypes,
		DependencyIndexes: file_InternalSerialization_proto_depIdxs,
		MessageInfos:      file_InternalSerialization_proto_msgTypes,
	}.Build()
	File_InternalSerialization_proto = out.File
	file_InternalSerialization_proto_rawDesc = nil
	file_InternalSerialization_proto_goTypes = nil
	file_InternalSerialization_proto_depIdxs = nil
}