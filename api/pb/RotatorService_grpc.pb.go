// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: RotatorService.proto

package pb

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

// RotatorClient is the client API for Rotator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RotatorClient interface {
	// Добавить новый баннер в ротацию в слоте
	AddBannerToSlot(ctx context.Context, in *AddBannerToSlotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Удалить баннер из ротации в слоте
	RemoveBannerFromSlot(ctx context.Context, in *RemoveBannerFromSlotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Засчитать переход
	CountTransition(ctx context.Context, in *CountTransitionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Выбрать баннер для показа
	ChooseBanner(ctx context.Context, in *ChooseBannerRequest, opts ...grpc.CallOption) (*ChooseBannerResponse, error)
}

type rotatorClient struct {
	cc grpc.ClientConnInterface
}

func NewRotatorClient(cc grpc.ClientConnInterface) RotatorClient {
	return &rotatorClient{cc}
}

func (c *rotatorClient) AddBannerToSlot(ctx context.Context, in *AddBannerToSlotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/AddBannerToSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) RemoveBannerFromSlot(ctx context.Context, in *RemoveBannerFromSlotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/RemoveBannerFromSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) CountTransition(ctx context.Context, in *CountTransitionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/CountTransition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) ChooseBanner(ctx context.Context, in *ChooseBannerRequest, opts ...grpc.CallOption) (*ChooseBannerResponse, error) {
	out := new(ChooseBannerResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/ChooseBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RotatorServer is the server API for Rotator service.
// All implementations must embed UnimplementedRotatorServer
// for forward compatibility
type RotatorServer interface {
	// Добавить новый баннер в ротацию в слоте
	AddBannerToSlot(context.Context, *AddBannerToSlotRequest) (*emptypb.Empty, error)
	// Удалить баннер из ротации в слоте
	RemoveBannerFromSlot(context.Context, *RemoveBannerFromSlotRequest) (*emptypb.Empty, error)
	// Засчитать переход
	CountTransition(context.Context, *CountTransitionRequest) (*emptypb.Empty, error)
	// Выбрать баннер для показа
	ChooseBanner(context.Context, *ChooseBannerRequest) (*ChooseBannerResponse, error)
	mustEmbedUnimplementedRotatorServer()
}

// UnimplementedRotatorServer must be embedded to have forward compatible implementations.
type UnimplementedRotatorServer struct {
}

func (UnimplementedRotatorServer) AddBannerToSlot(context.Context, *AddBannerToSlotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBannerToSlot not implemented")
}
func (UnimplementedRotatorServer) RemoveBannerFromSlot(context.Context, *RemoveBannerFromSlotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBannerFromSlot not implemented")
}
func (UnimplementedRotatorServer) CountTransition(context.Context, *CountTransitionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTransition not implemented")
}
func (UnimplementedRotatorServer) ChooseBanner(context.Context, *ChooseBannerRequest) (*ChooseBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChooseBanner not implemented")
}
func (UnimplementedRotatorServer) mustEmbedUnimplementedRotatorServer() {}

// UnsafeRotatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RotatorServer will
// result in compilation errors.
type UnsafeRotatorServer interface {
	mustEmbedUnimplementedRotatorServer()
}

func RegisterRotatorServer(s grpc.ServiceRegistrar, srv RotatorServer) {
	s.RegisterService(&Rotator_ServiceDesc, srv)
}

func _Rotator_AddBannerToSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBannerToSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).AddBannerToSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/AddBannerToSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).AddBannerToSlot(ctx, req.(*AddBannerToSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_RemoveBannerFromSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBannerFromSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).RemoveBannerFromSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/RemoveBannerFromSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).RemoveBannerFromSlot(ctx, req.(*RemoveBannerFromSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_CountTransition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTransitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).CountTransition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/CountTransition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).CountTransition(ctx, req.(*CountTransitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_ChooseBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChooseBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).ChooseBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/ChooseBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).ChooseBanner(ctx, req.(*ChooseBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rotator_ServiceDesc is the grpc.ServiceDesc for Rotator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rotator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rotator.Rotator",
	HandlerType: (*RotatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBannerToSlot",
			Handler:    _Rotator_AddBannerToSlot_Handler,
		},
		{
			MethodName: "RemoveBannerFromSlot",
			Handler:    _Rotator_RemoveBannerFromSlot_Handler,
		},
		{
			MethodName: "CountTransition",
			Handler:    _Rotator_CountTransition_Handler,
		},
		{
			MethodName: "ChooseBanner",
			Handler:    _Rotator_ChooseBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RotatorService.proto",
}