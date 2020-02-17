// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor/monitor.proto

package monitor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
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

type MonitorConfig_TargetType int32

const (
	MonitorConfig_SERIAL MonitorConfig_TargetType = 0
)

var MonitorConfig_TargetType_name = map[int32]string{
	0: "SERIAL",
}

var MonitorConfig_TargetType_value = map[string]int32{
	"SERIAL": 0,
}

func (x MonitorConfig_TargetType) String() string {
	return proto.EnumName(MonitorConfig_TargetType_name, int32(x))
}

func (MonitorConfig_TargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94d5950496a7550d, []int{1, 0}
}

// The top-level message sent by the client for the `StreamingOpen` method.
// Multiple `StreamingOpenReq` messages can be sent but the first message
// must contain a `monitor_config` message to initialize the monitor target.
// All subsequent messages must contain bytes to be sent to the target
// and must not contain a `monitor_config` message.
type StreamingOpenReq struct {
	// Content must be either a monitor config or data to be sent.
	//
	// Types that are valid to be assigned to Content:
	//	*StreamingOpenReq_MonitorConfig
	//	*StreamingOpenReq_Data
	Content              isStreamingOpenReq_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *StreamingOpenReq) Reset()         { *m = StreamingOpenReq{} }
func (m *StreamingOpenReq) String() string { return proto.CompactTextString(m) }
func (*StreamingOpenReq) ProtoMessage()    {}
func (*StreamingOpenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d5950496a7550d, []int{0}
}

func (m *StreamingOpenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingOpenReq.Unmarshal(m, b)
}
func (m *StreamingOpenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingOpenReq.Marshal(b, m, deterministic)
}
func (m *StreamingOpenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingOpenReq.Merge(m, src)
}
func (m *StreamingOpenReq) XXX_Size() int {
	return xxx_messageInfo_StreamingOpenReq.Size(m)
}
func (m *StreamingOpenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingOpenReq.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingOpenReq proto.InternalMessageInfo

type isStreamingOpenReq_Content interface {
	isStreamingOpenReq_Content()
}

type StreamingOpenReq_MonitorConfig struct {
	MonitorConfig *MonitorConfig `protobuf:"bytes,1,opt,name=monitorConfig,proto3,oneof"`
}

type StreamingOpenReq_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*StreamingOpenReq_MonitorConfig) isStreamingOpenReq_Content() {}

func (*StreamingOpenReq_Data) isStreamingOpenReq_Content() {}

func (m *StreamingOpenReq) GetContent() isStreamingOpenReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *StreamingOpenReq) GetMonitorConfig() *MonitorConfig {
	if x, ok := m.GetContent().(*StreamingOpenReq_MonitorConfig); ok {
		return x.MonitorConfig
	}
	return nil
}

func (m *StreamingOpenReq) GetData() []byte {
	if x, ok := m.GetContent().(*StreamingOpenReq_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamingOpenReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamingOpenReq_MonitorConfig)(nil),
		(*StreamingOpenReq_Data)(nil),
	}
}

// Tells the monitor which target to open and provides additional parameters
// that might be needed to configure the target or the monitor itself.
type MonitorConfig struct {
	Target               string                   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Type                 MonitorConfig_TargetType `protobuf:"varint,2,opt,name=type,proto3,enum=cc.arduino.cli.monitor.MonitorConfig_TargetType" json:"type,omitempty"`
	AdditionalConfig     *_struct.Struct          `protobuf:"bytes,3,opt,name=additionalConfig,proto3" json:"additionalConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MonitorConfig) Reset()         { *m = MonitorConfig{} }
func (m *MonitorConfig) String() string { return proto.CompactTextString(m) }
func (*MonitorConfig) ProtoMessage()    {}
func (*MonitorConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d5950496a7550d, []int{1}
}

func (m *MonitorConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorConfig.Unmarshal(m, b)
}
func (m *MonitorConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorConfig.Marshal(b, m, deterministic)
}
func (m *MonitorConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorConfig.Merge(m, src)
}
func (m *MonitorConfig) XXX_Size() int {
	return xxx_messageInfo_MonitorConfig.Size(m)
}
func (m *MonitorConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorConfig proto.InternalMessageInfo

func (m *MonitorConfig) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *MonitorConfig) GetType() MonitorConfig_TargetType {
	if m != nil {
		return m.Type
	}
	return MonitorConfig_SERIAL
}

func (m *MonitorConfig) GetAdditionalConfig() *_struct.Struct {
	if m != nil {
		return m.AdditionalConfig
	}
	return nil
}

//
type StreamingOpenResp struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingOpenResp) Reset()         { *m = StreamingOpenResp{} }
func (m *StreamingOpenResp) String() string { return proto.CompactTextString(m) }
func (*StreamingOpenResp) ProtoMessage()    {}
func (*StreamingOpenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d5950496a7550d, []int{2}
}

func (m *StreamingOpenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingOpenResp.Unmarshal(m, b)
}
func (m *StreamingOpenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingOpenResp.Marshal(b, m, deterministic)
}
func (m *StreamingOpenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingOpenResp.Merge(m, src)
}
func (m *StreamingOpenResp) XXX_Size() int {
	return xxx_messageInfo_StreamingOpenResp.Size(m)
}
func (m *StreamingOpenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingOpenResp.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingOpenResp proto.InternalMessageInfo

func (m *StreamingOpenResp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("cc.arduino.cli.monitor.MonitorConfig_TargetType", MonitorConfig_TargetType_name, MonitorConfig_TargetType_value)
	proto.RegisterType((*StreamingOpenReq)(nil), "cc.arduino.cli.monitor.StreamingOpenReq")
	proto.RegisterType((*MonitorConfig)(nil), "cc.arduino.cli.monitor.MonitorConfig")
	proto.RegisterType((*StreamingOpenResp)(nil), "cc.arduino.cli.monitor.StreamingOpenResp")
}

func init() { proto.RegisterFile("monitor/monitor.proto", fileDescriptor_94d5950496a7550d) }

var fileDescriptor_94d5950496a7550d = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x4e, 0xf2, 0x30,
	0x14, 0xc7, 0xb7, 0xef, 0x23, 0x10, 0x8e, 0x60, 0xb0, 0x51, 0x24, 0xc4, 0x0b, 0xb2, 0xc4, 0x38,
	0x8d, 0x76, 0x04, 0x9f, 0x40, 0xd0, 0x04, 0x13, 0x89, 0x49, 0xe1, 0xca, 0xbb, 0xd2, 0x95, 0x59,
	0x33, 0xda, 0x5a, 0xba, 0x0b, 0x6e, 0x7d, 0x3e, 0x1f, 0xca, 0x58, 0x4a, 0x14, 0xd4, 0x84, 0xab,
	0xa6, 0xd9, 0xf9, 0x9d, 0x73, 0xfe, 0xbf, 0x15, 0x8e, 0xe6, 0x4a, 0x0a, 0xab, 0x4c, 0xe2, 0x4f,
	0xac, 0x8d, 0xb2, 0x0a, 0x35, 0x19, 0xc3, 0xd4, 0xa4, 0x85, 0x90, 0x0a, 0xb3, 0x5c, 0x60, 0xff,
	0xb5, 0x7d, 0x92, 0x29, 0x95, 0xe5, 0x3c, 0x71, 0x55, 0xd3, 0x62, 0x96, 0x2c, 0xac, 0x29, 0x98,
	0x5d, 0x51, 0xd1, 0x5b, 0x08, 0x8d, 0xb1, 0x35, 0x9c, 0xce, 0x85, 0xcc, 0x1e, 0x35, 0x97, 0x84,
	0xbf, 0xa2, 0x11, 0xd4, 0x3d, 0x3d, 0x50, 0x72, 0x26, 0xb2, 0x56, 0xd8, 0x09, 0xe3, 0xbd, 0xde,
	0x29, 0xfe, 0x7d, 0x04, 0x1e, 0x7d, 0x2f, 0x1e, 0x06, 0x64, 0x93, 0x46, 0x87, 0x50, 0x4a, 0xa9,
	0xa5, 0xad, 0x7f, 0x9d, 0x30, 0xae, 0x0d, 0x03, 0xe2, 0x6e, 0xfd, 0x2a, 0x54, 0x98, 0x92, 0x96,
	0x4b, 0x1b, 0xbd, 0x87, 0x50, 0xdf, 0xe8, 0x81, 0x9a, 0x50, 0xb6, 0xd4, 0x64, 0xdc, 0xba, 0xd1,
	0x55, 0xe2, 0x6f, 0xe8, 0x16, 0x4a, 0x76, 0xa9, 0xb9, 0x6b, 0xb5, 0xdf, 0xeb, 0xee, 0xb4, 0x10,
	0x9e, 0x38, 0x76, 0xb2, 0xd4, 0x9c, 0x38, 0x1a, 0x0d, 0xa0, 0x41, 0xd3, 0x54, 0x58, 0xa1, 0x24,
	0xcd, 0x7d, 0xc4, 0xff, 0x2e, 0xe2, 0x31, 0x5e, 0xd9, 0xc2, 0x6b, 0x5b, 0x78, 0xec, 0x6c, 0x91,
	0x1f, 0x40, 0xd4, 0x02, 0xf8, 0x6a, 0x8c, 0x00, 0xca, 0xe3, 0x3b, 0x72, 0x7f, 0xf3, 0xd0, 0x08,
	0xa2, 0x33, 0x38, 0xd8, 0x52, 0xba, 0xd0, 0x08, 0x79, 0x09, 0x9f, 0x79, 0x6a, 0x2b, 0x05, 0xbd,
	0x02, 0x2a, 0x7e, 0x53, 0xf4, 0x02, 0xf5, 0x0d, 0x06, 0xc5, 0x7f, 0x65, 0xdb, 0xfe, 0x5b, 0xed,
	0xf3, 0x1d, 0x2b, 0x17, 0x3a, 0x0a, 0xe2, 0xb0, 0x1b, 0xf6, 0x2f, 0x9f, 0x2e, 0x32, 0x61, 0x9f,
	0x8b, 0x29, 0x66, 0x6a, 0x9e, 0x78, 0x72, 0x7d, 0x5e, 0xb1, 0x5c, 0x24, 0x46, 0xb3, 0xf5, 0xeb,
	0x9a, 0x96, 0x9d, 0x8a, 0xeb, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x6c, 0x5e, 0x00, 0x77,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorClient interface {
	StreamingOpen(ctx context.Context, opts ...grpc.CallOption) (Monitor_StreamingOpenClient, error)
}

type monitorClient struct {
	cc *grpc.ClientConn
}

func NewMonitorClient(cc *grpc.ClientConn) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) StreamingOpen(ctx context.Context, opts ...grpc.CallOption) (Monitor_StreamingOpenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Monitor_serviceDesc.Streams[0], "/cc.arduino.cli.monitor.Monitor/StreamingOpen", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorStreamingOpenClient{stream}
	return x, nil
}

type Monitor_StreamingOpenClient interface {
	Send(*StreamingOpenReq) error
	Recv() (*StreamingOpenResp, error)
	grpc.ClientStream
}

type monitorStreamingOpenClient struct {
	grpc.ClientStream
}

func (x *monitorStreamingOpenClient) Send(m *StreamingOpenReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *monitorStreamingOpenClient) Recv() (*StreamingOpenResp, error) {
	m := new(StreamingOpenResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorServer is the server API for Monitor service.
type MonitorServer interface {
	StreamingOpen(Monitor_StreamingOpenServer) error
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_StreamingOpen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MonitorServer).StreamingOpen(&monitorStreamingOpenServer{stream})
}

type Monitor_StreamingOpenServer interface {
	Send(*StreamingOpenResp) error
	Recv() (*StreamingOpenReq, error)
	grpc.ServerStream
}

type monitorStreamingOpenServer struct {
	grpc.ServerStream
}

func (x *monitorStreamingOpenServer) Send(m *StreamingOpenResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *monitorStreamingOpenServer) Recv() (*StreamingOpenReq, error) {
	m := new(StreamingOpenReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.arduino.cli.monitor.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOpen",
			Handler:       _Monitor_StreamingOpen_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "monitor/monitor.proto",
}
