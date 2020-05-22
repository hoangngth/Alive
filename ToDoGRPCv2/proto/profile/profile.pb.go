// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile/profile.proto

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile/profile.proto

It has these top-level messages:
	ProfileRequest
	ProfileResponse
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
	ReadRequest
	ReadResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
*/
package profile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type ProfileRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phone    int64  `protobuf:"varint,5,opt,name=phone" json:"phone,omitempty"`
	Address  string `protobuf:"bytes,6,opt,name=address" json:"address,omitempty"`
}

func (m *ProfileRequest) Reset()                    { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()               {}
func (*ProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProfileRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ProfileRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProfileRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ProfileRequest) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *ProfileRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ProfileResponse struct {
	Id          int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email       string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone       int64                      `protobuf:"varint,4,opt,name=phone" json:"phone,omitempty"`
	Address     string                     `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	CreatedDate *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created_date,json=createdDate" json:"created_date,omitempty"`
}

func (m *ProfileResponse) Reset()                    { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()               {}
func (*ProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProfileResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProfileResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProfileResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ProfileResponse) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *ProfileResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProfileResponse) GetCreatedDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedDate
	}
	return nil
}

type LoginRequest struct {
	Api      string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Api      string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id       int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	LoggedIn bool   `protobuf:"varint,3,opt,name=loggedIn" json:"loggedIn,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LoginResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginResponse) GetLoggedIn() bool {
	if m != nil {
		return m.LoggedIn
	}
	return false
}

type LogoutRequest struct {
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogoutRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LogoutRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LogoutResponse struct {
	Api       string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	LoggedOut bool   `protobuf:"varint,2,opt,name=loggedOut" json:"loggedOut,omitempty"`
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogoutResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LogoutResponse) GetLoggedOut() bool {
	if m != nil {
		return m.LoggedOut
	}
	return false
}

type ReadRequest struct {
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Api     string           `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Profile *ProfileResponse `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetProfile() *ProfileResponse {
	if m != nil {
		return m.Profile
	}
	return nil
}

type CreateRequest struct {
	Api     string          `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Profile *ProfileRequest `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetProfile() *ProfileRequest {
	if m != nil {
		return m.Profile
	}
	return nil
}

type CreateResponse struct {
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateRequest struct {
	Api     string          `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id      int64           `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Profile *ProfileRequest `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetProfile() *ProfileRequest {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateResponse struct {
	Api     string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Updated bool   `protobuf:"varint,2,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() bool {
	if m != nil {
		return m.Updated
	}
	return false
}

type DeleteRequest struct {
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Api     string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Deleted bool   `protobuf:"varint,2,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func init() {
	proto.RegisterType((*ProfileRequest)(nil), "profile.ProfileRequest")
	proto.RegisterType((*ProfileResponse)(nil), "profile.ProfileResponse")
	proto.RegisterType((*LoginRequest)(nil), "profile.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "profile.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "profile.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "profile.LogoutResponse")
	proto.RegisterType((*ReadRequest)(nil), "profile.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "profile.ReadResponse")
	proto.RegisterType((*CreateRequest)(nil), "profile.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "profile.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "profile.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "profile.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "profile.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "profile.DeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProfileService service

type ProfileServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/profile.ProfileService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/profile.ProfileService/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/profile.ProfileService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/profile.ProfileService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/profile.ProfileService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/profile.ProfileService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfileService service

type ProfileServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ProfileService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ProfileService_Logout_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ProfileService_Read_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ProfileService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProfileService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile/profile.proto",
}

func init() { proto.RegisterFile("profile/profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6a, 0xd4, 0x50,
	0x14, 0x65, 0x92, 0xb9, 0x75, 0xcf, 0x45, 0x39, 0x74, 0x6c, 0x08, 0x82, 0x25, 0x4f, 0x7d, 0xca,
	0xd0, 0x14, 0x84, 0x01, 0x05, 0xc1, 0xbe, 0x08, 0x8a, 0x12, 0x47, 0xf0, 0x4d, 0x4e, 0x7b, 0x76,
	0x63, 0x60, 0x26, 0x27, 0x26, 0x27, 0xfa, 0x3f, 0x7e, 0x86, 0x7f, 0xe4, 0x5f, 0x48, 0xce, 0x25,
	0xb7, 0x69, 0x86, 0xf1, 0xa9, 0xb3, 0x77, 0xf6, 0xda, 0x6b, 0x75, 0xed, 0x95, 0xc0, 0x2a, 0xcd,
	0xf8, 0x43, 0xbc, 0xc3, 0xb5, 0xfe, 0xeb, 0xa7, 0x19, 0x17, 0x9c, 0x4c, 0x74, 0xe9, 0xbe, 0x88,
	0x38, 0x8f, 0xd4, 0x63, 0xc1, 0xef, 0x8a, 0x87, 0xb5, 0x88, 0xf7, 0x98, 0x0b, 0xba, 0x4f, 0xd5,
	0xa4, 0xf7, 0x7b, 0x00, 0xcb, 0x4f, 0x6a, 0x38, 0xc4, 0x1f, 0x05, 0xe6, 0x82, 0xb8, 0x30, 0x2d,
	0x72, 0xcc, 0x12, 0xba, 0x47, 0x67, 0x70, 0x39, 0xb8, 0x3a, 0x0b, 0xab, 0xba, 0x7c, 0x96, 0xd2,
	0x3c, 0xff, 0xc5, 0x33, 0xe6, 0x58, 0xea, 0x99, 0xa9, 0x09, 0x81, 0xa1, 0xc4, 0xd8, 0xb2, 0x2f,
	0x7f, 0x93, 0x73, 0x18, 0xe1, 0x9e, 0xc6, 0x3b, 0x67, 0x28, 0x9b, 0xaa, 0x28, 0xbb, 0xe9, 0x77,
	0x9e, 0xa0, 0x33, 0xba, 0x1c, 0x5c, 0xd9, 0xa1, 0x2a, 0x88, 0x03, 0x13, 0xca, 0x58, 0x86, 0x79,
	0xee, 0x8c, 0xe5, 0xb4, 0x29, 0xbd, 0x3f, 0x03, 0x78, 0x52, 0x89, 0xcc, 0x53, 0x9e, 0xe4, 0x48,
	0x96, 0x60, 0xc5, 0x4c, 0xea, 0xb3, 0x43, 0x2b, 0xae, 0xd9, 0xad, 0xc7, 0xd8, 0xed, 0x47, 0xd9,
	0x87, 0x3d, 0xec, 0xa3, 0x16, 0x3b, 0x79, 0x0d, 0xf3, 0xfb, 0x0c, 0xa9, 0x40, 0xf6, 0x8d, 0x51,
	0x81, 0x52, 0xdc, 0x2c, 0x70, 0x7d, 0x65, 0xad, 0x6f, 0xac, 0xf5, 0xb7, 0xc6, 0xda, 0x70, 0xa6,
	0xe7, 0x6f, 0xa9, 0x40, 0xef, 0x2b, 0xcc, 0xdf, 0xf3, 0x28, 0x4e, 0x8c, 0xbd, 0x4f, 0xc1, 0xa6,
	0x69, 0xac, 0x9d, 0x2d, 0x7f, 0xb6, 0x0c, 0xb7, 0x8e, 0x18, 0x6e, 0xb7, 0x0d, 0xf7, 0x3e, 0xc0,
	0x42, 0x6f, 0xd6, 0x9e, 0x1c, 0xae, 0x56, 0x2e, 0x59, 0x95, 0x4b, 0x2e, 0x4c, 0x77, 0x3c, 0x8a,
	0x90, 0xbd, 0x4b, 0xe4, 0xba, 0x69, 0x58, 0xd5, 0xde, 0xb5, 0x5c, 0xc7, 0x0b, 0xd1, 0xaf, 0xb4,
	0xb3, 0xce, 0x7b, 0x03, 0x4b, 0x03, 0xe9, 0x95, 0xf0, 0x1c, 0xce, 0x14, 0xc5, 0xc7, 0x42, 0x48,
	0xe8, 0x34, 0xac, 0x1b, 0xde, 0x1a, 0x66, 0x21, 0x52, 0x76, 0x3a, 0xe5, 0x16, 0xe6, 0x0a, 0xd0,
	0x4b, 0x18, 0x80, 0x89, 0xbf, 0x84, 0xcd, 0x02, 0xc7, 0x37, 0x6f, 0x47, 0x27, 0x44, 0xa1, 0x19,
	0xf4, 0xb6, 0xb0, 0x78, 0x2b, 0x6f, 0xd6, 0x2f, 0xe4, 0xba, 0xbb, 0xf6, 0xe2, 0x70, 0xad, 0xc4,
	0xd6, 0x5b, 0x03, 0x58, 0x9a, 0xad, 0xa7, 0x5e, 0xc8, 0x63, 0xb0, 0xf8, 0x92, 0xb2, 0xa3, 0x4a,
	0xba, 0x47, 0x6d, 0x28, 0xb3, 0x4f, 0x54, 0xf6, 0x0a, 0x96, 0x86, 0xa5, 0x57, 0x99, 0x03, 0x93,
	0x42, 0xce, 0x30, 0x7d, 0x36, 0x53, 0x96, 0x49, 0xb9, 0xc5, 0x1d, 0xfe, 0x87, 0xc6, 0x92, 0xd0,
	0x40, 0x8e, 0x11, 0x32, 0x39, 0x53, 0x11, 0xea, 0x32, 0xf8, 0x6b, 0x55, 0x5f, 0xa9, 0xcf, 0x98,
	0xfd, 0x8c, 0xef, 0x91, 0xbc, 0x84, 0x91, 0x0c, 0x3f, 0x59, 0x55, 0xff, 0x6c, 0xf3, 0x35, 0x73,
	0x9f, 0x75, 0xdb, 0x9a, 0x76, 0x03, 0x63, 0x15, 0x59, 0xd2, 0x9a, 0xa8, 0x63, 0xef, 0x5e, 0x1c,
	0xf4, 0x35, 0xf4, 0x06, 0x86, 0x65, 0xf4, 0xc8, 0x79, 0x35, 0xd0, 0x88, 0xae, 0xbb, 0xea, 0x74,
	0x6b, 0x3e, 0x95, 0x81, 0x06, 0x5f, 0x2b, 0x6a, 0x0d, 0xbe, 0x4e, 0x58, 0x36, 0x30, 0x56, 0x47,
	0x6a, 0x40, 0x5b, 0xd9, 0x68, 0x40, 0x3b, 0xd7, 0xdc, 0xc0, 0x58, 0xd9, 0xdd, 0x80, 0xb6, 0x4e,
	0xd6, 0x80, 0xb6, 0xef, 0x72, 0x37, 0x96, 0x1f, 0xb4, 0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0xc8, 0xad, 0x87, 0x5b, 0x06, 0x00, 0x00,
}