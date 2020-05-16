// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	ToDoRequest
	ToDoResponse
	ReadRequest
	ReadResponse
	ReadAllRequest
	ReadAllResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ToDoRequest_Status int32

const (
	ToDoRequest_ONHOLD    ToDoRequest_Status = 0
	ToDoRequest_INPROCESS ToDoRequest_Status = 1
	ToDoRequest_DONE      ToDoRequest_Status = 2
)

var ToDoRequest_Status_name = map[int32]string{
	0: "ONHOLD",
	1: "INPROCESS",
	2: "DONE",
}
var ToDoRequest_Status_value = map[string]int32{
	"ONHOLD":    0,
	"INPROCESS": 1,
	"DONE":      2,
}

func (x ToDoRequest_Status) String() string {
	return proto1.EnumName(ToDoRequest_Status_name, int32(x))
}
func (ToDoRequest_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ToDoResponse_Status int32

const (
	ToDoResponse_ONHOLD    ToDoResponse_Status = 0
	ToDoResponse_INPROCESS ToDoResponse_Status = 1
	ToDoResponse_DONE      ToDoResponse_Status = 2
)

var ToDoResponse_Status_name = map[int32]string{
	0: "ONHOLD",
	1: "INPROCESS",
	2: "DONE",
}
var ToDoResponse_Status_value = map[string]int32{
	"ONHOLD":    0,
	"INPROCESS": 1,
	"DONE":      2,
}

func (x ToDoResponse_Status) String() string {
	return proto1.EnumName(ToDoResponse_Status_name, int32(x))
}
func (ToDoResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type ToDoRequest struct {
	Title       string             `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description string             `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Status      ToDoRequest_Status `protobuf:"varint,3,opt,name=status,enum=proto.ToDoRequest_Status" json:"status,omitempty"`
	Tag         string             `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
}

func (m *ToDoRequest) Reset()                    { *m = ToDoRequest{} }
func (m *ToDoRequest) String() string            { return proto1.CompactTextString(m) }
func (*ToDoRequest) ProtoMessage()               {}
func (*ToDoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ToDoRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDoRequest) GetStatus() ToDoRequest_Status {
	if m != nil {
		return m.Status
	}
	return ToDoRequest_ONHOLD
}

func (m *ToDoRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ToDoResponse struct {
	Id          int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId      int64                      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Title       string                     `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Description string                     `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Status      ToDoResponse_Status        `protobuf:"varint,5,opt,name=status,enum=proto.ToDoResponse_Status" json:"status,omitempty"`
	Tag         string                     `protobuf:"bytes,6,opt,name=tag" json:"tag,omitempty"`
	CreatedDate *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=created_date,json=createdDate" json:"created_date,omitempty"`
}

func (m *ToDoResponse) Reset()                    { *m = ToDoResponse{} }
func (m *ToDoResponse) String() string            { return proto1.CompactTextString(m) }
func (*ToDoResponse) ProtoMessage()               {}
func (*ToDoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ToDoResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDoResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ToDoResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDoResponse) GetStatus() ToDoResponse_Status {
	if m != nil {
		return m.Status
	}
	return ToDoResponse_ONHOLD
}

func (m *ToDoResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ToDoResponse) GetCreatedDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedDate
	}
	return nil
}

type ReadRequest struct {
	Api    string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id     int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	UserId int64  `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto1.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *ReadRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ReadResponse struct {
	Api  string        `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDo *ToDoResponse `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto1.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetToDo() *ToDoResponse {
	if m != nil {
		return m.ToDo
	}
	return nil
}

type ReadAllRequest struct {
	Api    string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
}

func (m *ReadAllRequest) Reset()                    { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string            { return proto1.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()               {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ReadAllResponse struct {
	Api   string          `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDos []*ToDoResponse `protobuf:"bytes,2,rep,name=toDos" json:"toDos,omitempty"`
}

func (m *ReadAllResponse) Reset()                    { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string            { return proto1.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()               {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetToDos() []*ToDoResponse {
	if m != nil {
		return m.ToDos
	}
	return nil
}

type CreateRequest struct {
	Api    string       `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDo   *ToDoRequest `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
	UserId int64        `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetToDo() *ToDoRequest {
	if m != nil {
		return m.ToDo
	}
	return nil
}

func (m *CreateRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateResponse struct {
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
	Api    string       `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id     int64        `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	ToDo   *ToDoRequest `protobuf:"bytes,3,opt,name=toDo" json:"toDo,omitempty"`
	UserId int64        `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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

func (m *UpdateRequest) GetToDo() *ToDoRequest {
	if m != nil {
		return m.ToDo
	}
	return nil
}

func (m *UpdateRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UpdateResponse struct {
	Api     string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Updated bool   `protobuf:"varint,2,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
	Api    string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id     int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	UserId int64  `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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

func (m *DeleteRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DeleteResponse struct {
	Api     string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Deleted bool   `protobuf:"varint,2,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
	proto1.RegisterType((*ToDoRequest)(nil), "proto.ToDoRequest")
	proto1.RegisterType((*ToDoResponse)(nil), "proto.ToDoResponse")
	proto1.RegisterType((*ReadRequest)(nil), "proto.ReadRequest")
	proto1.RegisterType((*ReadResponse)(nil), "proto.ReadResponse")
	proto1.RegisterType((*ReadAllRequest)(nil), "proto.ReadAllRequest")
	proto1.RegisterType((*ReadAllResponse)(nil), "proto.ReadAllResponse")
	proto1.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto1.RegisterType((*CreateResponse)(nil), "proto.CreateResponse")
	proto1.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto1.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
	proto1.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
	proto1.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
	proto1.RegisterEnum("proto.ToDoRequest_Status", ToDoRequest_Status_name, ToDoRequest_Status_value)
	proto1.RegisterEnum("proto.ToDoResponse_Status", ToDoResponse_Status_name, ToDoResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ToDoService service

type ToDoServiceClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type toDoServiceClient struct {
	cc *grpc.ClientConn
}

func NewToDoServiceClient(cc *grpc.ClientConn) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/proto.ToDoService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := grpc.Invoke(ctx, "/proto.ToDoService/ReadAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/proto.ToDoService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/proto.ToDoService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/proto.ToDoService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ToDoService service

type ToDoServiceServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterToDoServiceServer(s *grpc.Server, srv ToDoServiceServer) {
	s.RegisterService(&_ToDoService_serviceDesc, srv)
}

func _ToDoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToDoService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _ToDoService_Read_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ToDoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xfe, 0x7d, 0x89, 0xd3, 0x1e, 0x27, 0xfe, 0xa3, 0xe9, 0x45, 0xc6, 0x1b, 0x22, 0x2f, 0x20,
	0x2c, 0x70, 0x84, 0x51, 0x25, 0x84, 0xca, 0x02, 0xd5, 0x15, 0x44, 0x42, 0x09, 0x9a, 0x94, 0x35,
	0x72, 0xe3, 0x21, 0xb2, 0x94, 0x76, 0xdc, 0xcc, 0x84, 0xb7, 0x61, 0xc3, 0x5b, 0xf0, 0x76, 0x28,
	0x33, 0x63, 0xc7, 0x76, 0x3a, 0x11, 0x5d, 0xc5, 0x33, 0xfe, 0xce, 0x39, 0xdf, 0xe5, 0xc4, 0xd0,
	0x67, 0x64, 0xfd, 0x33, 0x5f, 0x90, 0xa8, 0x58, 0x53, 0x4e, 0x51, 0x47, 0xfc, 0x04, 0xcf, 0x97,
	0x94, 0x2e, 0x57, 0x64, 0x2c, 0x4e, 0xb7, 0x9b, 0x1f, 0x63, 0x9e, 0xdf, 0x11, 0xc6, 0xd3, 0xbb,
	0x42, 0xe2, 0xc2, 0x3f, 0x06, 0xb8, 0x37, 0x34, 0xa1, 0x98, 0x3c, 0x6c, 0x08, 0xe3, 0xe8, 0x14,
	0x3a, 0x3c, 0xe7, 0x2b, 0xe2, 0x1b, 0x43, 0x63, 0x74, 0x8c, 0xe5, 0x01, 0x0d, 0xc1, 0xcd, 0x08,
	0x5b, 0xac, 0xf3, 0x82, 0xe7, 0xf4, 0xde, 0x37, 0xc5, 0xbb, 0xfa, 0x15, 0x7a, 0x03, 0x0e, 0xe3,
	0x29, 0xdf, 0x30, 0xdf, 0x1a, 0x1a, 0x23, 0x2f, 0x7e, 0x26, 0xfb, 0x47, 0xb5, 0xde, 0xd1, 0x5c,
	0x00, 0xb0, 0x02, 0xa2, 0x01, 0x58, 0x3c, 0x5d, 0xfa, 0xb6, 0x68, 0xb6, 0x7d, 0x0c, 0x5f, 0x83,
	0x23, 0x31, 0x08, 0xc0, 0x99, 0x4d, 0x3f, 0xcf, 0xbe, 0x24, 0x83, 0xff, 0x50, 0x1f, 0x8e, 0x27,
	0xd3, 0xaf, 0x78, 0x76, 0x75, 0x3d, 0x9f, 0x0f, 0x0c, 0x74, 0x04, 0x76, 0x32, 0x9b, 0x5e, 0x0f,
	0xcc, 0xf0, 0xb7, 0x09, 0x3d, 0xd9, 0x9f, 0x15, 0xf4, 0x9e, 0x11, 0xe4, 0x81, 0x99, 0x67, 0x82,
	0xb9, 0x85, 0xcd, 0x3c, 0x43, 0xe7, 0xe0, 0x6c, 0x18, 0x59, 0x4f, 0x32, 0xc1, 0xd8, 0xc2, 0xea,
	0xb4, 0x13, 0x69, 0x1d, 0x10, 0x69, 0xef, 0x8b, 0x8c, 0x2b, 0x91, 0x1d, 0x21, 0x32, 0x68, 0x88,
	0x94, 0x24, 0x34, 0x2a, 0x9d, 0x4a, 0x25, 0xfa, 0x00, 0xbd, 0xc5, 0x9a, 0xa4, 0x9c, 0x64, 0xdf,
	0xb3, 0x94, 0x13, 0xbf, 0x3b, 0x34, 0x46, 0x6e, 0x1c, 0x44, 0x32, 0xaa, 0xa8, 0x8c, 0x2a, 0xba,
	0x29, 0xa3, 0xc2, 0xae, 0xc2, 0x27, 0x29, 0x27, 0x4f, 0x35, 0xe9, 0x13, 0xb8, 0x98, 0xa4, 0x59,
	0x99, 0xef, 0x00, 0xac, 0xb4, 0xc8, 0x55, 0xba, 0xdb, 0x47, 0x65, 0x9a, 0xf9, 0x88, 0x69, 0x56,
	0xdd, 0xb4, 0x70, 0x02, 0x3d, 0xd9, 0x48, 0x99, 0xbd, 0xdf, 0xe9, 0x25, 0xd8, 0x9c, 0x26, 0x54,
	0xf4, 0x72, 0xe3, 0x93, 0x47, 0xcc, 0xc1, 0x02, 0x10, 0xbe, 0x07, 0x6f, 0xdb, 0xea, 0xe3, 0x6a,
	0xa5, 0xa7, 0xa5, 0xc9, 0x2e, 0x9c, 0xc2, 0xff, 0x55, 0xad, 0x96, 0xc9, 0x2b, 0xe8, 0x6c, 0x07,
	0x31, 0xdf, 0x1c, 0x5a, 0x3a, 0x2a, 0x12, 0x11, 0xa6, 0xd0, 0xbf, 0x12, 0xee, 0xea, 0xa9, 0xbc,
	0x68, 0xe8, 0x42, 0xfb, 0x9b, 0x2d, 0x65, 0x69, 0x9d, 0x8b, 0xc1, 0x2b, 0x47, 0x68, 0x19, 0xb7,
	0x52, 0x08, 0x1f, 0xa0, 0xff, 0xad, 0xc8, 0x0e, 0xd2, 0x6a, 0x07, 0x57, 0xd2, 0xb4, 0xfe, 0x99,
	0xa6, 0xdd, 0xa0, 0x79, 0x09, 0x5e, 0x39, 0x52, 0x4b, 0xd3, 0x87, 0xee, 0x46, 0x60, 0xe4, 0xe0,
	0x23, 0x5c, 0x1e, 0xc3, 0x09, 0xf4, 0x13, 0xb2, 0x22, 0x4f, 0x21, 0xac, 0xf3, 0xeb, 0x12, 0xbc,
	0xb2, 0xd5, 0x21, 0x22, 0x99, 0xc0, 0x54, 0x44, 0xd4, 0x31, 0xfe, 0x65, 0xca, 0x2f, 0xda, 0x5c,
	0x7e, 0x0f, 0xd1, 0x18, 0xec, 0xed, 0xc2, 0xa0, 0xd2, 0x90, 0xda, 0xbf, 0x21, 0x38, 0x69, 0xdc,
	0xa9, 0x61, 0xef, 0xa0, 0xab, 0x36, 0x0c, 0x9d, 0xd5, 0xde, 0xef, 0xb6, 0x35, 0x38, 0x6f, 0x5f,
	0xab, 0xca, 0x0b, 0x70, 0x64, 0xd0, 0xe8, 0x54, 0x21, 0x1a, 0xab, 0x15, 0x9c, 0xb5, 0x6e, 0x77,
	0x65, 0xd2, 0xf8, 0xaa, 0xac, 0x11, 0x7d, 0x55, 0xd6, 0x4a, 0xe7, 0x02, 0x1c, 0x69, 0x53, 0x55,
	0xd6, 0x08, 0xa0, 0x2a, 0x6b, 0x7a, 0x79, 0xeb, 0x88, 0xdb, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0xe7, 0x12, 0x24, 0x31, 0x06, 0x00, 0x00,
}