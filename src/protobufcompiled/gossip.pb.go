// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: gossip.proto

package protobufcompiled

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

type Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignerPublicAddress string       `protobuf:"bytes,1,opt,name=signer_public_address,json=signerPublicAddress,proto3" json:"signer_public_address,omitempty"`
	CreaterdAt          uint64       `protobuf:"varint,2,opt,name=createrd_at,json=createrdAt,proto3" json:"createrd_at,omitempty"`
	Signature           []byte       `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Transaction         *Transaction `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Hash                []byte       `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	LeftParentHash      []byte       `protobuf:"bytes,6,opt,name=left_parent_hash,json=leftParentHash,proto3" json:"left_parent_hash,omitempty"`
	RightParentHash     []byte       `protobuf:"bytes,7,opt,name=right_parent_hash,json=rightParentHash,proto3" json:"right_parent_hash,omitempty"`
	Weight              uint64       `protobuf:"varint,8,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Vertex) Reset() {
	*x = Vertex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertex) ProtoMessage() {}

func (x *Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertex.ProtoReflect.Descriptor instead.
func (*Vertex) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{0}
}

func (x *Vertex) GetSignerPublicAddress() string {
	if x != nil {
		return x.SignerPublicAddress
	}
	return ""
}

func (x *Vertex) GetCreaterdAt() uint64 {
	if x != nil {
		return x.CreaterdAt
	}
	return 0
}

func (x *Vertex) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Vertex) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Vertex) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Vertex) GetLeftParentHash() []byte {
	if x != nil {
		return x.LeftParentHash
	}
	return nil
}

func (x *Vertex) GetRightParentHash() []byte {
	if x != nil {
		return x.RightParentHash
	}
	return nil
}

func (x *Vertex) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type Gossiper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Digest    []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Gossiper) Reset() {
	*x = Gossiper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gossiper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gossiper) ProtoMessage() {}

func (x *Gossiper) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gossiper.ProtoReflect.Descriptor instead.
func (*Gossiper) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *Gossiper) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Gossiper) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *Gossiper) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type VrxMsgGossip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vertex    *Vertex     `protobuf:"bytes,1,opt,name=vertex,proto3" json:"vertex,omitempty"`
	Gossipers []*Gossiper `protobuf:"bytes,2,rep,name=gossipers,proto3" json:"gossipers,omitempty"`
}

func (x *VrxMsgGossip) Reset() {
	*x = VrxMsgGossip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VrxMsgGossip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VrxMsgGossip) ProtoMessage() {}

func (x *VrxMsgGossip) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VrxMsgGossip.ProtoReflect.Descriptor instead.
func (*VrxMsgGossip) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{2}
}

func (x *VrxMsgGossip) GetVertex() *Vertex {
	if x != nil {
		return x.Vertex
	}
	return nil
}

func (x *VrxMsgGossip) GetGossipers() []*Gossiper {
	if x != nil {
		return x.Gossipers
	}
	return nil
}

type TrxMsgGossip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trx       *Transaction `protobuf:"bytes,1,opt,name=trx,proto3" json:"trx,omitempty"`
	Gossipers []*Gossiper  `protobuf:"bytes,2,rep,name=gossipers,proto3" json:"gossipers,omitempty"`
}

func (x *TrxMsgGossip) Reset() {
	*x = TrxMsgGossip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrxMsgGossip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrxMsgGossip) ProtoMessage() {}

func (x *TrxMsgGossip) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrxMsgGossip.ProtoReflect.Descriptor instead.
func (*TrxMsgGossip) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{3}
}

func (x *TrxMsgGossip) GetTrx() *Transaction {
	if x != nil {
		return x.Trx
	}
	return nil
}

func (x *TrxMsgGossip) GetGossipers() []*Gossiper {
	if x != nil {
		return x.Gossipers
	}
	return nil
}

type ConnectionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicAddress string `protobuf:"bytes,1,opt,name=public_address,json=publicAddress,proto3" json:"public_address,omitempty"`
	Url           string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt     uint64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Digest        []byte `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	Signature     []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ConnectionData) Reset() {
	*x = ConnectionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionData) ProtoMessage() {}

func (x *ConnectionData) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionData.ProtoReflect.Descriptor instead.
func (*ConnectionData) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectionData) GetPublicAddress() string {
	if x != nil {
		return x.PublicAddress
	}
	return ""
}

func (x *ConnectionData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ConnectionData) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ConnectionData) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *ConnectionData) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ConnectedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignerPublicAddress string            `protobuf:"bytes,1,opt,name=signer_public_address,json=signerPublicAddress,proto3" json:"signer_public_address,omitempty"`
	Connections         []*ConnectionData `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *ConnectedNodes) Reset() {
	*x = ConnectedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedNodes) ProtoMessage() {}

func (x *ConnectedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedNodes.ProtoReflect.Descriptor instead.
func (*ConnectedNodes) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectedNodes) GetSignerPublicAddress() string {
	if x != nil {
		return x.SignerPublicAddress
	}
	return ""
}

func (x *ConnectedNodes) GetConnections() []*ConnectionData {
	if x != nil {
		return x.Connections
	}
	return nil
}

var File_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x1a, 0x16, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb9, 0x02, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x15, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6c, 0x65, 0x66, 0x74, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5a, 0x0a, 0x08,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x70, 0x0a, 0x0c, 0x56, 0x72, 0x78, 0x4d,
	0x73, 0x67, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x06, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x09, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x65, 0x72, 0x52,
	0x09, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x65, 0x72, 0x73, 0x22, 0x6f, 0x0a, 0x0c, 0x54, 0x72,
	0x78, 0x4d, 0x73, 0x67, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x2a, 0x0a, 0x03, 0x74, 0x72,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x03, 0x74, 0x72, 0x78, 0x12, 0x33, 0x0a, 0x09, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x65, 0x72,
	0x52, 0x09, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x65, 0x72, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x83, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x32, 0x0a, 0x15, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0x91, 0x03, 0x0a, 0x09, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x50, 0x49,
	0x12, 0x39, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e,
	0x41, 0x6c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x4c,
	0x6f, 0x61, 0x64, 0x44, 0x61, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x56, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x56, 0x72, 0x78, 0x12,
	0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x56, 0x72,
	0x78, 0x4d, 0x73, 0x67, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x54, 0x72,
	0x78, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e,
	0x54, 0x72, 0x78, 0x4d, 0x73, 0x67, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x6f, 0x73, 0x73, 0x68, 0x2f, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gossip_proto_rawDescOnce sync.Once
	file_gossip_proto_rawDescData = file_gossip_proto_rawDesc
)

func file_gossip_proto_rawDescGZIP() []byte {
	file_gossip_proto_rawDescOnce.Do(func() {
		file_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_gossip_proto_rawDescData)
	})
	return file_gossip_proto_rawDescData
}

var file_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gossip_proto_goTypes = []interface{}{
	(*Vertex)(nil),         // 0: computantis.Vertex
	(*Gossiper)(nil),       // 1: computantis.Gossiper
	(*VrxMsgGossip)(nil),   // 2: computantis.VrxMsgGossip
	(*TrxMsgGossip)(nil),   // 3: computantis.TrxMsgGossip
	(*ConnectionData)(nil), // 4: computantis.ConnectionData
	(*ConnectedNodes)(nil), // 5: computantis.ConnectedNodes
	(*Transaction)(nil),    // 6: computantis.Transaction
	(*emptypb.Empty)(nil),  // 7: google.protobuf.Empty
	(*AliveData)(nil),      // 8: computantis.AliveData
}
var file_gossip_proto_depIdxs = []int32{
	6,  // 0: computantis.Vertex.transaction:type_name -> computantis.Transaction
	0,  // 1: computantis.VrxMsgGossip.vertex:type_name -> computantis.Vertex
	1,  // 2: computantis.VrxMsgGossip.gossipers:type_name -> computantis.Gossiper
	6,  // 3: computantis.TrxMsgGossip.trx:type_name -> computantis.Transaction
	1,  // 4: computantis.TrxMsgGossip.gossipers:type_name -> computantis.Gossiper
	4,  // 5: computantis.ConnectedNodes.connections:type_name -> computantis.ConnectionData
	7,  // 6: computantis.GossipAPI.Alive:input_type -> google.protobuf.Empty
	7,  // 7: computantis.GossipAPI.LoadDag:input_type -> google.protobuf.Empty
	4,  // 8: computantis.GossipAPI.Announce:input_type -> computantis.ConnectionData
	4,  // 9: computantis.GossipAPI.Discover:input_type -> computantis.ConnectionData
	2,  // 10: computantis.GossipAPI.GossipVrx:input_type -> computantis.VrxMsgGossip
	3,  // 11: computantis.GossipAPI.GossipTrx:input_type -> computantis.TrxMsgGossip
	8,  // 12: computantis.GossipAPI.Alive:output_type -> computantis.AliveData
	0,  // 13: computantis.GossipAPI.LoadDag:output_type -> computantis.Vertex
	7,  // 14: computantis.GossipAPI.Announce:output_type -> google.protobuf.Empty
	5,  // 15: computantis.GossipAPI.Discover:output_type -> computantis.ConnectedNodes
	7,  // 16: computantis.GossipAPI.GossipVrx:output_type -> google.protobuf.Empty
	7,  // 17: computantis.GossipAPI.GossipTrx:output_type -> google.protobuf.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_gossip_proto_init() }
func file_gossip_proto_init() {
	if File_gossip_proto != nil {
		return
	}
	file_computantistypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gossip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertex); i {
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
		file_gossip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gossiper); i {
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
		file_gossip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VrxMsgGossip); i {
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
		file_gossip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrxMsgGossip); i {
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
		file_gossip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionData); i {
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
		file_gossip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectedNodes); i {
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
			RawDescriptor: file_gossip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gossip_proto_goTypes,
		DependencyIndexes: file_gossip_proto_depIdxs,
		MessageInfos:      file_gossip_proto_msgTypes,
	}.Build()
	File_gossip_proto = out.File
	file_gossip_proto_rawDesc = nil
	file_gossip_proto_goTypes = nil
	file_gossip_proto_depIdxs = nil
}
