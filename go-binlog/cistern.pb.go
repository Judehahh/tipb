// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cistern.proto

package binlog

import (
	context "context"
	fmt "fmt"
	_ "gogoproto"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DumpBinlogReq struct {
	// beginCommitTS speicifies the position from which begin to dump binlogs.
	// note that actually the result of dump starts from the one next to beginCommitTS
	// it should be zero in case of the first request.
	BeginCommitTS        int64    `protobuf:"varint,1,opt,name=beginCommitTS,proto3" json:"beginCommitTS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpBinlogReq) Reset()         { *m = DumpBinlogReq{} }
func (m *DumpBinlogReq) String() string { return proto.CompactTextString(m) }
func (*DumpBinlogReq) ProtoMessage()    {}
func (*DumpBinlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1013ee3d8abade, []int{0}
}
func (m *DumpBinlogReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DumpBinlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DumpBinlogReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DumpBinlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpBinlogReq.Merge(m, src)
}
func (m *DumpBinlogReq) XXX_Size() int {
	return m.Size()
}
func (m *DumpBinlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpBinlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_DumpBinlogReq proto.InternalMessageInfo

func (m *DumpBinlogReq) GetBeginCommitTS() int64 {
	if m != nil {
		return m.BeginCommitTS
	}
	return 0
}

type DumpBinlogResp struct {
	// CommitTS specifies the commitTS of binlog
	CommitTS int64 `protobuf:"varint,1,opt,name=commitTS,proto3" json:"commitTS,omitempty"`
	// payloads is bytecodes encoded from binlog item
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// ddljob is json bytes marshaled from corresponding ddljob struct if payload is a DDL type of binlog
	Ddljob               []byte   `protobuf:"bytes,3,opt,name=ddljob,proto3" json:"ddljob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpBinlogResp) Reset()         { *m = DumpBinlogResp{} }
func (m *DumpBinlogResp) String() string { return proto.CompactTextString(m) }
func (*DumpBinlogResp) ProtoMessage()    {}
func (*DumpBinlogResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1013ee3d8abade, []int{1}
}
func (m *DumpBinlogResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DumpBinlogResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DumpBinlogResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DumpBinlogResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpBinlogResp.Merge(m, src)
}
func (m *DumpBinlogResp) XXX_Size() int {
	return m.Size()
}
func (m *DumpBinlogResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpBinlogResp.DiscardUnknown(m)
}

var xxx_messageInfo_DumpBinlogResp proto.InternalMessageInfo

func (m *DumpBinlogResp) GetCommitTS() int64 {
	if m != nil {
		return m.CommitTS
	}
	return 0
}

func (m *DumpBinlogResp) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DumpBinlogResp) GetDdljob() []byte {
	if m != nil {
		return m.Ddljob
	}
	return nil
}

type DumpDDLJobsReq struct {
	// beginCommitTS is the start point of drainer processing binlog, DumpDDLJobs() returns
	// all history DDL jobs before this position, then drainer will apply these DDL jobs
	// in order of job ID to restore the whole schema info at that moment.
	BeginCommitTS        int64    `protobuf:"varint,1,opt,name=beginCommitTS,proto3" json:"beginCommitTS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpDDLJobsReq) Reset()         { *m = DumpDDLJobsReq{} }
func (m *DumpDDLJobsReq) String() string { return proto.CompactTextString(m) }
func (*DumpDDLJobsReq) ProtoMessage()    {}
func (*DumpDDLJobsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1013ee3d8abade, []int{2}
}
func (m *DumpDDLJobsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DumpDDLJobsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DumpDDLJobsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DumpDDLJobsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpDDLJobsReq.Merge(m, src)
}
func (m *DumpDDLJobsReq) XXX_Size() int {
	return m.Size()
}
func (m *DumpDDLJobsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpDDLJobsReq.DiscardUnknown(m)
}

var xxx_messageInfo_DumpDDLJobsReq proto.InternalMessageInfo

func (m *DumpDDLJobsReq) GetBeginCommitTS() int64 {
	if m != nil {
		return m.BeginCommitTS
	}
	return 0
}

type DumpDDLJobsResp struct {
	// ddljobs is an array of JSON encoded history DDL jobs
	Ddljobs              [][]byte `protobuf:"bytes,1,rep,name=ddljobs,proto3" json:"ddljobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpDDLJobsResp) Reset()         { *m = DumpDDLJobsResp{} }
func (m *DumpDDLJobsResp) String() string { return proto.CompactTextString(m) }
func (*DumpDDLJobsResp) ProtoMessage()    {}
func (*DumpDDLJobsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1013ee3d8abade, []int{3}
}
func (m *DumpDDLJobsResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DumpDDLJobsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DumpDDLJobsResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DumpDDLJobsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpDDLJobsResp.Merge(m, src)
}
func (m *DumpDDLJobsResp) XXX_Size() int {
	return m.Size()
}
func (m *DumpDDLJobsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpDDLJobsResp.DiscardUnknown(m)
}

var xxx_messageInfo_DumpDDLJobsResp proto.InternalMessageInfo

func (m *DumpDDLJobsResp) GetDdljobs() [][]byte {
	if m != nil {
		return m.Ddljobs
	}
	return nil
}

type NotifyReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReq) Reset()         { *m = NotifyReq{} }
func (m *NotifyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReq) ProtoMessage()    {}
func (*NotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1013ee3d8abade, []int{4}
}
func (m *NotifyReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReq.Merge(m, src)
}
func (m *NotifyReq) XXX_Size() int {
	return m.Size()
}
func (m *NotifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReq proto.InternalMessageInfo

type NotifyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyResp) Reset()         { *m = NotifyResp{} }
func (m *NotifyResp) String() string { return proto.CompactTextString(m) }
func (*NotifyResp) ProtoMessage()    {}
func (*NotifyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1013ee3d8abade, []int{5}
}
func (m *NotifyResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyResp.Merge(m, src)
}
func (m *NotifyResp) XXX_Size() int {
	return m.Size()
}
func (m *NotifyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyResp.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DumpBinlogReq)(nil), "binlog.DumpBinlogReq")
	proto.RegisterType((*DumpBinlogResp)(nil), "binlog.DumpBinlogResp")
	proto.RegisterType((*DumpDDLJobsReq)(nil), "binlog.DumpDDLJobsReq")
	proto.RegisterType((*DumpDDLJobsResp)(nil), "binlog.DumpDDLJobsResp")
	proto.RegisterType((*NotifyReq)(nil), "binlog.NotifyReq")
	proto.RegisterType((*NotifyResp)(nil), "binlog.NotifyResp")
}

func init() { proto.RegisterFile("cistern.proto", fileDescriptor_da1013ee3d8abade) }

var fileDescriptor_da1013ee3d8abade = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x2c, 0x2e,
	0x49, 0x2d, 0xca, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xca, 0xcc, 0xcb, 0xc9,
	0x4f, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xe9, 0x83, 0x58, 0x10, 0x59, 0x25, 0x53,
	0x2e, 0x5e, 0x97, 0xd2, 0xdc, 0x02, 0x27, 0xb0, 0x9a, 0xa0, 0xd4, 0x42, 0x21, 0x15, 0x2e, 0xde,
	0xa4, 0xd4, 0xf4, 0xcc, 0x3c, 0xe7, 0xfc, 0xdc, 0xdc, 0xcc, 0x92, 0x90, 0x60, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xe6, 0x20, 0x54, 0x41, 0xa5, 0x38, 0x2e, 0x3e, 0x64, 0x6d, 0xc5, 0x05, 0x42, 0x52,
	0x5c, 0x1c, 0xc9, 0xa8, 0x5a, 0xe0, 0x7c, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc,
	0xc4, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0x48, 0x8c, 0x8b, 0x2d, 0x25,
	0x25, 0x27, 0x2b, 0x3f, 0x49, 0x82, 0x19, 0x2c, 0x01, 0xe5, 0x29, 0x99, 0x41, 0xcc, 0x77, 0x71,
	0xf1, 0xf1, 0xca, 0x4f, 0x2a, 0x26, 0xde, 0x5d, 0xda, 0x5c, 0xfc, 0x28, 0xfa, 0x8a, 0x0b, 0x40,
	0x96, 0x43, 0x0c, 0x2d, 0x96, 0x60, 0x54, 0x60, 0x06, 0x59, 0x0e, 0xe5, 0x2a, 0x71, 0x73, 0x71,
	0xfa, 0xe5, 0x97, 0x64, 0xa6, 0x55, 0x06, 0xa5, 0x16, 0x2a, 0xf1, 0x70, 0x71, 0xc1, 0x38, 0xc5,
	0x05, 0x46, 0xfb, 0x19, 0xb9, 0xd8, 0x9d, 0x21, 0xc1, 0x28, 0x64, 0xcf, 0xc5, 0x85, 0xf0, 0xab,
	0x90, 0xa8, 0x1e, 0x24, 0x3c, 0xf5, 0x50, 0x82, 0x4d, 0x4a, 0x0c, 0x9b, 0x70, 0x71, 0x81, 0x12,
	0x83, 0x01, 0xa3, 0x90, 0x03, 0x17, 0x37, 0x92, 0xa3, 0x84, 0x50, 0x94, 0x22, 0x7c, 0x28, 0x25,
	0x8e, 0x55, 0x1c, 0x64, 0x86, 0x90, 0x21, 0x17, 0x1b, 0xc4, 0x71, 0x42, 0x82, 0x30, 0x45, 0x70,
	0x97, 0x4b, 0x09, 0xa1, 0x0b, 0x81, 0xb4, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x63, 0xdc, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x2d, 0x8f, 0x8f, 0x20, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CisternClient is the client API for Cistern service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CisternClient interface {
	// DumpBinlog dumps continuous binlog items in a stream from a given position
	DumpBinlog(ctx context.Context, in *DumpBinlogReq, opts ...grpc.CallOption) (Cistern_DumpBinlogClient, error)
	// DumpDDLJobs dumps all history DDL jobs before a specified commitTS
	DumpDDLJobs(ctx context.Context, in *DumpDDLJobsReq, opts ...grpc.CallOption) (*DumpDDLJobsResp, error)
	// Notify notifies all living cisterns that a new pump is coming
	// the living cisterns can be queried from pd
	Notify(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyResp, error)
}

type cisternClient struct {
	cc *grpc.ClientConn
}

func NewCisternClient(cc *grpc.ClientConn) CisternClient {
	return &cisternClient{cc}
}

func (c *cisternClient) DumpBinlog(ctx context.Context, in *DumpBinlogReq, opts ...grpc.CallOption) (Cistern_DumpBinlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Cistern_serviceDesc.Streams[0], "/binlog.Cistern/DumpBinlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &cisternDumpBinlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cistern_DumpBinlogClient interface {
	Recv() (*DumpBinlogResp, error)
	grpc.ClientStream
}

type cisternDumpBinlogClient struct {
	grpc.ClientStream
}

func (x *cisternDumpBinlogClient) Recv() (*DumpBinlogResp, error) {
	m := new(DumpBinlogResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cisternClient) DumpDDLJobs(ctx context.Context, in *DumpDDLJobsReq, opts ...grpc.CallOption) (*DumpDDLJobsResp, error) {
	out := new(DumpDDLJobsResp)
	err := c.cc.Invoke(ctx, "/binlog.Cistern/DumpDDLJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cisternClient) Notify(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyResp, error) {
	out := new(NotifyResp)
	err := c.cc.Invoke(ctx, "/binlog.Cistern/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CisternServer is the server API for Cistern service.
type CisternServer interface {
	// DumpBinlog dumps continuous binlog items in a stream from a given position
	DumpBinlog(*DumpBinlogReq, Cistern_DumpBinlogServer) error
	// DumpDDLJobs dumps all history DDL jobs before a specified commitTS
	DumpDDLJobs(context.Context, *DumpDDLJobsReq) (*DumpDDLJobsResp, error)
	// Notify notifies all living cisterns that a new pump is coming
	// the living cisterns can be queried from pd
	Notify(context.Context, *NotifyReq) (*NotifyResp, error)
}

// UnimplementedCisternServer can be embedded to have forward compatible implementations.
type UnimplementedCisternServer struct {
}

func (*UnimplementedCisternServer) DumpBinlog(req *DumpBinlogReq, srv Cistern_DumpBinlogServer) error {
	return status.Errorf(codes.Unimplemented, "method DumpBinlog not implemented")
}
func (*UnimplementedCisternServer) DumpDDLJobs(ctx context.Context, req *DumpDDLJobsReq) (*DumpDDLJobsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpDDLJobs not implemented")
}
func (*UnimplementedCisternServer) Notify(ctx context.Context, req *NotifyReq) (*NotifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}

func RegisterCisternServer(s *grpc.Server, srv CisternServer) {
	s.RegisterService(&_Cistern_serviceDesc, srv)
}

func _Cistern_DumpBinlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpBinlogReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CisternServer).DumpBinlog(m, &cisternDumpBinlogServer{stream})
}

type Cistern_DumpBinlogServer interface {
	Send(*DumpBinlogResp) error
	grpc.ServerStream
}

type cisternDumpBinlogServer struct {
	grpc.ServerStream
}

func (x *cisternDumpBinlogServer) Send(m *DumpBinlogResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Cistern_DumpDDLJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpDDLJobsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CisternServer).DumpDDLJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binlog.Cistern/DumpDDLJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CisternServer).DumpDDLJobs(ctx, req.(*DumpDDLJobsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cistern_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CisternServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binlog.Cistern/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CisternServer).Notify(ctx, req.(*NotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cistern_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlog.Cistern",
	HandlerType: (*CisternServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DumpDDLJobs",
			Handler:    _Cistern_DumpDDLJobs_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Cistern_Notify_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DumpBinlog",
			Handler:       _Cistern_DumpBinlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cistern.proto",
}

func (m *DumpBinlogReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DumpBinlogReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DumpBinlogReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.BeginCommitTS != 0 {
		i = encodeVarintCistern(dAtA, i, uint64(m.BeginCommitTS))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DumpBinlogResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DumpBinlogResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DumpBinlogResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ddljob) > 0 {
		i -= len(m.Ddljob)
		copy(dAtA[i:], m.Ddljob)
		i = encodeVarintCistern(dAtA, i, uint64(len(m.Ddljob)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintCistern(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if m.CommitTS != 0 {
		i = encodeVarintCistern(dAtA, i, uint64(m.CommitTS))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DumpDDLJobsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DumpDDLJobsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DumpDDLJobsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.BeginCommitTS != 0 {
		i = encodeVarintCistern(dAtA, i, uint64(m.BeginCommitTS))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DumpDDLJobsResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DumpDDLJobsResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DumpDDLJobsResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ddljobs) > 0 {
		for iNdEx := len(m.Ddljobs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Ddljobs[iNdEx])
			copy(dAtA[i:], m.Ddljobs[iNdEx])
			i = encodeVarintCistern(dAtA, i, uint64(len(m.Ddljobs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *NotifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *NotifyResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintCistern(dAtA []byte, offset int, v uint64) int {
	offset -= sovCistern(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DumpBinlogReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BeginCommitTS != 0 {
		n += 1 + sovCistern(uint64(m.BeginCommitTS))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DumpBinlogResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommitTS != 0 {
		n += 1 + sovCistern(uint64(m.CommitTS))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovCistern(uint64(l))
	}
	l = len(m.Ddljob)
	if l > 0 {
		n += 1 + l + sovCistern(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DumpDDLJobsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BeginCommitTS != 0 {
		n += 1 + sovCistern(uint64(m.BeginCommitTS))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DumpDDLJobsResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ddljobs) > 0 {
		for _, b := range m.Ddljobs {
			l = len(b)
			n += 1 + l + sovCistern(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NotifyReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NotifyResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCistern(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCistern(x uint64) (n int) {
	return sovCistern(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DumpBinlogReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpBinlogReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpBinlogReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginCommitTS", wireType)
			}
			m.BeginCommitTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginCommitTS |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCistern
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DumpBinlogResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpBinlogResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpBinlogResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitTS", wireType)
			}
			m.CommitTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitTS |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCistern
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCistern
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ddljob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCistern
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCistern
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ddljob = append(m.Ddljob[:0], dAtA[iNdEx:postIndex]...)
			if m.Ddljob == nil {
				m.Ddljob = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCistern
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DumpDDLJobsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpDDLJobsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpDDLJobsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginCommitTS", wireType)
			}
			m.BeginCommitTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginCommitTS |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCistern
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DumpDDLJobsResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpDDLJobsResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpDDLJobsResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ddljobs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCistern
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCistern
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ddljobs = append(m.Ddljobs, make([]byte, postIndex-iNdEx))
			copy(m.Ddljobs[len(m.Ddljobs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCistern
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCistern
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCistern
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCistern(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCistern
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCistern
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCistern
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCistern        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCistern          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCistern = fmt.Errorf("proto: unexpected end of group")
)
