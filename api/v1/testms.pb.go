// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testms.proto

package testms

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

type Coordinates struct {
	Lat                  float32  `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float32  `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinates) Reset()         { *m = Coordinates{} }
func (m *Coordinates) String() string { return proto.CompactTextString(m) }
func (*Coordinates) ProtoMessage()    {}
func (*Coordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_testms_00b063260b8fd49f, []int{0}
}
func (m *Coordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinates.Unmarshal(m, b)
}
func (m *Coordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinates.Marshal(b, m, deterministic)
}
func (dst *Coordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinates.Merge(dst, src)
}
func (m *Coordinates) XXX_Size() int {
	return xxx_messageInfo_Coordinates.Size(m)
}
func (m *Coordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinates.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinates proto.InternalMessageInfo

func (m *Coordinates) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Coordinates) GetLng() float32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type SaveRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string       `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string       `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Coords               *Coordinates `protobuf:"bytes,5,opt,name=coords,proto3" json:"coords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testms_00b063260b8fd49f, []int{1}
}
func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (dst *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(dst, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SaveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SaveRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *SaveRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *SaveRequest) GetCoords() *Coordinates {
	if m != nil {
		return m.Coords
	}
	return nil
}

type SaveResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_testms_00b063260b8fd49f, []int{2}
}
func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (dst *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(dst, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testms_00b063260b8fd49f, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string       `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string       `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Coords               *Coordinates `protobuf:"bytes,5,opt,name=coords,proto3" json:"coords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_testms_00b063260b8fd49f, []int{4}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetResponse) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *GetResponse) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *GetResponse) GetCoords() *Coordinates {
	if m != nil {
		return m.Coords
	}
	return nil
}

func init() {
	proto.RegisterType((*Coordinates)(nil), "Coordinates")
	proto.RegisterType((*SaveRequest)(nil), "SaveRequest")
	proto.RegisterType((*SaveResponse)(nil), "SaveResponse")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortDomainServiceClient is the client API for PortDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortDomainServiceClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type portDomainServiceClient struct {
	cc *grpc.ClientConn
}

func NewPortDomainServiceClient(cc *grpc.ClientConn) PortDomainServiceClient {
	return &portDomainServiceClient{cc}
}

func (c *portDomainServiceClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/PortDomainService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/PortDomainService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServiceServer is the server API for PortDomainService service.
type PortDomainServiceServer interface {
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterPortDomainServiceServer(s *grpc.Server, srv PortDomainServiceServer) {
	s.RegisterService(&_PortDomainService_serviceDesc, srv)
}

func _PortDomainService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortDomainService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomainService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortDomainService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortDomainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PortDomainService",
	HandlerType: (*PortDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _PortDomainService_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PortDomainService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testms.proto",
}

func init() { proto.RegisterFile("testms.proto", fileDescriptor_testms_00b063260b8fd49f) }

var fileDescriptor_testms_00b063260b8fd49f = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0xbb, 0x7f, 0xac, 0x74, 0x76, 0x15, 0x9d, 0x53, 0x28, 0x1e, 0x4a, 0x50, 0xe8, 0x29,
	0x60, 0x7d, 0x04, 0x85, 0x5e, 0x65, 0x7b, 0x17, 0xe2, 0xee, 0x50, 0x02, 0x6e, 0x52, 0x93, 0x69,
	0xa1, 0x0f, 0xe0, 0x7b, 0x4b, 0xd2, 0xd5, 0xee, 0xc5, 0xb3, 0xb7, 0xc9, 0x37, 0x4c, 0xe6, 0xf7,
	0x25, 0x50, 0x33, 0x05, 0xee, 0x83, 0xda, 0x79, 0xc7, 0x4e, 0x3e, 0x42, 0xf5, 0xec, 0x9c, 0xef,
	0x8c, 0xd5, 0x4c, 0x01, 0x6f, 0xa0, 0xf8, 0xd0, 0x2c, 0xb2, 0x45, 0xb6, 0xcc, 0x9b, 0x58, 0x26,
	0x62, 0xb7, 0x22, 0x1f, 0x88, 0xdd, 0xca, 0xaf, 0x0c, 0xaa, 0x8d, 0x3e, 0x50, 0x43, 0x9f, 0x7b,
	0x0a, 0x8c, 0xd7, 0x90, 0x9b, 0x2e, 0x8d, 0xcc, 0x9a, 0xdc, 0x74, 0x88, 0x50, 0x5a, 0xdd, 0x53,
	0x1a, 0x99, 0x35, 0xa9, 0x8e, 0xac, 0x35, 0x7c, 0x14, 0xc5, 0x89, 0xc5, 0x1a, 0x05, 0x5c, 0xb6,
	0x6e, 0x6f, 0xd9, 0x1f, 0x45, 0x99, 0xf0, 0xcf, 0x11, 0xef, 0x61, 0xda, 0xc6, 0x50, 0x41, 0x5c,
	0x2c, 0xb2, 0x65, 0xb5, 0xaa, 0xd5, 0x28, 0x63, 0x33, 0xf4, 0xa4, 0x84, 0xfa, 0x14, 0x23, 0xec,
	0x9c, 0x0d, 0xf4, 0xbb, 0x37, 0x3b, 0xef, 0x95, 0x77, 0x00, 0x6b, 0xe2, 0x3f, 0x92, 0x26, 0x93,
	0xd4, 0x1e, 0x6e, 0xf8, 0x27, 0x93, 0xd5, 0x1b, 0xdc, 0xbe, 0x3a, 0xcf, 0x2f, 0xae, 0xd7, 0xc6,
	0x6e, 0xc8, 0x1f, 0x4c, 0x4b, 0xf8, 0x00, 0x65, 0xd4, 0xc3, 0x5a, 0x8d, 0x1e, 0x7b, 0x7e, 0xa5,
	0xc6, 0xce, 0x72, 0x82, 0x12, 0x8a, 0x35, 0x31, 0x56, 0xea, 0xec, 0x39, 0xaf, 0xd5, 0xc8, 0x4a,
	0x4e, 0xde, 0xa7, 0xe9, 0xaf, 0x9f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xdd, 0xeb, 0xae,
	0xfb, 0x01, 0x00, 0x00,
}
