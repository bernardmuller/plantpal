// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: plants/plantspb/plants.proto

package plants

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

type Plant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Common         string `protobuf:"bytes,2,opt,name=Common,proto3" json:"Common,omitempty"`
	Family         string `protobuf:"bytes,3,opt,name=Family,proto3" json:"Family,omitempty"`
	CreatedAt      string `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt      string `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Latin          string `protobuf:"bytes,6,opt,name=Latin,proto3" json:"Latin,omitempty"`
	Category       string `protobuf:"bytes,7,opt,name=Category,proto3" json:"Category,omitempty"`
	Origin         string `protobuf:"bytes,8,opt,name=Origin,proto3" json:"Origin,omitempty"`
	Climate        string `protobuf:"bytes,9,opt,name=Climate,proto3" json:"Climate,omitempty"`
	TempMax        string `protobuf:"bytes,10,opt,name=TempMax,proto3" json:"TempMax,omitempty"`
	TempMin        string `protobuf:"bytes,11,opt,name=TempMin,proto3" json:"TempMin,omitempty"`
	IdealLight     string `protobuf:"bytes,12,opt,name=IdealLight,proto3" json:"IdealLight,omitempty"`
	ToleratedLight string `protobuf:"bytes,13,opt,name=ToleratedLight,proto3" json:"ToleratedLight,omitempty"`
	Watering       string `protobuf:"bytes,14,opt,name=Watering,proto3" json:"Watering,omitempty"`
	Insects        string `protobuf:"bytes,15,opt,name=Insects,proto3" json:"Insects,omitempty"`
	Diseases       string `protobuf:"bytes,16,opt,name=Diseases,proto3" json:"Diseases,omitempty"`
	Soil           string `protobuf:"bytes,17,opt,name=Soil,proto3" json:"Soil,omitempty"`
	RepotPeriod    string `protobuf:"bytes,18,opt,name=RepotPeriod,proto3" json:"RepotPeriod,omitempty"`
	Use            string `protobuf:"bytes,19,opt,name=Use,proto3" json:"Use,omitempty"`
}

func (x *Plant) Reset() {
	*x = Plant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plants_plantspb_plants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plant) ProtoMessage() {}

func (x *Plant) ProtoReflect() protoreflect.Message {
	mi := &file_plants_plantspb_plants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plant.ProtoReflect.Descriptor instead.
func (*Plant) Descriptor() ([]byte, []int) {
	return file_plants_plantspb_plants_proto_rawDescGZIP(), []int{0}
}

func (x *Plant) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Plant) GetCommon() string {
	if x != nil {
		return x.Common
	}
	return ""
}

func (x *Plant) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *Plant) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Plant) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Plant) GetLatin() string {
	if x != nil {
		return x.Latin
	}
	return ""
}

func (x *Plant) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Plant) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Plant) GetClimate() string {
	if x != nil {
		return x.Climate
	}
	return ""
}

func (x *Plant) GetTempMax() string {
	if x != nil {
		return x.TempMax
	}
	return ""
}

func (x *Plant) GetTempMin() string {
	if x != nil {
		return x.TempMin
	}
	return ""
}

func (x *Plant) GetIdealLight() string {
	if x != nil {
		return x.IdealLight
	}
	return ""
}

func (x *Plant) GetToleratedLight() string {
	if x != nil {
		return x.ToleratedLight
	}
	return ""
}

func (x *Plant) GetWatering() string {
	if x != nil {
		return x.Watering
	}
	return ""
}

func (x *Plant) GetInsects() string {
	if x != nil {
		return x.Insects
	}
	return ""
}

func (x *Plant) GetDiseases() string {
	if x != nil {
		return x.Diseases
	}
	return ""
}

func (x *Plant) GetSoil() string {
	if x != nil {
		return x.Soil
	}
	return ""
}

func (x *Plant) GetRepotPeriod() string {
	if x != nil {
		return x.RepotPeriod
	}
	return ""
}

func (x *Plant) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

type CreatePlantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Common string `protobuf:"bytes,2,opt,name=Common,proto3" json:"Common,omitempty"`
	Family string `protobuf:"bytes,3,opt,name=Family,proto3" json:"Family,omitempty"`
}

func (x *CreatePlantRequest) Reset() {
	*x = CreatePlantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plants_plantspb_plants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlantRequest) ProtoMessage() {}

func (x *CreatePlantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plants_plantspb_plants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlantRequest.ProtoReflect.Descriptor instead.
func (*CreatePlantRequest) Descriptor() ([]byte, []int) {
	return file_plants_plantspb_plants_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlantRequest) GetCommon() string {
	if x != nil {
		return x.Common
	}
	return ""
}

func (x *CreatePlantRequest) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

type CreatePlantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Plant  *Plant `protobuf:"bytes,2,opt,name=plant,proto3" json:"plant,omitempty"`
}

func (x *CreatePlantResponse) Reset() {
	*x = CreatePlantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plants_plantspb_plants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlantResponse) ProtoMessage() {}

func (x *CreatePlantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plants_plantspb_plants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlantResponse.ProtoReflect.Descriptor instead.
func (*CreatePlantResponse) Descriptor() ([]byte, []int) {
	return file_plants_plantspb_plants_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePlantResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreatePlantResponse) GetPlant() *Plant {
	if x != nil {
		return x.Plant
	}
	return nil
}

type GetPlantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlantsRequest) Reset() {
	*x = GetPlantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plants_plantspb_plants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsRequest) ProtoMessage() {}

func (x *GetPlantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plants_plantspb_plants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsRequest.ProtoReflect.Descriptor instead.
func (*GetPlantsRequest) Descriptor() ([]byte, []int) {
	return file_plants_plantspb_plants_proto_rawDescGZIP(), []int{3}
}

type GetPlantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plants []*Plant `protobuf:"bytes,1,rep,name=plants,proto3" json:"plants,omitempty"`
}

func (x *GetPlantsResponse) Reset() {
	*x = GetPlantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plants_plantspb_plants_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantsResponse) ProtoMessage() {}

func (x *GetPlantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plants_plantspb_plants_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantsResponse.ProtoReflect.Descriptor instead.
func (*GetPlantsResponse) Descriptor() ([]byte, []int) {
	return file_plants_plantspb_plants_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlantsResponse) GetPlants() []*Plant {
	if x != nil {
		return x.Plants
	}
	return nil
}

var File_plants_plantspb_plants_proto protoreflect.FileDescriptor

var file_plants_plantspb_plants_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x70,
	0x62, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd,
	0x03, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x74, 0x69, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x74, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x65, 0x6d, 0x70,
	0x4d, 0x61, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x65, 0x6d, 0x70, 0x4d,
	0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x69, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x64, 0x65, 0x61, 0x6c, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x49, 0x64, 0x65, 0x61, 0x6c, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x54, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x65, 0x61, 0x73, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69,
	0x73, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x69, 0x6c, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6f, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6f, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x52, 0x65, 0x70, 0x6f, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x73, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x73, 0x65, 0x22, 0x44,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x22, 0x4b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x32, 0x81, 0x01, 0x0a, 0x0d, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x72,
	0x6e, 0x61, 0x72, 0x64, 0x6d, 0x75, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x70, 0x61, 0x6c, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_plants_plantspb_plants_proto_rawDescOnce sync.Once
	file_plants_plantspb_plants_proto_rawDescData = file_plants_plantspb_plants_proto_rawDesc
)

func file_plants_plantspb_plants_proto_rawDescGZIP() []byte {
	file_plants_plantspb_plants_proto_rawDescOnce.Do(func() {
		file_plants_plantspb_plants_proto_rawDescData = protoimpl.X.CompressGZIP(file_plants_plantspb_plants_proto_rawDescData)
	})
	return file_plants_plantspb_plants_proto_rawDescData
}

var file_plants_plantspb_plants_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plants_plantspb_plants_proto_goTypes = []interface{}{
	(*Plant)(nil),               // 0: Plant
	(*CreatePlantRequest)(nil),  // 1: CreatePlantRequest
	(*CreatePlantResponse)(nil), // 2: CreatePlantResponse
	(*GetPlantsRequest)(nil),    // 3: GetPlantsRequest
	(*GetPlantsResponse)(nil),   // 4: GetPlantsResponse
}
var file_plants_plantspb_plants_proto_depIdxs = []int32{
	0, // 0: CreatePlantResponse.plant:type_name -> Plant
	0, // 1: GetPlantsResponse.plants:type_name -> Plant
	1, // 2: PlantsService.CreatePlant:input_type -> CreatePlantRequest
	3, // 3: PlantsService.GetPlants:input_type -> GetPlantsRequest
	2, // 4: PlantsService.CreatePlant:output_type -> CreatePlantResponse
	4, // 5: PlantsService.GetPlants:output_type -> GetPlantsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plants_plantspb_plants_proto_init() }
func file_plants_plantspb_plants_proto_init() {
	if File_plants_plantspb_plants_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plants_plantspb_plants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plant); i {
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
		file_plants_plantspb_plants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlantRequest); i {
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
		file_plants_plantspb_plants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlantResponse); i {
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
		file_plants_plantspb_plants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantsRequest); i {
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
		file_plants_plantspb_plants_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantsResponse); i {
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
			RawDescriptor: file_plants_plantspb_plants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plants_plantspb_plants_proto_goTypes,
		DependencyIndexes: file_plants_plantspb_plants_proto_depIdxs,
		MessageInfos:      file_plants_plantspb_plants_proto_msgTypes,
	}.Build()
	File_plants_plantspb_plants_proto = out.File
	file_plants_plantspb_plants_proto_rawDesc = nil
	file_plants_plantspb_plants_proto_goTypes = nil
	file_plants_plantspb_plants_proto_depIdxs = nil
}
