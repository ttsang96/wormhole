// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/node/v1/node.proto

package nodev1

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

// NodePrivilegedServiceClient is the client API for NodePrivilegedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodePrivilegedServiceClient interface {
	// InjectGovernanceVAA injects a governance VAA into the guardian node.
	// The node will inject the VAA into the aggregator and sign/broadcast the VAA signature.
	//
	// A consensus majority of nodes on the network will have to inject the VAA within the
	// VAA timeout window for it to reach consensus.
	//
	InjectGovernanceVAA(ctx context.Context, in *InjectGovernanceVAARequest, opts ...grpc.CallOption) (*InjectGovernanceVAAResponse, error)
	// SendObservationRequest broadcasts a signed observation request to the gossip network
	// using the node's guardian key. The network rate limits these requests to one per second.
	// Requests at higher rates will fail silently.
	SendObservationRequest(ctx context.Context, in *SendObservationRequestRequest, opts ...grpc.CallOption) (*SendObservationRequestResponse, error)
}

type nodePrivilegedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodePrivilegedServiceClient(cc grpc.ClientConnInterface) NodePrivilegedServiceClient {
	return &nodePrivilegedServiceClient{cc}
}

func (c *nodePrivilegedServiceClient) InjectGovernanceVAA(ctx context.Context, in *InjectGovernanceVAARequest, opts ...grpc.CallOption) (*InjectGovernanceVAAResponse, error) {
	out := new(InjectGovernanceVAAResponse)
	err := c.cc.Invoke(ctx, "/node.v1.NodePrivilegedService/InjectGovernanceVAA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodePrivilegedServiceClient) SendObservationRequest(ctx context.Context, in *SendObservationRequestRequest, opts ...grpc.CallOption) (*SendObservationRequestResponse, error) {
	out := new(SendObservationRequestResponse)
	err := c.cc.Invoke(ctx, "/node.v1.NodePrivilegedService/SendObservationRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodePrivilegedServiceServer is the server API for NodePrivilegedService service.
// All implementations must embed UnimplementedNodePrivilegedServiceServer
// for forward compatibility
type NodePrivilegedServiceServer interface {
	// InjectGovernanceVAA injects a governance VAA into the guardian node.
	// The node will inject the VAA into the aggregator and sign/broadcast the VAA signature.
	//
	// A consensus majority of nodes on the network will have to inject the VAA within the
	// VAA timeout window for it to reach consensus.
	//
	InjectGovernanceVAA(context.Context, *InjectGovernanceVAARequest) (*InjectGovernanceVAAResponse, error)
	// SendObservationRequest broadcasts a signed observation request to the gossip network
	// using the node's guardian key. The network rate limits these requests to one per second.
	// Requests at higher rates will fail silently.
	SendObservationRequest(context.Context, *SendObservationRequestRequest) (*SendObservationRequestResponse, error)
	mustEmbedUnimplementedNodePrivilegedServiceServer()
}

// UnimplementedNodePrivilegedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodePrivilegedServiceServer struct {
}

func (UnimplementedNodePrivilegedServiceServer) InjectGovernanceVAA(context.Context, *InjectGovernanceVAARequest) (*InjectGovernanceVAAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InjectGovernanceVAA not implemented")
}
func (UnimplementedNodePrivilegedServiceServer) SendObservationRequest(context.Context, *SendObservationRequestRequest) (*SendObservationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendObservationRequest not implemented")
}
func (UnimplementedNodePrivilegedServiceServer) mustEmbedUnimplementedNodePrivilegedServiceServer() {}

// UnsafeNodePrivilegedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodePrivilegedServiceServer will
// result in compilation errors.
type UnsafeNodePrivilegedServiceServer interface {
	mustEmbedUnimplementedNodePrivilegedServiceServer()
}

func RegisterNodePrivilegedServiceServer(s grpc.ServiceRegistrar, srv NodePrivilegedServiceServer) {
	s.RegisterService(&NodePrivilegedService_ServiceDesc, srv)
}

func _NodePrivilegedService_InjectGovernanceVAA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InjectGovernanceVAARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePrivilegedServiceServer).InjectGovernanceVAA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.v1.NodePrivilegedService/InjectGovernanceVAA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePrivilegedServiceServer).InjectGovernanceVAA(ctx, req.(*InjectGovernanceVAARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodePrivilegedService_SendObservationRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendObservationRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodePrivilegedServiceServer).SendObservationRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.v1.NodePrivilegedService/SendObservationRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodePrivilegedServiceServer).SendObservationRequest(ctx, req.(*SendObservationRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodePrivilegedService_ServiceDesc is the grpc.ServiceDesc for NodePrivilegedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodePrivilegedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.v1.NodePrivilegedService",
	HandlerType: (*NodePrivilegedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InjectGovernanceVAA",
			Handler:    _NodePrivilegedService_InjectGovernanceVAA_Handler,
		},
		{
			MethodName: "SendObservationRequest",
			Handler:    _NodePrivilegedService_SendObservationRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/node/v1/node.proto",
}
