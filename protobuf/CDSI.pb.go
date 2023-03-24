// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: CDSI.proto

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

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each ACI/UAK pair is a 32-byte buffer, containing the 16-byte ACI followed
	// by its 16-byte UAK.
	AciUakPairs []byte `protobuf:"bytes,1,opt,name=aci_uak_pairs,json=aciUakPairs,proto3" json:"aci_uak_pairs,omitempty"`
	// Each E164 is an 8-byte big-endian number, as 8 bytes.
	PrevE164S    []byte `protobuf:"bytes,2,opt,name=prev_e164s,json=prevE164s,proto3" json:"prev_e164s,omitempty"`
	NewE164S     []byte `protobuf:"bytes,3,opt,name=new_e164s,json=newE164s,proto3" json:"new_e164s,omitempty"`
	DiscardE164S []byte `protobuf:"bytes,4,opt,name=discard_e164s,json=discardE164s,proto3" json:"discard_e164s,omitempty"`
	// If set, a token which allows rate limiting to discount the e164s in
	// the request's prev_e164s, only counting new_e164s.  If not set, then
	// rate limiting considers both prev_e164s' and new_e164s' size.
	Token []byte `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	// After receiving a new token from the server, send back a message just
	// containing a token_ack.
	TokenAck bool `protobuf:"varint,7,opt,name=token_ack,json=tokenAck,proto3" json:"token_ack,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CDSI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CDSI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_CDSI_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequest) GetAciUakPairs() []byte {
	if x != nil {
		return x.AciUakPairs
	}
	return nil
}

func (x *ClientRequest) GetPrevE164S() []byte {
	if x != nil {
		return x.PrevE164S
	}
	return nil
}

func (x *ClientRequest) GetNewE164S() []byte {
	if x != nil {
		return x.NewE164S
	}
	return nil
}

func (x *ClientRequest) GetDiscardE164S() []byte {
	if x != nil {
		return x.DiscardE164S
	}
	return nil
}

func (x *ClientRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *ClientRequest) GetTokenAck() bool {
	if x != nil {
		return x.TokenAck
	}
	return false
}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each triple is an 8-byte e164, a 16-byte PNI, and a 16-byte ACI.
	// If the e164 was not found, PNI and ACI are all zeros.  If the PNI
	// was found but the ACI was not, the PNI will be non-zero and the ACI
	// will be all zeros.  ACI will be returned if one of the returned
	// PNIs has an ACI/UAK pair that matches.
	//
	// Should the request be successful (IE: a successful status returned),
	// |e164_pni_aci_triple| will always equal |e164| of the request,
	// so the entire marshalled size of the response will be (2+32)*|e164|,
	// where the additional 2 bytes are the id/type/length additions of the
	// protobuf marshaling added to each byte array.  This avoids any data
	// leakage based on the size of the encrypted output.
	E164PniAciTriples []byte `protobuf:"bytes,1,opt,name=e164_pni_aci_triples,json=e164PniAciTriples,proto3" json:"e164_pni_aci_triples,omitempty"`
	// If the user has run out of quota for lookups, they will receive
	// a response with just the following field set, followed by a websocket
	// closure of type 4008 (RESOURCE_EXHAUSTED).  Should they retry exactly
	// the same request after the provided number of seconds has passed,
	// we expect it should work.
	RetryAfterSecs int32 `protobuf:"varint,2,opt,name=retry_after_secs,json=retryAfterSecs,proto3" json:"retry_after_secs,omitempty"`
	// A token which allows subsequent calls' rate limiting to discount the
	// e164s sent up in this request, only counting those in the next
	// request's new_e164s.
	Token []byte `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// On a successful response to a token_ack request, the number of permits
	// that were deducted from the user's rate-limit in order to process the
	// request
	DebugPermitsUsed int32 `protobuf:"varint,4,opt,name=debug_permits_used,json=debugPermitsUsed,proto3" json:"debug_permits_used,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CDSI_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CDSI_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_CDSI_proto_rawDescGZIP(), []int{1}
}

func (x *ClientResponse) GetE164PniAciTriples() []byte {
	if x != nil {
		return x.E164PniAciTriples
	}
	return nil
}

func (x *ClientResponse) GetRetryAfterSecs() int32 {
	if x != nil {
		return x.RetryAfterSecs
	}
	return 0
}

func (x *ClientResponse) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *ClientResponse) GetDebugPermitsUsed() int32 {
	if x != nil {
		return x.DebugPermitsUsed
	}
	return 0
}

type EnclaveLoad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, before loading any tuples entirely clear the current map,
	// zero'ing out all current data.
	ClearAll bool `protobuf:"varint,1,opt,name=clear_all,json=clearAll,proto3" json:"clear_all,omitempty"`
	// Each tuple is an 8-byte e164, a 16-byte PNI, a 16-byte ACI, and a
	// 16-byte UAK.  These should be loaded as a 48-byte value (PNI,ACI,UAK)
	// associated with an 8-byte key (e164).
	// ACI/PNI/UAK may all be zeros, in which case this is a delete of the e164.
	E164AciPniUakTuples []byte `protobuf:"bytes,2,opt,name=e164_aci_pni_uak_tuples,json=e164AciPniUakTuples,proto3" json:"e164_aci_pni_uak_tuples,omitempty"`
	// If non-empty, overwrite the shared token secret with this value.
	SharedTokenSecret []byte `protobuf:"bytes,3,opt,name=shared_token_secret,json=sharedTokenSecret,proto3" json:"shared_token_secret,omitempty"`
}

func (x *EnclaveLoad) Reset() {
	*x = EnclaveLoad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CDSI_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnclaveLoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnclaveLoad) ProtoMessage() {}

func (x *EnclaveLoad) ProtoReflect() protoreflect.Message {
	mi := &file_CDSI_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnclaveLoad.ProtoReflect.Descriptor instead.
func (*EnclaveLoad) Descriptor() ([]byte, []int) {
	return file_CDSI_proto_rawDescGZIP(), []int{2}
}

func (x *EnclaveLoad) GetClearAll() bool {
	if x != nil {
		return x.ClearAll
	}
	return false
}

func (x *EnclaveLoad) GetE164AciPniUakTuples() []byte {
	if x != nil {
		return x.E164AciPniUakTuples
	}
	return nil
}

func (x *EnclaveLoad) GetSharedTokenSecret() []byte {
	if x != nil {
		return x.SharedTokenSecret
	}
	return nil
}

type ClientHandshakeStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key associated with this server's enclave. For use in test-only
	// contexts where attestation is not available
	TestOnlyPubkey []byte `protobuf:"bytes,1,opt,name=test_only_pubkey,json=testOnlyPubkey,proto3" json:"test_only_pubkey,omitempty"`
	// Remote-attestation evidence associated with the public key
	Evidence []byte `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
	// Endorsements of remote-attestation evidence.
	Endorsement []byte `protobuf:"bytes,3,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
}

func (x *ClientHandshakeStart) Reset() {
	*x = ClientHandshakeStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CDSI_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientHandshakeStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientHandshakeStart) ProtoMessage() {}

func (x *ClientHandshakeStart) ProtoReflect() protoreflect.Message {
	mi := &file_CDSI_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientHandshakeStart.ProtoReflect.Descriptor instead.
func (*ClientHandshakeStart) Descriptor() ([]byte, []int) {
	return file_CDSI_proto_rawDescGZIP(), []int{3}
}

func (x *ClientHandshakeStart) GetTestOnlyPubkey() []byte {
	if x != nil {
		return x.TestOnlyPubkey
	}
	return nil
}

func (x *ClientHandshakeStart) GetEvidence() []byte {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *ClientHandshakeStart) GetEndorsement() []byte {
	if x != nil {
		return x.Endorsement
	}
	return nil
}

var File_CDSI_proto protoreflect.FileDescriptor

var file_CDSI_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x43, 0x44, 0x53, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6f, 0x72,
	0x67, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x64, 0x73, 0x69, 0x22, 0xc7, 0x01,
	0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0d, 0x61, 0x63, 0x69, 0x5f, 0x75, 0x61, 0x6b, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x69, 0x55, 0x61, 0x6b, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x65, 0x31, 0x36, 0x34,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x72, 0x65, 0x76, 0x45, 0x31, 0x36,
	0x34, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x31, 0x36, 0x34, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x31, 0x36, 0x34, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x31, 0x36, 0x34, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x45,
	0x31, 0x36, 0x34, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x6b, 0x22, 0xaf, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x65, 0x31,
	0x36, 0x34, 0x5f, 0x70, 0x6e, 0x69, 0x5f, 0x61, 0x63, 0x69, 0x5f, 0x74, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x65, 0x31, 0x36, 0x34, 0x50, 0x6e,
	0x69, 0x41, 0x63, 0x69, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x72, 0x79, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x73, 0x5f, 0x75, 0x73, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x65, 0x62, 0x75, 0x67, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x74, 0x73, 0x55, 0x73, 0x65, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x45, 0x6e,
	0x63, 0x6c, 0x61, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6c,
	0x65, 0x61, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x34, 0x0a, 0x17, 0x65, 0x31, 0x36, 0x34, 0x5f, 0x61,
	0x63, 0x69, 0x5f, 0x70, 0x6e, 0x69, 0x5f, 0x75, 0x61, 0x6b, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x65, 0x31, 0x36, 0x34, 0x41, 0x63, 0x69,
	0x50, 0x6e, 0x69, 0x55, 0x61, 0x6b, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x7e, 0x0a, 0x14,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e,
	0x74, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x2a, 0x0a, 0x15,
	0x6f, 0x72, 0x67, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x64, 0x73, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x3b, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CDSI_proto_rawDescOnce sync.Once
	file_CDSI_proto_rawDescData = file_CDSI_proto_rawDesc
)

func file_CDSI_proto_rawDescGZIP() []byte {
	file_CDSI_proto_rawDescOnce.Do(func() {
		file_CDSI_proto_rawDescData = protoimpl.X.CompressGZIP(file_CDSI_proto_rawDescData)
	})
	return file_CDSI_proto_rawDescData
}

var file_CDSI_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_CDSI_proto_goTypes = []interface{}{
	(*ClientRequest)(nil),        // 0: org.signal.cdsi.ClientRequest
	(*ClientResponse)(nil),       // 1: org.signal.cdsi.ClientResponse
	(*EnclaveLoad)(nil),          // 2: org.signal.cdsi.EnclaveLoad
	(*ClientHandshakeStart)(nil), // 3: org.signal.cdsi.ClientHandshakeStart
}
var file_CDSI_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CDSI_proto_init() }
func file_CDSI_proto_init() {
	if File_CDSI_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CDSI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_CDSI_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_CDSI_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnclaveLoad); i {
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
		file_CDSI_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientHandshakeStart); i {
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
			RawDescriptor: file_CDSI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CDSI_proto_goTypes,
		DependencyIndexes: file_CDSI_proto_depIdxs,
		MessageInfos:      file_CDSI_proto_msgTypes,
	}.Build()
	File_CDSI_proto = out.File
	file_CDSI_proto_rawDesc = nil
	file_CDSI_proto_goTypes = nil
	file_CDSI_proto_depIdxs = nil
}