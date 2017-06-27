// Code generated by protoc-gen-go.
// source: whatever.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	whatever.proto

It has these top-level messages:
	HiRequest
	HiReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HiRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *HiRequest) Reset()                    { *m = HiRequest{} }
func (m *HiRequest) String() string            { return proto.CompactTextString(m) }
func (*HiRequest) ProtoMessage()               {}
func (*HiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HiRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HiReply struct {
	V string `protobuf:"bytes,1,opt,name=v" json:"v,omitempty"`
}

func (m *HiReply) Reset()                    { *m = HiReply{} }
func (m *HiReply) String() string            { return proto.CompactTextString(m) }
func (*HiReply) ProtoMessage()               {}
func (*HiReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HiReply) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func init() {
	proto.RegisterType((*HiRequest)(nil), "pb.HiRequest")
	proto.RegisterType((*HiReply)(nil), "pb.HiReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	Hi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Hi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error) {
	out := new(HiReply)
	err := grpc.Invoke(ctx, "/pb.Hello/Hi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	Hi(context.Context, *HiRequest) (*HiReply, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Hi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Hi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hello/Hi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Hi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hi",
			Handler:    _Hello_Hi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "whatever.proto",
}

func init() { proto.RegisterFile("whatever.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xcf, 0x48, 0x2c,
	0x49, 0x2d, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92,
	0xe7, 0xe2, 0xf4, 0xc8, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xf1,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xc4, 0xb9, 0xd8,
	0x41, 0x0a, 0x0a, 0x72, 0x2a, 0x85, 0x78, 0xb8, 0x18, 0xcb, 0xa0, 0x72, 0x8c, 0x65, 0x46, 0xda,
	0x5c, 0xac, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x42, 0x4a, 0x5c, 0x4c, 0x1e, 0x99, 0x42, 0xbc, 0x7a,
	0x05, 0x49, 0x7a, 0x70, 0xa3, 0xa4, 0xb8, 0x61, 0xdc, 0x82, 0x9c, 0x4a, 0x25, 0x86, 0x24, 0x36,
	0xb0, 0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x20, 0x1a, 0xfe, 0x83, 0x00, 0x00,
	0x00,
}