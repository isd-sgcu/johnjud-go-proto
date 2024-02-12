// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: johnjud/file/image/v1/image.proto

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PetId     string `protobuf:"bytes,2,opt,name=petId,proto3" json:"petId,omitempty"`
	ImageUrl  string `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	ObjectKey string `protobuf:"bytes,4,opt,name=objectKey,proto3" json:"objectKey,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *Image) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Image) GetObjectKey() string {
	if x != nil {
		return x.ObjectKey
	}
	return ""
}

type FindAllImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllImageRequest) Reset() {
	*x = FindAllImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllImageRequest) ProtoMessage() {}

func (x *FindAllImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllImageRequest.ProtoReflect.Descriptor instead.
func (*FindAllImageRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{1}
}

type FindAllImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *FindAllImageResponse) Reset() {
	*x = FindAllImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllImageResponse) ProtoMessage() {}

func (x *FindAllImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllImageResponse.ProtoReflect.Descriptor instead.
func (*FindAllImageResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllImageResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	PetId    string `protobuf:"bytes,3,opt,name=petId,proto3" json:"petId,omitempty"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{3}
}

func (x *UploadImageRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadImageRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadImageRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{4}
}

func (x *UploadImageResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type FindImageByPetIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId string `protobuf:"bytes,1,opt,name=petId,proto3" json:"petId,omitempty"`
}

func (x *FindImageByPetIdRequest) Reset() {
	*x = FindImageByPetIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindImageByPetIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindImageByPetIdRequest) ProtoMessage() {}

func (x *FindImageByPetIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindImageByPetIdRequest.ProtoReflect.Descriptor instead.
func (*FindImageByPetIdRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{5}
}

func (x *FindImageByPetIdRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type FindImageByPetIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *FindImageByPetIdResponse) Reset() {
	*x = FindImageByPetIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindImageByPetIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindImageByPetIdResponse) ProtoMessage() {}

func (x *FindImageByPetIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindImageByPetIdResponse.ProtoReflect.Descriptor instead.
func (*FindImageByPetIdResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{6}
}

func (x *FindImageByPetIdResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type AssignPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	PetId string   `protobuf:"bytes,2,opt,name=petId,proto3" json:"petId,omitempty"`
}

func (x *AssignPetRequest) Reset() {
	*x = AssignPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignPetRequest) ProtoMessage() {}

func (x *AssignPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignPetRequest.ProtoReflect.Descriptor instead.
func (*AssignPetRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{7}
}

func (x *AssignPetRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AssignPetRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type AssignPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AssignPetResponse) Reset() {
	*x = AssignPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignPetResponse) ProtoMessage() {}

func (x *AssignPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignPetResponse.ProtoReflect.Descriptor instead.
func (*AssignPetResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{8}
}

func (x *AssignPetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteImageRequest) Reset() {
	*x = DeleteImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageRequest) ProtoMessage() {}

func (x *DeleteImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageRequest.ProtoReflect.Descriptor instead.
func (*DeleteImageRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteImageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteImageResponse) Reset() {
	*x = DeleteImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageResponse) ProtoMessage() {}

func (x *DeleteImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageResponse.ProtoReflect.Descriptor instead.
func (*DeleteImageResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteImageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteByPetIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId string `protobuf:"bytes,1,opt,name=petId,proto3" json:"petId,omitempty"`
}

func (x *DeleteByPetIdRequest) Reset() {
	*x = DeleteByPetIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByPetIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByPetIdRequest) ProtoMessage() {}

func (x *DeleteByPetIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByPetIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteByPetIdRequest) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteByPetIdRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type DeleteByPetIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteByPetIdResponse) Reset() {
	*x = DeleteByPetIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_johnjud_file_image_v1_image_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByPetIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByPetIdResponse) ProtoMessage() {}

func (x *DeleteByPetIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_johnjud_file_image_v1_image_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByPetIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteByPetIdResponse) Descriptor() ([]byte, []int) {
	return file_johnjud_file_image_v1_image_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteByPetIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_johnjud_file_image_v1_image_proto protoreflect.FileDescriptor

var file_johnjud_file_image_v1_image_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x67, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4b, 0x65, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x65, 0x74, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6f, 0x68,
	0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x2f, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x50, 0x65,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64,
	0x22, 0x50, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x50,
	0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a,
	0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x3a, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2d,
	0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79,
	0x50, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74,
	0x49, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x50, 0x65,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xfc, 0x04, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x12, 0x2a, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x06,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x29, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x70, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2e,
	0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0x79, 0x50, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0x79, 0x50, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x60, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x65, 0x74, 0x12, 0x27,
	0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75,
	0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e,
	0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a,
	0x75, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x79, 0x50, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75,
	0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x50, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x50, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x6a, 0x6f, 0x68, 0x6e, 0x6a, 0x75, 0x64, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_johnjud_file_image_v1_image_proto_rawDescOnce sync.Once
	file_johnjud_file_image_v1_image_proto_rawDescData = file_johnjud_file_image_v1_image_proto_rawDesc
)

func file_johnjud_file_image_v1_image_proto_rawDescGZIP() []byte {
	file_johnjud_file_image_v1_image_proto_rawDescOnce.Do(func() {
		file_johnjud_file_image_v1_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_johnjud_file_image_v1_image_proto_rawDescData)
	})
	return file_johnjud_file_image_v1_image_proto_rawDescData
}

var file_johnjud_file_image_v1_image_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_johnjud_file_image_v1_image_proto_goTypes = []interface{}{
	(*Image)(nil),                    // 0: johnjud.file.image.v1.Image
	(*FindAllImageRequest)(nil),      // 1: johnjud.file.image.v1.FindAllImageRequest
	(*FindAllImageResponse)(nil),     // 2: johnjud.file.image.v1.FindAllImageResponse
	(*UploadImageRequest)(nil),       // 3: johnjud.file.image.v1.UploadImageRequest
	(*UploadImageResponse)(nil),      // 4: johnjud.file.image.v1.UploadImageResponse
	(*FindImageByPetIdRequest)(nil),  // 5: johnjud.file.image.v1.FindImageByPetIdRequest
	(*FindImageByPetIdResponse)(nil), // 6: johnjud.file.image.v1.FindImageByPetIdResponse
	(*AssignPetRequest)(nil),         // 7: johnjud.file.image.v1.AssignPetRequest
	(*AssignPetResponse)(nil),        // 8: johnjud.file.image.v1.AssignPetResponse
	(*DeleteImageRequest)(nil),       // 9: johnjud.file.image.v1.DeleteImageRequest
	(*DeleteImageResponse)(nil),      // 10: johnjud.file.image.v1.DeleteImageResponse
	(*DeleteByPetIdRequest)(nil),     // 11: johnjud.file.image.v1.DeleteByPetIdRequest
	(*DeleteByPetIdResponse)(nil),    // 12: johnjud.file.image.v1.DeleteByPetIdResponse
}
var file_johnjud_file_image_v1_image_proto_depIdxs = []int32{
	0,  // 0: johnjud.file.image.v1.FindAllImageResponse.images:type_name -> johnjud.file.image.v1.Image
	0,  // 1: johnjud.file.image.v1.UploadImageResponse.image:type_name -> johnjud.file.image.v1.Image
	0,  // 2: johnjud.file.image.v1.FindImageByPetIdResponse.images:type_name -> johnjud.file.image.v1.Image
	1,  // 3: johnjud.file.image.v1.ImageService.FindAll:input_type -> johnjud.file.image.v1.FindAllImageRequest
	3,  // 4: johnjud.file.image.v1.ImageService.Upload:input_type -> johnjud.file.image.v1.UploadImageRequest
	5,  // 5: johnjud.file.image.v1.ImageService.FindByPetId:input_type -> johnjud.file.image.v1.FindImageByPetIdRequest
	7,  // 6: johnjud.file.image.v1.ImageService.AssignPet:input_type -> johnjud.file.image.v1.AssignPetRequest
	9,  // 7: johnjud.file.image.v1.ImageService.Delete:input_type -> johnjud.file.image.v1.DeleteImageRequest
	11, // 8: johnjud.file.image.v1.ImageService.DeleteByPetId:input_type -> johnjud.file.image.v1.DeleteByPetIdRequest
	2,  // 9: johnjud.file.image.v1.ImageService.FindAll:output_type -> johnjud.file.image.v1.FindAllImageResponse
	4,  // 10: johnjud.file.image.v1.ImageService.Upload:output_type -> johnjud.file.image.v1.UploadImageResponse
	6,  // 11: johnjud.file.image.v1.ImageService.FindByPetId:output_type -> johnjud.file.image.v1.FindImageByPetIdResponse
	8,  // 12: johnjud.file.image.v1.ImageService.AssignPet:output_type -> johnjud.file.image.v1.AssignPetResponse
	10, // 13: johnjud.file.image.v1.ImageService.Delete:output_type -> johnjud.file.image.v1.DeleteImageResponse
	12, // 14: johnjud.file.image.v1.ImageService.DeleteByPetId:output_type -> johnjud.file.image.v1.DeleteByPetIdResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_johnjud_file_image_v1_image_proto_init() }
func file_johnjud_file_image_v1_image_proto_init() {
	if File_johnjud_file_image_v1_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_johnjud_file_image_v1_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllImageRequest); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllImageResponse); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindImageByPetIdRequest); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindImageByPetIdResponse); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignPetRequest); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignPetResponse); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImageRequest); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImageResponse); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByPetIdRequest); i {
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
		file_johnjud_file_image_v1_image_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByPetIdResponse); i {
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
			RawDescriptor: file_johnjud_file_image_v1_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_johnjud_file_image_v1_image_proto_goTypes,
		DependencyIndexes: file_johnjud_file_image_v1_image_proto_depIdxs,
		MessageInfos:      file_johnjud_file_image_v1_image_proto_msgTypes,
	}.Build()
	File_johnjud_file_image_v1_image_proto = out.File
	file_johnjud_file_image_v1_image_proto_rawDesc = nil
	file_johnjud_file_image_v1_image_proto_goTypes = nil
	file_johnjud_file_image_v1_image_proto_depIdxs = nil
}
