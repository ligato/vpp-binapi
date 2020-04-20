// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/master/vppapi/plugins/ioam_vxlan_gpe.api.json

/*
Package ioam_vxlan_gpe is a generated VPP binary API for 'ioam_vxlan_gpe' module.

It consists of:
	  4 enums
	  5 aliases
	  6 types
	  1 union
	 12 messages
	  6 services
*/
package ioam_vxlan_gpe

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "ioam_vxlan_gpe"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xe60c2de2
)

// AddressFamily represents VPP binary API enum 'address_family'.
type AddressFamily uint8

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var AddressFamily_name = map[uint8]string{
	0: "ADDRESS_IP4",
	1: "ADDRESS_IP6",
}

var AddressFamily_value = map[string]uint8{
	"ADDRESS_IP4": 0,
	"ADDRESS_IP6": 1,
}

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint8(x)]
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
type IPProto uint8

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_ESP      IPProto = 50
	IP_API_PROTO_AH       IPProto = 51
	IP_API_PROTO_ICMP6    IPProto = 58
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var IPProto_name = map[uint8]string{
	0:   "IP_API_PROTO_HOPOPT",
	1:   "IP_API_PROTO_ICMP",
	2:   "IP_API_PROTO_IGMP",
	6:   "IP_API_PROTO_TCP",
	17:  "IP_API_PROTO_UDP",
	47:  "IP_API_PROTO_GRE",
	50:  "IP_API_PROTO_ESP",
	51:  "IP_API_PROTO_AH",
	58:  "IP_API_PROTO_ICMP6",
	88:  "IP_API_PROTO_EIGRP",
	89:  "IP_API_PROTO_OSPF",
	132: "IP_API_PROTO_SCTP",
	255: "IP_API_PROTO_RESERVED",
}

var IPProto_value = map[string]uint8{
	"IP_API_PROTO_HOPOPT":   0,
	"IP_API_PROTO_ICMP":     1,
	"IP_API_PROTO_IGMP":     2,
	"IP_API_PROTO_TCP":      6,
	"IP_API_PROTO_UDP":      17,
	"IP_API_PROTO_GRE":      47,
	"IP_API_PROTO_ESP":      50,
	"IP_API_PROTO_AH":       51,
	"IP_API_PROTO_ICMP6":    58,
	"IP_API_PROTO_EIGRP":    88,
	"IP_API_PROTO_OSPF":     89,
	"IP_API_PROTO_SCTP":     132,
	"IP_API_PROTO_RESERVED": 255,
}

func (x IPProto) String() string {
	s, ok := IPProto_name[uint8(x)]
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

// VxlanGpeIoamDisable represents VPP binary API message 'vxlan_gpe_ioam_disable'.
type VxlanGpeIoamDisable struct {
	ID uint16
}

func (m *VxlanGpeIoamDisable) Reset()                        { *m = VxlanGpeIoamDisable{} }
func (*VxlanGpeIoamDisable) GetMessageName() string          { return "vxlan_gpe_ioam_disable" }
func (*VxlanGpeIoamDisable) GetCrcString() string            { return "6b16a45e" }
func (*VxlanGpeIoamDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamDisableReply represents VPP binary API message 'vxlan_gpe_ioam_disable_reply'.
type VxlanGpeIoamDisableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamDisableReply) Reset()                        { *m = VxlanGpeIoamDisableReply{} }
func (*VxlanGpeIoamDisableReply) GetMessageName() string          { return "vxlan_gpe_ioam_disable_reply" }
func (*VxlanGpeIoamDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamEnable represents VPP binary API message 'vxlan_gpe_ioam_enable'.
type VxlanGpeIoamEnable struct {
	ID          uint16
	TracePpc    uint8
	PowEnable   bool
	TraceEnable bool
}

func (m *VxlanGpeIoamEnable) Reset()                        { *m = VxlanGpeIoamEnable{} }
func (*VxlanGpeIoamEnable) GetMessageName() string          { return "vxlan_gpe_ioam_enable" }
func (*VxlanGpeIoamEnable) GetCrcString() string            { return "2481bef7" }
func (*VxlanGpeIoamEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamEnableReply represents VPP binary API message 'vxlan_gpe_ioam_enable_reply'.
type VxlanGpeIoamEnableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamEnableReply) Reset()                        { *m = VxlanGpeIoamEnableReply{} }
func (*VxlanGpeIoamEnableReply) GetMessageName() string          { return "vxlan_gpe_ioam_enable_reply" }
func (*VxlanGpeIoamEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamTransitDisable represents VPP binary API message 'vxlan_gpe_ioam_transit_disable'.
type VxlanGpeIoamTransitDisable struct {
	OuterFibIndex uint32
	DstAddr       Address
}

func (m *VxlanGpeIoamTransitDisable) Reset()                        { *m = VxlanGpeIoamTransitDisable{} }
func (*VxlanGpeIoamTransitDisable) GetMessageName() string          { return "vxlan_gpe_ioam_transit_disable" }
func (*VxlanGpeIoamTransitDisable) GetCrcString() string            { return "553f5b7b" }
func (*VxlanGpeIoamTransitDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamTransitDisableReply represents VPP binary API message 'vxlan_gpe_ioam_transit_disable_reply'.
type VxlanGpeIoamTransitDisableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamTransitDisableReply) Reset() { *m = VxlanGpeIoamTransitDisableReply{} }
func (*VxlanGpeIoamTransitDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_disable_reply"
}
func (*VxlanGpeIoamTransitDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamTransitDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamTransitEnable represents VPP binary API message 'vxlan_gpe_ioam_transit_enable'.
type VxlanGpeIoamTransitEnable struct {
	OuterFibIndex uint32
	DstAddr       Address
}

func (m *VxlanGpeIoamTransitEnable) Reset()                        { *m = VxlanGpeIoamTransitEnable{} }
func (*VxlanGpeIoamTransitEnable) GetMessageName() string          { return "vxlan_gpe_ioam_transit_enable" }
func (*VxlanGpeIoamTransitEnable) GetCrcString() string            { return "553f5b7b" }
func (*VxlanGpeIoamTransitEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamTransitEnableReply represents VPP binary API message 'vxlan_gpe_ioam_transit_enable_reply'.
type VxlanGpeIoamTransitEnableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamTransitEnableReply) Reset() { *m = VxlanGpeIoamTransitEnableReply{} }
func (*VxlanGpeIoamTransitEnableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_enable_reply"
}
func (*VxlanGpeIoamTransitEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamTransitEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamVniDisable represents VPP binary API message 'vxlan_gpe_ioam_vni_disable'.
type VxlanGpeIoamVniDisable struct {
	Vni    uint32
	Local  Address
	Remote Address
}

func (m *VxlanGpeIoamVniDisable) Reset()                        { *m = VxlanGpeIoamVniDisable{} }
func (*VxlanGpeIoamVniDisable) GetMessageName() string          { return "vxlan_gpe_ioam_vni_disable" }
func (*VxlanGpeIoamVniDisable) GetCrcString() string            { return "997161fb" }
func (*VxlanGpeIoamVniDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamVniDisableReply represents VPP binary API message 'vxlan_gpe_ioam_vni_disable_reply'.
type VxlanGpeIoamVniDisableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamVniDisableReply) Reset() { *m = VxlanGpeIoamVniDisableReply{} }
func (*VxlanGpeIoamVniDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_disable_reply"
}
func (*VxlanGpeIoamVniDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamVniDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamVniEnable represents VPP binary API message 'vxlan_gpe_ioam_vni_enable'.
type VxlanGpeIoamVniEnable struct {
	Vni    uint32
	Local  Address
	Remote Address
}

func (m *VxlanGpeIoamVniEnable) Reset()                        { *m = VxlanGpeIoamVniEnable{} }
func (*VxlanGpeIoamVniEnable) GetMessageName() string          { return "vxlan_gpe_ioam_vni_enable" }
func (*VxlanGpeIoamVniEnable) GetCrcString() string            { return "997161fb" }
func (*VxlanGpeIoamVniEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamVniEnableReply represents VPP binary API message 'vxlan_gpe_ioam_vni_enable_reply'.
type VxlanGpeIoamVniEnableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamVniEnableReply) Reset()                        { *m = VxlanGpeIoamVniEnableReply{} }
func (*VxlanGpeIoamVniEnableReply) GetMessageName() string          { return "vxlan_gpe_ioam_vni_enable_reply" }
func (*VxlanGpeIoamVniEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamVniEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*VxlanGpeIoamDisable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamDisable")
	api.RegisterMessage((*VxlanGpeIoamDisableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamDisableReply")
	api.RegisterMessage((*VxlanGpeIoamEnable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamEnable")
	api.RegisterMessage((*VxlanGpeIoamEnableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamEnableReply")
	api.RegisterMessage((*VxlanGpeIoamTransitDisable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitDisable")
	api.RegisterMessage((*VxlanGpeIoamTransitDisableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitDisableReply")
	api.RegisterMessage((*VxlanGpeIoamTransitEnable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitEnable")
	api.RegisterMessage((*VxlanGpeIoamTransitEnableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitEnableReply")
	api.RegisterMessage((*VxlanGpeIoamVniDisable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniDisable")
	api.RegisterMessage((*VxlanGpeIoamVniDisableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniDisableReply")
	api.RegisterMessage((*VxlanGpeIoamVniEnable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniEnable")
	api.RegisterMessage((*VxlanGpeIoamVniEnableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniEnableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*VxlanGpeIoamDisable)(nil),
		(*VxlanGpeIoamDisableReply)(nil),
		(*VxlanGpeIoamEnable)(nil),
		(*VxlanGpeIoamEnableReply)(nil),
		(*VxlanGpeIoamTransitDisable)(nil),
		(*VxlanGpeIoamTransitDisableReply)(nil),
		(*VxlanGpeIoamTransitEnable)(nil),
		(*VxlanGpeIoamTransitEnableReply)(nil),
		(*VxlanGpeIoamVniDisable)(nil),
		(*VxlanGpeIoamVniDisableReply)(nil),
		(*VxlanGpeIoamVniEnable)(nil),
		(*VxlanGpeIoamVniEnableReply)(nil),
	}
}

// RPCService represents RPC service API for ioam_vxlan_gpe module.
type RPCService interface {
	VxlanGpeIoamDisable(ctx context.Context, in *VxlanGpeIoamDisable) (*VxlanGpeIoamDisableReply, error)
	VxlanGpeIoamEnable(ctx context.Context, in *VxlanGpeIoamEnable) (*VxlanGpeIoamEnableReply, error)
	VxlanGpeIoamTransitDisable(ctx context.Context, in *VxlanGpeIoamTransitDisable) (*VxlanGpeIoamTransitDisableReply, error)
	VxlanGpeIoamTransitEnable(ctx context.Context, in *VxlanGpeIoamTransitEnable) (*VxlanGpeIoamTransitEnableReply, error)
	VxlanGpeIoamVniDisable(ctx context.Context, in *VxlanGpeIoamVniDisable) (*VxlanGpeIoamVniDisableReply, error)
	VxlanGpeIoamVniEnable(ctx context.Context, in *VxlanGpeIoamVniEnable) (*VxlanGpeIoamVniEnableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) VxlanGpeIoamDisable(ctx context.Context, in *VxlanGpeIoamDisable) (*VxlanGpeIoamDisableReply, error) {
	out := new(VxlanGpeIoamDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamEnable(ctx context.Context, in *VxlanGpeIoamEnable) (*VxlanGpeIoamEnableReply, error) {
	out := new(VxlanGpeIoamEnableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamTransitDisable(ctx context.Context, in *VxlanGpeIoamTransitDisable) (*VxlanGpeIoamTransitDisableReply, error) {
	out := new(VxlanGpeIoamTransitDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamTransitEnable(ctx context.Context, in *VxlanGpeIoamTransitEnable) (*VxlanGpeIoamTransitEnableReply, error) {
	out := new(VxlanGpeIoamTransitEnableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamVniDisable(ctx context.Context, in *VxlanGpeIoamVniDisable) (*VxlanGpeIoamVniDisableReply, error) {
	out := new(VxlanGpeIoamVniDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamVniEnable(ctx context.Context, in *VxlanGpeIoamVniEnable) (*VxlanGpeIoamVniEnableReply, error) {
	out := new(VxlanGpeIoamVniEnableReply)
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
