// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: savourrpc/keylocker.proto

package keylocker

import (
	common "./proto/common"
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

type SocialKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SocialKey) Reset() {
	*x = SocialKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialKey) ProtoMessage() {}

func (x *SocialKey) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialKey.ProtoReflect.Descriptor instead.
func (*SocialKey) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{0}
}

func (x *SocialKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SocialKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SupportChainReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Network       string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *SupportChainReq) Reset() {
	*x = SupportChainReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportChainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChainReq) ProtoMessage() {}

func (x *SupportChainReq) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChainReq.ProtoReflect.Descriptor instead.
func (*SupportChainReq) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{1}
}

func (x *SupportChainReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SupportChainReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SupportChainReq) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type SupportChainRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    common.ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.ReturnCode" json:"code,omitempty"`
	Msg     string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Support bool              `protobuf:"varint,3,opt,name=support,proto3" json:"support,omitempty"`
}

func (x *SupportChainRep) Reset() {
	*x = SupportChainRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportChainRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChainRep) ProtoMessage() {}

func (x *SupportChainRep) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChainRep.ProtoReflect.Descriptor instead.
func (*SupportChainRep) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{2}
}

func (x *SupportChainRep) GetCode() common.ReturnCode {
	if x != nil {
		return x.Code
	}
	return common.ReturnCode_SUCCESS
}

func (x *SupportChainRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SupportChainRep) GetSupport() bool {
	if x != nil {
		return x.Support
	}
	return false
}

type SetSocialKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Uuid          string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Key           string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Password      string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SetSocialKeyReq) Reset() {
	*x = SetSocialKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSocialKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSocialKeyReq) ProtoMessage() {}

func (x *SetSocialKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSocialKeyReq.ProtoReflect.Descriptor instead.
func (*SetSocialKeyReq) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{3}
}

func (x *SetSocialKeyReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SetSocialKeyReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SetSocialKeyReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SetSocialKeyReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetSocialKeyReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SetSocialKeyRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      common.ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.ReturnCode" json:"code,omitempty"`
	Msg       string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Pub       string            `protobuf:"bytes,3,opt,name=pub,proto3" json:"pub,omitempty"`
	Priv      string            `protobuf:"bytes,4,opt,name=priv,proto3" json:"priv,omitempty"`
	CryptoWay string            `protobuf:"bytes,5,opt,name=crypto_way,json=cryptoWay,proto3" json:"crypto_way,omitempty"`
	FileCid   string            `protobuf:"bytes,6,opt,name=file_cid,json=fileCid,proto3" json:"file_cid,omitempty"`
	Contract  string            `protobuf:"bytes,7,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *SetSocialKeyRep) Reset() {
	*x = SetSocialKeyRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSocialKeyRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSocialKeyRep) ProtoMessage() {}

func (x *SetSocialKeyRep) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSocialKeyRep.ProtoReflect.Descriptor instead.
func (*SetSocialKeyRep) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{4}
}

func (x *SetSocialKeyRep) GetCode() common.ReturnCode {
	if x != nil {
		return x.Code
	}
	return common.ReturnCode_SUCCESS
}

func (x *SetSocialKeyRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SetSocialKeyRep) GetPub() string {
	if x != nil {
		return x.Pub
	}
	return ""
}

func (x *SetSocialKeyRep) GetPriv() string {
	if x != nil {
		return x.Priv
	}
	return ""
}

func (x *SetSocialKeyRep) GetCryptoWay() string {
	if x != nil {
		return x.CryptoWay
	}
	return ""
}

func (x *SetSocialKeyRep) GetFileCid() string {
	if x != nil {
		return x.FileCid
	}
	return ""
}

func (x *SetSocialKeyRep) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type GetSocialKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Uuid          string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FileCid       string `protobuf:"bytes,4,opt,name=file_cid,json=fileCid,proto3" json:"file_cid,omitempty"`
	Contract      string `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *GetSocialKeyReq) Reset() {
	*x = GetSocialKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSocialKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSocialKeyReq) ProtoMessage() {}

func (x *GetSocialKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSocialKeyReq.ProtoReflect.Descriptor instead.
func (*GetSocialKeyReq) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{5}
}

func (x *GetSocialKeyReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *GetSocialKeyReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetSocialKeyReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetSocialKeyReq) GetFileCid() string {
	if x != nil {
		return x.FileCid
	}
	return ""
}

func (x *GetSocialKeyReq) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type GetSocialKeyRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    common.ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.ReturnCode" json:"code,omitempty"`
	Msg     string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	KeyList []*SocialKey      `protobuf:"bytes,3,rep,name=key_list,json=keyList,proto3" json:"key_list,omitempty"`
}

func (x *GetSocialKeyRep) Reset() {
	*x = GetSocialKeyRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_keylocker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSocialKeyRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSocialKeyRep) ProtoMessage() {}

func (x *GetSocialKeyRep) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_keylocker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSocialKeyRep.ProtoReflect.Descriptor instead.
func (*GetSocialKeyRep) Descriptor() ([]byte, []int) {
	return file_savourrpc_keylocker_proto_rawDescGZIP(), []int{6}
}

func (x *GetSocialKeyRep) GetCode() common.ReturnCode {
	if x != nil {
		return x.Code
	}
	return common.ReturnCode_SUCCESS
}

func (x *GetSocialKeyRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetSocialKeyRep) GetKeyList() []*SocialKey {
	if x != nil {
		return x.KeyList
	}
	return nil
}

var File_savourrpc_keylocker_proto protoreflect.FileDescriptor

var file_savourrpc_keylocker_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x6b, 0x65, 0x79, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x1a, 0x16, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x09, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x68, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x22, 0x68, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xca,
	0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x75,
	0x62, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x69, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x72, 0x69, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x77, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x57, 0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x39, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0xaf, 0x02, 0x0a, 0x10, 0x4c, 0x65, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x2e, 0x73, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65,
	0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0c, 0x73, 0x65, 0x74,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x53, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x2b, 0x0a, 0x16, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5a,
	0x11, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_savourrpc_keylocker_proto_rawDescOnce sync.Once
	file_savourrpc_keylocker_proto_rawDescData = file_savourrpc_keylocker_proto_rawDesc
)

func file_savourrpc_keylocker_proto_rawDescGZIP() []byte {
	file_savourrpc_keylocker_proto_rawDescOnce.Do(func() {
		file_savourrpc_keylocker_proto_rawDescData = protoimpl.X.CompressGZIP(file_savourrpc_keylocker_proto_rawDescData)
	})
	return file_savourrpc_keylocker_proto_rawDescData
}

var file_savourrpc_keylocker_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_savourrpc_keylocker_proto_goTypes = []interface{}{
	(*SocialKey)(nil),       // 0: savourrpc.keylocker.SocialKey
	(*SupportChainReq)(nil), // 1: savourrpc.keylocker.SupportChainReq
	(*SupportChainRep)(nil), // 2: savourrpc.keylocker.SupportChainRep
	(*SetSocialKeyReq)(nil), // 3: savourrpc.keylocker.SetSocialKeyReq
	(*SetSocialKeyRep)(nil), // 4: savourrpc.keylocker.SetSocialKeyRep
	(*GetSocialKeyReq)(nil), // 5: savourrpc.keylocker.GetSocialKeyReq
	(*GetSocialKeyRep)(nil), // 6: savourrpc.keylocker.GetSocialKeyRep
	(common.ReturnCode)(0),  // 7: savourrpc.ReturnCode
}
var file_savourrpc_keylocker_proto_depIdxs = []int32{
	7, // 0: savourrpc.keylocker.SupportChainRep.code:type_name -> savourrpc.ReturnCode
	7, // 1: savourrpc.keylocker.SetSocialKeyRep.code:type_name -> savourrpc.ReturnCode
	7, // 2: savourrpc.keylocker.GetSocialKeyRep.code:type_name -> savourrpc.ReturnCode
	0, // 3: savourrpc.keylocker.GetSocialKeyRep.key_list:type_name -> savourrpc.keylocker.SocialKey
	1, // 4: savourrpc.keylocker.LeyLockerService.getSupportChain:input_type -> savourrpc.keylocker.SupportChainReq
	3, // 5: savourrpc.keylocker.LeyLockerService.setSocialKey:input_type -> savourrpc.keylocker.SetSocialKeyReq
	5, // 6: savourrpc.keylocker.LeyLockerService.getSocialKey:input_type -> savourrpc.keylocker.GetSocialKeyReq
	2, // 7: savourrpc.keylocker.LeyLockerService.getSupportChain:output_type -> savourrpc.keylocker.SupportChainRep
	4, // 8: savourrpc.keylocker.LeyLockerService.setSocialKey:output_type -> savourrpc.keylocker.SetSocialKeyRep
	6, // 9: savourrpc.keylocker.LeyLockerService.getSocialKey:output_type -> savourrpc.keylocker.GetSocialKeyRep
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_savourrpc_keylocker_proto_init() }
func file_savourrpc_keylocker_proto_init() {
	if File_savourrpc_keylocker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_savourrpc_keylocker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialKey); i {
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
		file_savourrpc_keylocker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportChainReq); i {
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
		file_savourrpc_keylocker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportChainRep); i {
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
		file_savourrpc_keylocker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSocialKeyReq); i {
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
		file_savourrpc_keylocker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSocialKeyRep); i {
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
		file_savourrpc_keylocker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSocialKeyReq); i {
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
		file_savourrpc_keylocker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSocialKeyRep); i {
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
			RawDescriptor: file_savourrpc_keylocker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_savourrpc_keylocker_proto_goTypes,
		DependencyIndexes: file_savourrpc_keylocker_proto_depIdxs,
		MessageInfos:      file_savourrpc_keylocker_proto_msgTypes,
	}.Build()
	File_savourrpc_keylocker_proto = out.File
	file_savourrpc_keylocker_proto_rawDesc = nil
	file_savourrpc_keylocker_proto_goTypes = nil
	file_savourrpc_keylocker_proto_depIdxs = nil
}