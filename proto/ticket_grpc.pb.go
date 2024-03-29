// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0--rc1
// source: ticket.proto

package proto

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

// TrainBookingClient is the client API for TrainBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainBookingClient interface {
	BookTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	GetReceipt(ctx context.Context, in *TicketReceipt, opts ...grpc.CallOption) (*TicketDetails, error)
	DeleteTicket(ctx context.Context, in *TicketReceipt, opts ...grpc.CallOption) (*TicketStatus, error)
	GetAllUsers(ctx context.Context, in *TicketSection, opts ...grpc.CallOption) (TrainBooking_GetAllUsersClient, error)
	ModifySeat(ctx context.Context, in *TicketReceipt, opts ...grpc.CallOption) (*TicketResponse, error)
}

type trainBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainBookingClient(cc grpc.ClientConnInterface) TrainBookingClient {
	return &trainBookingClient{cc}
}

func (c *trainBookingClient) BookTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, "/train_ticket.TrainBooking/BookTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainBookingClient) GetReceipt(ctx context.Context, in *TicketReceipt, opts ...grpc.CallOption) (*TicketDetails, error) {
	out := new(TicketDetails)
	err := c.cc.Invoke(ctx, "/train_ticket.TrainBooking/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainBookingClient) DeleteTicket(ctx context.Context, in *TicketReceipt, opts ...grpc.CallOption) (*TicketStatus, error) {
	out := new(TicketStatus)
	err := c.cc.Invoke(ctx, "/train_ticket.TrainBooking/DeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainBookingClient) GetAllUsers(ctx context.Context, in *TicketSection, opts ...grpc.CallOption) (TrainBooking_GetAllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrainBooking_ServiceDesc.Streams[0], "/train_ticket.TrainBooking/GetAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &trainBookingGetAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TrainBooking_GetAllUsersClient interface {
	Recv() (*TicketDetails, error)
	grpc.ClientStream
}

type trainBookingGetAllUsersClient struct {
	grpc.ClientStream
}

func (x *trainBookingGetAllUsersClient) Recv() (*TicketDetails, error) {
	m := new(TicketDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trainBookingClient) ModifySeat(ctx context.Context, in *TicketReceipt, opts ...grpc.CallOption) (*TicketResponse, error) {
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, "/train_ticket.TrainBooking/ModifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainBookingServer is the server API for TrainBooking service.
// All implementations must embed UnimplementedTrainBookingServer
// for forward compatibility
type TrainBookingServer interface {
	BookTicket(context.Context, *TicketRequest) (*TicketResponse, error)
	GetReceipt(context.Context, *TicketReceipt) (*TicketDetails, error)
	DeleteTicket(context.Context, *TicketReceipt) (*TicketStatus, error)
	GetAllUsers(*TicketSection, TrainBooking_GetAllUsersServer) error
	ModifySeat(context.Context, *TicketReceipt) (*TicketResponse, error)
	mustEmbedUnimplementedTrainBookingServer()
}

// UnimplementedTrainBookingServer must be embedded to have forward compatible implementations.
type UnimplementedTrainBookingServer struct {
}

func (UnimplementedTrainBookingServer) BookTicket(context.Context, *TicketRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}
func (UnimplementedTrainBookingServer) GetReceipt(context.Context, *TicketReceipt) (*TicketDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTrainBookingServer) DeleteTicket(context.Context, *TicketReceipt) (*TicketStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTrainBookingServer) GetAllUsers(*TicketSection, TrainBooking_GetAllUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedTrainBookingServer) ModifySeat(context.Context, *TicketReceipt) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainBookingServer) mustEmbedUnimplementedTrainBookingServer() {}

// UnsafeTrainBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainBookingServer will
// result in compilation errors.
type UnsafeTrainBookingServer interface {
	mustEmbedUnimplementedTrainBookingServer()
}

func RegisterTrainBookingServer(s grpc.ServiceRegistrar, srv TrainBookingServer) {
	s.RegisterService(&TrainBooking_ServiceDesc, srv)
}

func _TrainBooking_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainBookingServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train_ticket.TrainBooking/BookTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainBookingServer).BookTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainBooking_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketReceipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainBookingServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train_ticket.TrainBooking/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainBookingServer).GetReceipt(ctx, req.(*TicketReceipt))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainBooking_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketReceipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainBookingServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train_ticket.TrainBooking/DeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainBookingServer).DeleteTicket(ctx, req.(*TicketReceipt))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainBooking_GetAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TicketSection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrainBookingServer).GetAllUsers(m, &trainBookingGetAllUsersServer{stream})
}

type TrainBooking_GetAllUsersServer interface {
	Send(*TicketDetails) error
	grpc.ServerStream
}

type trainBookingGetAllUsersServer struct {
	grpc.ServerStream
}

func (x *trainBookingGetAllUsersServer) Send(m *TicketDetails) error {
	return x.ServerStream.SendMsg(m)
}

func _TrainBooking_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketReceipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainBookingServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/train_ticket.TrainBooking/ModifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainBookingServer).ModifySeat(ctx, req.(*TicketReceipt))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainBooking_ServiceDesc is the grpc.ServiceDesc for TrainBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "train_ticket.TrainBooking",
	HandlerType: (*TrainBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookTicket",
			Handler:    _TrainBooking_BookTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TrainBooking_GetReceipt_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TrainBooking_DeleteTicket_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainBooking_ModifySeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllUsers",
			Handler:       _TrainBooking_GetAllUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ticket.proto",
}
