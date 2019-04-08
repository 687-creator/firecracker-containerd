// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firecracker.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CreateVMRequest specifies creation parameters for a new FC instance
type CreateVMRequest struct {
	// Specifies the machine configuration for the VM
	MachineCfg *FirecrackerMachineConfiguration `protobuf:"bytes,1,opt,name=MachineCfg,proto3" json:"MachineCfg,omitempty"`
	// Specifies the file path where the kernel image is located
	KernelImagePath string `protobuf:"bytes,2,opt,name=KernelImagePath,proto3" json:"KernelImagePath,omitempty"`
	// Specifies the commandline arguments that should be passed to the kernel
	KernelArgs string `protobuf:"bytes,3,opt,name=KernelArgs,proto3" json:"KernelArgs,omitempty"`
	// Specifies the root block device for the VM
	RootDrive *FirecrackerDrive `protobuf:"bytes,4,opt,name=RootDrive,proto3" json:"RootDrive,omitempty"`
	// Specifies the additional block device config for the VM.
	AdditionalDrives []*FirecrackerDrive `protobuf:"bytes,5,rep,name=AdditionalDrives,proto3" json:"AdditionalDrives,omitempty"`
	// Specifies the networking configuration for a VM
	NetworkInterfaces []*FirecrackerNetworkInterface `protobuf:"bytes,6,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	// The number of dummy drives to reserve in advance before running FC instance.
	ContainerCount       int32    `protobuf:"varint,7,opt,name=ContainerCount,proto3" json:"ContainerCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMRequest) Reset()         { *m = CreateVMRequest{} }
func (m *CreateVMRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVMRequest) ProtoMessage()    {}
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_73f6b188d2d4dd0e, []int{0}
}
func (m *CreateVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMRequest.Unmarshal(m, b)
}
func (m *CreateVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMRequest.Merge(dst, src)
}
func (m *CreateVMRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVMRequest.Size(m)
}
func (m *CreateVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMRequest proto.InternalMessageInfo

func (m *CreateVMRequest) GetMachineCfg() *FirecrackerMachineConfiguration {
	if m != nil {
		return m.MachineCfg
	}
	return nil
}

func (m *CreateVMRequest) GetKernelImagePath() string {
	if m != nil {
		return m.KernelImagePath
	}
	return ""
}

func (m *CreateVMRequest) GetKernelArgs() string {
	if m != nil {
		return m.KernelArgs
	}
	return ""
}

func (m *CreateVMRequest) GetRootDrive() *FirecrackerDrive {
	if m != nil {
		return m.RootDrive
	}
	return nil
}

func (m *CreateVMRequest) GetAdditionalDrives() []*FirecrackerDrive {
	if m != nil {
		return m.AdditionalDrives
	}
	return nil
}

func (m *CreateVMRequest) GetNetworkInterfaces() []*FirecrackerNetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *CreateVMRequest) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

type CreateVMResponse struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMResponse) Reset()         { *m = CreateVMResponse{} }
func (m *CreateVMResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVMResponse) ProtoMessage()    {}
func (*CreateVMResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_73f6b188d2d4dd0e, []int{1}
}
func (m *CreateVMResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMResponse.Unmarshal(m, b)
}
func (m *CreateVMResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMResponse.Marshal(b, m, deterministic)
}
func (dst *CreateVMResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMResponse.Merge(dst, src)
}
func (m *CreateVMResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVMResponse.Size(m)
}
func (m *CreateVMResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMResponse proto.InternalMessageInfo

func (m *CreateVMResponse) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type StopVMRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopVMRequest) Reset()         { *m = StopVMRequest{} }
func (m *StopVMRequest) String() string { return proto.CompactTextString(m) }
func (*StopVMRequest) ProtoMessage()    {}
func (*StopVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_73f6b188d2d4dd0e, []int{2}
}
func (m *StopVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopVMRequest.Unmarshal(m, b)
}
func (m *StopVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopVMRequest.Marshal(b, m, deterministic)
}
func (dst *StopVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopVMRequest.Merge(dst, src)
}
func (m *StopVMRequest) XXX_Size() int {
	return xxx_messageInfo_StopVMRequest.Size(m)
}
func (m *StopVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopVMRequest proto.InternalMessageInfo

func (m *StopVMRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMAddressRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMAddressRequest) Reset()         { *m = GetVMAddressRequest{} }
func (m *GetVMAddressRequest) String() string { return proto.CompactTextString(m) }
func (*GetVMAddressRequest) ProtoMessage()    {}
func (*GetVMAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_73f6b188d2d4dd0e, []int{3}
}
func (m *GetVMAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMAddressRequest.Unmarshal(m, b)
}
func (m *GetVMAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMAddressRequest.Marshal(b, m, deterministic)
}
func (dst *GetVMAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMAddressRequest.Merge(dst, src)
}
func (m *GetVMAddressRequest) XXX_Size() int {
	return xxx_messageInfo_GetVMAddressRequest.Size(m)
}
func (m *GetVMAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMAddressRequest proto.InternalMessageInfo

func (m *GetVMAddressRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMAddressResponse struct {
	SocketPath           string   `protobuf:"bytes,1,opt,name=SocketPath,proto3" json:"SocketPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMAddressResponse) Reset()         { *m = GetVMAddressResponse{} }
func (m *GetVMAddressResponse) String() string { return proto.CompactTextString(m) }
func (*GetVMAddressResponse) ProtoMessage()    {}
func (*GetVMAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_73f6b188d2d4dd0e, []int{4}
}
func (m *GetVMAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMAddressResponse.Unmarshal(m, b)
}
func (m *GetVMAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMAddressResponse.Marshal(b, m, deterministic)
}
func (dst *GetVMAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMAddressResponse.Merge(dst, src)
}
func (m *GetVMAddressResponse) XXX_Size() int {
	return xxx_messageInfo_GetVMAddressResponse.Size(m)
}
func (m *GetVMAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMAddressResponse proto.InternalMessageInfo

func (m *GetVMAddressResponse) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVMRequest)(nil), "CreateVMRequest")
	proto.RegisterType((*CreateVMResponse)(nil), "CreateVMResponse")
	proto.RegisterType((*StopVMRequest)(nil), "StopVMRequest")
	proto.RegisterType((*GetVMAddressRequest)(nil), "GetVMAddressRequest")
	proto.RegisterType((*GetVMAddressResponse)(nil), "GetVMAddressResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FirecrackerClient is the client API for Firecracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirecrackerClient interface {
	// Runs new Firecracker VM instance
	CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*CreateVMResponse, error)
	// Stops existing Firecracker instance by VM ID
	StopVM(ctx context.Context, in *StopVMRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets VM instance socket file location
	GetVMAddress(ctx context.Context, in *GetVMAddressRequest, opts ...grpc.CallOption) (*GetVMAddressResponse, error)
}

type firecrackerClient struct {
	cc *grpc.ClientConn
}

func NewFirecrackerClient(cc *grpc.ClientConn) FirecrackerClient {
	return &firecrackerClient{cc}
}

func (c *firecrackerClient) CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*CreateVMResponse, error) {
	out := new(CreateVMResponse)
	err := c.cc.Invoke(ctx, "/Firecracker/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) StopVM(ctx context.Context, in *StopVMRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Firecracker/StopVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) GetVMAddress(ctx context.Context, in *GetVMAddressRequest, opts ...grpc.CallOption) (*GetVMAddressResponse, error) {
	out := new(GetVMAddressResponse)
	err := c.cc.Invoke(ctx, "/Firecracker/GetVMAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirecrackerServer is the server API for Firecracker service.
type FirecrackerServer interface {
	// Runs new Firecracker VM instance
	CreateVM(context.Context, *CreateVMRequest) (*CreateVMResponse, error)
	// Stops existing Firecracker instance by VM ID
	StopVM(context.Context, *StopVMRequest) (*empty.Empty, error)
	// Gets VM instance socket file location
	GetVMAddress(context.Context, *GetVMAddressRequest) (*GetVMAddressResponse, error)
}

func RegisterFirecrackerServer(s *grpc.Server, srv FirecrackerServer) {
	s.RegisterService(&_Firecracker_serviceDesc, srv)
}

func _Firecracker_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).CreateVM(ctx, req.(*CreateVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_StopVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).StopVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/StopVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).StopVM(ctx, req.(*StopVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_GetVMAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVMAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).GetVMAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/GetVMAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).GetVMAddress(ctx, req.(*GetVMAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Firecracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Firecracker",
	HandlerType: (*FirecrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVM",
			Handler:    _Firecracker_CreateVM_Handler,
		},
		{
			MethodName: "StopVM",
			Handler:    _Firecracker_StopVM_Handler,
		},
		{
			MethodName: "GetVMAddress",
			Handler:    _Firecracker_GetVMAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firecracker.proto",
}

func init() { proto.RegisterFile("firecracker.proto", fileDescriptor_firecracker_73f6b188d2d4dd0e) }

var fileDescriptor_firecracker_73f6b188d2d4dd0e = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xe8, 0xda, 0xd1, 0x5b, 0xd8, 0x5a, 0x33, 0x50, 0x14, 0xd0, 0x14, 0x15, 0x69, 0x0a,
	0x2f, 0x0e, 0x1a, 0x12, 0x2f, 0x08, 0x89, 0x92, 0x01, 0x2a, 0x28, 0x08, 0x79, 0x52, 0x1f, 0x78,
	0x33, 0xc9, 0x4d, 0x16, 0xb5, 0xb3, 0x83, 0xed, 0x80, 0xf6, 0xc8, 0x07, 0xf1, 0x8f, 0x68, 0xce,
	0xaa, 0xa4, 0x69, 0xc4, 0x53, 0xe2, 0x73, 0xcf, 0xb1, 0xcf, 0xb9, 0xf7, 0xc2, 0x2c, 0x2b, 0x14,
	0x26, 0x8a, 0x27, 0x6b, 0x54, 0xb4, 0x54, 0xd2, 0x48, 0xef, 0x69, 0x2e, 0x65, 0xbe, 0xc1, 0xd0,
	0x9e, 0x7e, 0x54, 0x59, 0x88, 0xd7, 0xa5, 0xb9, 0xb9, 0x2b, 0x4e, 0xcc, 0x4d, 0x89, 0xba, 0x3e,
	0xcc, 0xff, 0x0c, 0xe0, 0x38, 0x52, 0xc8, 0x0d, 0xae, 0x62, 0x86, 0x3f, 0x2b, 0xd4, 0x86, 0xbc,
	0x03, 0x88, 0x79, 0x72, 0x55, 0x08, 0x8c, 0xb2, 0xdc, 0x75, 0x7c, 0x27, 0x98, 0x9c, 0xfb, 0xf4,
	0x63, 0xf3, 0xca, 0xb6, 0x2a, 0x45, 0x56, 0xe4, 0x95, 0xe2, 0xa6, 0x90, 0x82, 0xb5, 0x34, 0x24,
	0x80, 0xe3, 0x2f, 0xa8, 0x04, 0x6e, 0x96, 0xd7, 0x3c, 0xc7, 0x6f, 0xdc, 0x5c, 0xb9, 0xf7, 0x7c,
	0x27, 0x18, 0xb3, 0x2e, 0x4c, 0x4e, 0x01, 0x6a, 0x68, 0xa1, 0x72, 0xed, 0x0e, 0x2c, 0xa9, 0x85,
	0x90, 0x10, 0xc6, 0x4c, 0x4a, 0x73, 0xa1, 0x8a, 0x5f, 0xe8, 0x1e, 0x58, 0x2b, 0xb3, 0xb6, 0x15,
	0x5b, 0x60, 0x0d, 0x87, 0xbc, 0x85, 0xe9, 0x22, 0x4d, 0x8b, 0x5b, 0x4b, 0x7c, 0x63, 0x21, 0xed,
	0x0e, 0xfd, 0x41, 0xbf, 0x6e, 0x8f, 0x4a, 0x3e, 0xc3, 0xec, 0x2b, 0x9a, 0xdf, 0x52, 0xad, 0x97,
	0xc2, 0xa0, 0xca, 0x78, 0x82, 0xda, 0x1d, 0x59, 0xfd, 0xb3, 0xb6, 0xbe, 0x4b, 0x62, 0xfb, 0x32,
	0x72, 0x06, 0x47, 0x91, 0x14, 0x86, 0x17, 0x02, 0x55, 0x24, 0x2b, 0x61, 0xdc, 0x43, 0xdf, 0x09,
	0x86, 0xac, 0x83, 0xce, 0xcf, 0x60, 0xda, 0x8c, 0x40, 0x97, 0x52, 0x68, 0x24, 0x04, 0x0e, 0x56,
	0xf1, 0xf2, 0xc2, 0x76, 0x7f, 0xcc, 0xec, 0xff, 0xfc, 0x39, 0x3c, 0xbc, 0x34, 0xb2, 0x6c, 0x06,
	0xd5, 0x47, 0x7a, 0x01, 0x8f, 0x3e, 0xa1, 0x59, 0xc5, 0x8b, 0x34, 0x55, 0xa8, 0xf5, 0xff, 0xa8,
	0xaf, 0xe1, 0x64, 0x97, 0x7a, 0xf7, 0xf6, 0x29, 0xc0, 0xa5, 0x4c, 0xd6, 0x68, 0xec, 0xe0, 0x6a,
	0x45, 0x0b, 0x39, 0xff, 0xeb, 0xc0, 0xa4, 0xd5, 0x0a, 0x12, 0xc2, 0xfd, 0xad, 0x7f, 0x32, 0xa5,
	0x9d, 0x6d, 0xf2, 0x66, 0x74, 0x2f, 0xdc, 0x4b, 0x18, 0xd5, 0x41, 0xc8, 0x11, 0xdd, 0x49, 0xe4,
	0x3d, 0xa1, 0xf5, 0xe6, 0xd2, 0xed, 0xe6, 0xd2, 0x0f, 0xb7, 0x9b, 0x4b, 0xde, 0xc0, 0x83, 0xb6,
	0x55, 0x72, 0x42, 0x7b, 0x42, 0x7a, 0x8f, 0x69, 0x5f, 0x9e, 0xf7, 0x87, 0xdf, 0x87, 0xf5, 0x75,
	0x23, 0xfb, 0x79, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x36, 0x3c, 0x79, 0x32, 0x03, 0x00,
	0x00,
}