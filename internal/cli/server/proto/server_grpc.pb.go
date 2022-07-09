// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: command/server/proto/server.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetraClient is the client API for Metra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetraClient interface {
	Pprof(ctx context.Context, in *PprofRequest, opts ...grpc.CallOption) (Metra_PprofClient, error)
	PeersAdd(ctx context.Context, in *PeersAddRequest, opts ...grpc.CallOption) (*PeersAddResponse, error)
	PeersRemove(ctx context.Context, in *PeersRemoveRequest, opts ...grpc.CallOption) (*PeersRemoveResponse, error)
	PeersList(ctx context.Context, in *PeersListRequest, opts ...grpc.CallOption) (*PeersListResponse, error)
	PeersStatus(ctx context.Context, in *PeersStatusRequest, opts ...grpc.CallOption) (*PeersStatusResponse, error)
	ChainSetHead(ctx context.Context, in *ChainSetHeadRequest, opts ...grpc.CallOption) (*ChainSetHeadResponse, error)
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	ChainWatch(ctx context.Context, in *ChainWatchRequest, opts ...grpc.CallOption) (Metra_ChainWatchClient, error)
}

type metraClient struct {
	cc grpc.ClientConnInterface
}

func NewMetraClient(cc grpc.ClientConnInterface) MetraClient {
	return &metraClient{cc}
}

func (c *metraClient) Pprof(ctx context.Context, in *PprofRequest, opts ...grpc.CallOption) (Metra_PprofClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metra_ServiceDesc.Streams[0], "/proto.Metra/Pprof", opts...)
	if err != nil {
		return nil, err
	}
	x := &metraPprofClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metra_PprofClient interface {
	Recv() (*PprofResponse, error)
	grpc.ClientStream
}

type metraPprofClient struct {
	grpc.ClientStream
}

func (x *metraPprofClient) Recv() (*PprofResponse, error) {
	m := new(PprofResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metraClient) PeersAdd(ctx context.Context, in *PeersAddRequest, opts ...grpc.CallOption) (*PeersAddResponse, error) {
	out := new(PeersAddResponse)
	err := c.cc.Invoke(ctx, "/proto.Metra/PeersAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metraClient) PeersRemove(ctx context.Context, in *PeersRemoveRequest, opts ...grpc.CallOption) (*PeersRemoveResponse, error) {
	out := new(PeersRemoveResponse)
	err := c.cc.Invoke(ctx, "/proto.Metra/PeersRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metraClient) PeersList(ctx context.Context, in *PeersListRequest, opts ...grpc.CallOption) (*PeersListResponse, error) {
	out := new(PeersListResponse)
	err := c.cc.Invoke(ctx, "/proto.Metra/PeersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metraClient) PeersStatus(ctx context.Context, in *PeersStatusRequest, opts ...grpc.CallOption) (*PeersStatusResponse, error) {
	out := new(PeersStatusResponse)
	err := c.cc.Invoke(ctx, "/proto.Metra/PeersStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metraClient) ChainSetHead(ctx context.Context, in *ChainSetHeadRequest, opts ...grpc.CallOption) (*ChainSetHeadResponse, error) {
	out := new(ChainSetHeadResponse)
	err := c.cc.Invoke(ctx, "/proto.Metra/ChainSetHead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metraClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.Metra/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metraClient) ChainWatch(ctx context.Context, in *ChainWatchRequest, opts ...grpc.CallOption) (Metra_ChainWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metra_ServiceDesc.Streams[1], "/proto.Metra/ChainWatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &metraChainWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metra_ChainWatchClient interface {
	Recv() (*ChainWatchResponse, error)
	grpc.ClientStream
}

type metraChainWatchClient struct {
	grpc.ClientStream
}

func (x *metraChainWatchClient) Recv() (*ChainWatchResponse, error) {
	m := new(ChainWatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetraServer is the server API for Metra service.
// All implementations must embed UnimplementedMetraServer
// for forward compatibility
type MetraServer interface {
	Pprof(*PprofRequest, Metra_PprofServer) error
	PeersAdd(context.Context, *PeersAddRequest) (*PeersAddResponse, error)
	PeersRemove(context.Context, *PeersRemoveRequest) (*PeersRemoveResponse, error)
	PeersList(context.Context, *PeersListRequest) (*PeersListResponse, error)
	PeersStatus(context.Context, *PeersStatusRequest) (*PeersStatusResponse, error)
	ChainSetHead(context.Context, *ChainSetHeadRequest) (*ChainSetHeadResponse, error)
	Status(context.Context, *emptypb.Empty) (*StatusResponse, error)
	ChainWatch(*ChainWatchRequest, Metra_ChainWatchServer) error
	mustEmbedUnimplementedMetraServer()
}

// UnimplementedMetraServer must be embedded to have forward compatible implementations.
type UnimplementedMetraServer struct {
}

func (UnimplementedMetraServer) Pprof(*PprofRequest, Metra_PprofServer) error {
	return status.Errorf(codes.Unimplemented, "method Pprof not implemented")
}
func (UnimplementedMetraServer) PeersAdd(context.Context, *PeersAddRequest) (*PeersAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersAdd not implemented")
}
func (UnimplementedMetraServer) PeersRemove(context.Context, *PeersRemoveRequest) (*PeersRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersRemove not implemented")
}
func (UnimplementedMetraServer) PeersList(context.Context, *PeersListRequest) (*PeersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersList not implemented")
}
func (UnimplementedMetraServer) PeersStatus(context.Context, *PeersStatusRequest) (*PeersStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersStatus not implemented")
}
func (UnimplementedMetraServer) ChainSetHead(context.Context, *ChainSetHeadRequest) (*ChainSetHeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainSetHead not implemented")
}
func (UnimplementedMetraServer) Status(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedMetraServer) ChainWatch(*ChainWatchRequest, Metra_ChainWatchServer) error {
	return status.Errorf(codes.Unimplemented, "method ChainWatch not implemented")
}
func (UnimplementedMetraServer) mustEmbedUnimplementedMetraServer() {}

// UnsafeMetraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetraServer will
// result in compilation errors.
type UnsafeMetraServer interface {
	mustEmbedUnimplementedMetraServer()
}

func RegisterMetraServer(s grpc.ServiceRegistrar, srv MetraServer) {
	s.RegisterService(&Metra_ServiceDesc, srv)
}

func _Metra_Pprof_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PprofRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetraServer).Pprof(m, &metraPprofServer{stream})
}

type Metra_PprofServer interface {
	Send(*PprofResponse) error
	grpc.ServerStream
}

type metraPprofServer struct {
	grpc.ServerStream
}

func (x *metraPprofServer) Send(m *PprofResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Metra_PeersAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetraServer).PeersAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metra/PeersAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetraServer).PeersAdd(ctx, req.(*PeersAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metra_PeersRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetraServer).PeersRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metra/PeersRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetraServer).PeersRemove(ctx, req.(*PeersRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metra_PeersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetraServer).PeersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metra/PeersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetraServer).PeersList(ctx, req.(*PeersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metra_PeersStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetraServer).PeersStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metra/PeersStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetraServer).PeersStatus(ctx, req.(*PeersStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metra_ChainSetHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainSetHeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetraServer).ChainSetHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metra/ChainSetHead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetraServer).ChainSetHead(ctx, req.(*ChainSetHeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metra_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetraServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metra/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetraServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metra_ChainWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChainWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetraServer).ChainWatch(m, &metraChainWatchServer{stream})
}

type Metra_ChainWatchServer interface {
	Send(*ChainWatchResponse) error
	grpc.ServerStream
}

type metraChainWatchServer struct {
	grpc.ServerStream
}

func (x *metraChainWatchServer) Send(m *ChainWatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Metra_ServiceDesc is the grpc.ServiceDesc for Metra service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metra_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Metra",
	HandlerType: (*MetraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PeersAdd",
			Handler:    _Metra_PeersAdd_Handler,
		},
		{
			MethodName: "PeersRemove",
			Handler:    _Metra_PeersRemove_Handler,
		},
		{
			MethodName: "PeersList",
			Handler:    _Metra_PeersList_Handler,
		},
		{
			MethodName: "PeersStatus",
			Handler:    _Metra_PeersStatus_Handler,
		},
		{
			MethodName: "ChainSetHead",
			Handler:    _Metra_ChainSetHead_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Metra_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pprof",
			Handler:       _Metra_Pprof_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChainWatch",
			Handler:       _Metra_ChainWatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "command/server/proto/server.proto",
}
