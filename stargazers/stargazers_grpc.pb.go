// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stargazers

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StargazersClient is the client API for Stargazers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StargazersClient interface {
	GetStargazers(ctx context.Context, in *StargazerRequest, opts ...grpc.CallOption) (*StargazersResponse, error)
}

type stargazersClient struct {
	cc grpc.ClientConnInterface
}

func NewStargazersClient(cc grpc.ClientConnInterface) StargazersClient {
	return &stargazersClient{cc}
}

func (c *stargazersClient) GetStargazers(ctx context.Context, in *StargazerRequest, opts ...grpc.CallOption) (*StargazersResponse, error) {
	out := new(StargazersResponse)
	err := c.cc.Invoke(ctx, "/stargazers.Stargazers/GetStargazers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StargazersServer is the server API for Stargazers service.
// All implementations must embed UnimplementedStargazersServer
// for forward compatibility
type StargazersServer interface {
	GetStargazers(context.Context, *StargazerRequest) (*StargazersResponse, error)
	mustEmbedUnimplementedStargazersServer()
}

// UnimplementedStargazersServer must be embedded to have forward compatible implementations.
type UnimplementedStargazersServer struct {
}

func (UnimplementedStargazersServer) GetStargazers(context.Context, *StargazerRequest) (*StargazersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStargazers not implemented")
}
func (UnimplementedStargazersServer) mustEmbedUnimplementedStargazersServer() {}

// UnsafeStargazersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StargazersServer will
// result in compilation errors.
type UnsafeStargazersServer interface {
	mustEmbedUnimplementedStargazersServer()
}

func RegisterStargazersServer(s *grpc.Server, srv StargazersServer) {
	s.RegisterService(&_Stargazers_serviceDesc, srv)
}

func _Stargazers_GetStargazers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StargazerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StargazersServer).GetStargazers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargazers.Stargazers/GetStargazers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StargazersServer).GetStargazers(ctx, req.(*StargazerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stargazers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stargazers.Stargazers",
	HandlerType: (*StargazersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStargazers",
			Handler:    _Stargazers_GetStargazers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargazers.proto",
}
