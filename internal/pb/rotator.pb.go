// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/rotator/rotator.proto

package pb

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

type CreateBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateBannerRequest) Reset() {
	*x = CreateBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBannerRequest) ProtoMessage() {}

func (x *CreateBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBannerRequest.ProtoReflect.Descriptor instead.
func (*CreateBannerRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBannerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBannerRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBannerResponse) Reset() {
	*x = CreateBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBannerResponse) ProtoMessage() {}

func (x *CreateBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBannerResponse.ProtoReflect.Descriptor instead.
func (*CreateBannerResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBannerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateSlotRequest) Reset() {
	*x = CreateSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlotRequest) ProtoMessage() {}

func (x *CreateSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlotRequest.ProtoReflect.Descriptor instead.
func (*CreateSlotRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSlotRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateSlotRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateSlotResponse) Reset() {
	*x = CreateSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlotResponse) ProtoMessage() {}

func (x *CreateSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlotResponse.ProtoReflect.Descriptor instead.
func (*CreateSlotResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSlotResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGroupResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int32 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   int32 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *AddBannerRequest) Reset() {
	*x = AddBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerRequest) ProtoMessage() {}

func (x *AddBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerRequest.ProtoReflect.Descriptor instead.
func (*AddBannerRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{6}
}

func (x *AddBannerRequest) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *AddBannerRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type AddBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BannerId int32 `protobuf:"varint,2,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   int32 `protobuf:"varint,3,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *AddBannerResponse) Reset() {
	*x = AddBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerResponse) ProtoMessage() {}

func (x *AddBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerResponse.ProtoReflect.Descriptor instead.
func (*AddBannerResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{7}
}

func (x *AddBannerResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddBannerResponse) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *AddBannerResponse) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type DeleteBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int32 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   int32 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *DeleteBannerRequest) Reset() {
	*x = DeleteBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBannerRequest) ProtoMessage() {}

func (x *DeleteBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBannerRequest.ProtoReflect.Descriptor instead.
func (*DeleteBannerRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBannerRequest) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *DeleteBannerRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type DeleteBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBannerResponse) Reset() {
	*x = DeleteBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBannerResponse) ProtoMessage() {}

func (x *DeleteBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBannerResponse.ProtoReflect.Descriptor instead.
func (*DeleteBannerResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteBannerResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddClickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int32 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   int32 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	GroupId  int32 `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *AddClickRequest) Reset() {
	*x = AddClickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClickRequest) ProtoMessage() {}

func (x *AddClickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClickRequest.ProtoReflect.Descriptor instead.
func (*AddClickRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{10}
}

func (x *AddClickRequest) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *AddClickRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *AddClickRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type AddClickResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddClickResponse) Reset() {
	*x = AddClickResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClickResponse) ProtoMessage() {}

func (x *AddClickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClickResponse.ProtoReflect.Descriptor instead.
func (*AddClickResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{11}
}

func (x *AddClickResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddBannerDisplayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId  int32 `protobuf:"varint,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	GroupId int32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *AddBannerDisplayRequest) Reset() {
	*x = AddBannerDisplayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerDisplayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerDisplayRequest) ProtoMessage() {}

func (x *AddBannerDisplayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerDisplayRequest.ProtoReflect.Descriptor instead.
func (*AddBannerDisplayRequest) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{12}
}

func (x *AddBannerDisplayRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *AddBannerDisplayRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type AddBannerDisplayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddBannerDisplayResponse) Reset() {
	*x = AddBannerDisplayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rotator_rotator_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerDisplayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerDisplayResponse) ProtoMessage() {}

func (x *AddBannerDisplayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rotator_rotator_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerDisplayResponse.ProtoReflect.Descriptor instead.
func (*AddBannerDisplayResponse) Descriptor() ([]byte, []int) {
	return file_api_rotator_rotator_proto_rawDescGZIP(), []int{13}
}

func (x *AddBannerDisplayResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_rotator_rotator_proto protoreflect.FileDescriptor

var file_api_rotator_rotator_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x48, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6c, 0x6f,
	0x74, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6c,
	0x6f, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x92,
	0x04, 0x0a, 0x07, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x2e,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x20,
	0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rotator_rotator_proto_rawDescOnce sync.Once
	file_api_rotator_rotator_proto_rawDescData = file_api_rotator_rotator_proto_rawDesc
)

func file_api_rotator_rotator_proto_rawDescGZIP() []byte {
	file_api_rotator_rotator_proto_rawDescOnce.Do(func() {
		file_api_rotator_rotator_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rotator_rotator_proto_rawDescData)
	})
	return file_api_rotator_rotator_proto_rawDescData
}

var file_api_rotator_rotator_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_rotator_rotator_proto_goTypes = []interface{}{
	(*CreateBannerRequest)(nil),      // 0: rotator.CreateBannerRequest
	(*CreateBannerResponse)(nil),     // 1: rotator.CreateBannerResponse
	(*CreateSlotRequest)(nil),        // 2: rotator.CreateSlotRequest
	(*CreateSlotResponse)(nil),       // 3: rotator.CreateSlotResponse
	(*CreateGroupRequest)(nil),       // 4: rotator.CreateGroupRequest
	(*CreateGroupResponse)(nil),      // 5: rotator.CreateGroupResponse
	(*AddBannerRequest)(nil),         // 6: rotator.AddBannerRequest
	(*AddBannerResponse)(nil),        // 7: rotator.AddBannerResponse
	(*DeleteBannerRequest)(nil),      // 8: rotator.DeleteBannerRequest
	(*DeleteBannerResponse)(nil),     // 9: rotator.DeleteBannerResponse
	(*AddClickRequest)(nil),          // 10: rotator.AddClickRequest
	(*AddClickResponse)(nil),         // 11: rotator.AddClickResponse
	(*AddBannerDisplayRequest)(nil),  // 12: rotator.AddBannerDisplayRequest
	(*AddBannerDisplayResponse)(nil), // 13: rotator.AddBannerDisplayResponse
}
var file_api_rotator_rotator_proto_depIdxs = []int32{
	0,  // 0: rotator.Rotator.CreateBanner:input_type -> rotator.CreateBannerRequest
	2,  // 1: rotator.Rotator.CreateSlot:input_type -> rotator.CreateSlotRequest
	4,  // 2: rotator.Rotator.CreateGroup:input_type -> rotator.CreateGroupRequest
	6,  // 3: rotator.Rotator.AddBanner:input_type -> rotator.AddBannerRequest
	8,  // 4: rotator.Rotator.DeleteBanner:input_type -> rotator.DeleteBannerRequest
	10, // 5: rotator.Rotator.AddClick:input_type -> rotator.AddClickRequest
	12, // 6: rotator.Rotator.AddBannerDisplay:input_type -> rotator.AddBannerDisplayRequest
	1,  // 7: rotator.Rotator.CreateBanner:output_type -> rotator.CreateBannerResponse
	3,  // 8: rotator.Rotator.CreateSlot:output_type -> rotator.CreateSlotResponse
	5,  // 9: rotator.Rotator.CreateGroup:output_type -> rotator.CreateGroupResponse
	7,  // 10: rotator.Rotator.AddBanner:output_type -> rotator.AddBannerResponse
	9,  // 11: rotator.Rotator.DeleteBanner:output_type -> rotator.DeleteBannerResponse
	11, // 12: rotator.Rotator.AddClick:output_type -> rotator.AddClickResponse
	13, // 13: rotator.Rotator.AddBannerDisplay:output_type -> rotator.AddBannerDisplayResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_rotator_rotator_proto_init() }
func file_api_rotator_rotator_proto_init() {
	if File_api_rotator_rotator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rotator_rotator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBannerRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBannerResponse); i {
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
		file_api_rotator_rotator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlotRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlotResponse); i {
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
		file_api_rotator_rotator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_api_rotator_rotator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerResponse); i {
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
		file_api_rotator_rotator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBannerRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBannerResponse); i {
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
		file_api_rotator_rotator_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClickRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClickResponse); i {
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
		file_api_rotator_rotator_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerDisplayRequest); i {
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
		file_api_rotator_rotator_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerDisplayResponse); i {
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
			RawDescriptor: file_api_rotator_rotator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rotator_rotator_proto_goTypes,
		DependencyIndexes: file_api_rotator_rotator_proto_depIdxs,
		MessageInfos:      file_api_rotator_rotator_proto_msgTypes,
	}.Build()
	File_api_rotator_rotator_proto = out.File
	file_api_rotator_rotator_proto_rawDesc = nil
	file_api_rotator_rotator_proto_goTypes = nil
	file_api_rotator_rotator_proto_depIdxs = nil
}