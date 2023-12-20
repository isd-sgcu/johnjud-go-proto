// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: johnjud/backend/pet/v1/pet.proto

package v1

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

type PetStatus int32

const (
	PetStatus_FINDHOME PetStatus = 0
	PetStatus_ADOPTED  PetStatus = 1
)

// Enum value maps for PetStatus.
var (
	PetStatus_name = map[int32]string{
		0: "FINDHOME",
		1: "ADOPTED",
	}
	PetStatus_value = map[string]int32{
		"FINDHOME": 0,
		"ADOPTED":  1,
	}
)

func (x PetStatus) Enum() *PetStatus {
	p := new(PetStatus)
	*p = x
	return p
}

func (x PetStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PetStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_johnjud_backend_pet_v1_pet_proto_enumTypes[0].Descriptor()
}

func (PetStatus) Type() protoreflect.EnumType {
	return &file_johnjud_backend_pet_v1_pet_proto_enumTypes[0]
}

func (x PetStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PetStatus.Descriptor instead.
func (PetStatus) EnumDescriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{0}
}

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_johnjud_backend_pet_v1_pet_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_johnjud_backend_pet_v1_pet_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{1}
}

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type       string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Species    string    `protobuf:"bytes,3,opt,name=species,proto3" json:"species,omitempty"`
	Name       string    `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Birthdate  string    `protobuf:"bytes,5,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Gender     Gender    `protobuf:"varint,6,opt,name=gender,proto3,enum=johnjud.backend.pet.v1.Gender" json:"gender,omitempty"`
	Sterile    bool      `protobuf:"varint,7,opt,name=sterile,proto3" json:"sterile,omitempty"`
	Vaccine    bool      `protobuf:"varint,8,opt,name=vaccine,proto3" json:"vaccine,omitempty"`
	Status     PetStatus `protobuf:"varint,9,opt,name=status,proto3,enum=johnjud.backend.pet.v1.PetStatus" json:"status,omitempty"`
	Habit      string    `protobuf:"bytes,10,opt,name=habit,proto3" json:"habit,omitempty"`
	Caption    string    `protobuf:"bytes,11,opt,name=caption,proto3" json:"caption,omitempty"`
	Visible    bool      `protobuf:"varint,12,opt,name=visible,proto3" json:"visible,omitempty"`
	ImageUrl   string    `protobuf:"bytes,13,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	IsClubPet  bool      `protobuf:"varint,14,opt,name=isClubPet,proto3" json:"isClubPet,omitempty"`
	Background string    `protobuf:"bytes,15,opt,name=background,proto3" json:"background,omitempty"`
	Address    string    `protobuf:"bytes,16,opt,name=address,proto3" json:"address,omitempty"`
	Contact    string    `protobuf:"bytes,17,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Pet) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetBirthdate() string {
	if x != nil {
		return x.Birthdate
	}
	return ""
}

func (x *Pet) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

func (x *Pet) GetSterile() bool {
	if x != nil {
		return x.Sterile
	}
	return false
}

func (x *Pet) GetVaccine() bool {
	if x != nil {
		return x.Vaccine
	}
	return false
}

func (x *Pet) GetStatus() PetStatus {
	if x != nil {
		return x.Status
	}
	return PetStatus_FINDHOME
}

func (x *Pet) GetHabit() string {
	if x != nil {
		return x.Habit
	}
	return ""
}

func (x *Pet) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Pet) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Pet) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Pet) GetIsClubPet() bool {
	if x != nil {
		return x.IsClubPet
	}
	return false
}

func (x *Pet) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *Pet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Pet) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type FindAllPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllPetRequest) Reset() {
	*x = FindAllPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllPetRequest) ProtoMessage() {}

func (x *FindAllPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllPetRequest.ProtoReflect.Descriptor instead.
func (*FindAllPetRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{1}
}

type FindAllPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=Pets,proto3" json:"Pets,omitempty"`
}

func (x *FindAllPetResponse) Reset() {
	*x = FindAllPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllPetResponse) ProtoMessage() {}

func (x *FindAllPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllPetResponse.ProtoReflect.Descriptor instead.
func (*FindAllPetResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllPetResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

type FindOnePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOnePetRequest) Reset() {
	*x = FindOnePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOnePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOnePetRequest) ProtoMessage() {}

func (x *FindOnePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOnePetRequest.ProtoReflect.Descriptor instead.
func (*FindOnePetRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{3}
}

func (x *FindOnePetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOnePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=Pet,proto3" json:"Pet,omitempty"`
}

func (x *FindOnePetResponse) Reset() {
	*x = FindOnePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOnePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOnePetResponse) ProtoMessage() {}

func (x *FindOnePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOnePetResponse.ProtoReflect.Descriptor instead.
func (*FindOnePetResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{4}
}

func (x *FindOnePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type CreatePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=Pet,proto3" json:"Pet,omitempty"`
}

func (x *CreatePetRequest) Reset() {
	*x = CreatePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetRequest) ProtoMessage() {}

func (x *CreatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetRequest.ProtoReflect.Descriptor instead.
func (*CreatePetRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePetRequest) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type CreatePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=Pet,proto3" json:"Pet,omitempty"`
}

func (x *CreatePetResponse) Reset() {
	*x = CreatePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetResponse) ProtoMessage() {}

func (x *CreatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetResponse.ProtoReflect.Descriptor instead.
func (*CreatePetResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type UpdatePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=Pet,proto3" json:"Pet,omitempty"`
}

func (x *UpdatePetRequest) Reset() {
	*x = UpdatePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetRequest) ProtoMessage() {}

func (x *UpdatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePetRequest) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type UpdatePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=Pet,proto3" json:"Pet,omitempty"`
}

func (x *UpdatePetResponse) Reset() {
	*x = UpdatePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetResponse) ProtoMessage() {}

func (x *UpdatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetResponse.ProtoReflect.Descriptor instead.
func (*UpdatePetResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type DeletePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePetRequest) Reset() {
	*x = DeletePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetRequest) ProtoMessage() {}

func (x *DeletePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetRequest.ProtoReflect.Descriptor instead.
func (*DeletePetRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePetResponse) Reset() {
	*x = DeletePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetResponse) ProtoMessage() {}

func (x *DeletePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_backend_pet_v1_pet_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetResponse.ProtoReflect.Descriptor instead.
func (*DeletePetResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_johnjud_backend_pet_v1_pet_proto protoreflect.FileDescriptor

var file_johnjud_backend_pet_v1_pet_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xf4, 0x03, 0x0a, 0x03, 0x50,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x65, 0x72, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x65,
	0x72, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x63, 0x63, 0x69, 0x6e, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x61, 0x63, 0x63, 0x69, 0x6e, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x62,
	0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x61, 0x62, 0x69, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x65, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x65, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x22, 0x13, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04,
	0x50, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x68,
	0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x50, 0x65, 0x74, 0x73, 0x22, 0x23, 0x0a,
	0x11, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x74, 0x52, 0x03, 0x50, 0x65, 0x74, 0x22, 0x41, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x50,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x50, 0x65, 0x74, 0x22, 0x42, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a,
	0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x50, 0x65, 0x74, 0x22, 0x41,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x50, 0x65,
	0x74, 0x22, 0x42, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74,
	0x52, 0x03, 0x50, 0x65, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x26, 0x0a, 0x09, 0x50, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x44, 0x48, 0x4f, 0x4d,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x4f, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x2a, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01,
	0x32, 0xf7, 0x03, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x62, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x29, 0x2e, 0x6a, 0x6f, 0x68,
	0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x29,
	0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6a, 0x6f, 0x68, 0x6e,
	0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6a, 0x6f,
	0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x28, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6a,
	0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x6a, 0x6f,
	0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x65,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_johnjud_backend_pet_v1_pet_proto_rawDescOnce sync.Once
	file_johnjud_backend_pet_v1_pet_proto_rawDescData = file_johnjud_backend_pet_v1_pet_proto_rawDesc
)

func file_johnjud_backend_pet_v1_pet_proto_rawDescGZIP() []byte {
	file_johnjud_backend_pet_v1_pet_proto_rawDescOnce.Do(func() {
		file_johnjud_backend_pet_v1_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_johnjud_backend_pet_v1_pet_proto_rawDescData)
	})
	return file_johnjud_backend_pet_v1_pet_proto_rawDescData
}

var file_johnjud_backend_pet_v1_pet_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_johnjud_backend_pet_v1_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_johnjud_backend_pet_v1_pet_proto_goTypes = []interface{}{
	(PetStatus)(0),             // 0: johnjud.backend.pet.v1.PetStatus
	(Gender)(0),                // 1: johnjud.backend.pet.v1.Gender
	(*Pet)(nil),                // 2: johnjud.backend.pet.v1.Pet
	(*FindAllPetRequest)(nil),  // 3: johnjud.backend.pet.v1.FindAllPetRequest
	(*FindAllPetResponse)(nil), // 4: johnjud.backend.pet.v1.FindAllPetResponse
	(*FindOnePetRequest)(nil),  // 5: johnjud.backend.pet.v1.FindOnePetRequest
	(*FindOnePetResponse)(nil), // 6: johnjud.backend.pet.v1.FindOnePetResponse
	(*CreatePetRequest)(nil),   // 7: johnjud.backend.pet.v1.CreatePetRequest
	(*CreatePetResponse)(nil),  // 8: johnjud.backend.pet.v1.CreatePetResponse
	(*UpdatePetRequest)(nil),   // 9: johnjud.backend.pet.v1.UpdatePetRequest
	(*UpdatePetResponse)(nil),  // 10: johnjud.backend.pet.v1.UpdatePetResponse
	(*DeletePetRequest)(nil),   // 11: johnjud.backend.pet.v1.DeletePetRequest
	(*DeletePetResponse)(nil),  // 12: johnjud.backend.pet.v1.DeletePetResponse
}
var file_johnjud_backend_pet_v1_pet_proto_depIdxs = []int32{
	1,  // 0: johnjud.backend.pet.v1.Pet.gender:type_name -> johnjud.backend.pet.v1.Gender
	0,  // 1: johnjud.backend.pet.v1.Pet.status:type_name -> johnjud.backend.pet.v1.PetStatus
	2,  // 2: johnjud.backend.pet.v1.FindAllPetResponse.Pets:type_name -> johnjud.backend.pet.v1.Pet
	2,  // 3: johnjud.backend.pet.v1.FindOnePetResponse.Pet:type_name -> johnjud.backend.pet.v1.Pet
	2,  // 4: johnjud.backend.pet.v1.CreatePetRequest.Pet:type_name -> johnjud.backend.pet.v1.Pet
	2,  // 5: johnjud.backend.pet.v1.CreatePetResponse.Pet:type_name -> johnjud.backend.pet.v1.Pet
	2,  // 6: johnjud.backend.pet.v1.UpdatePetRequest.Pet:type_name -> johnjud.backend.pet.v1.Pet
	2,  // 7: johnjud.backend.pet.v1.UpdatePetResponse.Pet:type_name -> johnjud.backend.pet.v1.Pet
	3,  // 8: johnjud.backend.pet.v1.PetService.FindAll:input_type -> johnjud.backend.pet.v1.FindAllPetRequest
	5,  // 9: johnjud.backend.pet.v1.PetService.FindOne:input_type -> johnjud.backend.pet.v1.FindOnePetRequest
	7,  // 10: johnjud.backend.pet.v1.PetService.Create:input_type -> johnjud.backend.pet.v1.CreatePetRequest
	9,  // 11: johnjud.backend.pet.v1.PetService.Update:input_type -> johnjud.backend.pet.v1.UpdatePetRequest
	11, // 12: johnjud.backend.pet.v1.PetService.Delete:input_type -> johnjud.backend.pet.v1.DeletePetRequest
	4,  // 13: johnjud.backend.pet.v1.PetService.FindAll:output_type -> johnjud.backend.pet.v1.FindAllPetResponse
	6,  // 14: johnjud.backend.pet.v1.PetService.FindOne:output_type -> johnjud.backend.pet.v1.FindOnePetResponse
	8,  // 15: johnjud.backend.pet.v1.PetService.Create:output_type -> johnjud.backend.pet.v1.CreatePetResponse
	10, // 16: johnjud.backend.pet.v1.PetService.Update:output_type -> johnjud.backend.pet.v1.UpdatePetResponse
	12, // 17: johnjud.backend.pet.v1.PetService.Delete:output_type -> johnjud.backend.pet.v1.DeletePetResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_johnjud_backend_pet_v1_pet_proto_init() }
func file_johnjud_backend_pet_v1_pet_proto_init() {
	if File_johnjud_backend_pet_v1_pet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllPetRequest); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllPetResponse); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOnePetRequest); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOnePetResponse); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetRequest); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetResponse); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetRequest); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetResponse); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetRequest); i {
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
		file_johnjud_backend_pet_v1_pet_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetResponse); i {
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
			RawDescriptor: file_johnjud_backend_pet_v1_pet_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_johnjud_backend_pet_v1_pet_proto_goTypes,
		DependencyIndexes: file_johnjud_backend_pet_v1_pet_proto_depIdxs,
		EnumInfos:         file_johnjud_backend_pet_v1_pet_proto_enumTypes,
		MessageInfos:      file_johnjud_backend_pet_v1_pet_proto_msgTypes,
	}.Build()
	File_johnjud_backend_pet_v1_pet_proto = out.File
	file_johnjud_backend_pet_v1_pet_proto_rawDesc = nil
	file_johnjud_backend_pet_v1_pet_proto_goTypes = nil
	file_johnjud_backend_pet_v1_pet_proto_depIdxs = nil
}
