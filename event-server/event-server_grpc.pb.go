// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: event-server.proto

package __

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
	EventServer_EventBus_FullMethodName = "/server.EventServer/EventBus"
)

// EventServerClient is the client API for EventServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServerClient interface {
	EventBus(ctx context.Context, in *EventBusRequest, opts ...grpc.CallOption) (*EventBusResponse, error)
}

type eventServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServerClient(cc grpc.ClientConnInterface) EventServerClient {
	return &eventServerClient{cc}
}

func (c *eventServerClient) EventBus(ctx context.Context, in *EventBusRequest, opts ...grpc.CallOption) (*EventBusResponse, error) {
	out := new(EventBusResponse)
	err := c.cc.Invoke(ctx, EventServer_EventBus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServerServer is the server API for EventServer service.
// All implementations must embed UnimplementedEventServerServer
// for forward compatibility
type EventServerServer interface {
	EventBus(context.Context, *EventBusRequest) (*EventBusResponse, error)
	mustEmbedUnimplementedEventServerServer()
}

// UnimplementedEventServerServer must be embedded to have forward compatible implementations.
type UnimplementedEventServerServer struct {
}

func (UnimplementedEventServerServer) EventBus(context.Context, *EventBusRequest) (*EventBusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventBus not implemented")
}
func (UnimplementedEventServerServer) mustEmbedUnimplementedEventServerServer() {}

// UnsafeEventServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServerServer will
// result in compilation errors.
type UnsafeEventServerServer interface {
	mustEmbedUnimplementedEventServerServer()
}

func RegisterEventServerServer(s grpc.ServiceRegistrar, srv EventServerServer) {
	s.RegisterService(&EventServer_ServiceDesc, srv)
}

func _EventServer_EventBus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventBusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServerServer).EventBus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventServer_EventBus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServerServer).EventBus(ctx, req.(*EventBusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventServer_ServiceDesc is the grpc.ServiceDesc for EventServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.EventServer",
	HandlerType: (*EventServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventBus",
			Handler:    _EventServer_EventBus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event-server.proto",
}