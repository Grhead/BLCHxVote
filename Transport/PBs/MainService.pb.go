// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: MainService.proto

package __

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CandidateListWithBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidatePK   string `protobuf:"bytes,1,opt,name=candidatePK,proto3" json:"candidatePK,omitempty"`
	CandidateName string `protobuf:"bytes,2,opt,name=candidateName,proto3" json:"candidateName,omitempty"`
	Balance       string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *CandidateListWithBalance) Reset() {
	*x = CandidateListWithBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateListWithBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateListWithBalance) ProtoMessage() {}

func (x *CandidateListWithBalance) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateListWithBalance.ProtoReflect.Descriptor instead.
func (*CandidateListWithBalance) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{0}
}

func (x *CandidateListWithBalance) GetCandidatePK() string {
	if x != nil {
		return x.CandidatePK
	}
	return ""
}

func (x *CandidateListWithBalance) GetCandidateName() string {
	if x != nil {
		return x.CandidateName
	}
	return ""
}

func (x *CandidateListWithBalance) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type RegData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passport string `protobuf:"bytes,1,opt,name=Passport,proto3" json:"Passport,omitempty"`
	PublicK  string `protobuf:"bytes,2,opt,name=PublicK,proto3" json:"PublicK,omitempty"`
	Salt     string `protobuf:"bytes,3,opt,name=Salt,proto3" json:"Salt,omitempty"`
}

func (x *RegData) Reset() {
	*x = RegData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegData) ProtoMessage() {}

func (x *RegData) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegData.ProtoReflect.Descriptor instead.
func (*RegData) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{1}
}

func (x *RegData) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *RegData) GetPublicK() string {
	if x != nil {
		return x.PublicK
	}
	return ""
}

func (x *RegData) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicK  string `protobuf:"bytes,1,opt,name=PublicK,proto3" json:"PublicK,omitempty"`
	PrivateK string `protobuf:"bytes,2,opt,name=PrivateK,proto3" json:"PrivateK,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{2}
}

func (x *AuthData) GetPublicK() string {
	if x != nil {
		return x.PublicK
	}
	return ""
}

func (x *AuthData) GetPrivateK() string {
	if x != nil {
		return x.PrivateK
	}
	return ""
}

type AuthRegResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distortion string `protobuf:"bytes,1,opt,name=Distortion,proto3" json:"Distortion,omitempty"`
}

func (x *AuthRegResult) Reset() {
	*x = AuthRegResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRegResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRegResult) ProtoMessage() {}

func (x *AuthRegResult) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRegResult.ProtoReflect.Descriptor instead.
func (*AuthRegResult) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{3}
}

func (x *AuthRegResult) GetDistortion() string {
	if x != nil {
		return x.Distortion
	}
	return ""
}

type ResponseSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size string `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ResponseSize) Reset() {
	*x = ResponseSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSize) ProtoMessage() {}

func (x *ResponseSize) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSize.ProtoReflect.Descriptor instead.
func (*ResponseSize) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseSize) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Useradrr string `protobuf:"bytes,1,opt,name=useradrr,proto3" json:"useradrr,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{5}
}

func (x *Address) GetUseradrr() string {
	if x != nil {
		return x.Useradrr
	}
	return ""
}

type Lanb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance string `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Lanb) Reset() {
	*x = Lanb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lanb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lanb) ProtoMessage() {}

func (x *Lanb) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lanb.ProtoReflect.Descriptor instead.
func (*Lanb) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{6}
}

func (x *Lanb) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InBlock []string `protobuf:"bytes,1,rep,name=InBlock,proto3" json:"InBlock,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{7}
}

func (x *Chain) GetInBlock() []string {
	if x != nil {
		return x.InBlock
	}
	return nil
}

type CandidateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidatePK   string `protobuf:"bytes,1,opt,name=candidatePK,proto3" json:"candidatePK,omitempty"`
	CandidateName string `protobuf:"bytes,2,opt,name=candidateName,proto3" json:"candidateName,omitempty"`
}

func (x *CandidateList) Reset() {
	*x = CandidateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateList) ProtoMessage() {}

func (x *CandidateList) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateList.ProtoReflect.Descriptor instead.
func (*CandidateList) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{8}
}

func (x *CandidateList) GetCandidatePK() string {
	if x != nil {
		return x.CandidatePK
	}
	return ""
}

func (x *CandidateList) GetCandidateName() string {
	if x != nil {
		return x.CandidateName
	}
	return ""
}

type LowData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCandidate string `protobuf:"bytes,1,opt,name=UserCandidate,proto3" json:"UserCandidate,omitempty"`
	Num           uint64 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Private       string `protobuf:"bytes,3,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *LowData) Reset() {
	*x = LowData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LowData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LowData) ProtoMessage() {}

func (x *LowData) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LowData.ProtoReflect.Descriptor instead.
func (*LowData) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{9}
}

func (x *LowData) GetUserCandidate() string {
	if x != nil {
		return x.UserCandidate
	}
	return ""
}

func (x *LowData) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *LowData) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

type LowDataChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCandidate string `protobuf:"bytes,1,opt,name=UserCandidate,proto3" json:"UserCandidate,omitempty"`
	Num           uint64 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *LowDataChain) Reset() {
	*x = LowDataChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LowDataChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LowDataChain) ProtoMessage() {}

func (x *LowDataChain) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LowDataChain.ProtoReflect.Descriptor instead.
func (*LowDataChain) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{10}
}

func (x *LowDataChain) GetUserCandidate() string {
	if x != nil {
		return x.UserCandidate
	}
	return ""
}

func (x *LowDataChain) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type IsComplited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ic bool `protobuf:"varint,1,opt,name=ic,proto3" json:"ic,omitempty"`
}

func (x *IsComplited) Reset() {
	*x = IsComplited{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsComplited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsComplited) ProtoMessage() {}

func (x *IsComplited) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsComplited.ProtoReflect.Descriptor instead.
func (*IsComplited) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{11}
}

func (x *IsComplited) GetIc() bool {
	if x != nil {
		return x.Ic
	}
	return false
}

type IsComplitedVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ic string `protobuf:"bytes,1,opt,name=ic,proto3" json:"ic,omitempty"`
}

func (x *IsComplitedVote) Reset() {
	*x = IsComplitedVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsComplitedVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsComplitedVote) ProtoMessage() {}

func (x *IsComplitedVote) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsComplitedVote.ProtoReflect.Descriptor instead.
func (*IsComplitedVote) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{12}
}

func (x *IsComplitedVote) GetIc() string {
	if x != nil {
		return x.Ic
	}
	return ""
}

type BlockDataGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNum string `protobuf:"bytes,1,opt,name=BlockNum,proto3" json:"BlockNum,omitempty"`
}

func (x *BlockDataGet) Reset() {
	*x = BlockDataGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockDataGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockDataGet) ProtoMessage() {}

func (x *BlockDataGet) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockDataGet.ProtoReflect.Descriptor instead.
func (*BlockDataGet) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{13}
}

func (x *BlockDataGet) GetBlockNum() string {
	if x != nil {
		return x.BlockNum
	}
	return ""
}

type TimeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndTime string `protobuf:"bytes,1,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *TimeData) Reset() {
	*x = TimeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MainService_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeData) ProtoMessage() {}

func (x *TimeData) ProtoReflect() protoreflect.Message {
	mi := &file_MainService_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeData.ProtoReflect.Descriptor instead.
func (*TimeData) Descriptor() ([]byte, []int) {
	return file_MainService_proto_rawDescGZIP(), []int{14}
}

func (x *TimeData) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

var File_MainService_proto protoreflect.FileDescriptor

var file_MainService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7c, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x4b, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x53,
	0x0a, 0x07, 0x52, 0x65, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x12,
	0x12, 0x0a, 0x04, 0x53, 0x61, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53,
	0x61, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x22, 0x2f, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x25, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x72,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x72,
	0x72, 0x22, 0x20, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x49,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x57, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x4b, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x5b, 0x0a, 0x07, 0x4c, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x0c,
	0x4c, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x69, 0x63, 0x22, 0x21, 0x0a, 0x0f, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74,
	0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x63, 0x22, 0x2a, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x22, 0x24, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x96, 0x04, 0x0a, 0x0d, 0x42, 0x4c, 0x43,
	0x48, 0x5f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x0c, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x09, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1a, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x08, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x05, 0x2e, 0x4c, 0x61, 0x6e, 0x62, 0x12, 0x3a, 0x0a, 0x0e,
	0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x30, 0x01, 0x12, 0x27, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x1a, 0x0c, 0x2e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x4c, 0x6f, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x65,
	0x64, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x44, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x53,
	0x6f, 0x6c, 0x6f, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0e, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MainService_proto_rawDescOnce sync.Once
	file_MainService_proto_rawDescData = file_MainService_proto_rawDesc
)

func file_MainService_proto_rawDescGZIP() []byte {
	file_MainService_proto_rawDescOnce.Do(func() {
		file_MainService_proto_rawDescData = protoimpl.X.CompressGZIP(file_MainService_proto_rawDescData)
	})
	return file_MainService_proto_rawDescData
}

var file_MainService_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_MainService_proto_goTypes = []interface{}{
	(*CandidateListWithBalance)(nil), // 0: CandidateListWithBalance
	(*RegData)(nil),                  // 1: RegData
	(*AuthData)(nil),                 // 2: AuthData
	(*AuthRegResult)(nil),            // 3: AuthRegResult
	(*ResponseSize)(nil),             // 4: ResponseSize
	(*Address)(nil),                  // 5: Address
	(*Lanb)(nil),                     // 6: Lanb
	(*Chain)(nil),                    // 7: Chain
	(*CandidateList)(nil),            // 8: CandidateList
	(*LowData)(nil),                  // 9: LowData
	(*LowDataChain)(nil),             // 10: LowDataChain
	(*IsComplited)(nil),              // 11: IsComplited
	(*IsComplitedVote)(nil),          // 12: IsComplitedVote
	(*BlockDataGet)(nil),             // 13: BlockDataGet
	(*TimeData)(nil),                 // 14: TimeData
	(*empty.Empty)(nil),              // 15: google.protobuf.Empty
}
var file_MainService_proto_depIdxs = []int32{
	1,  // 0: BLCH_Contract.AuthRegister:input_type -> RegData
	2,  // 1: BLCH_Contract.AuthLogin:input_type -> AuthData
	15, // 2: BLCH_Contract.ChainSize:input_type -> google.protobuf.Empty
	5,  // 3: BLCH_Contract.Balance:input_type -> Address
	15, // 4: BLCH_Contract.ViewCandidates:input_type -> google.protobuf.Empty
	10, // 5: BLCH_Contract.Transfer:input_type -> LowDataChain
	9,  // 6: BLCH_Contract.Vote:input_type -> LowData
	15, // 7: BLCH_Contract.TimeBlock:input_type -> google.protobuf.Empty
	15, // 8: BLCH_Contract.ChainPrint:input_type -> google.protobuf.Empty
	15, // 9: BLCH_Contract.ResultsWinner:input_type -> google.protobuf.Empty
	15, // 10: BLCH_Contract.SoloWinner:input_type -> google.protobuf.Empty
	3,  // 11: BLCH_Contract.AuthRegister:output_type -> AuthRegResult
	3,  // 12: BLCH_Contract.AuthLogin:output_type -> AuthRegResult
	4,  // 13: BLCH_Contract.ChainSize:output_type -> ResponseSize
	6,  // 14: BLCH_Contract.Balance:output_type -> Lanb
	8,  // 15: BLCH_Contract.ViewCandidates:output_type -> CandidateList
	11, // 16: BLCH_Contract.Transfer:output_type -> IsComplited
	12, // 17: BLCH_Contract.Vote:output_type -> IsComplitedVote
	14, // 18: BLCH_Contract.TimeBlock:output_type -> TimeData
	7,  // 19: BLCH_Contract.ChainPrint:output_type -> Chain
	0,  // 20: BLCH_Contract.ResultsWinner:output_type -> CandidateListWithBalance
	8,  // 21: BLCH_Contract.SoloWinner:output_type -> CandidateList
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_MainService_proto_init() }
func file_MainService_proto_init() {
	if File_MainService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MainService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateListWithBalance); i {
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
		file_MainService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegData); i {
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
		file_MainService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_MainService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRegResult); i {
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
		file_MainService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSize); i {
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
		file_MainService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_MainService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lanb); i {
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
		file_MainService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain); i {
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
		file_MainService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateList); i {
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
		file_MainService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LowData); i {
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
		file_MainService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LowDataChain); i {
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
		file_MainService_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsComplited); i {
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
		file_MainService_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsComplitedVote); i {
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
		file_MainService_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockDataGet); i {
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
		file_MainService_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeData); i {
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
			RawDescriptor: file_MainService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MainService_proto_goTypes,
		DependencyIndexes: file_MainService_proto_depIdxs,
		MessageInfos:      file_MainService_proto_msgTypes,
	}.Build()
	File_MainService_proto = out.File
	file_MainService_proto_rawDesc = nil
	file_MainService_proto_goTypes = nil
	file_MainService_proto_depIdxs = nil
}