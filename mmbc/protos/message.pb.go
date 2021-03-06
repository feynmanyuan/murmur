// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/message.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/message.proto

It has these top-level messages:
	GossipMessage
	Empty
	ConnEstablish
	Envelope
	Member
	PeerTime
	AliveMessage
*/
package protos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GossipMessage_Tag int32

const (
	GossipMessage_UNDEFINED    GossipMessage_Tag = 0
	GossipMessage_EMPTY        GossipMessage_Tag = 1
	GossipMessage_ORG_ONLY     GossipMessage_Tag = 2
	GossipMessage_CHAN_ONLY    GossipMessage_Tag = 3
	GossipMessage_CHAN_AND_ORG GossipMessage_Tag = 4
	GossipMessage_CHAN_OR_ORG  GossipMessage_Tag = 5
)

var GossipMessage_Tag_name = map[int32]string{
	0: "UNDEFINED",
	1: "EMPTY",
	2: "ORG_ONLY",
	3: "CHAN_ONLY",
	4: "CHAN_AND_ORG",
	5: "CHAN_OR_ORG",
}
var GossipMessage_Tag_value = map[string]int32{
	"UNDEFINED":    0,
	"EMPTY":        1,
	"ORG_ONLY":     2,
	"CHAN_ONLY":    3,
	"CHAN_AND_ORG": 4,
	"CHAN_OR_ORG":  5,
}

func (x GossipMessage_Tag) String() string {
	return proto.EnumName(GossipMessage_Tag_name, int32(x))
}
func (GossipMessage_Tag) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 0} }

type GossipMessage struct {
	Nonce   uint64            `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Channel []byte            `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Tag     GossipMessage_Tag `protobuf:"varint,3,opt,name=tag,proto3,enum=gossip.GossipMessage_Tag" json:"tag,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*GossipMessage_AliveMsg
	//	*GossipMessage_Conn
	Content isGossipMessage_Content `protobuf_oneof:"content"`
}

func (m *GossipMessage) Reset()                    { *m = GossipMessage{} }
func (m *GossipMessage) String() string            { return proto.CompactTextString(m) }
func (*GossipMessage) ProtoMessage()               {}
func (*GossipMessage) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type isGossipMessage_Content interface {
	isGossipMessage_Content()
}

type GossipMessage_AliveMsg struct {
	AliveMsg *AliveMessage `protobuf:"bytes,5,opt,name=alive_msg,json=aliveMsg,oneof"`
}
type GossipMessage_Conn struct {
	Conn *ConnEstablish `protobuf:"bytes,6,opt,name=conn,oneof"`
}

func (*GossipMessage_AliveMsg) isGossipMessage_Content() {}
func (*GossipMessage_Conn) isGossipMessage_Content()     {}

func (m *GossipMessage) GetContent() isGossipMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GossipMessage) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *GossipMessage) GetChannel() []byte {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *GossipMessage) GetTag() GossipMessage_Tag {
	if m != nil {
		return m.Tag
	}
	return GossipMessage_UNDEFINED
}

func (m *GossipMessage) GetAliveMsg() *AliveMessage {
	if x, ok := m.GetContent().(*GossipMessage_AliveMsg); ok {
		return x.AliveMsg
	}
	return nil
}

func (m *GossipMessage) GetConn() *ConnEstablish {
	if x, ok := m.GetContent().(*GossipMessage_Conn); ok {
		return x.Conn
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GossipMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GossipMessage_OneofMarshaler, _GossipMessage_OneofUnmarshaler, _GossipMessage_OneofSizer, []interface{}{
		(*GossipMessage_AliveMsg)(nil),
		(*GossipMessage_Conn)(nil),
	}
}

func _GossipMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GossipMessage)
	// content
	switch x := m.Content.(type) {
	case *GossipMessage_AliveMsg:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AliveMsg); err != nil {
			return err
		}
	case *GossipMessage_Conn:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Conn); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GossipMessage.Content has unexpected type %T", x)
	}
	return nil
}

func _GossipMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GossipMessage)
	switch tag {
	case 5: // content.alive_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AliveMessage)
		err := b.DecodeMessage(msg)
		m.Content = &GossipMessage_AliveMsg{msg}
		return true, err
	case 6: // content.conn
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConnEstablish)
		err := b.DecodeMessage(msg)
		m.Content = &GossipMessage_Conn{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GossipMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GossipMessage)
	// content
	switch x := m.Content.(type) {
	case *GossipMessage_AliveMsg:
		s := proto.Size(x.AliveMsg)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GossipMessage_Conn:
		s := proto.Size(x.Conn)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

type ConnEstablish struct {
	PkiId       []byte `protobuf:"bytes,1,opt,name=pki_id,json=pkiId,proto3" json:"pki_id,omitempty"`
	Identity    []byte `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	TlsCertHash []byte `protobuf:"bytes,3,opt,name=tls_cert_hash,json=tlsCertHash,proto3" json:"tls_cert_hash,omitempty"`
}

func (m *ConnEstablish) Reset()                    { *m = ConnEstablish{} }
func (m *ConnEstablish) String() string            { return proto.CompactTextString(m) }
func (*ConnEstablish) ProtoMessage()               {}
func (*ConnEstablish) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{2} }

func (m *ConnEstablish) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}

func (m *ConnEstablish) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *ConnEstablish) GetTlsCertHash() []byte {
	if m != nil {
		return m.TlsCertHash
	}
	return nil
}

type Envelope struct {
	Payload   []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{3} }

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Member struct {
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	PkiId    []byte `protobuf:"bytes,3,opt,name=pki_id,json=pkiId,proto3" json:"pki_id,omitempty"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{4} }

func (m *Member) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Member) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Member) GetPkiId() []byte {
	if m != nil {
		return m.PkiId
	}
	return nil
}

type PeerTime struct {
	IncNum uint64 `protobuf:"varint,1,opt,name=inc_num,json=incNum,proto3" json:"inc_num,omitempty"`
	SeqNum uint64 `protobuf:"varint,2,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
}

func (m *PeerTime) Reset()                    { *m = PeerTime{} }
func (m *PeerTime) String() string            { return proto.CompactTextString(m) }
func (*PeerTime) ProtoMessage()               {}
func (*PeerTime) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{5} }

func (m *PeerTime) GetIncNum() uint64 {
	if m != nil {
		return m.IncNum
	}
	return 0
}

func (m *PeerTime) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

type AliveMessage struct {
	Member        *Member   `protobuf:"bytes,1,opt,name=member" json:"member,omitempty"`
	Timestamp     *PeerTime `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	IsDeclaration bool      `protobuf:"varint,3,opt,name=is_declaration,json=isDeclaration,proto3" json:"is_declaration,omitempty"`
}

func (m *AliveMessage) Reset()                    { *m = AliveMessage{} }
func (m *AliveMessage) String() string            { return proto.CompactTextString(m) }
func (*AliveMessage) ProtoMessage()               {}
func (*AliveMessage) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{6} }

func (m *AliveMessage) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *AliveMessage) GetTimestamp() *PeerTime {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *AliveMessage) GetIsDeclaration() bool {
	if m != nil {
		return m.IsDeclaration
	}
	return false
}

func init() {
	proto.RegisterType((*GossipMessage)(nil), "gossip.GossipMessage")
	proto.RegisterType((*Empty)(nil), "gossip.Empty")
	proto.RegisterType((*ConnEstablish)(nil), "gossip.ConnEstablish")
	proto.RegisterType((*Envelope)(nil), "gossip.Envelope")
	proto.RegisterType((*Member)(nil), "gossip.Member")
	proto.RegisterType((*PeerTime)(nil), "gossip.PeerTime")
	proto.RegisterType((*AliveMessage)(nil), "gossip.AliveMessage")
	proto.RegisterEnum("gossip.GossipMessage_Tag", GossipMessage_Tag_name, GossipMessage_Tag_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gossip service

type GossipClient interface {
	GossipStream(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipStreamClient, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type gossipClient struct {
	cc *grpc.ClientConn
}

func NewGossipClient(cc *grpc.ClientConn) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) GossipStream(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Gossip_serviceDesc.Streams[0], c.cc, "/gossip.Gossip/GossipStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gossipGossipStreamClient{stream}
	return x, nil
}

type Gossip_GossipStreamClient interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type gossipGossipStreamClient struct {
	grpc.ClientStream
}

func (x *gossipGossipStreamClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gossipGossipStreamClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gossipClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/gossip.Gossip/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gossip service

type GossipServer interface {
	GossipStream(Gossip_GossipStreamServer) error
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterGossipServer(s *grpc.Server, srv GossipServer) {
	s.RegisterService(&_Gossip_serviceDesc, srv)
}

func _Gossip_GossipStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GossipServer).GossipStream(&gossipGossipStreamServer{stream})
}

type Gossip_GossipStreamServer interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type gossipGossipStreamServer struct {
	grpc.ServerStream
}

func (x *gossipGossipStreamServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gossipGossipStreamServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Gossip_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gossip.Gossip/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gossip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gossip.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Gossip_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GossipStream",
			Handler:       _Gossip_GossipStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/message.proto",
}

func init() { proto.RegisterFile("protos/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5f, 0x6f, 0xda, 0x3e,
	0x14, 0x25, 0xfc, 0x09, 0xe4, 0x12, 0xfa, 0x43, 0x56, 0xab, 0x1f, 0x43, 0x7d, 0x40, 0x91, 0x56,
	0x21, 0x55, 0xa2, 0x13, 0x7d, 0xd9, 0xc3, 0x5e, 0xda, 0xc2, 0x4a, 0xa5, 0x91, 0x56, 0x1e, 0xd3,
	0xd4, 0xbd, 0x44, 0x26, 0xdc, 0x05, 0xaf, 0xb1, 0x93, 0xc6, 0xa6, 0x12, 0x1f, 0x62, 0x5f, 0x68,
	0x9f, 0x6e, 0x8a, 0x43, 0x68, 0xd1, 0x9e, 0xe2, 0x73, 0xcf, 0xf1, 0xf1, 0xb1, 0xef, 0x0d, 0x1c,
	0xa7, 0x59, 0xa2, 0x13, 0x75, 0x21, 0x50, 0x29, 0x16, 0xe1, 0xc8, 0x40, 0x62, 0x47, 0x89, 0x52,
	0x3c, 0xf5, 0xfe, 0x54, 0xa1, 0x73, 0x6b, 0x96, 0xf3, 0x82, 0x27, 0xc7, 0xd0, 0x90, 0x89, 0x0c,
	0xb1, 0x67, 0x0d, 0xac, 0x61, 0x9d, 0x16, 0x80, 0xf4, 0xa0, 0x19, 0xae, 0x99, 0x94, 0x18, 0xf7,
	0xaa, 0x03, 0x6b, 0xe8, 0xd2, 0x12, 0x92, 0x73, 0xa8, 0x69, 0x16, 0xf5, 0x6a, 0x03, 0x6b, 0x78,
	0x34, 0x7e, 0x37, 0x2a, 0x7c, 0x47, 0x07, 0x9e, 0xa3, 0x05, 0x8b, 0x68, 0xae, 0x22, 0x97, 0xe0,
	0xb0, 0x98, 0xbf, 0x60, 0x20, 0x54, 0xd4, 0x6b, 0x0c, 0xac, 0x61, 0x7b, 0x7c, 0x5c, 0x6e, 0xb9,
	0xca, 0x89, 0xdd, 0x8e, 0x59, 0x85, 0xb6, 0x8c, 0x70, 0xae, 0x22, 0x72, 0x0e, 0xf5, 0x30, 0x91,
	0xb2, 0x67, 0x1b, 0xfd, 0x49, 0xa9, 0xbf, 0x49, 0xa4, 0x9c, 0x2a, 0xcd, 0x96, 0x31, 0x57, 0xeb,
	0x59, 0x85, 0x1a, 0x91, 0x17, 0x40, 0x6d, 0xc1, 0x22, 0xd2, 0x01, 0xe7, 0x9b, 0x3f, 0x99, 0x7e,
	0xbe, 0xf3, 0xa7, 0x93, 0x6e, 0x85, 0x38, 0xd0, 0x98, 0xce, 0x1f, 0x16, 0x8f, 0x5d, 0x8b, 0xb8,
	0xd0, 0xba, 0xa7, 0xb7, 0xc1, 0xbd, 0xff, 0xe5, 0xb1, 0x5b, 0xcd, 0x75, 0x37, 0xb3, 0x2b, 0xbf,
	0x80, 0x35, 0xd2, 0x05, 0xd7, 0xc0, 0x2b, 0x7f, 0x12, 0xdc, 0xd3, 0xdb, 0x6e, 0x9d, 0xfc, 0x07,
	0xed, 0x42, 0x40, 0x4d, 0xa1, 0x71, 0xed, 0x40, 0x33, 0x4c, 0xa4, 0x46, 0xa9, 0xbd, 0x26, 0x34,
	0xa6, 0x22, 0xd5, 0x5b, 0xef, 0x27, 0x74, 0x0e, 0xd2, 0x90, 0x13, 0xb0, 0xd3, 0x27, 0x1e, 0xf0,
	0x95, 0x79, 0x45, 0x97, 0x36, 0xd2, 0x27, 0x7e, 0xb7, 0x22, 0x7d, 0x68, 0xf1, 0x15, 0x4a, 0xcd,
	0xf5, 0x76, 0xf7, 0x8c, 0x7b, 0x4c, 0x3c, 0xe8, 0xe8, 0x58, 0x05, 0x21, 0x66, 0x3a, 0x58, 0x33,
	0xb5, 0x36, 0x2f, 0xea, 0xd2, 0xb6, 0x8e, 0xd5, 0x0d, 0x66, 0x7a, 0xc6, 0xd4, 0xda, 0xbb, 0x86,
	0xd6, 0x54, 0xbe, 0x60, 0x9c, 0xa4, 0xa6, 0x23, 0x29, 0xdb, 0xc6, 0x09, 0x2b, 0xcf, 0x28, 0x21,
	0x39, 0x05, 0x47, 0xf1, 0x48, 0x32, 0xbd, 0xc9, 0x70, 0x77, 0xcc, 0x6b, 0xc1, 0xfb, 0x0e, 0xf6,
	0x1c, 0xc5, 0x12, 0xb3, 0x3c, 0x0d, 0xca, 0x55, 0x9a, 0x70, 0xa9, 0x8d, 0x85, 0x43, 0xf7, 0x38,
	0xe7, 0x04, 0x6a, 0xb6, 0x62, 0x9a, 0x95, 0x49, 0x4b, 0xfc, 0xe6, 0x72, 0xb5, 0x37, 0x97, 0xf3,
	0x3e, 0x41, 0xeb, 0x01, 0x31, 0x5b, 0x70, 0x81, 0xe4, 0x7f, 0x68, 0x72, 0x19, 0x06, 0x72, 0x23,
	0x76, 0x63, 0x64, 0x73, 0x19, 0xfa, 0x1b, 0x91, 0x13, 0x0a, 0x9f, 0x0d, 0x51, 0x2d, 0x08, 0x85,
	0xcf, 0xfe, 0x46, 0x78, 0xbf, 0x2d, 0x70, 0xdf, 0x4e, 0x00, 0x39, 0x03, 0x5b, 0x98, 0x9c, 0xc6,
	0xa1, 0x3d, 0x3e, 0x2a, 0xfb, 0x5e, 0xa4, 0xa7, 0x3b, 0x96, 0x8c, 0xc0, 0xd1, 0x5c, 0xa0, 0xd2,
	0x4c, 0xa4, 0xc6, 0xb3, 0x3d, 0xee, 0x96, 0xd2, 0x32, 0x0f, 0x7d, 0x95, 0x90, 0xf7, 0x70, 0xc4,
	0x55, 0xb0, 0xc2, 0x30, 0x66, 0x19, 0xd3, 0x3c, 0x91, 0xe6, 0x16, 0x2d, 0xda, 0xe1, 0x6a, 0xf2,
	0x5a, 0x1c, 0xff, 0x02, 0xbb, 0x98, 0x61, 0xf2, 0x11, 0xdc, 0x62, 0xf5, 0x55, 0x67, 0xc8, 0x04,
	0xd9, 0xbb, 0x97, 0xad, 0xe8, 0xff, 0x53, 0xf1, 0x2a, 0x43, 0xeb, 0x83, 0x45, 0xce, 0xa0, 0xfe,
	0xc0, 0x65, 0x44, 0x3a, 0x7b, 0x3e, 0x9f, 0x96, 0xfe, 0x21, 0xf4, 0x2a, 0xd7, 0xa7, 0x3f, 0xfa,
	0xcb, 0xad, 0x46, 0xb1, 0xc9, 0xc4, 0x26, 0x1b, 0x85, 0x89, 0xb8, 0x10, 0x62, 0x19, 0x5e, 0x14,
	0x3f, 0xee, 0xd2, 0x36, 0xdf, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x26, 0xcc, 0x9d,
	0xc9, 0x03, 0x00, 0x00,
}
