// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/repocreds/repocreds.proto

// Repository Service
//
// Repository Service API performs CRUD actions against repository resources

package repocreds

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/vathsalashetty96/argo-cd/pkg/apis/application/v1alpha1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "k8s.io/api/core/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RepoCredsQuery is a query for RepoCreds resources
type RepoCredsQuery struct {
	// Repo URL for query
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoCredsQuery) Reset()         { *m = RepoCredsQuery{} }
func (m *RepoCredsQuery) String() string { return proto.CompactTextString(m) }
func (*RepoCredsQuery) ProtoMessage()    {}
func (*RepoCredsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b5fce4710a8821, []int{0}
}
func (m *RepoCredsQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepoCredsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepoCredsQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepoCredsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoCredsQuery.Merge(m, src)
}
func (m *RepoCredsQuery) XXX_Size() int {
	return m.Size()
}
func (m *RepoCredsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoCredsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RepoCredsQuery proto.InternalMessageInfo

func (m *RepoCredsQuery) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type RepoCredsDeleteRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoCredsDeleteRequest) Reset()         { *m = RepoCredsDeleteRequest{} }
func (m *RepoCredsDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*RepoCredsDeleteRequest) ProtoMessage()    {}
func (*RepoCredsDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b5fce4710a8821, []int{1}
}
func (m *RepoCredsDeleteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepoCredsDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepoCredsDeleteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepoCredsDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoCredsDeleteRequest.Merge(m, src)
}
func (m *RepoCredsDeleteRequest) XXX_Size() int {
	return m.Size()
}
func (m *RepoCredsDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoCredsDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepoCredsDeleteRequest proto.InternalMessageInfo

func (m *RepoCredsDeleteRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// RepoCredsResponse is a response to most repository credentials requests
type RepoCredsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoCredsResponse) Reset()         { *m = RepoCredsResponse{} }
func (m *RepoCredsResponse) String() string { return proto.CompactTextString(m) }
func (*RepoCredsResponse) ProtoMessage()    {}
func (*RepoCredsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b5fce4710a8821, []int{2}
}
func (m *RepoCredsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepoCredsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepoCredsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepoCredsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoCredsResponse.Merge(m, src)
}
func (m *RepoCredsResponse) XXX_Size() int {
	return m.Size()
}
func (m *RepoCredsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoCredsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepoCredsResponse proto.InternalMessageInfo

// RepoCreateRequest is a request for creating repository credentials config
type RepoCredsCreateRequest struct {
	// Repository definition
	Creds *v1alpha1.RepoCreds `protobuf:"bytes,1,opt,name=creds,proto3" json:"creds,omitempty"`
	// Whether to create in upsert mode
	Upsert               bool     `protobuf:"varint,2,opt,name=upsert,proto3" json:"upsert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoCredsCreateRequest) Reset()         { *m = RepoCredsCreateRequest{} }
func (m *RepoCredsCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RepoCredsCreateRequest) ProtoMessage()    {}
func (*RepoCredsCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b5fce4710a8821, []int{3}
}
func (m *RepoCredsCreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepoCredsCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepoCredsCreateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepoCredsCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoCredsCreateRequest.Merge(m, src)
}
func (m *RepoCredsCreateRequest) XXX_Size() int {
	return m.Size()
}
func (m *RepoCredsCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoCredsCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepoCredsCreateRequest proto.InternalMessageInfo

func (m *RepoCredsCreateRequest) GetCreds() *v1alpha1.RepoCreds {
	if m != nil {
		return m.Creds
	}
	return nil
}

func (m *RepoCredsCreateRequest) GetUpsert() bool {
	if m != nil {
		return m.Upsert
	}
	return false
}

// RepoCredsUpdateRequest is a request for updating existing repository credentials config
type RepoCredsUpdateRequest struct {
	Creds                *v1alpha1.RepoCreds `protobuf:"bytes,1,opt,name=creds,proto3" json:"creds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RepoCredsUpdateRequest) Reset()         { *m = RepoCredsUpdateRequest{} }
func (m *RepoCredsUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*RepoCredsUpdateRequest) ProtoMessage()    {}
func (*RepoCredsUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0b5fce4710a8821, []int{4}
}
func (m *RepoCredsUpdateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RepoCredsUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RepoCredsUpdateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RepoCredsUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoCredsUpdateRequest.Merge(m, src)
}
func (m *RepoCredsUpdateRequest) XXX_Size() int {
	return m.Size()
}
func (m *RepoCredsUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoCredsUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepoCredsUpdateRequest proto.InternalMessageInfo

func (m *RepoCredsUpdateRequest) GetCreds() *v1alpha1.RepoCreds {
	if m != nil {
		return m.Creds
	}
	return nil
}

func init() {
	proto.RegisterType((*RepoCredsQuery)(nil), "repocreds.RepoCredsQuery")
	proto.RegisterType((*RepoCredsDeleteRequest)(nil), "repocreds.RepoCredsDeleteRequest")
	proto.RegisterType((*RepoCredsResponse)(nil), "repocreds.RepoCredsResponse")
	proto.RegisterType((*RepoCredsCreateRequest)(nil), "repocreds.RepoCredsCreateRequest")
	proto.RegisterType((*RepoCredsUpdateRequest)(nil), "repocreds.RepoCredsUpdateRequest")
}

func init() { proto.RegisterFile("server/repocreds/repocreds.proto", fileDescriptor_b0b5fce4710a8821) }

var fileDescriptor_b0b5fce4710a8821 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x49, 0xa5, 0xc5, 0x46, 0x90, 0x76, 0x2a, 0xb5, 0x3b, 0x5b, 0xd7, 0x35, 0x07, 0x29,
	0x45, 0x13, 0xb6, 0xbd, 0x88, 0x37, 0x6d, 0x0f, 0x0a, 0x5e, 0x1c, 0xf1, 0xd2, 0x8b, 0xa4, 0x33,
	0x8f, 0x69, 0xdc, 0x71, 0x12, 0x93, 0xcc, 0x40, 0x11, 0x11, 0x3c, 0x7b, 0xeb, 0xd1, 0x2f, 0xe0,
	0x07, 0xf0, 0x43, 0x78, 0x14, 0xfc, 0x02, 0xb2, 0xf8, 0x41, 0x24, 0xd9, 0x9d, 0x99, 0x1d, 0x76,
	0x0e, 0x7b, 0x58, 0x7a, 0x7b, 0x93, 0xbc, 0xf9, 0xbf, 0xdf, 0xcb, 0x3f, 0x2f, 0x78, 0x68, 0x40,
	0x97, 0xa0, 0x99, 0x06, 0x25, 0x63, 0x0d, 0x89, 0x69, 0x22, 0xaa, 0xb4, 0xb4, 0x32, 0xd8, 0xac,
	0x17, 0xc2, 0x3b, 0xa9, 0x4c, 0xa5, 0x5f, 0x65, 0x2e, 0x9a, 0x26, 0x84, 0xfb, 0xa9, 0x94, 0x69,
	0x06, 0x8c, 0x2b, 0xc1, 0x78, 0x9e, 0x4b, 0xcb, 0xad, 0x90, 0xf9, 0xec, 0xf7, 0x90, 0x8c, 0x9f,
	0x18, 0x2a, 0xa4, 0xdf, 0x8d, 0xa5, 0x06, 0x56, 0x8e, 0x58, 0x0a, 0x39, 0x68, 0x6e, 0x21, 0x99,
	0xe5, 0xbc, 0x4c, 0x85, 0xbd, 0x28, 0xce, 0x69, 0x2c, 0x3f, 0x30, 0xae, 0x7d, 0x89, 0xf7, 0x3e,
	0x78, 0x1c, 0x27, 0x4c, 0x8d, 0x53, 0xf7, 0xb3, 0x61, 0x5c, 0xa9, 0x4c, 0xc4, 0x5e, 0x9c, 0x95,
	0x23, 0x9e, 0xa9, 0x0b, 0xbe, 0x20, 0x45, 0x08, 0xbe, 0x1d, 0x81, 0x92, 0x27, 0x8e, 0xf7, 0x75,
	0x01, 0xfa, 0x32, 0xd8, 0xc2, 0x37, 0x0a, 0x9d, 0xed, 0xa1, 0x21, 0x3a, 0xd8, 0x8c, 0x5c, 0x48,
	0x0e, 0xf1, 0x6e, 0x9d, 0x73, 0x0a, 0x19, 0x58, 0x88, 0xe0, 0x63, 0x01, 0xc6, 0x76, 0xe4, 0xee,
	0xe0, 0xed, 0x3a, 0x37, 0x02, 0xa3, 0x64, 0x6e, 0x80, 0x7c, 0x43, 0x73, 0x0a, 0x27, 0x1a, 0x78,
	0xa3, 0x70, 0x86, 0xd7, 0xfd, 0x59, 0x79, 0x8d, 0x5b, 0x47, 0xa7, 0xb4, 0x69, 0x8d, 0x56, 0xad,
	0xf9, 0xe0, 0x5d, 0x9c, 0x50, 0x35, 0x4e, 0xa9, 0x6b, 0x8d, 0xce, 0xb5, 0x46, 0xab, 0xd6, 0x68,
	0x53, 0x77, 0x2a, 0x19, 0xec, 0xe2, 0x8d, 0x42, 0x19, 0xd0, 0x76, 0x6f, 0x6d, 0x88, 0x0e, 0x6e,
	0x46, 0xb3, 0x2f, 0x62, 0xe7, 0x68, 0xde, 0xaa, 0xe4, 0x7a, 0x68, 0x8e, 0xae, 0xd6, 0xf1, 0x56,
	0xbd, 0xf8, 0x06, 0x74, 0x29, 0x62, 0x08, 0xbe, 0x23, 0xdc, 0x7b, 0x25, 0x8c, 0x75, 0x1b, 0x46,
	0x58, 0xa9, 0x2f, 0xdd, 0x36, 0xe4, 0x56, 0xf0, 0xcc, 0x04, 0x3d, 0xda, 0x5c, 0xae, 0xb6, 0x4b,
	0xe1, 0x8b, 0x55, 0xa0, 0xb9, 0xca, 0xa4, 0xf7, 0xf5, 0xcf, 0xbf, 0xab, 0xb5, 0x9d, 0x60, 0xdb,
	0x5f, 0xb9, 0x72, 0xd4, 0x5c, 0xe8, 0xe0, 0x07, 0xc2, 0xfd, 0xca, 0xae, 0x2e, 0xbe, 0x07, 0x5d,
	0x7c, 0x2d, 0x7f, 0xc3, 0x95, 0x1c, 0x21, 0x19, 0x7a, 0xc6, 0x90, 0x2c, 0x32, 0x3e, 0x9d, 0x79,
	0xfd, 0x13, 0xe1, 0x7e, 0xe5, 0xe5, 0xd2, 0xa8, 0x2d, 0xf3, 0x57, 0x84, 0xfa, 0xc8, 0xa3, 0x3e,
	0x0c, 0xef, 0x2d, 0xa0, 0xb2, 0x4f, 0xd3, 0xf2, 0x85, 0xce, 0x3e, 0x57, 0xd8, 0x5f, 0x70, 0xbf,
	0x9a, 0xa8, 0xa5, 0xa9, 0x5b, 0x23, 0x18, 0xee, 0x77, 0xa5, 0xd4, 0x93, 0x77, 0xdf, 0xd3, 0xf4,
	0x0e, 0xef, 0x76, 0xd0, 0x38, 0x8e, 0xe7, 0xcf, 0x7e, 0x4d, 0x06, 0xe8, 0xf7, 0x64, 0x80, 0xfe,
	0x4e, 0x06, 0xe8, 0xec, 0x78, 0x89, 0x87, 0x25, 0xce, 0x04, 0xe4, 0xb6, 0x11, 0x3a, 0xdf, 0xf0,
	0x2f, 0xc9, 0xf1, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe4, 0x90, 0x83, 0x1b, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepoCredsServiceClient is the client API for RepoCredsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepoCredsServiceClient interface {
	// ListRepositoryCredentials gets a list of all configured repository credential sets
	ListRepositoryCredentials(ctx context.Context, in *RepoCredsQuery, opts ...grpc.CallOption) (*v1alpha1.RepoCredsList, error)
	// CreateRepositoryCredentials creates a new repository credential set
	CreateRepositoryCredentials(ctx context.Context, in *RepoCredsCreateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error)
	// UpdateRepositoryCredentials updates a repository credential set
	UpdateRepositoryCredentials(ctx context.Context, in *RepoCredsUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error)
	// DeleteRepositoryCredentials deletes a repository credential set from the configuration
	DeleteRepositoryCredentials(ctx context.Context, in *RepoCredsDeleteRequest, opts ...grpc.CallOption) (*RepoCredsResponse, error)
}

type repoCredsServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepoCredsServiceClient(cc *grpc.ClientConn) RepoCredsServiceClient {
	return &repoCredsServiceClient{cc}
}

func (c *repoCredsServiceClient) ListRepositoryCredentials(ctx context.Context, in *RepoCredsQuery, opts ...grpc.CallOption) (*v1alpha1.RepoCredsList, error) {
	out := new(v1alpha1.RepoCredsList)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/ListRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoCredsServiceClient) CreateRepositoryCredentials(ctx context.Context, in *RepoCredsCreateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error) {
	out := new(v1alpha1.RepoCreds)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/CreateRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoCredsServiceClient) UpdateRepositoryCredentials(ctx context.Context, in *RepoCredsUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error) {
	out := new(v1alpha1.RepoCreds)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/UpdateRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoCredsServiceClient) DeleteRepositoryCredentials(ctx context.Context, in *RepoCredsDeleteRequest, opts ...grpc.CallOption) (*RepoCredsResponse, error) {
	out := new(RepoCredsResponse)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/DeleteRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoCredsServiceServer is the server API for RepoCredsService service.
type RepoCredsServiceServer interface {
	// ListRepositoryCredentials gets a list of all configured repository credential sets
	ListRepositoryCredentials(context.Context, *RepoCredsQuery) (*v1alpha1.RepoCredsList, error)
	// CreateRepositoryCredentials creates a new repository credential set
	CreateRepositoryCredentials(context.Context, *RepoCredsCreateRequest) (*v1alpha1.RepoCreds, error)
	// UpdateRepositoryCredentials updates a repository credential set
	UpdateRepositoryCredentials(context.Context, *RepoCredsUpdateRequest) (*v1alpha1.RepoCreds, error)
	// DeleteRepositoryCredentials deletes a repository credential set from the configuration
	DeleteRepositoryCredentials(context.Context, *RepoCredsDeleteRequest) (*RepoCredsResponse, error)
}

// UnimplementedRepoCredsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepoCredsServiceServer struct {
}

func (*UnimplementedRepoCredsServiceServer) ListRepositoryCredentials(ctx context.Context, req *RepoCredsQuery) (*v1alpha1.RepoCredsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositoryCredentials not implemented")
}
func (*UnimplementedRepoCredsServiceServer) CreateRepositoryCredentials(ctx context.Context, req *RepoCredsCreateRequest) (*v1alpha1.RepoCreds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepositoryCredentials not implemented")
}
func (*UnimplementedRepoCredsServiceServer) UpdateRepositoryCredentials(ctx context.Context, req *RepoCredsUpdateRequest) (*v1alpha1.RepoCreds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepositoryCredentials not implemented")
}
func (*UnimplementedRepoCredsServiceServer) DeleteRepositoryCredentials(ctx context.Context, req *RepoCredsDeleteRequest) (*RepoCredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepositoryCredentials not implemented")
}

func RegisterRepoCredsServiceServer(s *grpc.Server, srv RepoCredsServiceServer) {
	s.RegisterService(&_RepoCredsService_serviceDesc, srv)
}

func _RepoCredsService_ListRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).ListRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/ListRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).ListRepositoryCredentials(ctx, req.(*RepoCredsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoCredsService_CreateRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).CreateRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/CreateRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).CreateRepositoryCredentials(ctx, req.(*RepoCredsCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoCredsService_UpdateRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).UpdateRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/UpdateRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).UpdateRepositoryCredentials(ctx, req.(*RepoCredsUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoCredsService_DeleteRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).DeleteRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/DeleteRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).DeleteRepositoryCredentials(ctx, req.(*RepoCredsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoCredsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repocreds.RepoCredsService",
	HandlerType: (*RepoCredsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepositoryCredentials",
			Handler:    _RepoCredsService_ListRepositoryCredentials_Handler,
		},
		{
			MethodName: "CreateRepositoryCredentials",
			Handler:    _RepoCredsService_CreateRepositoryCredentials_Handler,
		},
		{
			MethodName: "UpdateRepositoryCredentials",
			Handler:    _RepoCredsService_UpdateRepositoryCredentials_Handler,
		},
		{
			MethodName: "DeleteRepositoryCredentials",
			Handler:    _RepoCredsService_DeleteRepositoryCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/repocreds/repocreds.proto",
}

func (m *RepoCredsQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoCredsQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepoCredsQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintRepocreds(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RepoCredsDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoCredsDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepoCredsDeleteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintRepocreds(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RepoCredsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoCredsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepoCredsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *RepoCredsCreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoCredsCreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepoCredsCreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Upsert {
		i--
		if m.Upsert {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Creds != nil {
		{
			size, err := m.Creds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRepocreds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RepoCredsUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepoCredsUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RepoCredsUpdateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Creds != nil {
		{
			size, err := m.Creds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRepocreds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRepocreds(dAtA []byte, offset int, v uint64) int {
	offset -= sovRepocreds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RepoCredsQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovRepocreds(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RepoCredsDeleteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovRepocreds(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RepoCredsResponse) Size() (n int) {
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

func (m *RepoCredsCreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Creds != nil {
		l = m.Creds.Size()
		n += 1 + l + sovRepocreds(uint64(l))
	}
	if m.Upsert {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RepoCredsUpdateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Creds != nil {
		l = m.Creds.Size()
		n += 1 + l + sovRepocreds(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRepocreds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRepocreds(x uint64) (n int) {
	return sovRepocreds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RepoCredsQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepocreds
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
			return fmt.Errorf("proto: RepoCredsQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoCredsQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepocreds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRepocreds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRepocreds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRepocreds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepocreds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRepocreds
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
func (m *RepoCredsDeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepocreds
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
			return fmt.Errorf("proto: RepoCredsDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoCredsDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepocreds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRepocreds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRepocreds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRepocreds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepocreds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRepocreds
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
func (m *RepoCredsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepocreds
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
			return fmt.Errorf("proto: RepoCredsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoCredsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRepocreds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepocreds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRepocreds
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
func (m *RepoCredsCreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepocreds
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
			return fmt.Errorf("proto: RepoCredsCreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoCredsCreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepocreds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRepocreds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRepocreds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Creds == nil {
				m.Creds = &v1alpha1.RepoCreds{}
			}
			if err := m.Creds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upsert", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepocreds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Upsert = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRepocreds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepocreds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRepocreds
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
func (m *RepoCredsUpdateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepocreds
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
			return fmt.Errorf("proto: RepoCredsUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepoCredsUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepocreds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRepocreds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRepocreds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Creds == nil {
				m.Creds = &v1alpha1.RepoCreds{}
			}
			if err := m.Creds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRepocreds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRepocreds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRepocreds
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
func skipRepocreds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRepocreds
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
					return 0, ErrIntOverflowRepocreds
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
					return 0, ErrIntOverflowRepocreds
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
				return 0, ErrInvalidLengthRepocreds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRepocreds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRepocreds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRepocreds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRepocreds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRepocreds = fmt.Errorf("proto: unexpected end of group")
)
