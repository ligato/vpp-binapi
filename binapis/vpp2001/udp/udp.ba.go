// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2001/.vppapi/core/udp.api.json

/*
Package udp is a generated VPP binary API for 'udp' module.

It consists of:
	  4 enums
	  5 aliases
	  7 types
	  1 union
	  6 messages
	  3 services
*/
package udp

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "udp"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xa6b4c475
)

// AddressFamily represents VPP binary API enum 'address_family'.
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var AddressFamily_name = map[uint32]string{
	0: "ADDRESS_IP4",
	1: "ADDRESS_IP6",
}

var AddressFamily_value = map[string]uint32{
	"ADDRESS_IP4": 0,
	"ADDRESS_IP6": 1,
}

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPDscp represents VPP binary API enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var IPDscp_name = map[uint8]string{
	0:  "IP_API_DSCP_CS0",
	8:  "IP_API_DSCP_CS1",
	10: "IP_API_DSCP_AF11",
	12: "IP_API_DSCP_AF12",
	14: "IP_API_DSCP_AF13",
	16: "IP_API_DSCP_CS2",
	18: "IP_API_DSCP_AF21",
	20: "IP_API_DSCP_AF22",
	22: "IP_API_DSCP_AF23",
	24: "IP_API_DSCP_CS3",
	26: "IP_API_DSCP_AF31",
	28: "IP_API_DSCP_AF32",
	30: "IP_API_DSCP_AF33",
	32: "IP_API_DSCP_CS4",
	34: "IP_API_DSCP_AF41",
	36: "IP_API_DSCP_AF42",
	38: "IP_API_DSCP_AF43",
	40: "IP_API_DSCP_CS5",
	46: "IP_API_DSCP_EF",
	48: "IP_API_DSCP_CS6",
	50: "IP_API_DSCP_CS7",
}

var IPDscp_value = map[string]uint8{
	"IP_API_DSCP_CS0":  0,
	"IP_API_DSCP_CS1":  8,
	"IP_API_DSCP_AF11": 10,
	"IP_API_DSCP_AF12": 12,
	"IP_API_DSCP_AF13": 14,
	"IP_API_DSCP_CS2":  16,
	"IP_API_DSCP_AF21": 18,
	"IP_API_DSCP_AF22": 20,
	"IP_API_DSCP_AF23": 22,
	"IP_API_DSCP_CS3":  24,
	"IP_API_DSCP_AF31": 26,
	"IP_API_DSCP_AF32": 28,
	"IP_API_DSCP_AF33": 30,
	"IP_API_DSCP_CS4":  32,
	"IP_API_DSCP_AF41": 34,
	"IP_API_DSCP_AF42": 36,
	"IP_API_DSCP_AF43": 38,
	"IP_API_DSCP_CS5":  40,
	"IP_API_DSCP_EF":   46,
	"IP_API_DSCP_CS6":  48,
	"IP_API_DSCP_CS7":  50,
}

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPEcn represents VPP binary API enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var IPEcn_name = map[uint8]string{
	0: "IP_API_ECN_NONE",
	1: "IP_API_ECN_ECT0",
	2: "IP_API_ECN_ECT1",
	3: "IP_API_ECN_CE",
}

var IPEcn_value = map[string]uint8{
	"IP_API_ECN_NONE": 0,
	"IP_API_ECN_ECT0": 1,
	"IP_API_ECN_ECT1": 2,
	"IP_API_ECN_CE":   3,
}

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPProto represents VPP binary API enum 'ip_proto'.
type IPProto uint32

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_AH       IPProto = 50
	IP_API_PROTO_ESP      IPProto = 51
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var IPProto_name = map[uint32]string{
	0:   "IP_API_PROTO_HOPOPT",
	1:   "IP_API_PROTO_ICMP",
	2:   "IP_API_PROTO_IGMP",
	6:   "IP_API_PROTO_TCP",
	17:  "IP_API_PROTO_UDP",
	47:  "IP_API_PROTO_GRE",
	50:  "IP_API_PROTO_AH",
	51:  "IP_API_PROTO_ESP",
	88:  "IP_API_PROTO_EIGRP",
	89:  "IP_API_PROTO_OSPF",
	132: "IP_API_PROTO_SCTP",
	255: "IP_API_PROTO_RESERVED",
}

var IPProto_value = map[string]uint32{
	"IP_API_PROTO_HOPOPT":   0,
	"IP_API_PROTO_ICMP":     1,
	"IP_API_PROTO_IGMP":     2,
	"IP_API_PROTO_TCP":      6,
	"IP_API_PROTO_UDP":      17,
	"IP_API_PROTO_GRE":      47,
	"IP_API_PROTO_AH":       50,
	"IP_API_PROTO_ESP":      51,
	"IP_API_PROTO_EIGRP":    88,
	"IP_API_PROTO_OSPF":     89,
	"IP_API_PROTO_SCTP":     132,
	"IP_API_PROTO_RESERVED": 255,
}

func (x IPProto) String() string {
	s, ok := IPProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// AddressWithPrefix represents VPP binary API alias 'address_with_prefix'.
type AddressWithPrefix Prefix

// IP4Address represents VPP binary API alias 'ip4_address'.
type IP4Address [4]uint8

// IP4AddressWithPrefix represents VPP binary API alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address represents VPP binary API alias 'ip6_address'.
type IP6Address [16]uint8

// IP6AddressWithPrefix represents VPP binary API alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// Address represents VPP binary API type 'address'.
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string { return "address" }

// IP4Prefix represents VPP binary API type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address
	Len     uint8
}

func (*IP4Prefix) GetTypeName() string { return "ip4_prefix" }

// IP6Prefix represents VPP binary API type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address
	Len     uint8
}

func (*IP6Prefix) GetTypeName() string { return "ip6_prefix" }

// Mprefix represents VPP binary API type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string { return "mprefix" }

// Prefix represents VPP binary API type 'prefix'.
type Prefix struct {
	Address Address
	Len     uint8
}

func (*Prefix) GetTypeName() string { return "prefix" }

// PrefixMatcher represents VPP binary API type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8
	Ge uint8
}

func (*PrefixMatcher) GetTypeName() string { return "prefix_matcher" }

// UDPEncap represents VPP binary API type 'udp_encap'.
type UDPEncap struct {
	TableID uint32
	SrcPort uint16
	DstPort uint16
	SrcIP   Address
	DstIP   Address
	ID      uint32
}

func (*UDPEncap) GetTypeName() string { return "udp_encap" }

// AddressUnion represents VPP binary API union 'address_union'.
type AddressUnion struct {
	XXX_UnionData [16]byte
}

func (*AddressUnion) GetTypeName() string { return "address_union" }

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

// UDPEncapAdd represents VPP binary API message 'udp_encap_add'.
type UDPEncapAdd struct {
	UDPEncap UDPEncap
}

func (m *UDPEncapAdd) Reset()                        { *m = UDPEncapAdd{} }
func (*UDPEncapAdd) GetMessageName() string          { return "udp_encap_add" }
func (*UDPEncapAdd) GetCrcString() string            { return "61d5fc48" }
func (*UDPEncapAdd) GetMessageType() api.MessageType { return api.RequestMessage }

// UDPEncapAddReply represents VPP binary API message 'udp_encap_add_reply'.
type UDPEncapAddReply struct {
	Retval int32
	ID     uint32
}

func (m *UDPEncapAddReply) Reset()                        { *m = UDPEncapAddReply{} }
func (*UDPEncapAddReply) GetMessageName() string          { return "udp_encap_add_reply" }
func (*UDPEncapAddReply) GetCrcString() string            { return "e2fc8294" }
func (*UDPEncapAddReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// UDPEncapDel represents VPP binary API message 'udp_encap_del'.
type UDPEncapDel struct {
	ID uint32
}

func (m *UDPEncapDel) Reset()                        { *m = UDPEncapDel{} }
func (*UDPEncapDel) GetMessageName() string          { return "udp_encap_del" }
func (*UDPEncapDel) GetCrcString() string            { return "3a91bde5" }
func (*UDPEncapDel) GetMessageType() api.MessageType { return api.RequestMessage }

// UDPEncapDelReply represents VPP binary API message 'udp_encap_del_reply'.
type UDPEncapDelReply struct {
	Retval int32
}

func (m *UDPEncapDelReply) Reset()                        { *m = UDPEncapDelReply{} }
func (*UDPEncapDelReply) GetMessageName() string          { return "udp_encap_del_reply" }
func (*UDPEncapDelReply) GetCrcString() string            { return "e8d4e804" }
func (*UDPEncapDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// UDPEncapDetails represents VPP binary API message 'udp_encap_details'.
type UDPEncapDetails struct {
	UDPEncap UDPEncap
}

func (m *UDPEncapDetails) Reset()                        { *m = UDPEncapDetails{} }
func (*UDPEncapDetails) GetMessageName() string          { return "udp_encap_details" }
func (*UDPEncapDetails) GetCrcString() string            { return "87c82821" }
func (*UDPEncapDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// UDPEncapDump represents VPP binary API message 'udp_encap_dump'.
type UDPEncapDump struct{}

func (m *UDPEncapDump) Reset()                        { *m = UDPEncapDump{} }
func (*UDPEncapDump) GetMessageName() string          { return "udp_encap_dump" }
func (*UDPEncapDump) GetCrcString() string            { return "51077d14" }
func (*UDPEncapDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*UDPEncapAdd)(nil), "udp.UDPEncapAdd")
	api.RegisterMessage((*UDPEncapAddReply)(nil), "udp.UDPEncapAddReply")
	api.RegisterMessage((*UDPEncapDel)(nil), "udp.UDPEncapDel")
	api.RegisterMessage((*UDPEncapDelReply)(nil), "udp.UDPEncapDelReply")
	api.RegisterMessage((*UDPEncapDetails)(nil), "udp.UDPEncapDetails")
	api.RegisterMessage((*UDPEncapDump)(nil), "udp.UDPEncapDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*UDPEncapAdd)(nil),
		(*UDPEncapAddReply)(nil),
		(*UDPEncapDel)(nil),
		(*UDPEncapDelReply)(nil),
		(*UDPEncapDetails)(nil),
		(*UDPEncapDump)(nil),
	}
}

// RPCService represents RPC service API for udp module.
type RPCService interface {
	DumpUDPEncap(ctx context.Context, in *UDPEncapDump) (RPCService_DumpUDPEncapClient, error)
	UDPEncapAdd(ctx context.Context, in *UDPEncapAdd) (*UDPEncapAddReply, error)
	UDPEncapDel(ctx context.Context, in *UDPEncapDel) (*UDPEncapDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpUDPEncap(ctx context.Context, in *UDPEncapDump) (RPCService_DumpUDPEncapClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpUDPEncapClient{stream}
	return x, nil
}

type RPCService_DumpUDPEncapClient interface {
	Recv() (*UDPEncapDetails, error)
}

type serviceClient_DumpUDPEncapClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpUDPEncapClient) Recv() (*UDPEncapDetails, error) {
	m := new(UDPEncapDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) UDPEncapAdd(ctx context.Context, in *UDPEncapAdd) (*UDPEncapAddReply, error) {
	out := new(UDPEncapAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UDPEncapDel(ctx context.Context, in *UDPEncapDel) (*UDPEncapDelReply, error) {
	out := new(UDPEncapDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
