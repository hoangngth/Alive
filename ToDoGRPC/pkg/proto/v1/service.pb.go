// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package v1 is a generated protocol buffer package.

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
package v1

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
	return proto.EnumName(ToDoRequest_Status_name, int32(x))
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
	return proto.EnumName(ToDoResponse_Status_name, int32(x))
}
func (ToDoResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type ToDoRequest struct {
	Title       string             `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description string             `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Status      ToDoRequest_Status `protobuf:"varint,3,opt,name=status,enum=v1.ToDoRequest_Status" json:"status,omitempty"`
	Tag         string             `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
}

func (m *ToDoRequest) Reset()                    { *m = ToDoRequest{} }
func (m *ToDoRequest) String() string            { return proto.CompactTextString(m) }
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
	Title       string                     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Status      ToDoResponse_Status        `protobuf:"varint,4,opt,name=status,enum=v1.ToDoResponse_Status" json:"status,omitempty"`
	Tag         string                     `protobuf:"bytes,5,opt,name=tag" json:"tag,omitempty"`
	CreatedDate *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created_date,json=createdDate" json:"created_date,omitempty"`
}

func (m *ToDoResponse) Reset()                    { *m = ToDoResponse{} }
func (m *ToDoResponse) String() string            { return proto.CompactTextString(m) }
func (*ToDoResponse) ProtoMessage()               {}
func (*ToDoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ToDoResponse) GetId() int64 {
	if m != nil {
		return m.Id
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
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
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

type ReadResponse struct {
	Api  string        `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDo *ToDoResponse `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
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
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
}

func (m *ReadAllRequest) Reset()                    { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()               {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type ReadAllResponse struct {
	Api   string          `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDos []*ToDoResponse `protobuf:"bytes,2,rep,name=toDos" json:"toDos,omitempty"`
}

func (m *ReadAllResponse) Reset()                    { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string            { return proto.CompactTextString(m) }
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
	Api  string       `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDo *ToDoRequest `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
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

type CreateResponse struct {
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
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
	Api  string       `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	ToDo *ToDoRequest `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetToDo() *ToDoRequest {
	if m != nil {
		return m.ToDo
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
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
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

type DeleteResponse struct {
	Api     string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	Deleted bool   `protobuf:"varint,2,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
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
	proto.RegisterType((*ToDoRequest)(nil), "v1.ToDoRequest")
	proto.RegisterType((*ToDoResponse)(nil), "v1.ToDoResponse")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "v1.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "v1.ReadAllResponse")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v1.DeleteResponse")
	proto.RegisterEnum("v1.ToDoRequest_Status", ToDoRequest_Status_name, ToDoRequest_Status_value)
	proto.RegisterEnum("v1.ToDoResponse_Status", ToDoResponse_Status_name, ToDoResponse_Status_value)
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
	err := grpc.Invoke(ctx, "/v1.ToDoService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/ReadAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Delete", in, out, c.cc, opts...)
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
		FullMethod: "/v1.ToDoService/Read",
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
		FullMethod: "/v1.ToDoService/ReadAll",
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
		FullMethod: "/v1.ToDoService/Create",
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
		FullMethod: "/v1.ToDoService/Update",
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
		FullMethod: "/v1.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ToDoService",
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

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0x15, 0x68, 0xd9, 0xdd, 0x4b, 0x61, 0x71, 0x34, 0x4a, 0x78, 0x91, 0xa0, 0x31, 0xf5, 0x41,
	0x48, 0xf1, 0x75, 0x7d, 0x30, 0xcb, 0x6e, 0x34, 0x9a, 0xd6, 0xd0, 0xf5, 0xd9, 0xb0, 0x65, 0x6c,
	0x48, 0xd8, 0x1d, 0xec, 0x4c, 0xfb, 0x17, 0xfc, 0x47, 0xfe, 0x3d, 0x0d, 0xf3, 0x51, 0x4a, 0x2b,
	0x4d, 0xcc, 0xbe, 0x31, 0xf7, 0xde, 0x73, 0xef, 0x3d, 0xe7, 0x0c, 0x03, 0x36, 0xc5, 0xab, 0x4d,
	0xb9, 0xc0, 0x51, 0xbd, 0x22, 0x8c, 0x20, 0x7d, 0x33, 0xf1, 0x5f, 0x2c, 0x09, 0x59, 0x56, 0x38,
	0xe6, 0x91, 0xdb, 0xf5, 0x8f, 0x98, 0x95, 0x77, 0x98, 0xb2, 0xfc, 0xae, 0x16, 0x45, 0xe1, 0x6f,
	0x0d, 0xac, 0x1b, 0x92, 0x92, 0x0c, 0xff, 0x5c, 0x63, 0xca, 0xd0, 0x53, 0x18, 0xb2, 0x92, 0x55,
	0xd8, 0xd3, 0x02, 0x6d, 0x7c, 0x96, 0x89, 0x03, 0x0a, 0xc0, 0x2a, 0x30, 0x5d, 0xac, 0xca, 0x9a,
	0x95, 0xe4, 0xde, 0xd3, 0x79, 0x6e, 0x37, 0x84, 0x22, 0x30, 0x29, 0xcb, 0xd9, 0x9a, 0x7a, 0x46,
	0xa0, 0x8d, 0x9d, 0xe4, 0x59, 0xb4, 0x99, 0x44, 0x3b, 0x8d, 0xa3, 0x39, 0xcf, 0x66, 0xb2, 0x0a,
	0xb9, 0x60, 0xb0, 0x7c, 0xe9, 0x0d, 0x78, 0xa7, 0xe6, 0x33, 0x7c, 0x0b, 0xa6, 0xa8, 0x41, 0x00,
	0xe6, 0x6c, 0xfa, 0x71, 0xf6, 0x25, 0x75, 0x1f, 0x21, 0x1b, 0xce, 0x3e, 0x4d, 0xbf, 0x66, 0xb3,
	0xcb, 0xab, 0xf9, 0xdc, 0xd5, 0xd0, 0x29, 0x0c, 0xd2, 0xd9, 0xf4, 0xca, 0xd5, 0xc3, 0x5f, 0x3a,
	0x8c, 0x44, 0x7f, 0x5a, 0x93, 0x7b, 0x8a, 0x91, 0x03, 0x7a, 0x59, 0xf0, 0xb5, 0x8d, 0x4c, 0x2f,
	0x8b, 0x96, 0x89, 0x7e, 0x84, 0x89, 0x71, 0xc8, 0x24, 0xde, 0x32, 0x19, 0x70, 0x26, 0xcf, 0x5b,
	0x26, 0x62, 0x52, 0x0f, 0x95, 0xe1, 0x96, 0x0a, 0x7a, 0x0f, 0xa3, 0xc5, 0x0a, 0xe7, 0x0c, 0x17,
	0xdf, 0x8b, 0x9c, 0x61, 0xcf, 0x0c, 0xb4, 0xb1, 0x95, 0xf8, 0x91, 0x30, 0x23, 0x52, 0x66, 0x44,
	0x37, 0xca, 0x8c, 0xcc, 0x92, 0xf5, 0x69, 0xce, 0xf0, 0xff, 0x2a, 0x11, 0x83, 0x95, 0xe1, 0xbc,
	0x50, 0x0e, 0xba, 0x60, 0xe4, 0x75, 0x29, 0xfd, 0x6b, 0x3e, 0xa5, 0x32, 0xba, 0x52, 0x26, 0xbc,
	0x86, 0x91, 0x00, 0x48, 0xe5, 0x0e, 0x11, 0xaf, 0x60, 0xc0, 0x48, 0x4a, 0x38, 0xc6, 0x4a, 0xdc,
	0x7d, 0x05, 0x32, 0x9e, 0x0d, 0x43, 0x70, 0x9a, 0x3e, 0x1f, 0xaa, 0xaa, 0x77, 0x76, 0xf8, 0x19,
	0xce, 0xb7, 0x35, 0xbd, 0xe3, 0x5e, 0xc3, 0xb0, 0x69, 0x48, 0x3d, 0x3d, 0x30, 0xfe, 0x39, 0x4f,
	0xa4, 0xc3, 0x6b, 0xb0, 0x2f, 0xb9, 0x4e, 0xfd, 0x5c, 0x5f, 0x76, 0x36, 0x3f, 0xdf, 0xbb, 0x85,
	0x72, 0xf1, 0x04, 0x1c, 0xd5, 0xa7, 0x77, 0xa7, 0x43, 0xd1, 0xec, 0x6f, 0x75, 0xf1, 0xf0, 0xd9,
	0x17, 0xe0, 0xa8, 0x3e, 0xbd, 0xb3, 0x3d, 0x38, 0x59, 0xf3, 0x1a, 0xb1, 0xc0, 0x69, 0xa6, 0x8e,
	0xe1, 0x04, 0xec, 0x14, 0x57, 0xf8, 0xd8, 0x16, 0xfb, 0x8b, 0x5f, 0x80, 0xa3, 0x20, 0xc7, 0x06,
	0x16, 0xbc, 0x66, 0x3b, 0x50, 0x1e, 0x93, 0x3f, 0xf2, 0x7d, 0x98, 0x8b, 0xa7, 0x05, 0xbd, 0x81,
	0x41, 0xe3, 0x27, 0xe2, 0xec, 0x76, 0xae, 0x9d, 0xef, 0xb6, 0x01, 0x39, 0x26, 0x81, 0x13, 0x69,
	0x3d, 0x42, 0x2a, 0xd9, 0xde, 0x15, 0xff, 0x49, 0x27, 0x26, 0x31, 0x31, 0x98, 0xc2, 0x19, 0xf4,
	0xb8, 0x49, 0x77, 0xdc, 0xf6, 0xd1, 0x6e, 0xa8, 0x05, 0x08, 0x39, 0x05, 0xa0, 0x63, 0x91, 0x00,
	0xec, 0xa9, 0x1d, 0x83, 0x29, 0xe4, 0x10, 0x80, 0x8e, 0x9a, 0x02, 0xd0, 0x55, 0xeb, 0xd6, 0xe4,
	0xbf, 0xeb, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x2e, 0x3f, 0x6d, 0x5e, 0x05, 0x00,
	0x00,
}
