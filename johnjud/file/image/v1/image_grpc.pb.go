// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: johnjud/file/image/v1/image.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ImageService_Upload_FullMethodName      = "/johnjud.file.image.v1.ImageService/Upload"
	ImageService_FindByPetId_FullMethodName = "/johnjud.file.image.v1.ImageService/FindByPetId"
	ImageService_AssignPet_FullMethodName   = "/johnjud.file.image.v1.ImageService/AssignPet"
	ImageService_Delete_FullMethodName      = "/johnjud.file.image.v1.ImageService/Delete"
)

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	Upload(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error)
	FindByPetId(ctx context.Context, in *FindImageByPetIdRequest, opts ...grpc.CallOption) (*FindImageByPetIdResponse, error)
	AssignPet(ctx context.Context, in *AssignPetRequest, opts ...grpc.CallOption) (*AssignPetResponse, error)
	Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) Upload(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error) {
	out := new(UploadImageResponse)
	err := c.cc.Invoke(ctx, ImageService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) FindByPetId(ctx context.Context, in *FindImageByPetIdRequest, opts ...grpc.CallOption) (*FindImageByPetIdResponse, error) {
	out := new(FindImageByPetIdResponse)
	err := c.cc.Invoke(ctx, ImageService_FindByPetId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) AssignPet(ctx context.Context, in *AssignPetRequest, opts ...grpc.CallOption) (*AssignPetResponse, error) {
	out := new(AssignPetResponse)
	err := c.cc.Invoke(ctx, ImageService_AssignPet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error) {
	out := new(DeleteImageResponse)
	err := c.cc.Invoke(ctx, ImageService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	Upload(context.Context, *UploadImageRequest) (*UploadImageResponse, error)
	FindByPetId(context.Context, *FindImageByPetIdRequest) (*FindImageByPetIdResponse, error)
	AssignPet(context.Context, *AssignPetRequest) (*AssignPetResponse, error)
	Delete(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) Upload(context.Context, *UploadImageRequest) (*UploadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedImageServiceServer) FindByPetId(context.Context, *FindImageByPetIdRequest) (*FindImageByPetIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByPetId not implemented")
}
func (UnimplementedImageServiceServer) AssignPet(context.Context, *AssignPetRequest) (*AssignPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignPet not implemented")
}
func (UnimplementedImageServiceServer) Delete(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Upload(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_FindByPetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindImageByPetIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).FindByPetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_FindByPetId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).FindByPetId(ctx, req.(*FindImageByPetIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_AssignPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).AssignPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_AssignPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).AssignPet(ctx, req.(*AssignPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Delete(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "johnjud.file.image.v1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _ImageService_Upload_Handler,
		},
		{
			MethodName: "FindByPetId",
			Handler:    _ImageService_FindByPetId_Handler,
		},
		{
			MethodName: "AssignPet",
			Handler:    _ImageService_AssignPet_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ImageService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "johnjud/file/image/v1/image.proto",
}
