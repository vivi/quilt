// Code generated by protoc-gen-go.
// source: pb/pb.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/pb.proto

It has these top-level messages:
	DBQuery
	Machine
	MachineReply
	Container
	ContainerReply
	RunRequest
	RunResult
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

type DBQuery struct {
}

func (m *DBQuery) Reset()                    { *m = DBQuery{} }
func (m *DBQuery) String() string            { return proto.CompactTextString(m) }
func (*DBQuery) ProtoMessage()               {}
func (*DBQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Machine struct {
	ID        int32    `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Role      string   `protobuf:"bytes,2,opt,name=Role,json=role" json:"Role,omitempty"`
	Provider  string   `protobuf:"bytes,3,opt,name=Provider,json=provider" json:"Provider,omitempty"`
	Region    string   `protobuf:"bytes,4,opt,name=Region,json=region" json:"Region,omitempty"`
	Size      string   `protobuf:"bytes,5,opt,name=Size,json=size" json:"Size,omitempty"`
	DiskSize  int32    `protobuf:"varint,6,opt,name=DiskSize,json=diskSize" json:"DiskSize,omitempty"`
	SSHKeys   []string `protobuf:"bytes,7,rep,name=SSHKeys,json=sSHKeys" json:"SSHKeys,omitempty"`
	CloudID   string   `protobuf:"bytes,8,opt,name=CloudID,json=cloudID" json:"CloudID,omitempty"`
	PublicIP  string   `protobuf:"bytes,9,opt,name=PublicIP,json=publicIP" json:"PublicIP,omitempty"`
	PrivateIP string   `protobuf:"bytes,10,opt,name=PrivateIP,json=privateIP" json:"PrivateIP,omitempty"`
	Connected bool     `protobuf:"varint,11,opt,name=Connected,json=connected" json:"Connected,omitempty"`
}

func (m *Machine) Reset()                    { *m = Machine{} }
func (m *Machine) String() string            { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()               {}
func (*Machine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MachineReply struct {
	Machines []*Machine `protobuf:"bytes,3,rep,name=Machines,json=machines" json:"Machines,omitempty"`
}

func (m *MachineReply) Reset()                    { *m = MachineReply{} }
func (m *MachineReply) String() string            { return proto.CompactTextString(m) }
func (*MachineReply) ProtoMessage()               {}
func (*MachineReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MachineReply) GetMachines() []*Machine {
	if m != nil {
		return m.Machines
	}
	return nil
}

type Container struct {
	ID       int32    `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	DockerID string   `protobuf:"bytes,4,opt,name=DockerID,json=dockerID" json:"DockerID,omitempty"`
	Image    string   `protobuf:"bytes,5,opt,name=Image,json=image" json:"Image,omitempty"`
	Command  []string `protobuf:"bytes,6,rep,name=Command,json=command" json:"Command,omitempty"`
	Labels   []string `protobuf:"bytes,7,rep,name=Labels,json=labels" json:"Labels,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ContainerReply struct {
	Containers []*Container `protobuf:"bytes,3,rep,name=Containers,json=containers" json:"Containers,omitempty"`
}

func (m *ContainerReply) Reset()                    { *m = ContainerReply{} }
func (m *ContainerReply) String() string            { return proto.CompactTextString(m) }
func (*ContainerReply) ProtoMessage()               {}
func (*ContainerReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ContainerReply) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type RunRequest struct {
	Stitch string `protobuf:"bytes,1,opt,name=Stitch,json=stitch" json:"Stitch,omitempty"`
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RunResult struct {
	Success bool   `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *RunResult) Reset()                    { *m = RunResult{} }
func (m *RunResult) String() string            { return proto.CompactTextString(m) }
func (*RunResult) ProtoMessage()               {}
func (*RunResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*DBQuery)(nil), "DBQuery")
	proto.RegisterType((*Machine)(nil), "Machine")
	proto.RegisterType((*MachineReply)(nil), "MachineReply")
	proto.RegisterType((*Container)(nil), "Container")
	proto.RegisterType((*ContainerReply)(nil), "ContainerReply")
	proto.RegisterType((*RunRequest)(nil), "RunRequest")
	proto.RegisterType((*RunResult)(nil), "RunResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for API service

type APIClient interface {
	QueryMachines(ctx context.Context, in *DBQuery, opts ...grpc.CallOption) (*MachineReply, error)
	QueryContainers(ctx context.Context, in *DBQuery, opts ...grpc.CallOption) (*ContainerReply, error)
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResult, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) QueryMachines(ctx context.Context, in *DBQuery, opts ...grpc.CallOption) (*MachineReply, error) {
	out := new(MachineReply)
	err := grpc.Invoke(ctx, "/API/QueryMachines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryContainers(ctx context.Context, in *DBQuery, opts ...grpc.CallOption) (*ContainerReply, error) {
	out := new(ContainerReply)
	err := grpc.Invoke(ctx, "/API/QueryContainers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResult, error) {
	out := new(RunResult)
	err := grpc.Invoke(ctx, "/API/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	QueryMachines(context.Context, *DBQuery) (*MachineReply, error)
	QueryContainers(context.Context, *DBQuery) (*ContainerReply, error)
	Run(context.Context, *RunRequest) (*RunResult, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_QueryMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/QueryMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryMachines(ctx, req.(*DBQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/QueryContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryContainers(ctx, req.(*DBQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryMachines",
			Handler:    _API_QueryMachines_Handler,
		},
		{
			MethodName: "QueryContainers",
			Handler:    _API_QueryContainers_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _API_Run_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xb7, 0x7f, 0x92, 0x38, 0x53, 0x76, 0x57, 0xb2, 0x10, 0x8a, 0x2a, 0x0e, 0x95, 0xb5,
	0x87, 0x15, 0x07, 0x23, 0x2d, 0xdc, 0xe0, 0x02, 0x5b, 0x24, 0x2a, 0x40, 0x2a, 0xe9, 0x13, 0x24,
	0x8e, 0xb5, 0x6b, 0x6d, 0x1a, 0x07, 0x3b, 0x59, 0xa9, 0x5c, 0x38, 0xf0, 0x1c, 0xbc, 0x2b, 0xf6,
	0xc4, 0x69, 0x17, 0x71, 0xcb, 0x6f, 0xbe, 0xe9, 0xf8, 0x9b, 0xcf, 0x2e, 0x2c, 0xda, 0xf2, 0x75,
	0x5b, 0xf2, 0xd6, 0xe8, 0x4e, 0xb3, 0x14, 0x92, 0xf5, 0xc7, 0xef, 0xbd, 0x34, 0x07, 0xf6, 0x67,
	0x0a, 0xc9, 0xb7, 0x42, 0xdc, 0xab, 0x46, 0xd2, 0x0b, 0x98, 0x6e, 0xd6, 0xd9, 0x64, 0x35, 0xb9,
	0x8e, 0xf2, 0xa9, 0x5a, 0x53, 0x0a, 0xf3, 0x5c, 0xd7, 0x32, 0x9b, 0xba, 0x4a, 0x9a, 0xcf, 0x8d,
	0xfb, 0xa6, 0x4b, 0x20, 0x5b, 0xa3, 0x1f, 0x55, 0x25, 0x4d, 0x36, 0xc3, 0x3a, 0x69, 0x03, 0xd3,
	0x17, 0x10, 0xe7, 0xf2, 0x4e, 0xe9, 0x26, 0x9b, 0xa3, 0x12, 0x1b, 0x24, 0x3f, 0x67, 0xa7, 0x7e,
	0xca, 0x2c, 0x1a, 0xe6, 0x58, 0xf7, 0xed, 0xe7, 0xac, 0x95, 0x7d, 0xc0, 0x7a, 0x8c, 0x27, 0x92,
	0x2a, 0x30, 0xcd, 0x20, 0xd9, 0xed, 0x3e, 0x7f, 0x91, 0x07, 0x9b, 0x25, 0xab, 0x99, 0xfb, 0x49,
	0x62, 0x07, 0xf4, 0xca, 0x6d, 0xad, 0xfb, 0xca, 0xd9, 0x24, 0x38, 0x2c, 0x11, 0x03, 0xa2, 0xaf,
	0xbe, 0xac, 0x95, 0xd8, 0x6c, 0xb3, 0x34, 0xf8, 0x0a, 0x4c, 0x5f, 0x42, 0xba, 0x35, 0xea, 0xb1,
	0xe8, 0xa4, 0x13, 0x01, 0xc5, 0xb4, 0x1d, 0x0b, 0x5e, 0xbd, 0xd5, 0x4d, 0x23, 0x45, 0x27, 0xab,
	0x6c, 0xe1, 0x54, 0x92, 0xa7, 0x62, 0x2c, 0xb0, 0xb7, 0xf0, 0x2c, 0xc4, 0x93, 0xcb, 0xb6, 0x3e,
	0xd0, 0x2b, 0x20, 0x81, 0xad, 0xdb, 0x7f, 0x76, 0xbd, 0xb8, 0x21, 0x7c, 0x6c, 0x20, 0xfb, 0xa0,
	0xb0, 0x5f, 0x38, 0xb3, 0x2b, 0x1c, 0x98, 0xff, 0x62, 0xf5, 0xab, 0x6b, 0xf1, 0x20, 0x8d, 0xab,
	0x0e, 0x41, 0x91, 0x2a, 0x30, 0x7d, 0x0e, 0xd1, 0x66, 0x5f, 0xdc, 0x8d, 0x59, 0x45, 0xca, 0x03,
	0xae, 0xad, 0xf7, 0xfb, 0xa2, 0xa9, 0x5c, 0x56, 0x18, 0x88, 0x18, 0xd0, 0x47, 0xfe, 0xb5, 0x28,
	0x65, 0x3d, 0x26, 0x15, 0xd7, 0x48, 0xec, 0x3d, 0x5c, 0x1c, 0x0d, 0x0c, 0xc6, 0x5f, 0x01, 0x1c,
	0x2b, 0xa3, 0x75, 0xe0, 0xa7, 0x26, 0x10, 0x47, 0x95, 0x5d, 0x01, 0xe4, 0x7d, 0x93, 0xcb, 0x1f,
	0xbd, 0xb4, 0x9d, 0x3f, 0x63, 0xd7, 0xa9, 0x4e, 0xdc, 0xe3, 0x0e, 0xee, 0x0c, 0x8b, 0xc4, 0xde,
	0x41, 0x8a, 0x5d, 0xb6, 0xaf, 0x3b, 0xbc, 0xb3, 0x5e, 0x08, 0x69, 0x2d, 0x76, 0x11, 0x77, 0x67,
	0x03, 0xfa, 0x95, 0x3e, 0x19, 0xa3, 0x4d, 0x78, 0x46, 0x91, 0xf4, 0x70, 0xf3, 0x7b, 0x02, 0xb3,
	0x0f, 0xdb, 0x8d, 0xb3, 0x75, 0x8e, 0x0f, 0x71, 0x0c, 0x95, 0x12, 0x1e, 0x9e, 0xe6, 0xf2, 0x9c,
	0x3f, 0x4d, 0x9e, 0x9d, 0x51, 0x0e, 0x97, 0xa8, 0x9c, 0xf6, 0x78, 0xd2, 0x7d, 0xc9, 0xff, 0x5d,
	0xd8, 0xf5, 0xaf, 0x60, 0xe6, 0x0c, 0xd2, 0x05, 0x3f, 0x2d, 0xb3, 0x04, 0x7e, 0xf4, 0xcc, 0xce,
	0xca, 0x18, 0xff, 0x0f, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x2e, 0x6c, 0x99, 0x1e,
	0x03, 0x00, 0x00,
}
