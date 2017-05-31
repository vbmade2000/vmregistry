// Code generated by protoc-gen-go.
// source: vmregistry.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	vmregistry.proto

It has these top-level messages:
	VM
	ListVMRequest
	ListVMReply
	FindRequest
*/
package api

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

type FindRequest_FindBy int32

const (
	FindRequest_UNSPECIFIED FindRequest_FindBy = 0
	FindRequest_IP          FindRequest_FindBy = 1
	FindRequest_MAC         FindRequest_FindBy = 2
)

var FindRequest_FindBy_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "IP",
	2: "MAC",
}
var FindRequest_FindBy_value = map[string]int32{
	"UNSPECIFIED": 0,
	"IP":          1,
	"MAC":         2,
}

func (x FindRequest_FindBy) String() string {
	return proto.EnumName(FindRequest_FindBy_name, int32(x))
}
func (FindRequest_FindBy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type VM struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mac  string `protobuf:"bytes,2,opt,name=mac" json:"mac,omitempty"`
	Ip   string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
}

func (m *VM) Reset()                    { *m = VM{} }
func (m *VM) String() string            { return proto.CompactTextString(m) }
func (*VM) ProtoMessage()               {}
func (*VM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VM) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VM) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *VM) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ListVMRequest struct {
}

func (m *ListVMRequest) Reset()                    { *m = ListVMRequest{} }
func (m *ListVMRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVMRequest) ProtoMessage()               {}
func (*ListVMRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListVMReply struct {
	Vms []*VM `protobuf:"bytes,1,rep,name=vms" json:"vms,omitempty"`
}

func (m *ListVMReply) Reset()                    { *m = ListVMReply{} }
func (m *ListVMReply) String() string            { return proto.CompactTextString(m) }
func (*ListVMReply) ProtoMessage()               {}
func (*ListVMReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListVMReply) GetVms() []*VM {
	if m != nil {
		return m.Vms
	}
	return nil
}

type FindRequest struct {
	FindBy FindRequest_FindBy `protobuf:"varint,1,opt,name=find_by,json=findBy,enum=api.FindRequest_FindBy" json:"find_by,omitempty"`
	Value  string             `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *FindRequest) Reset()                    { *m = FindRequest{} }
func (m *FindRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()               {}
func (*FindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FindRequest) GetFindBy() FindRequest_FindBy {
	if m != nil {
		return m.FindBy
	}
	return FindRequest_UNSPECIFIED
}

func (m *FindRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*VM)(nil), "api.VM")
	proto.RegisterType((*ListVMRequest)(nil), "api.ListVMRequest")
	proto.RegisterType((*ListVMReply)(nil), "api.ListVMReply")
	proto.RegisterType((*FindRequest)(nil), "api.FindRequest")
	proto.RegisterEnum("api.FindRequest_FindBy", FindRequest_FindBy_name, FindRequest_FindBy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VMRegistry service

type VMRegistryClient interface {
	List(ctx context.Context, in *ListVMRequest, opts ...grpc.CallOption) (*ListVMReply, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*VM, error)
}

type vMRegistryClient struct {
	cc *grpc.ClientConn
}

func NewVMRegistryClient(cc *grpc.ClientConn) VMRegistryClient {
	return &vMRegistryClient{cc}
}

func (c *vMRegistryClient) List(ctx context.Context, in *ListVMRequest, opts ...grpc.CallOption) (*ListVMReply, error) {
	out := new(ListVMReply)
	err := grpc.Invoke(ctx, "/api.VMRegistry/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMRegistryClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*VM, error) {
	out := new(VM)
	err := grpc.Invoke(ctx, "/api.VMRegistry/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VMRegistry service

type VMRegistryServer interface {
	List(context.Context, *ListVMRequest) (*ListVMReply, error)
	Find(context.Context, *FindRequest) (*VM, error)
}

func RegisterVMRegistryServer(s *grpc.Server, srv VMRegistryServer) {
	s.RegisterService(&_VMRegistry_serviceDesc, srv)
}

func _VMRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VMRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMRegistryServer).List(ctx, req.(*ListVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMRegistry_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMRegistryServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VMRegistry/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMRegistryServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VMRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.VMRegistry",
	HandlerType: (*VMRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _VMRegistry_List_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _VMRegistry_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vmregistry.proto",
}

func init() { proto.RegisterFile("vmregistry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd7, 0xa4, 0xb6, 0xf8, 0x8a, 0x5b, 0x78, 0x08, 0xd6, 0x9d, 0x46, 0xbc, 0x14, 0x0f,
	0x45, 0xea, 0xcd, 0x9b, 0xce, 0x0d, 0x0a, 0x56, 0x46, 0xc5, 0x5e, 0x25, 0x73, 0x9d, 0x04, 0xda,
	0x2e, 0xb6, 0x5d, 0xa1, 0x47, 0xff, 0x73, 0x49, 0xba, 0x81, 0xf3, 0xd4, 0xaf, 0x3f, 0xbe, 0xf0,
	0x7e, 0xef, 0x01, 0xeb, 0xca, 0x3a, 0xff, 0x92, 0x4d, 0x5b, 0xf7, 0xa1, 0xaa, 0x77, 0xed, 0x0e,
	0xa9, 0x50, 0x92, 0x3f, 0x00, 0xc9, 0x12, 0x44, 0xb0, 0x2b, 0x51, 0xe6, 0xbe, 0x35, 0xb3, 0x82,
	0xf3, 0xd4, 0x64, 0x64, 0x40, 0x4b, 0xf1, 0xe9, 0x13, 0x83, 0x74, 0xc4, 0x31, 0x10, 0xa9, 0x7c,
	0x6a, 0x00, 0x91, 0x8a, 0x4f, 0xe0, 0xe2, 0x45, 0x36, 0x6d, 0x96, 0xa4, 0xf9, 0xf7, 0x3e, 0x6f,
	0x5a, 0x1e, 0x80, 0x77, 0x04, 0xaa, 0xe8, 0xf1, 0x1a, 0x68, 0x57, 0x36, 0xbe, 0x35, 0xa3, 0x81,
	0x17, 0xb9, 0xa1, 0x50, 0x32, 0xcc, 0x92, 0x54, 0x33, 0xfe, 0x63, 0x81, 0xb7, 0x94, 0xd5, 0xe6,
	0xf0, 0x12, 0xef, 0xc0, 0xdd, 0xca, 0x6a, 0xf3, 0xb1, 0xee, 0x8d, 0xc3, 0x38, 0xba, 0x32, 0xf5,
	0x3f, 0x15, 0x93, 0x9f, 0xfa, 0xd4, 0xd9, 0x9a, 0x2f, 0x5e, 0xc2, 0x59, 0x27, 0x8a, 0x7d, 0x7e,
	0x10, 0x1c, 0x7e, 0xf8, 0x2d, 0x38, 0x43, 0x0f, 0x27, 0xe0, 0xbd, 0xbf, 0xbe, 0xad, 0x16, 0xf3,
	0x78, 0x19, 0x2f, 0x9e, 0xd9, 0x08, 0x1d, 0x20, 0xf1, 0x8a, 0x59, 0xe8, 0x02, 0x4d, 0x1e, 0xe7,
	0x8c, 0x44, 0x02, 0x40, 0x9b, 0x0e, 0x37, 0xc1, 0x10, 0x6c, 0xed, 0x8e, 0x68, 0x06, 0x9f, 0xec,
	0x35, 0x65, 0x27, 0x4c, 0x15, 0x3d, 0x1f, 0xe1, 0x0d, 0xd8, 0x7a, 0x12, 0xb2, 0xff, 0xa2, 0xd3,
	0xe3, 0xa6, 0x7c, 0xb4, 0x76, 0xcc, 0xa5, 0xef, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xeb,
	0xc8, 0x93, 0x7d, 0x01, 0x00, 0x00,
}