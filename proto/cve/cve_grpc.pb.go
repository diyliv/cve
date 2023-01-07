// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/cve/cve.proto

package cvepb

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

// CveServiceClient is the client API for CveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CveServiceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	FindVulnerabilities(ctx context.Context, in *FindVulnerabilitiesReq, opts ...grpc.CallOption) (*FindVulnerabilitiesResp, error)
}

type cveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCveServiceClient(cc grpc.ClientConnInterface) CveServiceClient {
	return &cveServiceClient{cc}
}

func (c *cveServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/cve.CveService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cveServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/cve.CveService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cveServiceClient) FindVulnerabilities(ctx context.Context, in *FindVulnerabilitiesReq, opts ...grpc.CallOption) (*FindVulnerabilitiesResp, error) {
	out := new(FindVulnerabilitiesResp)
	err := c.cc.Invoke(ctx, "/cve.CveService/FindVulnerabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CveServiceServer is the server API for CveService service.
// All implementations must embed UnimplementedCveServiceServer
// for forward compatibility
type CveServiceServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	FindVulnerabilities(context.Context, *FindVulnerabilitiesReq) (*FindVulnerabilitiesResp, error)
	mustEmbedUnimplementedCveServiceServer()
}

// UnimplementedCveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCveServiceServer struct {
}

func (UnimplementedCveServiceServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCveServiceServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCveServiceServer) FindVulnerabilities(context.Context, *FindVulnerabilitiesReq) (*FindVulnerabilitiesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVulnerabilities not implemented")
}
func (UnimplementedCveServiceServer) mustEmbedUnimplementedCveServiceServer() {}

// UnsafeCveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CveServiceServer will
// result in compilation errors.
type UnsafeCveServiceServer interface {
	mustEmbedUnimplementedCveServiceServer()
}

func RegisterCveServiceServer(s grpc.ServiceRegistrar, srv CveServiceServer) {
	s.RegisterService(&CveService_ServiceDesc, srv)
}

func _CveService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CveServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cve.CveService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CveServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CveService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CveServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cve.CveService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CveServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CveService_FindVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVulnerabilitiesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CveServiceServer).FindVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cve.CveService/FindVulnerabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CveServiceServer).FindVulnerabilities(ctx, req.(*FindVulnerabilitiesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CveService_ServiceDesc is the grpc.ServiceDesc for CveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cve.CveService",
	HandlerType: (*CveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _CveService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CveService_Login_Handler,
		},
		{
			MethodName: "FindVulnerabilities",
			Handler:    _CveService_FindVulnerabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cve/cve.proto",
}