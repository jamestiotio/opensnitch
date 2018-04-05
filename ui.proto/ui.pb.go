// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ui.proto

/*
Package ui is a generated protocol buffer package.

It is generated from these files:
	ui.proto

It has these top-level messages:
	PingRequest
	PingReply
	RuleRequest
	RuleReply
*/
package ui

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

type PingRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PingReply struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RuleRequest struct {
	Protocol    string   `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	SrcIp       string   `protobuf:"bytes,2,opt,name=src_ip,json=srcIp" json:"src_ip,omitempty"`
	SrcPort     uint32   `protobuf:"varint,3,opt,name=src_port,json=srcPort" json:"src_port,omitempty"`
	DstIp       string   `protobuf:"bytes,4,opt,name=dst_ip,json=dstIp" json:"dst_ip,omitempty"`
	DstHost     string   `protobuf:"bytes,5,opt,name=dst_host,json=dstHost" json:"dst_host,omitempty"`
	DstPort     uint32   `protobuf:"varint,6,opt,name=dst_port,json=dstPort" json:"dst_port,omitempty"`
	UserId      uint32   `protobuf:"varint,7,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ProcessId   uint32   `protobuf:"varint,8,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	ProcessPath string   `protobuf:"bytes,9,opt,name=process_path,json=processPath" json:"process_path,omitempty"`
	ProcessArgs []string `protobuf:"bytes,10,rep,name=process_args,json=processArgs" json:"process_args,omitempty"`
}

func (m *RuleRequest) Reset()                    { *m = RuleRequest{} }
func (m *RuleRequest) String() string            { return proto.CompactTextString(m) }
func (*RuleRequest) ProtoMessage()               {}
func (*RuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RuleRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *RuleRequest) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

func (m *RuleRequest) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *RuleRequest) GetDstIp() string {
	if m != nil {
		return m.DstIp
	}
	return ""
}

func (m *RuleRequest) GetDstHost() string {
	if m != nil {
		return m.DstHost
	}
	return ""
}

func (m *RuleRequest) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *RuleRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RuleRequest) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *RuleRequest) GetProcessPath() string {
	if m != nil {
		return m.ProcessPath
	}
	return ""
}

func (m *RuleRequest) GetProcessArgs() []string {
	if m != nil {
		return m.ProcessArgs
	}
	return nil
}

type RuleReply struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Duration string `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	What     string `protobuf:"bytes,4,opt,name=what" json:"what,omitempty"`
	Value    string `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
}

func (m *RuleReply) Reset()                    { *m = RuleReply{} }
func (m *RuleReply) String() string            { return proto.CompactTextString(m) }
func (*RuleReply) ProtoMessage()               {}
func (*RuleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RuleReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RuleReply) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RuleReply) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *RuleReply) GetWhat() string {
	if m != nil {
		return m.What
	}
	return ""
}

func (m *RuleReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "ui.PingRequest")
	proto.RegisterType((*PingReply)(nil), "ui.PingReply")
	proto.RegisterType((*RuleRequest)(nil), "ui.RuleRequest")
	proto.RegisterType((*RuleReply)(nil), "ui.RuleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UI service

type UIClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	AskRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleReply, error)
}

type uIClient struct {
	cc *grpc.ClientConn
}

func NewUIClient(cc *grpc.ClientConn) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/ui.UI/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) AskRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleReply, error) {
	out := new(RuleReply)
	err := grpc.Invoke(ctx, "/ui.UI/AskRule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UI service

type UIServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	AskRule(context.Context, *RuleRequest) (*RuleReply, error)
}

func RegisterUIServer(s *grpc.Server, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.UI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_AskRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).AskRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.UI/AskRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).AskRule(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ui.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UI_Ping_Handler,
		},
		{
			MethodName: "AskRule",
			Handler:    _UI_AskRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ui.proto",
}

func init() { proto.RegisterFile("ui.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4f, 0xe3, 0x30,
	0x10, 0xc5, 0xdb, 0x34, 0xcd, 0x9f, 0xe9, 0x76, 0x57, 0xb2, 0x76, 0x97, 0x50, 0x54, 0xa9, 0xe4,
	0x54, 0x09, 0xa9, 0x07, 0xf8, 0x04, 0xbd, 0x91, 0x5b, 0x15, 0x89, 0x13, 0x87, 0xca, 0xc4, 0x51,
	0x63, 0x11, 0x6a, 0xe3, 0xb1, 0x41, 0x3d, 0xf0, 0x75, 0xf8, 0x9c, 0xc8, 0x76, 0x52, 0x22, 0x71,
	0xf3, 0xcc, 0xef, 0xbd, 0xd7, 0xe6, 0x0d, 0x24, 0x86, 0x6f, 0xa4, 0x12, 0x5a, 0x90, 0xc0, 0xf0,
	0x7c, 0x09, 0xb3, 0x1d, 0x3f, 0x1e, 0xca, 0xfa, 0xd5, 0xd4, 0xa8, 0xc9, 0x6f, 0x08, 0x38, 0xcb,
	0xc6, 0xab, 0xf1, 0x3a, 0x2c, 0x03, 0xce, 0xf2, 0x2b, 0x48, 0x3d, 0x96, 0xed, 0xe9, 0x07, 0xfc,
	0x0c, 0x60, 0x56, 0x9a, 0xb6, 0xee, 0xcd, 0x0b, 0x48, 0x5c, 0x70, 0x25, 0x5a, 0xa7, 0x4a, 0xcb,
	0xf3, 0x4c, 0xfe, 0x41, 0x84, 0xaa, 0xda, 0x73, 0x99, 0x05, 0x8e, 0x4c, 0x51, 0x55, 0x85, 0x24,
	0x97, 0x90, 0xd8, 0xb5, 0x14, 0x4a, 0x67, 0x93, 0xd5, 0x78, 0x3d, 0x2f, 0x63, 0x54, 0xd5, 0x4e,
	0x28, 0x6d, 0x1d, 0x0c, 0xb5, 0x75, 0x84, 0xde, 0xc1, 0x50, 0x7b, 0x87, 0x5d, 0x37, 0x02, 0x75,
	0x36, 0x75, 0x20, 0x66, 0xa8, 0xef, 0x05, 0xea, 0x1e, 0xb9, 0xb0, 0xc8, 0x87, 0x31, 0xd4, 0x2e,
	0xec, 0x02, 0x62, 0x83, 0xb5, 0xda, 0x73, 0x96, 0xc5, 0x8e, 0x44, 0x76, 0x2c, 0x18, 0x59, 0x02,
	0x48, 0x25, 0xaa, 0x1a, 0xd1, 0xb2, 0xc4, 0xb1, 0xb4, 0xdb, 0x14, 0x8c, 0x5c, 0xc3, 0xaf, 0x1e,
	0x4b, 0xaa, 0x9b, 0x2c, 0x75, 0xbf, 0x38, 0xeb, 0x76, 0x3b, 0xaa, 0x9b, 0xa1, 0x84, 0xaa, 0x03,
	0x66, 0xb0, 0x9a, 0x0c, 0x24, 0x5b, 0x75, 0xc0, 0xfc, 0x03, 0x52, 0xdf, 0x93, 0x6d, 0x91, 0x40,
	0x78, 0xa4, 0x2f, 0x75, 0xd7, 0x90, 0x7b, 0x93, 0xff, 0x10, 0xd1, 0x4a, 0x73, 0x71, 0xec, 0xda,
	0xe9, 0x26, 0xdb, 0x28, 0x33, 0x8a, 0x3a, 0x32, 0xf1, 0x8d, 0xf6, 0xb3, 0xcd, 0x79, 0x6f, 0xa8,
	0xee, 0xda, 0x71, 0x6f, 0xf2, 0x17, 0xa6, 0x6f, 0xb4, 0x35, 0x75, 0xd7, 0x8c, 0x1f, 0x6e, 0x1f,
	0x21, 0x78, 0x28, 0xc8, 0x1a, 0x42, 0x7b, 0x4a, 0xf2, 0x67, 0x63, 0xf8, 0x66, 0x70, 0xf3, 0xc5,
	0xfc, 0x7b, 0x21, 0xdb, 0x53, 0x3e, 0x22, 0x37, 0x10, 0x6f, 0xf1, 0xd9, 0xfe, 0x63, 0x2f, 0x1e,
	0xdc, 0xd8, 0x8b, 0xcf, 0x1f, 0x93, 0x8f, 0x9e, 0x22, 0x77, 0xe2, 0xbb, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x46, 0x06, 0x72, 0xda, 0x57, 0x02, 0x00, 0x00,
}
