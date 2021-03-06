// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2001/.vppapi/core/l2tp.api.json

/*
Package l2tp is a generated VPP binary API for 'l2tp' module.

It consists of:
	 11 enums
	  7 aliases
	  6 types
	  1 union
	 10 messages
	  5 services
*/
package l2tp

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
	ModuleName = "l2tp"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x1ecf7730
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

// IfStatusFlags represents VPP binary API enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var IfStatusFlags_name = map[uint32]string{
	1: "IF_STATUS_API_FLAG_ADMIN_UP",
	2: "IF_STATUS_API_FLAG_LINK_UP",
}

var IfStatusFlags_value = map[string]uint32{
	"IF_STATUS_API_FLAG_ADMIN_UP": 1,
	"IF_STATUS_API_FLAG_LINK_UP":  2,
}

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfType represents VPP binary API enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var IfType_name = map[uint32]string{
	1: "IF_API_TYPE_HARDWARE",
	2: "IF_API_TYPE_SUB",
	3: "IF_API_TYPE_P2P",
	4: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 1,
	"IF_API_TYPE_SUB":      2,
	"IF_API_TYPE_P2P":      3,
	"IF_API_TYPE_PIPE":     4,
}

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
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

// L2tLookupKey represents VPP binary API enum 'l2t_lookup_key'.
type L2tLookupKey uint8

const (
	L2T_LOOKUP_KEY_API_SRC_ADDR   L2tLookupKey = 0
	L2T_LOOKUP_KEY_API_DST_ADDR   L2tLookupKey = 1
	L2T_LOOKUP_KEY_API_SESSION_ID L2tLookupKey = 2
)

var L2tLookupKey_name = map[uint8]string{
	0: "L2T_LOOKUP_KEY_API_SRC_ADDR",
	1: "L2T_LOOKUP_KEY_API_DST_ADDR",
	2: "L2T_LOOKUP_KEY_API_SESSION_ID",
}

var L2tLookupKey_value = map[string]uint8{
	"L2T_LOOKUP_KEY_API_SRC_ADDR":   0,
	"L2T_LOOKUP_KEY_API_DST_ADDR":   1,
	"L2T_LOOKUP_KEY_API_SESSION_ID": 2,
}

func (x L2tLookupKey) String() string {
	s, ok := L2tLookupKey_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LinkDuplex represents VPP binary API enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var LinkDuplex_name = map[uint32]string{
	0: "LINK_DUPLEX_API_UNKNOWN",
	1: "LINK_DUPLEX_API_HALF",
	2: "LINK_DUPLEX_API_FULL",
}

var LinkDuplex_value = map[string]uint32{
	"LINK_DUPLEX_API_UNKNOWN": 0,
	"LINK_DUPLEX_API_HALF":    1,
	"LINK_DUPLEX_API_FULL":    2,
}

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var MtuProto_name = map[uint32]string{
	1: "MTU_PROTO_API_L3",
	2: "MTU_PROTO_API_IP4",
	3: "MTU_PROTO_API_IP6",
	4: "MTU_PROTO_API_MPLS",
	5: "MTU_PROTO_API_N",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   1,
	"MTU_PROTO_API_IP4":  2,
	"MTU_PROTO_API_IP6":  3,
	"MTU_PROTO_API_MPLS": 4,
	"MTU_PROTO_API_N":    5,
}

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RxMode represents VPP binary API enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var RxMode_name = map[uint32]string{
	0: "RX_MODE_API_UNKNOWN",
	1: "RX_MODE_API_POLLING",
	2: "RX_MODE_API_INTERRUPT",
	3: "RX_MODE_API_ADAPTIVE",
	4: "RX_MODE_API_DEFAULT",
}

var RxMode_value = map[string]uint32{
	"RX_MODE_API_UNKNOWN":   0,
	"RX_MODE_API_POLLING":   1,
	"RX_MODE_API_INTERRUPT": 2,
	"RX_MODE_API_ADAPTIVE":  3,
	"RX_MODE_API_DEFAULT":   4,
}

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SubIfFlags represents VPP binary API enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var SubIfFlags_name = map[uint32]string{
	1:   "SUB_IF_API_FLAG_NO_TAGS",
	2:   "SUB_IF_API_FLAG_ONE_TAG",
	4:   "SUB_IF_API_FLAG_TWO_TAGS",
	8:   "SUB_IF_API_FLAG_DOT1AD",
	16:  "SUB_IF_API_FLAG_EXACT_MATCH",
	32:  "SUB_IF_API_FLAG_DEFAULT",
	64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
	128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
	254: "SUB_IF_API_FLAG_MASK_VNET",
	256: "SUB_IF_API_FLAG_DOT1AH",
}

var SubIfFlags_value = map[string]uint32{
	"SUB_IF_API_FLAG_NO_TAGS":           1,
	"SUB_IF_API_FLAG_ONE_TAG":           2,
	"SUB_IF_API_FLAG_TWO_TAGS":          4,
	"SUB_IF_API_FLAG_DOT1AD":            8,
	"SUB_IF_API_FLAG_EXACT_MATCH":       16,
	"SUB_IF_API_FLAG_DEFAULT":           32,
	"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
	"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
	"SUB_IF_API_FLAG_MASK_VNET":         254,
	"SUB_IF_API_FLAG_DOT1AH":            256,
}

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// AddressWithPrefix represents VPP binary API alias 'address_with_prefix'.
type AddressWithPrefix Prefix

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// IP4Address represents VPP binary API alias 'ip4_address'.
type IP4Address [4]uint8

// IP4AddressWithPrefix represents VPP binary API alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address represents VPP binary API alias 'ip6_address'.
type IP6Address [16]uint8

// IP6AddressWithPrefix represents VPP binary API alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// MacAddress represents VPP binary API alias 'mac_address'.
type MacAddress [6]uint8

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

// L2tpv3CreateTunnel represents VPP binary API message 'l2tpv3_create_tunnel'.
type L2tpv3CreateTunnel struct {
	ClientAddress     Address
	OurAddress        Address
	LocalSessionID    uint32
	RemoteSessionID   uint32
	LocalCookie       uint64
	RemoteCookie      uint64
	L2SublayerPresent bool
	EncapVrfID        uint32
}

func (m *L2tpv3CreateTunnel) Reset()                        { *m = L2tpv3CreateTunnel{} }
func (*L2tpv3CreateTunnel) GetMessageName() string          { return "l2tpv3_create_tunnel" }
func (*L2tpv3CreateTunnel) GetCrcString() string            { return "596892cb" }
func (*L2tpv3CreateTunnel) GetMessageType() api.MessageType { return api.RequestMessage }

// L2tpv3CreateTunnelReply represents VPP binary API message 'l2tpv3_create_tunnel_reply'.
type L2tpv3CreateTunnelReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *L2tpv3CreateTunnelReply) Reset()                        { *m = L2tpv3CreateTunnelReply{} }
func (*L2tpv3CreateTunnelReply) GetMessageName() string          { return "l2tpv3_create_tunnel_reply" }
func (*L2tpv3CreateTunnelReply) GetCrcString() string            { return "5383d31f" }
func (*L2tpv3CreateTunnelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// L2tpv3InterfaceEnableDisable represents VPP binary API message 'l2tpv3_interface_enable_disable'.
type L2tpv3InterfaceEnableDisable struct {
	EnableDisable bool
	SwIfIndex     InterfaceIndex
}

func (m *L2tpv3InterfaceEnableDisable) Reset() { *m = L2tpv3InterfaceEnableDisable{} }
func (*L2tpv3InterfaceEnableDisable) GetMessageName() string {
	return "l2tpv3_interface_enable_disable"
}
func (*L2tpv3InterfaceEnableDisable) GetCrcString() string            { return "3865946c" }
func (*L2tpv3InterfaceEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// L2tpv3InterfaceEnableDisableReply represents VPP binary API message 'l2tpv3_interface_enable_disable_reply'.
type L2tpv3InterfaceEnableDisableReply struct {
	Retval int32
}

func (m *L2tpv3InterfaceEnableDisableReply) Reset() { *m = L2tpv3InterfaceEnableDisableReply{} }
func (*L2tpv3InterfaceEnableDisableReply) GetMessageName() string {
	return "l2tpv3_interface_enable_disable_reply"
}
func (*L2tpv3InterfaceEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*L2tpv3InterfaceEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// L2tpv3SetLookupKey represents VPP binary API message 'l2tpv3_set_lookup_key'.
type L2tpv3SetLookupKey struct {
	Key L2tLookupKey
}

func (m *L2tpv3SetLookupKey) Reset()                        { *m = L2tpv3SetLookupKey{} }
func (*L2tpv3SetLookupKey) GetMessageName() string          { return "l2tpv3_set_lookup_key" }
func (*L2tpv3SetLookupKey) GetCrcString() string            { return "c9892c86" }
func (*L2tpv3SetLookupKey) GetMessageType() api.MessageType { return api.RequestMessage }

// L2tpv3SetLookupKeyReply represents VPP binary API message 'l2tpv3_set_lookup_key_reply'.
type L2tpv3SetLookupKeyReply struct {
	Retval int32
}

func (m *L2tpv3SetLookupKeyReply) Reset()                        { *m = L2tpv3SetLookupKeyReply{} }
func (*L2tpv3SetLookupKeyReply) GetMessageName() string          { return "l2tpv3_set_lookup_key_reply" }
func (*L2tpv3SetLookupKeyReply) GetCrcString() string            { return "e8d4e804" }
func (*L2tpv3SetLookupKeyReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// L2tpv3SetTunnelCookies represents VPP binary API message 'l2tpv3_set_tunnel_cookies'.
type L2tpv3SetTunnelCookies struct {
	SwIfIndex       InterfaceIndex
	NewLocalCookie  uint64
	NewRemoteCookie uint64
}

func (m *L2tpv3SetTunnelCookies) Reset()                        { *m = L2tpv3SetTunnelCookies{} }
func (*L2tpv3SetTunnelCookies) GetMessageName() string          { return "l2tpv3_set_tunnel_cookies" }
func (*L2tpv3SetTunnelCookies) GetCrcString() string            { return "b3f4faf7" }
func (*L2tpv3SetTunnelCookies) GetMessageType() api.MessageType { return api.RequestMessage }

// L2tpv3SetTunnelCookiesReply represents VPP binary API message 'l2tpv3_set_tunnel_cookies_reply'.
type L2tpv3SetTunnelCookiesReply struct {
	Retval int32
}

func (m *L2tpv3SetTunnelCookiesReply) Reset()                        { *m = L2tpv3SetTunnelCookiesReply{} }
func (*L2tpv3SetTunnelCookiesReply) GetMessageName() string          { return "l2tpv3_set_tunnel_cookies_reply" }
func (*L2tpv3SetTunnelCookiesReply) GetCrcString() string            { return "e8d4e804" }
func (*L2tpv3SetTunnelCookiesReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwIfL2tpv3TunnelDetails represents VPP binary API message 'sw_if_l2tpv3_tunnel_details'.
type SwIfL2tpv3TunnelDetails struct {
	SwIfIndex         InterfaceIndex
	InterfaceName     string `struc:"[64]byte"`
	ClientAddress     Address
	OurAddress        Address
	LocalSessionID    uint32
	RemoteSessionID   uint32
	LocalCookie       []uint64 `struc:"[2]uint64"`
	RemoteCookie      uint64
	L2SublayerPresent bool
}

func (m *SwIfL2tpv3TunnelDetails) Reset()                        { *m = SwIfL2tpv3TunnelDetails{} }
func (*SwIfL2tpv3TunnelDetails) GetMessageName() string          { return "sw_if_l2tpv3_tunnel_details" }
func (*SwIfL2tpv3TunnelDetails) GetCrcString() string            { return "1dab5c7e" }
func (*SwIfL2tpv3TunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwIfL2tpv3TunnelDump represents VPP binary API message 'sw_if_l2tpv3_tunnel_dump'.
type SwIfL2tpv3TunnelDump struct{}

func (m *SwIfL2tpv3TunnelDump) Reset()                        { *m = SwIfL2tpv3TunnelDump{} }
func (*SwIfL2tpv3TunnelDump) GetMessageName() string          { return "sw_if_l2tpv3_tunnel_dump" }
func (*SwIfL2tpv3TunnelDump) GetCrcString() string            { return "51077d14" }
func (*SwIfL2tpv3TunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*L2tpv3CreateTunnel)(nil), "l2tp.L2tpv3CreateTunnel")
	api.RegisterMessage((*L2tpv3CreateTunnelReply)(nil), "l2tp.L2tpv3CreateTunnelReply")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisable)(nil), "l2tp.L2tpv3InterfaceEnableDisable")
	api.RegisterMessage((*L2tpv3InterfaceEnableDisableReply)(nil), "l2tp.L2tpv3InterfaceEnableDisableReply")
	api.RegisterMessage((*L2tpv3SetLookupKey)(nil), "l2tp.L2tpv3SetLookupKey")
	api.RegisterMessage((*L2tpv3SetLookupKeyReply)(nil), "l2tp.L2tpv3SetLookupKeyReply")
	api.RegisterMessage((*L2tpv3SetTunnelCookies)(nil), "l2tp.L2tpv3SetTunnelCookies")
	api.RegisterMessage((*L2tpv3SetTunnelCookiesReply)(nil), "l2tp.L2tpv3SetTunnelCookiesReply")
	api.RegisterMessage((*SwIfL2tpv3TunnelDetails)(nil), "l2tp.SwIfL2tpv3TunnelDetails")
	api.RegisterMessage((*SwIfL2tpv3TunnelDump)(nil), "l2tp.SwIfL2tpv3TunnelDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*L2tpv3CreateTunnel)(nil),
		(*L2tpv3CreateTunnelReply)(nil),
		(*L2tpv3InterfaceEnableDisable)(nil),
		(*L2tpv3InterfaceEnableDisableReply)(nil),
		(*L2tpv3SetLookupKey)(nil),
		(*L2tpv3SetLookupKeyReply)(nil),
		(*L2tpv3SetTunnelCookies)(nil),
		(*L2tpv3SetTunnelCookiesReply)(nil),
		(*SwIfL2tpv3TunnelDetails)(nil),
		(*SwIfL2tpv3TunnelDump)(nil),
	}
}

// RPCService represents RPC service API for l2tp module.
type RPCService interface {
	DumpSwIfL2tpv3Tunnel(ctx context.Context, in *SwIfL2tpv3TunnelDump) (RPCService_DumpSwIfL2tpv3TunnelClient, error)
	L2tpv3CreateTunnel(ctx context.Context, in *L2tpv3CreateTunnel) (*L2tpv3CreateTunnelReply, error)
	L2tpv3InterfaceEnableDisable(ctx context.Context, in *L2tpv3InterfaceEnableDisable) (*L2tpv3InterfaceEnableDisableReply, error)
	L2tpv3SetLookupKey(ctx context.Context, in *L2tpv3SetLookupKey) (*L2tpv3SetLookupKeyReply, error)
	L2tpv3SetTunnelCookies(ctx context.Context, in *L2tpv3SetTunnelCookies) (*L2tpv3SetTunnelCookiesReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwIfL2tpv3Tunnel(ctx context.Context, in *SwIfL2tpv3TunnelDump) (RPCService_DumpSwIfL2tpv3TunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwIfL2tpv3TunnelClient{stream}
	return x, nil
}

type RPCService_DumpSwIfL2tpv3TunnelClient interface {
	Recv() (*SwIfL2tpv3TunnelDetails, error)
}

type serviceClient_DumpSwIfL2tpv3TunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwIfL2tpv3TunnelClient) Recv() (*SwIfL2tpv3TunnelDetails, error) {
	m := new(SwIfL2tpv3TunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) L2tpv3CreateTunnel(ctx context.Context, in *L2tpv3CreateTunnel) (*L2tpv3CreateTunnelReply, error) {
	out := new(L2tpv3CreateTunnelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) L2tpv3InterfaceEnableDisable(ctx context.Context, in *L2tpv3InterfaceEnableDisable) (*L2tpv3InterfaceEnableDisableReply, error) {
	out := new(L2tpv3InterfaceEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) L2tpv3SetLookupKey(ctx context.Context, in *L2tpv3SetLookupKey) (*L2tpv3SetLookupKeyReply, error) {
	out := new(L2tpv3SetLookupKeyReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) L2tpv3SetTunnelCookies(ctx context.Context, in *L2tpv3SetTunnelCookies) (*L2tpv3SetTunnelCookiesReply, error) {
	out := new(L2tpv3SetTunnelCookiesReply)
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
