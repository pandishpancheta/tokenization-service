// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: pkg/pb/tokenization.proto

package tokenization

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

// TokenizationServiceClient is the client API for TokenizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenizationServiceClient interface {
	Tokenize(ctx context.Context, in *TokenizationRequest, opts ...grpc.CallOption) (*TokenizationResponse, error)
}

type tokenizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenizationServiceClient(cc grpc.ClientConnInterface) TokenizationServiceClient {
	return &tokenizationServiceClient{cc}
}

func (c *tokenizationServiceClient) Tokenize(ctx context.Context, in *TokenizationRequest, opts ...grpc.CallOption) (*TokenizationResponse, error) {
	out := new(TokenizationResponse)
	err := c.cc.Invoke(ctx, "/tokenization.TokenizationService/Tokenize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenizationServiceServer is the server API for TokenizationService service.
// All implementations must embed UnimplementedTokenizationServiceServer
// for forward compatibility
type TokenizationServiceServer interface {
	Tokenize(context.Context, *TokenizationRequest) (*TokenizationResponse, error)
	mustEmbedUnimplementedTokenizationServiceServer()
}

// UnimplementedTokenizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenizationServiceServer struct {
}

func (UnimplementedTokenizationServiceServer) Tokenize(context.Context, *TokenizationRequest) (*TokenizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tokenize not implemented")
}
func (UnimplementedTokenizationServiceServer) mustEmbedUnimplementedTokenizationServiceServer() {}

// UnsafeTokenizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenizationServiceServer will
// result in compilation errors.
type UnsafeTokenizationServiceServer interface {
	mustEmbedUnimplementedTokenizationServiceServer()
}

func RegisterTokenizationServiceServer(s grpc.ServiceRegistrar, srv TokenizationServiceServer) {
	s.RegisterService(&TokenizationService_ServiceDesc, srv)
}

func _TokenizationService_Tokenize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenizationServiceServer).Tokenize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenization.TokenizationService/Tokenize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenizationServiceServer).Tokenize(ctx, req.(*TokenizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenizationService_ServiceDesc is the grpc.ServiceDesc for TokenizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenization.TokenizationService",
	HandlerType: (*TokenizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tokenize",
			Handler:    _TokenizationService_Tokenize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/tokenization.proto",
}
