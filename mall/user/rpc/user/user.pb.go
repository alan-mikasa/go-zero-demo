// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type IdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserResponse struct {
	// 用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 用户性别
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func init() {
	proto.RegisterType((*IdRequest)(nil), "user.IdRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xa4, 0xb9, 0x38, 0x3d, 0x53,
	0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0xbc, 0xb8, 0x78, 0x42, 0x8b, 0x53, 0x8b, 0x82,
	0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xd1, 0xe5, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x98, 0xc0, 0x22, 0x60, 0xb6, 0x90, 0x18, 0x17, 0x5b, 0x7a, 0x6a, 0x5e, 0x4a, 0x6a,
	0x91, 0x04, 0x33, 0x58, 0x14, 0xca, 0x33, 0x32, 0xe3, 0x62, 0x01, 0x99, 0x25, 0xa4, 0xc7, 0xc5,
	0x9e, 0x9e, 0x5a, 0x02, 0x66, 0xf2, 0xeb, 0x81, 0x9d, 0x03, 0xb7, 0x5f, 0x4a, 0x08, 0x22, 0x80,
	0x6c, 0x67, 0x12, 0x1b, 0xd8, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x4a, 0xfd,
	0xe5, 0xbb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.User/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUser(context.Context, *IdRequest) (*UserResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUser(ctx context.Context, req *IdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUser",
			Handler:    _User_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
