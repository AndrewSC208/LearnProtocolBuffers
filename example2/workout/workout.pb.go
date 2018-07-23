// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workout.proto

package workout

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

type Item struct {
	Id                   int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Workout              *WorkoutRequest `protobuf:"bytes,2,opt,name=workout,proto3" json:"workout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_workout_9ac8e22db1f5b690, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetWorkout() *WorkoutRequest {
	if m != nil {
		return m.Workout
	}
	return nil
}

type WorkoutRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	StartDateTime        int64    `protobuf:"varint,3,opt,name=startDateTime,proto3" json:"startDateTime,omitempty"`
	Duration             int64    `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkoutRequest) Reset()         { *m = WorkoutRequest{} }
func (m *WorkoutRequest) String() string { return proto.CompactTextString(m) }
func (*WorkoutRequest) ProtoMessage()    {}
func (*WorkoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workout_9ac8e22db1f5b690, []int{1}
}
func (m *WorkoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkoutRequest.Unmarshal(m, b)
}
func (m *WorkoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkoutRequest.Marshal(b, m, deterministic)
}
func (dst *WorkoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkoutRequest.Merge(dst, src)
}
func (m *WorkoutRequest) XXX_Size() int {
	return xxx_messageInfo_WorkoutRequest.Size(m)
}
func (m *WorkoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkoutRequest proto.InternalMessageInfo

func (m *WorkoutRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *WorkoutRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *WorkoutRequest) GetStartDateTime() int64 {
	if m != nil {
		return m.StartDateTime
	}
	return 0
}

func (m *WorkoutRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type ItemRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequest) Reset()         { *m = ItemRequest{} }
func (m *ItemRequest) String() string { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()    {}
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workout_9ac8e22db1f5b690, []int{2}
}
func (m *ItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequest.Unmarshal(m, b)
}
func (m *ItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequest.Marshal(b, m, deterministic)
}
func (dst *ItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequest.Merge(dst, src)
}
func (m *ItemRequest) XXX_Size() int {
	return xxx_messageInfo_ItemRequest.Size(m)
}
func (m *ItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequest proto.InternalMessageInfo

func (m *ItemRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "workout.Item")
	proto.RegisterType((*WorkoutRequest)(nil), "workout.WorkoutRequest")
	proto.RegisterType((*ItemRequest)(nil), "workout.ItemRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkoutClient is the client API for Workout service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkoutClient interface {
	// Creates a workout and returns
	CreateWorkout(ctx context.Context, in *WorkoutRequest, opts ...grpc.CallOption) (*Item, error)
	// GetWorkout, get a workout by id
	ReadWorkout(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*Item, error)
}

type workoutClient struct {
	cc *grpc.ClientConn
}

func NewWorkoutClient(cc *grpc.ClientConn) WorkoutClient {
	return &workoutClient{cc}
}

func (c *workoutClient) CreateWorkout(ctx context.Context, in *WorkoutRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/workout.Workout/CreateWorkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workoutClient) ReadWorkout(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/workout.Workout/ReadWorkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkoutServer is the server API for Workout service.
type WorkoutServer interface {
	// Creates a workout and returns
	CreateWorkout(context.Context, *WorkoutRequest) (*Item, error)
	// GetWorkout, get a workout by id
	ReadWorkout(context.Context, *ItemRequest) (*Item, error)
}

func RegisterWorkoutServer(s *grpc.Server, srv WorkoutServer) {
	s.RegisterService(&_Workout_serviceDesc, srv)
}

func _Workout_CreateWorkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkoutServer).CreateWorkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workout.Workout/CreateWorkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkoutServer).CreateWorkout(ctx, req.(*WorkoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workout_ReadWorkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkoutServer).ReadWorkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workout.Workout/ReadWorkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkoutServer).ReadWorkout(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workout_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workout.Workout",
	HandlerType: (*WorkoutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkout",
			Handler:    _Workout_CreateWorkout_Handler,
		},
		{
			MethodName: "ReadWorkout",
			Handler:    _Workout_ReadWorkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workout.proto",
}

func init() { proto.RegisterFile("workout.proto", fileDescriptor_workout_9ac8e22db1f5b690) }

var fileDescriptor_workout_9ac8e22db1f5b690 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x4d, 0xd7, 0x39, 0xf7, 0x4a, 0x76, 0x78, 0x0c, 0x0c, 0x03, 0xa1, 0x14, 0x0f, 0x3d,
	0x0d, 0x9c, 0x5e, 0x3c, 0xeb, 0x65, 0xd7, 0x20, 0x78, 0x8e, 0xe4, 0x1d, 0x82, 0xce, 0xcc, 0xec,
	0x15, 0x51, 0xfc, 0xe3, 0xa5, 0x69, 0x53, 0xe8, 0xa0, 0xa7, 0xe4, 0x7b, 0xdf, 0x7b, 0xbf, 0xe4,
	0x7b, 0x20, 0xbf, 0x7d, 0x78, 0xf7, 0x0d, 0x6f, 0x8f, 0xc1, 0xb3, 0xc7, 0x45, 0x2f, 0xab, 0x3d,
	0xe4, 0x7b, 0xa6, 0x03, 0xae, 0x20, 0x73, 0x56, 0x89, 0x52, 0xd4, 0x73, 0x9d, 0x39, 0x8b, 0x77,
	0x90, 0x5a, 0x54, 0x56, 0x8a, 0xba, 0xd8, 0x5d, 0x6f, 0x13, 0xe1, 0xb5, 0x3b, 0x35, 0x7d, 0x35,
	0x74, 0x62, 0x3d, 0xa0, 0xfe, 0x60, 0x35, 0xb6, 0x70, 0x0d, 0x73, 0x76, 0xfc, 0x41, 0x91, 0xbb,
	0xd4, 0x9d, 0x40, 0x84, 0x9c, 0x7f, 0x8e, 0x14, 0xb9, 0x4b, 0x1d, 0xef, 0x78, 0x0b, 0xf2, 0xc4,
	0x26, 0xf0, 0xb3, 0x61, 0x7a, 0x71, 0x07, 0x52, 0xb3, 0x52, 0xd4, 0x33, 0x3d, 0x2e, 0xe2, 0x06,
	0xae, 0x6c, 0x13, 0x0c, 0x3b, 0xff, 0xa9, 0xf2, 0xd8, 0x30, 0xe8, 0xea, 0x06, 0x8a, 0x36, 0x48,
	0x7a, 0xfa, 0x2c, 0xcf, 0xee, 0x17, 0x16, 0xfd, 0xe7, 0xf0, 0x11, 0xe4, 0x53, 0x20, 0xc3, 0x94,
	0x0a, 0x53, 0xd1, 0x36, 0x72, 0x30, 0x5a, 0x74, 0x75, 0x81, 0x0f, 0x50, 0x68, 0x32, 0x36, 0x0d,
	0xae, 0x47, 0xfe, 0xd4, 0xd4, 0xdb, 0x65, 0xdc, 0xf9, 0xfd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0x70, 0x83, 0x98, 0x84, 0x01, 0x00, 0x00,
}
