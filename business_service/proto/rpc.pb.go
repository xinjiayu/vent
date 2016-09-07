// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Register service

type RegisterClient interface {
	Register(ctx context.Context, in *RegisterC2S, opts ...grpc.CallOption) (*RegisterS2C, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *RegisterC2S, opts ...grpc.CallOption) (*RegisterS2C, error) {
	out := new(RegisterS2C)
	err := grpc.Invoke(ctx, "/proto.Register/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Register service

type RegisterServer interface {
	Register(context.Context, *RegisterC2S) (*RegisterS2C, error)
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Register/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*RegisterC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for Login service

type LoginClient interface {
	Login(ctx context.Context, in *LoginC2S, opts ...grpc.CallOption) (*LoginS2C, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginC2S, opts ...grpc.CallOption) (*LoginS2C, error) {
	out := new(LoginS2C)
	err := grpc.Invoke(ctx, "/proto.Login/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginServer interface {
	Login(context.Context, *LoginC2S) (*LoginS2C, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto1.RegisterFile("rpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x3c, 0xc9, 0x39, 0x99, 0xa9, 0x79,
	0x25, 0x10, 0x41, 0x23, 0x07, 0x2e, 0x8e, 0xa0, 0xd4, 0xf4, 0xcc, 0xe2, 0x92, 0xd4, 0x22, 0x21,
	0x13, 0x24, 0xb6, 0x10, 0x44, 0x5e, 0x0f, 0x26, 0xe0, 0x6c, 0x54, 0x2c, 0x85, 0x2e, 0x16, 0x6c,
	0x94, 0xac, 0xc4, 0x60, 0x64, 0xc2, 0xc5, 0xea, 0x93, 0x9f, 0x9e, 0x99, 0x27, 0xa4, 0x0d, 0x63,
	0xf0, 0x43, 0xd5, 0x81, 0x79, 0x20, 0x8d, 0x28, 0x02, 0x60, 0x5d, 0x49, 0x6c, 0x60, 0x11, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe0, 0x3d, 0x42, 0xa0, 0x00, 0x00, 0x00,
}