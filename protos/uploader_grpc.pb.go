// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cliente

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ChunksUploadClient is the client API for ChunksUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunksUploadClient interface {
	UploadChunk(ctx context.Context, opts ...grpc.CallOption) (ChunksUpload_UploadChunkClient, error)
	Propuesta(ctx context.Context, in *Prop, opts ...grpc.CallOption) (*Accept, error)
	SendChunk(ctx context.Context, opts ...grpc.CallOption) (ChunksUpload_SendChunkClient, error)
	SendLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Accept, error)
	RequestAdress(ctx context.Context, in *Prop, opts ...grpc.CallOption) (*Adress, error)
	DownloadChunk(ctx context.Context, in *Prop, opts ...grpc.CallOption) (*Chunk, error)
	PropuestaCentralizada(ctx context.Context, in *Adress, opts ...grpc.CallOption) (*Adress, error)
}

type chunksUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewChunksUploadClient(cc grpc.ClientConnInterface) ChunksUploadClient {
	return &chunksUploadClient{cc}
}

func (c *chunksUploadClient) UploadChunk(ctx context.Context, opts ...grpc.CallOption) (ChunksUpload_UploadChunkClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChunksUpload_ServiceDesc.Streams[0], "/cliente.ChunksUpload/UploadChunk", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunksUploadUploadChunkClient{stream}
	return x, nil
}

type ChunksUpload_UploadChunkClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type chunksUploadUploadChunkClient struct {
	grpc.ClientStream
}

func (x *chunksUploadUploadChunkClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chunksUploadUploadChunkClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chunksUploadClient) Propuesta(ctx context.Context, in *Prop, opts ...grpc.CallOption) (*Accept, error) {
	out := new(Accept)
	err := c.cc.Invoke(ctx, "/cliente.ChunksUpload/Propuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunksUploadClient) SendChunk(ctx context.Context, opts ...grpc.CallOption) (ChunksUpload_SendChunkClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChunksUpload_ServiceDesc.Streams[1], "/cliente.ChunksUpload/SendChunk", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunksUploadSendChunkClient{stream}
	return x, nil
}

type ChunksUpload_SendChunkClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type chunksUploadSendChunkClient struct {
	grpc.ClientStream
}

func (x *chunksUploadSendChunkClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chunksUploadSendChunkClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chunksUploadClient) SendLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Accept, error) {
	out := new(Accept)
	err := c.cc.Invoke(ctx, "/cliente.ChunksUpload/SendLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunksUploadClient) RequestAdress(ctx context.Context, in *Prop, opts ...grpc.CallOption) (*Adress, error) {
	out := new(Adress)
	err := c.cc.Invoke(ctx, "/cliente.ChunksUpload/RequestAdress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunksUploadClient) DownloadChunk(ctx context.Context, in *Prop, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/cliente.ChunksUpload/DownloadChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunksUploadClient) PropuestaCentralizada(ctx context.Context, in *Adress, opts ...grpc.CallOption) (*Adress, error) {
	out := new(Adress)
	err := c.cc.Invoke(ctx, "/cliente.ChunksUpload/PropuestaCentralizada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunksUploadServer is the server API for ChunksUpload service.
// All implementations must embed UnimplementedChunksUploadServer
// for forward compatibility
type ChunksUploadServer interface {
	UploadChunk(ChunksUpload_UploadChunkServer) error
	Propuesta(context.Context, *Prop) (*Accept, error)
	SendChunk(ChunksUpload_SendChunkServer) error
	SendLog(context.Context, *Log) (*Accept, error)
	RequestAdress(context.Context, *Prop) (*Adress, error)
	DownloadChunk(context.Context, *Prop) (*Chunk, error)
	PropuestaCentralizada(context.Context, *Adress) (*Adress, error)
	mustEmbedUnimplementedChunksUploadServer()
}

// UnimplementedChunksUploadServer must be embedded to have forward compatible implementations.
type UnimplementedChunksUploadServer struct {
}

func (UnimplementedChunksUploadServer) UploadChunk(ChunksUpload_UploadChunkServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadChunk not implemented")
}
func (UnimplementedChunksUploadServer) Propuesta(context.Context, *Prop) (*Accept, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Propuesta not implemented")
}
func (UnimplementedChunksUploadServer) SendChunk(ChunksUpload_SendChunkServer) error {
	return status.Errorf(codes.Unimplemented, "method SendChunk not implemented")
}
func (UnimplementedChunksUploadServer) SendLog(context.Context, *Log) (*Accept, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLog not implemented")
}
func (UnimplementedChunksUploadServer) RequestAdress(context.Context, *Prop) (*Adress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAdress not implemented")
}
func (UnimplementedChunksUploadServer) DownloadChunk(context.Context, *Prop) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadChunk not implemented")
}
func (UnimplementedChunksUploadServer) PropuestaCentralizada(context.Context, *Adress) (*Adress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropuestaCentralizada not implemented")
}
func (UnimplementedChunksUploadServer) mustEmbedUnimplementedChunksUploadServer() {}

// UnsafeChunksUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunksUploadServer will
// result in compilation errors.
type UnsafeChunksUploadServer interface {
	mustEmbedUnimplementedChunksUploadServer()
}

func RegisterChunksUploadServer(s grpc.ServiceRegistrar, srv ChunksUploadServer) {
	s.RegisterService(&ChunksUpload_ServiceDesc, srv)
}

func _ChunksUpload_UploadChunk_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChunksUploadServer).UploadChunk(&chunksUploadUploadChunkServer{stream})
}

type ChunksUpload_UploadChunkServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type chunksUploadUploadChunkServer struct {
	grpc.ServerStream
}

func (x *chunksUploadUploadChunkServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chunksUploadUploadChunkServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChunksUpload_Propuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunksUploadServer).Propuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.ChunksUpload/Propuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunksUploadServer).Propuesta(ctx, req.(*Prop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunksUpload_SendChunk_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChunksUploadServer).SendChunk(&chunksUploadSendChunkServer{stream})
}

type ChunksUpload_SendChunkServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type chunksUploadSendChunkServer struct {
	grpc.ServerStream
}

func (x *chunksUploadSendChunkServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chunksUploadSendChunkServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChunksUpload_SendLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunksUploadServer).SendLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.ChunksUpload/SendLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunksUploadServer).SendLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunksUpload_RequestAdress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunksUploadServer).RequestAdress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.ChunksUpload/RequestAdress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunksUploadServer).RequestAdress(ctx, req.(*Prop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunksUpload_DownloadChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunksUploadServer).DownloadChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.ChunksUpload/DownloadChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunksUploadServer).DownloadChunk(ctx, req.(*Prop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunksUpload_PropuestaCentralizada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Adress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunksUploadServer).PropuestaCentralizada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.ChunksUpload/PropuestaCentralizada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunksUploadServer).PropuestaCentralizada(ctx, req.(*Adress))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunksUpload_ServiceDesc is the grpc.ServiceDesc for ChunksUpload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunksUpload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cliente.ChunksUpload",
	HandlerType: (*ChunksUploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Propuesta",
			Handler:    _ChunksUpload_Propuesta_Handler,
		},
		{
			MethodName: "SendLog",
			Handler:    _ChunksUpload_SendLog_Handler,
		},
		{
			MethodName: "RequestAdress",
			Handler:    _ChunksUpload_RequestAdress_Handler,
		},
		{
			MethodName: "DownloadChunk",
			Handler:    _ChunksUpload_DownloadChunk_Handler,
		},
		{
			MethodName: "PropuestaCentralizada",
			Handler:    _ChunksUpload_PropuestaCentralizada_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadChunk",
			Handler:       _ChunksUpload_UploadChunk_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendChunk",
			Handler:       _ChunksUpload_SendChunk_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "uploader.proto",
}
