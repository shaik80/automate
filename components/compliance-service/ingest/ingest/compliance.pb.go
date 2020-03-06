// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/compliance-service/ingest/ingest/compliance.proto

package ingest

import (
	context "context"
	fmt "fmt"
	event "github.com/chef/automate/api/interservice/event"
	compliance "github.com/chef/automate/components/compliance-service/ingest/events/compliance"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type ProjectUpdateStatusReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectUpdateStatusReq) Reset()         { *m = ProjectUpdateStatusReq{} }
func (m *ProjectUpdateStatusReq) String() string { return proto.CompactTextString(m) }
func (*ProjectUpdateStatusReq) ProtoMessage()    {}
func (*ProjectUpdateStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_26326bddb420a177, []int{0}
}

func (m *ProjectUpdateStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectUpdateStatusReq.Unmarshal(m, b)
}
func (m *ProjectUpdateStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectUpdateStatusReq.Marshal(b, m, deterministic)
}
func (m *ProjectUpdateStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectUpdateStatusReq.Merge(m, src)
}
func (m *ProjectUpdateStatusReq) XXX_Size() int {
	return xxx_messageInfo_ProjectUpdateStatusReq.Size(m)
}
func (m *ProjectUpdateStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectUpdateStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectUpdateStatusReq proto.InternalMessageInfo

type ProjectUpdateStatusResp struct {
	State                 string               `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	EstimatedTimeComplete *timestamp.Timestamp `protobuf:"bytes,2,opt,name=estimated_time_complete,json=estimatedTimeComplete,proto3" json:"estimated_time_complete,omitempty"`
	PercentageComplete    float32              `protobuf:"fixed32,3,opt,name=percentage_complete,json=percentageComplete,proto3" json:"percentage_complete,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *ProjectUpdateStatusResp) Reset()         { *m = ProjectUpdateStatusResp{} }
func (m *ProjectUpdateStatusResp) String() string { return proto.CompactTextString(m) }
func (*ProjectUpdateStatusResp) ProtoMessage()    {}
func (*ProjectUpdateStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_26326bddb420a177, []int{1}
}

func (m *ProjectUpdateStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectUpdateStatusResp.Unmarshal(m, b)
}
func (m *ProjectUpdateStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectUpdateStatusResp.Marshal(b, m, deterministic)
}
func (m *ProjectUpdateStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectUpdateStatusResp.Merge(m, src)
}
func (m *ProjectUpdateStatusResp) XXX_Size() int {
	return xxx_messageInfo_ProjectUpdateStatusResp.Size(m)
}
func (m *ProjectUpdateStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectUpdateStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectUpdateStatusResp proto.InternalMessageInfo

func (m *ProjectUpdateStatusResp) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ProjectUpdateStatusResp) GetEstimatedTimeComplete() *timestamp.Timestamp {
	if m != nil {
		return m.EstimatedTimeComplete
	}
	return nil
}

func (m *ProjectUpdateStatusResp) GetPercentageComplete() float32 {
	if m != nil {
		return m.PercentageComplete
	}
	return 0
}

func init() {
	proto.RegisterType((*ProjectUpdateStatusReq)(nil), "chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusReq")
	proto.RegisterType((*ProjectUpdateStatusResp)(nil), "chef.automate.domain.compliance.ingest.ingest.ProjectUpdateStatusResp")
}

func init() {
	proto.RegisterFile("components/compliance-service/ingest/ingest/compliance.proto", fileDescriptor_26326bddb420a177)
}

var fileDescriptor_26326bddb420a177 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x55, 0xfa, 0x04, 0x12, 0x7e, 0x9b, 0x1f, 0xbc, 0x57, 0x02, 0x82, 0x2a, 0x62, 0x88, 0x90,
	0x6a, 0x4b, 0x8f, 0x0d, 0x31, 0x20, 0x1e, 0x7d, 0xc0, 0x80, 0x84, 0x02, 0x2c, 0x2c, 0x95, 0x9b,
	0xdc, 0xa6, 0x46, 0x8d, 0x6d, 0xec, 0x9b, 0x4a, 0xac, 0xfc, 0x02, 0x13, 0x3b, 0x9f, 0xc0, 0xce,
	0x47, 0xf0, 0x0b, 0x7c, 0x08, 0x72, 0x9c, 0xb4, 0x55, 0x9b, 0xa1, 0x95, 0x58, 0xe2, 0x38, 0xf7,
	0x9c, 0xe3, 0x73, 0xef, 0x71, 0xc8, 0xb3, 0x5c, 0x57, 0x46, 0x2b, 0x50, 0xe8, 0xb8, 0x7f, 0x5d,
	0x4a, 0xa1, 0x72, 0x18, 0x3b, 0xb0, 0x2b, 0x99, 0x03, 0x97, 0xaa, 0x04, 0x87, 0xdd, 0xb2, 0x01,
	0x30, 0x63, 0x35, 0x6a, 0x3a, 0xce, 0x17, 0x30, 0x67, 0xa2, 0x46, 0x5d, 0x09, 0x04, 0x56, 0xe8,
	0x4a, 0x48, 0xc5, 0xb6, 0x60, 0x81, 0xd8, 0x2e, 0xf1, 0xfd, 0x52, 0xeb, 0x72, 0x09, 0x5c, 0x18,
	0xc9, 0x85, 0x52, 0x1a, 0x05, 0x4a, 0xad, 0x5c, 0x10, 0x8b, 0x5f, 0x1d, 0x64, 0x05, 0x56, 0x3b,
	0x80, 0x3d, 0x57, 0xf1, 0xbd, 0xf6, 0x98, 0x66, 0x37, 0xab, 0xe7, 0x1c, 0x2a, 0x83, 0x5f, 0xdb,
	0x62, 0xe2, 0x0f, 0x97, 0x0a, 0xc1, 0x76, 0xc2, 0x8d, 0x62, 0x78, 0xb6, 0x98, 0x87, 0xbb, 0x02,
	0x28, 0x2b, 0x70, 0x28, 0x2a, 0x13, 0x00, 0xc9, 0x90, 0x9c, 0xbf, 0xb3, 0xfa, 0x33, 0xe4, 0xf8,
	0xd1, 0x14, 0x02, 0xe1, 0x3d, 0x0a, 0xac, 0x5d, 0x06, 0x5f, 0x92, 0x5f, 0x11, 0xb9, 0xe8, 0x2d,
	0x39, 0x43, 0x6f, 0x93, 0x1b, 0x0e, 0x05, 0xc2, 0x30, 0x1a, 0x45, 0xe9, 0xad, 0x2c, 0x6c, 0x68,
	0x46, 0x2e, 0xc0, 0xa1, 0xf4, 0x03, 0x2c, 0xa6, 0xfe, 0xa0, 0x69, 0xd3, 0x10, 0x20, 0x0c, 0x07,
	0xa3, 0x28, 0x3d, 0xbd, 0x8c, 0x59, 0xb0, 0xc3, 0x3a, 0x3b, 0xec, 0x43, 0x67, 0x27, 0xbb, 0xb3,
	0xa6, 0xfa, 0x6f, 0x57, 0x2d, 0x91, 0x72, 0x72, 0x66, 0xc0, 0xe6, 0xa0, 0x50, 0x94, 0x5b, 0x7a,
	0x27, 0xa3, 0x28, 0x1d, 0x64, 0x74, 0x53, 0xea, 0x08, 0x97, 0xbf, 0x4f, 0x08, 0xbd, 0x5a, 0xcf,
	0xf1, 0x4d, 0x33, 0x6a, 0xb0, 0xf4, 0x47, 0xe8, 0x26, 0x07, 0xe7, 0x36, 0xd5, 0x0c, 0x8c, 0xb6,
	0x48, 0x9f, 0xb3, 0x03, 0xc3, 0x0f, 0x89, 0x6d, 0x17, 0x82, 0x42, 0x7c, 0xbe, 0xd7, 0xd8, 0xc4,
	0x07, 0x95, 0x3c, 0xfa, 0xf6, 0xe7, 0xef, 0xf7, 0xc1, 0x83, 0xe4, 0x6e, 0x4f, 0xd8, 0xb6, 0xa1,
	0x3e, 0x8d, 0x1e, 0xd3, 0x39, 0x39, 0x7d, 0x2d, 0x54, 0xb1, 0x84, 0x89, 0x07, 0xd1, 0xb4, 0xdf,
	0x4e, 0x88, 0x55, 0x18, 0xc9, 0x1a, 0xd8, 0x5b, 0x57, 0xc6, 0xe3, 0x83, 0x90, 0x3e, 0x31, 0xad,
	0x1c, 0xd0, 0x9f, 0x11, 0x39, 0xeb, 0x49, 0x94, 0x4e, 0xd8, 0x51, 0x97, 0x9f, 0xf5, 0x5f, 0x98,
	0xf8, 0xfa, 0x7f, 0xc8, 0x38, 0xf3, 0xe2, 0xfa, 0xd3, 0xcb, 0x52, 0xe2, 0xa2, 0x9e, 0x79, 0x2a,
	0xf7, 0x9a, 0xbc, 0xd3, 0xe4, 0x47, 0xfc, 0xe3, 0xb3, 0x9b, 0x4d, 0x18, 0x4f, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0xb5, 0x57, 0xa9, 0x19, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComplianceIngesterClient is the client API for ComplianceIngester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComplianceIngesterClient interface {
	ProcessComplianceReport(ctx context.Context, in *compliance.Report, opts ...grpc.CallOption) (*empty.Empty, error)
	HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error)
	ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error)
}

type complianceIngesterClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceIngesterClient(cc grpc.ClientConnInterface) ComplianceIngesterClient {
	return &complianceIngesterClient{cc}
}

func (c *complianceIngesterClient) ProcessComplianceReport(ctx context.Context, in *compliance.Report, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngester/ProcessComplianceReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceIngesterClient) HandleEvent(ctx context.Context, in *event.EventMsg, opts ...grpc.CallOption) (*event.EventResponse, error) {
	out := new(event.EventResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngester/HandleEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceIngesterClient) ProjectUpdateStatus(ctx context.Context, in *ProjectUpdateStatusReq, opts ...grpc.CallOption) (*ProjectUpdateStatusResp, error) {
	out := new(ProjectUpdateStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngester/ProjectUpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceIngesterServer is the server API for ComplianceIngester service.
type ComplianceIngesterServer interface {
	ProcessComplianceReport(context.Context, *compliance.Report) (*empty.Empty, error)
	HandleEvent(context.Context, *event.EventMsg) (*event.EventResponse, error)
	ProjectUpdateStatus(context.Context, *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error)
}

// UnimplementedComplianceIngesterServer can be embedded to have forward compatible implementations.
type UnimplementedComplianceIngesterServer struct {
}

func (*UnimplementedComplianceIngesterServer) ProcessComplianceReport(ctx context.Context, req *compliance.Report) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessComplianceReport not implemented")
}
func (*UnimplementedComplianceIngesterServer) HandleEvent(ctx context.Context, req *event.EventMsg) (*event.EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleEvent not implemented")
}
func (*UnimplementedComplianceIngesterServer) ProjectUpdateStatus(ctx context.Context, req *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectUpdateStatus not implemented")
}

func RegisterComplianceIngesterServer(s *grpc.Server, srv ComplianceIngesterServer) {
	s.RegisterService(&_ComplianceIngester_serviceDesc, srv)
}

func _ComplianceIngester_ProcessComplianceReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(compliance.Report)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIngesterServer).ProcessComplianceReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngester/ProcessComplianceReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIngesterServer).ProcessComplianceReport(ctx, req.(*compliance.Report))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceIngester_HandleEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(event.EventMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIngesterServer).HandleEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngester/HandleEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIngesterServer).HandleEvent(ctx, req.(*event.EventMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceIngester_ProjectUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectUpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIngesterServer).ProjectUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.ingest.ingest.ComplianceIngester/ProjectUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIngesterServer).ProjectUpdateStatus(ctx, req.(*ProjectUpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComplianceIngester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.compliance.ingest.ingest.ComplianceIngester",
	HandlerType: (*ComplianceIngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessComplianceReport",
			Handler:    _ComplianceIngester_ProcessComplianceReport_Handler,
		},
		{
			MethodName: "HandleEvent",
			Handler:    _ComplianceIngester_HandleEvent_Handler,
		},
		{
			MethodName: "ProjectUpdateStatus",
			Handler:    _ComplianceIngester_ProjectUpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/compliance-service/ingest/ingest/compliance.proto",
}