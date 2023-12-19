// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: johnjud/file/file/v1/file.proto

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
	FileService_Upload_FullMethodName       = "/johnjud.file.file.v1.FileService/Upload"
	FileService_GetSignedUrl_FullMethodName = "/johnjud.file.file.v1.FileService/GetSignedUrl"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	GetSignedUrl(ctx context.Context, in *GetSignedUrlRequest, opts ...grpc.CallOption) (*GetSignedUrlResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, FileService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetSignedUrl(ctx context.Context, in *GetSignedUrlRequest, opts ...grpc.CallOption) (*GetSignedUrlResponse, error) {
	out := new(GetSignedUrlResponse)
	err := c.cc.Invoke(ctx, FileService_GetSignedUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	GetSignedUrl(context.Context, *GetSignedUrlRequest) (*GetSignedUrlResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileServiceServer) GetSignedUrl(context.Context, *GetSignedUrlRequest) (*GetSignedUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignedUrl not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetSignedUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetSignedUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetSignedUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetSignedUrl(ctx, req.(*GetSignedUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "johnjud.file.file.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _FileService_Upload_Handler,
		},
		{
			MethodName: "GetSignedUrl",
			Handler:    _FileService_GetSignedUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "johnjud/file/file/v1/file.proto",
}
