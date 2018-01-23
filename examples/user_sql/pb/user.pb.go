// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/user.proto

It has these top-level messages:
	User
	Friends
	SliceStringParam
	FriendsReq
	Empty
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/tcncloud/protoc-gen-persist/persist"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type User struct {
	Id        int64                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Friends   *Friends                    `protobuf:"bytes,3,opt,name=friends" json:"friends,omitempty"`
	CreatedOn *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=created_on,json=createdOn" json:"created_on,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetFriends() *Friends {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *User) GetCreatedOn() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

type Friends struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *Friends) Reset()                    { *m = Friends{} }
func (m *Friends) String() string            { return proto.CompactTextString(m) }
func (*Friends) ProtoMessage()               {}
func (*Friends) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Friends) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type SliceStringParam struct {
	Slice []string `protobuf:"bytes,1,rep,name=slice" json:"slice,omitempty"`
}

func (m *SliceStringParam) Reset()                    { *m = SliceStringParam{} }
func (m *SliceStringParam) String() string            { return proto.CompactTextString(m) }
func (*SliceStringParam) ProtoMessage()               {}
func (*SliceStringParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SliceStringParam) GetSlice() []string {
	if m != nil {
		return m.Slice
	}
	return nil
}

type FriendsReq struct {
	Names *SliceStringParam `protobuf:"bytes,2,opt,name=names" json:"names,omitempty"`
}

func (m *FriendsReq) Reset()                    { *m = FriendsReq{} }
func (m *FriendsReq) String() string            { return proto.CompactTextString(m) }
func (*FriendsReq) ProtoMessage()               {}
func (*FriendsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FriendsReq) GetNames() *SliceStringParam {
	if m != nil {
		return m.Names
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Friends)(nil), "pb.Friends")
	proto.RegisterType((*SliceStringParam)(nil), "pb.SliceStringParam")
	proto.RegisterType((*FriendsReq)(nil), "pb.FriendsReq")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UServ service

type UServClient interface {
	CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	InsertUsers(ctx context.Context, opts ...grpc.CallOption) (UServ_InsertUsersClient, error)
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_GetAllUsersClient, error)
	SelectUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserNames(ctx context.Context, opts ...grpc.CallOption) (UServ_UpdateUserNamesClient, error)
	UpdateNameToFoo(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	UpdateAllNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_UpdateAllNamesClient, error)
	GetFriends(ctx context.Context, in *FriendsReq, opts ...grpc.CallOption) (UServ_GetFriendsClient, error)
	DropTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type uServClient struct {
	cc *grpc.ClientConn
}

func NewUServClient(cc *grpc.ClientConn) UServClient {
	return &uServClient{cc}
}

func (c *uServClient) CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.UServ/CreateTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uServClient) InsertUsers(ctx context.Context, opts ...grpc.CallOption) (UServ_InsertUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UServ_serviceDesc.Streams[0], c.cc, "/pb.UServ/InsertUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServInsertUsersClient{stream}
	return x, nil
}

type UServ_InsertUsersClient interface {
	Send(*User) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type uServInsertUsersClient struct {
	grpc.ClientStream
}

func (x *uServInsertUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uServInsertUsersClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_GetAllUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UServ_serviceDesc.Streams[1], c.cc, "/pb.UServ/GetAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServGetAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_GetAllUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServGetAllUsersClient struct {
	grpc.ClientStream
}

func (x *uServGetAllUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) SelectUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/pb.UServ/SelectUserById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uServClient) UpdateUserNames(ctx context.Context, opts ...grpc.CallOption) (UServ_UpdateUserNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UServ_serviceDesc.Streams[2], c.cc, "/pb.UServ/UpdateUserNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServUpdateUserNamesClient{stream}
	return x, nil
}

type UServ_UpdateUserNamesClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type uServUpdateUserNamesClient struct {
	grpc.ClientStream
}

func (x *uServUpdateUserNamesClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uServUpdateUserNamesClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) UpdateNameToFoo(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.UServ/UpdateNameToFoo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uServClient) UpdateAllNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_UpdateAllNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UServ_serviceDesc.Streams[3], c.cc, "/pb.UServ/UpdateAllNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServUpdateAllNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_UpdateAllNamesClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServUpdateAllNamesClient struct {
	grpc.ClientStream
}

func (x *uServUpdateAllNamesClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) GetFriends(ctx context.Context, in *FriendsReq, opts ...grpc.CallOption) (UServ_GetFriendsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UServ_serviceDesc.Streams[4], c.cc, "/pb.UServ/GetFriends", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServGetFriendsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_GetFriendsClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServGetFriendsClient struct {
	grpc.ClientStream
}

func (x *uServGetFriendsClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) DropTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.UServ/DropTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UServ service

type UServServer interface {
	CreateTable(context.Context, *Empty) (*Empty, error)
	InsertUsers(UServ_InsertUsersServer) error
	GetAllUsers(*Empty, UServ_GetAllUsersServer) error
	SelectUserById(context.Context, *User) (*User, error)
	UpdateUserNames(UServ_UpdateUserNamesServer) error
	UpdateNameToFoo(context.Context, *User) (*Empty, error)
	UpdateAllNames(*Empty, UServ_UpdateAllNamesServer) error
	GetFriends(*FriendsReq, UServ_GetFriendsServer) error
	DropTable(context.Context, *Empty) (*Empty, error)
}

func RegisterUServServer(s *grpc.Server, srv UServServer) {
	s.RegisterService(&_UServ_serviceDesc, srv)
}

func _UServ_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).CreateTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UServ_InsertUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UServServer).InsertUsers(&uServInsertUsersServer{stream})
}

type UServ_InsertUsersServer interface {
	SendAndClose(*Empty) error
	Recv() (*User, error)
	grpc.ServerStream
}

type uServInsertUsersServer struct {
	grpc.ServerStream
}

func (x *uServInsertUsersServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uServInsertUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UServ_GetAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).GetAllUsers(m, &uServGetAllUsersServer{stream})
}

type UServ_GetAllUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServGetAllUsersServer struct {
	grpc.ServerStream
}

func (x *uServGetAllUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_SelectUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).SelectUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/SelectUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).SelectUserById(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UServ_UpdateUserNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UServServer).UpdateUserNames(&uServUpdateUserNamesServer{stream})
}

type UServ_UpdateUserNamesServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type uServUpdateUserNamesServer struct {
	grpc.ServerStream
}

func (x *uServUpdateUserNamesServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uServUpdateUserNamesServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UServ_UpdateNameToFoo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).UpdateNameToFoo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/UpdateNameToFoo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).UpdateNameToFoo(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UServ_UpdateAllNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).UpdateAllNames(m, &uServUpdateAllNamesServer{stream})
}

type UServ_UpdateAllNamesServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServUpdateAllNamesServer struct {
	grpc.ServerStream
}

func (x *uServUpdateAllNamesServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_GetFriends_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FriendsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).GetFriends(m, &uServGetFriendsServer{stream})
}

type UServ_GetFriendsServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServGetFriendsServer struct {
	grpc.ServerStream
}

func (x *uServGetFriendsServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_DropTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).DropTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/DropTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).DropTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UServ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UServ",
	HandlerType: (*UServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _UServ_CreateTable_Handler,
		},
		{
			MethodName: "SelectUserById",
			Handler:    _UServ_SelectUserById_Handler,
		},
		{
			MethodName: "UpdateNameToFoo",
			Handler:    _UServ_UpdateNameToFoo_Handler,
		},
		{
			MethodName: "DropTable",
			Handler:    _UServ_DropTable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InsertUsers",
			Handler:       _UServ_InsertUsers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetAllUsers",
			Handler:       _UServ_GetAllUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateUserNames",
			Handler:       _UServ_UpdateUserNames_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateAllNames",
			Handler:       _UServ_UpdateAllNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFriends",
			Handler:       _UServ_GetFriends_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/user.proto",
}

func init() { proto.RegisterFile("pb/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x2d, 0x95, 0x64, 0x99, 0xaf, 0xb1, 0x2c, 0x20, 0x32, 0xd4, 0xf0, 0x4b, 0x09, 0xc1, 0x43,
	0xe4, 0xce, 0xf5, 0x57, 0xb7, 0x61, 0x1f, 0xe8, 0x83, 0xe2, 0x28, 0x89, 0xd0, 0xc4, 0x29, 0x68,
	0xb9, 0x83, 0x81, 0xa1, 0x85, 0x3e, 0x68, 0x4f, 0xa8, 0x2c, 0x2a, 0x12, 0x5d, 0xac, 0xc0, 0x9e,
	0xfa, 0x38, 0x0c, 0x43, 0xf3, 0x4f, 0xf6, 0x34, 0xe4, 0x97, 0xec, 0x6d, 0x0f, 0xfb, 0x27, 0x03,
	0x45, 0x2b, 0xd6, 0x06, 0xcf, 0xc0, 0x9a, 0x37, 0x92, 0x3a, 0x3a, 0xe7, 0xdc, 0x7b, 0x0f, 0x09,
	0x1f, 0x25, 0x5e, 0x67, 0x91, 0xb1, 0xb4, 0x9d, 0xa4, 0x5c, 0x70, 0xac, 0x25, 0x5e, 0xfd, 0x93,
	0x84, 0xa5, 0x59, 0x98, 0x89, 0x0e, 0x4f, 0x44, 0xc8, 0xe3, 0x4c, 0x7d, 0xaa, 0x3f, 0x98, 0x71,
	0x3e, 0x8b, 0x58, 0x27, 0xdf, 0x79, 0x8b, 0x69, 0x47, 0x84, 0x73, 0x96, 0x09, 0x77, 0x9e, 0x28,
	0x80, 0xfe, 0x0b, 0x82, 0xed, 0x71, 0xc6, 0x52, 0xbc, 0x07, 0x5a, 0x18, 0xd4, 0x10, 0x41, 0xc6,
	0x16, 0xd5, 0xc2, 0x00, 0x63, 0xd8, 0x8e, 0xdd, 0x39, 0xab, 0x69, 0x04, 0x19, 0x15, 0x9a, 0xaf,
	0xf1, 0xa7, 0xb0, 0x3b, 0x4d, 0x43, 0x16, 0x07, 0x59, 0x6d, 0x8b, 0x20, 0xa3, 0xda, 0xaf, 0xb6,
	0x13, 0xaf, 0x7d, 0xa2, 0x8e, 0x68, 0xf1, 0x0d, 0x7f, 0x0d, 0xe0, 0xa7, 0xcc, 0x15, 0x2c, 0x78,
	0xc9, 0xe3, 0xda, 0x76, 0x8e, 0xac, 0xb7, 0x95, 0x93, 0x76, 0xe1, 0xa4, 0xed, 0x14, 0x4e, 0x68,
	0x65, 0x89, 0xbe, 0x8c, 0xf5, 0x07, 0xb0, 0xbb, 0xa4, 0xc3, 0x07, 0xb0, 0x23, 0x45, 0xb3, 0x1a,
	0x22, 0x5b, 0x46, 0x85, 0xaa, 0x8d, 0x6e, 0xc0, 0xfe, 0x28, 0x0a, 0x7d, 0x36, 0x12, 0x69, 0x18,
	0xcf, 0x9e, 0xb9, 0xa9, 0x3b, 0x97, 0xc8, 0x4c, 0x9e, 0x15, 0xc8, 0x7c, 0xa3, 0x7f, 0x05, 0x50,
	0x38, 0x63, 0x57, 0xf8, 0x61, 0xc1, 0xa6, 0xe5, 0x76, 0x0e, 0xa4, 0xf1, 0x7f, 0x13, 0x15, 0x1a,
	0xbb, 0xb0, 0x63, 0xcd, 0x13, 0xf1, 0xa6, 0xff, 0x16, 0x60, 0x67, 0x3c, 0x62, 0xe9, 0x6b, 0xfc,
	0x0e, 0x41, 0x75, 0x90, 0xbb, 0x74, 0x5c, 0x2f, 0x62, 0xb8, 0x22, 0xff, 0xcf, 0x41, 0xf5, 0xd5,
	0x52, 0x7f, 0xf5, 0xf6, 0xe6, 0x5a, 0x9b, 0xc2, 0xd3, 0x01, 0xb5, 0x4c, 0xc7, 0x22, 0x8e, 0x79,
	0x74, 0x6e, 0x11, 0x39, 0xa8, 0xcc, 0x08, 0x03, 0x12, 0xc6, 0x82, 0xcd, 0x58, 0x4a, 0x9e, 0x51,
	0xfb, 0xc2, 0xa4, 0x13, 0xf2, 0xd4, 0x9a, 0xb4, 0x88, 0x54, 0x24, 0xcf, 0x4d, 0x3a, 0x38, 0x33,
	0xa9, 0xf1, 0x45, 0xb7, 0xd9, 0x22, 0xcb, 0x2e, 0x92, 0xa3, 0x89, 0x63, 0x99, 0x2d, 0xb8, 0xbf,
	0x6a, 0x65, 0x19, 0xd7, 0xc4, 0x7f, 0x22, 0xa8, 0xda, 0x71, 0xc6, 0x52, 0x21, 0xe7, 0x97, 0xe1,
	0x0f, 0xa5, 0x0f, 0xb9, 0x2c, 0x3b, 0xfa, 0x1d, 0x49, 0x4b, 0xbf, 0x21, 0xb0, 0xed, 0xe1, 0xc8,
	0xa2, 0x0e, 0xb1, 0x87, 0xce, 0xa5, 0xb2, 0x44, 0x8c, 0x30, 0x50, 0xfa, 0xb7, 0x9a, 0x2d, 0xb2,
	0x12, 0x6b, 0x92, 0xe7, 0xe6, 0xf9, 0xd8, 0x1a, 0x11, 0xa3, 0xd1, 0x6b, 0x91, 0x46, 0xbf, 0x45,
	0x1a, 0x8f, 0x5b, 0xa4, 0xf1, 0x79, 0x13, 0xcb, 0x74, 0xa8, 0x3c, 0xdc, 0x4e, 0xbc, 0x34, 0x6f,
	0x7a, 0x0c, 0x3b, 0x76, 0xec, 0xdb, 0x01, 0xfe, 0x76, 0x16, 0x8a, 0x1f, 0x16, 0x5e, 0xdb, 0xe7,
	0xf3, 0x8e, 0xf0, 0x63, 0x3f, 0xe2, 0x8b, 0x40, 0x05, 0xd1, 0x7f, 0x34, 0x63, 0xf1, 0xa3, 0x22,
	0xb1, 0xec, 0x47, 0x77, 0x9e, 0x44, 0x2c, 0xcb, 0x23, 0xfd, 0x32, 0xbb, 0x8a, 0x3a, 0x89, 0x67,
	0x20, 0xfc, 0x3d, 0x54, 0x4f, 0x99, 0x30, 0xa3, 0x48, 0x95, 0x57, 0xea, 0xf8, 0x6d, 0xa5, 0xfa,
	0x97, 0xb2, 0xba, 0x1e, 0x74, 0x46, 0xd6, 0xb9, 0x35, 0x70, 0xc8, 0xc6, 0x82, 0xc8, 0x09, 0xbd,
	0xbc, 0x50, 0xa5, 0x77, 0x11, 0x8e, 0x61, 0x6f, 0xc4, 0x22, 0xe6, 0xe7, 0xcd, 0x3b, 0x7a, 0x63,
	0x07, 0xa5, 0xfe, 0xad, 0xf8, 0xcf, 0x24, 0xff, 0x00, 0x9e, 0xfc, 0x4f, 0x7e, 0xf2, 0xdd, 0x99,
	0x45, 0x2d, 0x12, 0x06, 0xe4, 0x09, 0x69, 0xf4, 0xf2, 0x86, 0xfd, 0x8a, 0xe0, 0xe3, 0x71, 0x12,
	0xb8, 0x82, 0x49, 0xe2, 0xa1, 0xcc, 0xd9, 0x5a, 0x45, 0x5f, 0x2a, 0xbe, 0x80, 0xa6, 0xc2, 0x2e,
	0xe9, 0x32, 0x26, 0x54, 0x50, 0x24, 0xdd, 0x3f, 0xc8, 0xfb, 0x04, 0x0e, 0xa9, 0xe5, 0x8c, 0xe9,
	0xd0, 0x1e, 0x9e, 0x6e, 0xf6, 0xb7, 0x9c, 0x99, 0x16, 0x06, 0x06, 0xea, 0x22, 0xfc, 0xa2, 0xf0,
	0x23, 0xbd, 0x38, 0xfc, 0x84, 0xf3, 0xf5, 0x09, 0x5a, 0xb6, 0xf8, 0xb3, 0xff, 0x32, 0x74, 0x38,
	0xe5, 0xfc, 0x70, 0x4d, 0xc1, 0x4d, 0xd8, 0x53, 0xbf, 0x98, 0x51, 0xa4, 0xca, 0x5d, 0x3b, 0xc1,
	0x7b, 0x5d, 0x84, 0x7f, 0x02, 0x38, 0x65, 0xa2, 0xb8, 0xf7, 0x7b, 0xe5, 0x37, 0x85, 0x5d, 0x95,
	0xb0, 0x54, 0x5a, 0xb9, 0x80, 0xe3, 0xf7, 0x9b, 0xc6, 0xd2, 0xaf, 0x39, 0x9c, 0x18, 0x8d, 0x5e,
	0x13, 0xab, 0xbb, 0xde, 0x45, 0xf8, 0x1b, 0xa8, 0x1c, 0xa7, 0x3c, 0xd9, 0x74, 0xaf, 0xef, 0x4b,
	0x61, 0x0c, 0xfb, 0x41, 0xca, 0x13, 0x22, 0x24, 0x50, 0x31, 0xd7, 0xff, 0x42, 0x3f, 0xdf, 0x5c,
	0x6b, 0x7f, 0x20, 0x98, 0xc1, 0x86, 0x07, 0x6e, 0xbf, 0xaa, 0x83, 0xdc, 0xa8, 0xa7, 0xe6, 0xe1,
	0x5d, 0xee, 0x05, 0xbc, 0x82, 0xb5, 0x4f, 0xd7, 0x7e, 0x55, 0x3f, 0x28, 0x9d, 0x0d, 0x78, 0xfc,
	0x9a, 0xa5, 0x82, 0xa5, 0x77, 0x12, 0x7b, 0x77, 0x73, 0xad, 0xdd, 0xf3, 0x3e, 0xc8, 0xd1, 0x8f,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xb0, 0xbb, 0xc5, 0x75, 0x06, 0x00, 0x00,
}