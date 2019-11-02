// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package accounts

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateAccountRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateAccountResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(m, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "accounts.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "accounts.CreateAccountResponse")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x62,
	0xb8, 0x44, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x1d, 0x21, 0x22, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x60, 0xb6, 0x90, 0x14, 0x17, 0x47, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04,
	0x13, 0x58, 0x1c, 0xce, 0x17, 0x12, 0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x06,
	0x4b, 0x40, 0x38, 0x4a, 0xea, 0x5c, 0xa2, 0x68, 0xa6, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x40, 0x0d, 0x67, 0xca, 0x4c, 0x31, 0x8a, 0xe3, 0xe2, 0x80, 0x2a,
	0x29, 0x16, 0x0a, 0xe2, 0xe2, 0x45, 0xd1, 0x24, 0x24, 0xa7, 0x07, 0x77, 0x3e, 0x36, 0xb7, 0x4a,
	0xc9, 0xe3, 0x94, 0x87, 0xd8, 0xa6, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0xb7, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x10, 0xa8, 0x65, 0x70, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/accounts.Accounts/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
}

// UnimplementedAccountsServer can be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (*UnimplementedAccountsServer) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounts/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Accounts_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}