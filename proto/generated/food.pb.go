// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: food.proto

package proto

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

type FoodCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl       string  `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Description    string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Id             string  `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	RestarauntName string  `protobuf:"bytes,5,opt,name=restarauntName,proto3" json:"restarauntName,omitempty"`
	Price          float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	RestarauntUrl  string  `protobuf:"bytes,7,opt,name=restarauntUrl,proto3" json:"restarauntUrl,omitempty"`
}

func (x *FoodCard) Reset() {
	*x = FoodCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodCard) ProtoMessage() {}

func (x *FoodCard) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodCard.ProtoReflect.Descriptor instead.
func (*FoodCard) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *FoodCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodCard) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *FoodCard) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FoodCard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FoodCard) GetRestarauntName() string {
	if x != nil {
		return x.RestarauntName
	}
	return ""
}

func (x *FoodCard) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FoodCard) GetRestarauntUrl() string {
	if x != nil {
		return x.RestarauntUrl
	}
	return ""
}

type FoodFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KitchenWhiteList    []string `protobuf:"bytes,1,rep,name=kitchenWhiteList,proto3" json:"kitchenWhiteList,omitempty"`
	RestarauntWhiteList []string `protobuf:"bytes,2,rep,name=restarauntWhiteList,proto3" json:"restarauntWhiteList,omitempty"`
	KitchenBlackList    []string `protobuf:"bytes,3,rep,name=kitchenBlackList,proto3" json:"kitchenBlackList,omitempty"`
	RestarauntBlackList []string `protobuf:"bytes,4,rep,name=restarauntBlackList,proto3" json:"restarauntBlackList,omitempty"`
}

func (x *FoodFilter) Reset() {
	*x = FoodFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodFilter) ProtoMessage() {}

func (x *FoodFilter) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodFilter.ProtoReflect.Descriptor instead.
func (*FoodFilter) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *FoodFilter) GetKitchenWhiteList() []string {
	if x != nil {
		return x.KitchenWhiteList
	}
	return nil
}

func (x *FoodFilter) GetRestarauntWhiteList() []string {
	if x != nil {
		return x.RestarauntWhiteList
	}
	return nil
}

func (x *FoodFilter) GetKitchenBlackList() []string {
	if x != nil {
		return x.KitchenBlackList
	}
	return nil
}

func (x *FoodFilter) GetRestarauntBlackList() []string {
	if x != nil {
		return x.RestarauntBlackList
	}
	return nil
}

type FoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardsNum   int64       `protobuf:"varint,1,opt,name=cardsNum,proto3" json:"cardsNum,omitempty"`
	FoodFilter *FoodFilter `protobuf:"bytes,2,opt,name=foodFilter,proto3" json:"foodFilter,omitempty"`
	Longitude  float32     `protobuf:"fixed32,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude   float32     `protobuf:"fixed32,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *FoodRequest) Reset() {
	*x = FoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodRequest) ProtoMessage() {}

func (x *FoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodRequest.ProtoReflect.Descriptor instead.
func (*FoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *FoodRequest) GetCardsNum() int64 {
	if x != nil {
		return x.CardsNum
	}
	return 0
}

func (x *FoodRequest) GetFoodFilter() *FoodFilter {
	if x != nil {
		return x.FoodFilter
	}
	return nil
}

func (x *FoodRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *FoodRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type FoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodCards []*FoodCard `protobuf:"bytes,1,rep,name=foodCards,proto3" json:"foodCards,omitempty"`
	Succeed   bool        `protobuf:"varint,2,opt,name=succeed,proto3" json:"succeed,omitempty"`
}

func (x *FoodResponse) Reset() {
	*x = FoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodResponse) ProtoMessage() {}

func (x *FoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodResponse.ProtoReflect.Descriptor instead.
func (*FoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *FoodResponse) GetFoodCards() []*FoodCard {
	if x != nil {
		return x.FoodCards
	}
	return nil
}

func (x *FoodResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

type FoodFromCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardsNum int64   `protobuf:"varint,1,opt,name=cardsNum,proto3" json:"cardsNum,omitempty"`
	FoodIds  []int64 `protobuf:"varint,2,rep,packed,name=foodIds,proto3" json:"foodIds,omitempty"`
}

func (x *FoodFromCollectionRequest) Reset() {
	*x = FoodFromCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodFromCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodFromCollectionRequest) ProtoMessage() {}

func (x *FoodFromCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodFromCollectionRequest.ProtoReflect.Descriptor instead.
func (*FoodFromCollectionRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *FoodFromCollectionRequest) GetCardsNum() int64 {
	if x != nil {
		return x.CardsNum
	}
	return 0
}

func (x *FoodFromCollectionRequest) GetFoodIds() []int64 {
	if x != nil {
		return x.FoodIds
	}
	return nil
}

type FoodFromCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodCards []*FoodCard `protobuf:"bytes,1,rep,name=foodCards,proto3" json:"foodCards,omitempty"`
	Succeed   bool        `protobuf:"varint,2,opt,name=succeed,proto3" json:"succeed,omitempty"`
}

func (x *FoodFromCollectionResponse) Reset() {
	*x = FoodFromCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodFromCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodFromCollectionResponse) ProtoMessage() {}

func (x *FoodFromCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodFromCollectionResponse.ProtoReflect.Descriptor instead.
func (*FoodFromCollectionResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *FoodFromCollectionResponse) GetFoodCards() []*FoodCard {
	if x != nil {
		return x.FoodCards
	}
	return nil
}

func (x *FoodFromCollectionResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x22, 0xd0, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x61, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x55, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x61, 0x75,
	0x6e, 0x74, 0x55, 0x72, 0x6c, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x57, 0x68,
	0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x6b, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x6e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x95, 0x01, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x30, 0x0a, 0x0a,
	0x66, 0x6f, 0x6f, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x56, 0x0a, 0x0c, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x64,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6f,
	0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x09, 0x66, 0x6f, 0x6f,
	0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64,
	0x22, 0x51, 0x0a, 0x19, 0x46, 0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x72, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x61, 0x72, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6f,
	0x64, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x66, 0x6f, 0x6f, 0x64,
	0x49, 0x64, 0x73, 0x22, 0x64, 0x0a, 0x1a, 0x46, 0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x32, 0xa8, 0x01, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x11, 0x2e, 0x66, 0x6f,
	0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x46,
	0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x46, 0x72, 0x6f,
	0x6d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_food_proto_goTypes = []interface{}{
	(*FoodCard)(nil),                   // 0: food.FoodCard
	(*FoodFilter)(nil),                 // 1: food.FoodFilter
	(*FoodRequest)(nil),                // 2: food.FoodRequest
	(*FoodResponse)(nil),               // 3: food.FoodResponse
	(*FoodFromCollectionRequest)(nil),  // 4: food.FoodFromCollectionRequest
	(*FoodFromCollectionResponse)(nil), // 5: food.FoodFromCollectionResponse
}
var file_food_proto_depIdxs = []int32{
	1, // 0: food.FoodRequest.foodFilter:type_name -> food.FoodFilter
	0, // 1: food.FoodResponse.foodCards:type_name -> food.FoodCard
	0, // 2: food.FoodFromCollectionResponse.foodCards:type_name -> food.FoodCard
	2, // 3: food.DeliveryClub.GetRandomFood:input_type -> food.FoodRequest
	4, // 4: food.DeliveryClub.GetRandomFoodFromCollection:input_type -> food.FoodFromCollectionRequest
	3, // 5: food.DeliveryClub.GetRandomFood:output_type -> food.FoodResponse
	5, // 6: food.DeliveryClub.GetRandomFoodFromCollection:output_type -> food.FoodFromCollectionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodCard); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodFilter); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodRequest); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodResponse); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodFromCollectionRequest); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodFromCollectionResponse); i {
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
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}