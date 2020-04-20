// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2001/.vppapi/plugins/ikev2.api.json

/*
Package ikev2 is a generated VPP binary API for 'ikev2' module.

It consists of:
	 10 enums
	  6 aliases
	  6 types
	  1 union
	 28 messages
	 14 services
*/
package ikev2

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
	ModuleName = "ikev2"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x4eab5d22
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

// Ikev2InitiateDelChildSa represents VPP binary API message 'ikev2_initiate_del_child_sa'.
type Ikev2InitiateDelChildSa struct {
	Ispi uint32
}

func (m *Ikev2InitiateDelChildSa) Reset()                        { *m = Ikev2InitiateDelChildSa{} }
func (*Ikev2InitiateDelChildSa) GetMessageName() string          { return "ikev2_initiate_del_child_sa" }
func (*Ikev2InitiateDelChildSa) GetCrcString() string            { return "7f004d2e" }
func (*Ikev2InitiateDelChildSa) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2InitiateDelChildSaReply represents VPP binary API message 'ikev2_initiate_del_child_sa_reply'.
type Ikev2InitiateDelChildSaReply struct {
	Retval int32
}

func (m *Ikev2InitiateDelChildSaReply) Reset() { *m = Ikev2InitiateDelChildSaReply{} }
func (*Ikev2InitiateDelChildSaReply) GetMessageName() string {
	return "ikev2_initiate_del_child_sa_reply"
}
func (*Ikev2InitiateDelChildSaReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2InitiateDelChildSaReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2InitiateDelIkeSa represents VPP binary API message 'ikev2_initiate_del_ike_sa'.
type Ikev2InitiateDelIkeSa struct {
	Ispi uint64
}

func (m *Ikev2InitiateDelIkeSa) Reset()                        { *m = Ikev2InitiateDelIkeSa{} }
func (*Ikev2InitiateDelIkeSa) GetMessageName() string          { return "ikev2_initiate_del_ike_sa" }
func (*Ikev2InitiateDelIkeSa) GetCrcString() string            { return "8d125bdd" }
func (*Ikev2InitiateDelIkeSa) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2InitiateDelIkeSaReply represents VPP binary API message 'ikev2_initiate_del_ike_sa_reply'.
type Ikev2InitiateDelIkeSaReply struct {
	Retval int32
}

func (m *Ikev2InitiateDelIkeSaReply) Reset()                        { *m = Ikev2InitiateDelIkeSaReply{} }
func (*Ikev2InitiateDelIkeSaReply) GetMessageName() string          { return "ikev2_initiate_del_ike_sa_reply" }
func (*Ikev2InitiateDelIkeSaReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2InitiateDelIkeSaReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2InitiateRekeyChildSa represents VPP binary API message 'ikev2_initiate_rekey_child_sa'.
type Ikev2InitiateRekeyChildSa struct {
	Ispi uint32
}

func (m *Ikev2InitiateRekeyChildSa) Reset()                        { *m = Ikev2InitiateRekeyChildSa{} }
func (*Ikev2InitiateRekeyChildSa) GetMessageName() string          { return "ikev2_initiate_rekey_child_sa" }
func (*Ikev2InitiateRekeyChildSa) GetCrcString() string            { return "7f004d2e" }
func (*Ikev2InitiateRekeyChildSa) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2InitiateRekeyChildSaReply represents VPP binary API message 'ikev2_initiate_rekey_child_sa_reply'.
type Ikev2InitiateRekeyChildSaReply struct {
	Retval int32
}

func (m *Ikev2InitiateRekeyChildSaReply) Reset() { *m = Ikev2InitiateRekeyChildSaReply{} }
func (*Ikev2InitiateRekeyChildSaReply) GetMessageName() string {
	return "ikev2_initiate_rekey_child_sa_reply"
}
func (*Ikev2InitiateRekeyChildSaReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2InitiateRekeyChildSaReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2InitiateSaInit represents VPP binary API message 'ikev2_initiate_sa_init'.
type Ikev2InitiateSaInit struct {
	Name string `struc:"[64]byte"`
}

func (m *Ikev2InitiateSaInit) Reset()                        { *m = Ikev2InitiateSaInit{} }
func (*Ikev2InitiateSaInit) GetMessageName() string          { return "ikev2_initiate_sa_init" }
func (*Ikev2InitiateSaInit) GetCrcString() string            { return "ebf79a66" }
func (*Ikev2InitiateSaInit) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2InitiateSaInitReply represents VPP binary API message 'ikev2_initiate_sa_init_reply'.
type Ikev2InitiateSaInitReply struct {
	Retval int32
}

func (m *Ikev2InitiateSaInitReply) Reset()                        { *m = Ikev2InitiateSaInitReply{} }
func (*Ikev2InitiateSaInitReply) GetMessageName() string          { return "ikev2_initiate_sa_init_reply" }
func (*Ikev2InitiateSaInitReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2InitiateSaInitReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2PluginGetVersion represents VPP binary API message 'ikev2_plugin_get_version'.
type Ikev2PluginGetVersion struct{}

func (m *Ikev2PluginGetVersion) Reset()                        { *m = Ikev2PluginGetVersion{} }
func (*Ikev2PluginGetVersion) GetMessageName() string          { return "ikev2_plugin_get_version" }
func (*Ikev2PluginGetVersion) GetCrcString() string            { return "51077d14" }
func (*Ikev2PluginGetVersion) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2PluginGetVersionReply represents VPP binary API message 'ikev2_plugin_get_version_reply'.
type Ikev2PluginGetVersionReply struct {
	Major uint32
	Minor uint32
}

func (m *Ikev2PluginGetVersionReply) Reset()                        { *m = Ikev2PluginGetVersionReply{} }
func (*Ikev2PluginGetVersionReply) GetMessageName() string          { return "ikev2_plugin_get_version_reply" }
func (*Ikev2PluginGetVersionReply) GetCrcString() string            { return "9b32cf86" }
func (*Ikev2PluginGetVersionReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2ProfileAddDel represents VPP binary API message 'ikev2_profile_add_del'.
type Ikev2ProfileAddDel struct {
	Name  string `struc:"[64]byte"`
	IsAdd bool
}

func (m *Ikev2ProfileAddDel) Reset()                        { *m = Ikev2ProfileAddDel{} }
func (*Ikev2ProfileAddDel) GetMessageName() string          { return "ikev2_profile_add_del" }
func (*Ikev2ProfileAddDel) GetCrcString() string            { return "2c925b55" }
func (*Ikev2ProfileAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2ProfileAddDelReply represents VPP binary API message 'ikev2_profile_add_del_reply'.
type Ikev2ProfileAddDelReply struct {
	Retval int32
}

func (m *Ikev2ProfileAddDelReply) Reset()                        { *m = Ikev2ProfileAddDelReply{} }
func (*Ikev2ProfileAddDelReply) GetMessageName() string          { return "ikev2_profile_add_del_reply" }
func (*Ikev2ProfileAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2ProfileAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2ProfileSetAuth represents VPP binary API message 'ikev2_profile_set_auth'.
type Ikev2ProfileSetAuth struct {
	Name       string `struc:"[64]byte"`
	AuthMethod uint8
	IsHex      bool
	DataLen    uint32 `struc:"sizeof=Data"`
	Data       []byte
}

func (m *Ikev2ProfileSetAuth) Reset()                        { *m = Ikev2ProfileSetAuth{} }
func (*Ikev2ProfileSetAuth) GetMessageName() string          { return "ikev2_profile_set_auth" }
func (*Ikev2ProfileSetAuth) GetCrcString() string            { return "642c97cd" }
func (*Ikev2ProfileSetAuth) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2ProfileSetAuthReply represents VPP binary API message 'ikev2_profile_set_auth_reply'.
type Ikev2ProfileSetAuthReply struct {
	Retval int32
}

func (m *Ikev2ProfileSetAuthReply) Reset()                        { *m = Ikev2ProfileSetAuthReply{} }
func (*Ikev2ProfileSetAuthReply) GetMessageName() string          { return "ikev2_profile_set_auth_reply" }
func (*Ikev2ProfileSetAuthReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2ProfileSetAuthReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2ProfileSetID represents VPP binary API message 'ikev2_profile_set_id'.
type Ikev2ProfileSetID struct {
	Name    string `struc:"[64]byte"`
	IsLocal bool
	IDType  uint8
	DataLen uint32 `struc:"sizeof=Data"`
	Data    []byte
}

func (m *Ikev2ProfileSetID) Reset()                        { *m = Ikev2ProfileSetID{} }
func (*Ikev2ProfileSetID) GetMessageName() string          { return "ikev2_profile_set_id" }
func (*Ikev2ProfileSetID) GetCrcString() string            { return "4d7e2418" }
func (*Ikev2ProfileSetID) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2ProfileSetIDReply represents VPP binary API message 'ikev2_profile_set_id_reply'.
type Ikev2ProfileSetIDReply struct {
	Retval int32
}

func (m *Ikev2ProfileSetIDReply) Reset()                        { *m = Ikev2ProfileSetIDReply{} }
func (*Ikev2ProfileSetIDReply) GetMessageName() string          { return "ikev2_profile_set_id_reply" }
func (*Ikev2ProfileSetIDReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2ProfileSetIDReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2ProfileSetTs represents VPP binary API message 'ikev2_profile_set_ts'.
type Ikev2ProfileSetTs struct {
	Name      string `struc:"[64]byte"`
	IsLocal   bool
	Proto     uint8
	StartPort uint16
	EndPort   uint16
	StartAddr uint32
	EndAddr   uint32
}

func (m *Ikev2ProfileSetTs) Reset()                        { *m = Ikev2ProfileSetTs{} }
func (*Ikev2ProfileSetTs) GetMessageName() string          { return "ikev2_profile_set_ts" }
func (*Ikev2ProfileSetTs) GetCrcString() string            { return "64d55c16" }
func (*Ikev2ProfileSetTs) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2ProfileSetTsReply represents VPP binary API message 'ikev2_profile_set_ts_reply'.
type Ikev2ProfileSetTsReply struct {
	Retval int32
}

func (m *Ikev2ProfileSetTsReply) Reset()                        { *m = Ikev2ProfileSetTsReply{} }
func (*Ikev2ProfileSetTsReply) GetMessageName() string          { return "ikev2_profile_set_ts_reply" }
func (*Ikev2ProfileSetTsReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2ProfileSetTsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2SetEspTransforms represents VPP binary API message 'ikev2_set_esp_transforms'.
type Ikev2SetEspTransforms struct {
	Name          string `struc:"[64]byte"`
	CryptoAlg     uint32
	CryptoKeySize uint32
	IntegAlg      uint32
	DhGroup       uint32
}

func (m *Ikev2SetEspTransforms) Reset()                        { *m = Ikev2SetEspTransforms{} }
func (*Ikev2SetEspTransforms) GetMessageName() string          { return "ikev2_set_esp_transforms" }
func (*Ikev2SetEspTransforms) GetCrcString() string            { return "936a1a37" }
func (*Ikev2SetEspTransforms) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2SetEspTransformsReply represents VPP binary API message 'ikev2_set_esp_transforms_reply'.
type Ikev2SetEspTransformsReply struct {
	Retval int32
}

func (m *Ikev2SetEspTransformsReply) Reset()                        { *m = Ikev2SetEspTransformsReply{} }
func (*Ikev2SetEspTransformsReply) GetMessageName() string          { return "ikev2_set_esp_transforms_reply" }
func (*Ikev2SetEspTransformsReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2SetEspTransformsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2SetIkeTransforms represents VPP binary API message 'ikev2_set_ike_transforms'.
type Ikev2SetIkeTransforms struct {
	Name          string `struc:"[64]byte"`
	CryptoAlg     uint32
	CryptoKeySize uint32
	IntegAlg      uint32
	DhGroup       uint32
}

func (m *Ikev2SetIkeTransforms) Reset()                        { *m = Ikev2SetIkeTransforms{} }
func (*Ikev2SetIkeTransforms) GetMessageName() string          { return "ikev2_set_ike_transforms" }
func (*Ikev2SetIkeTransforms) GetCrcString() string            { return "936a1a37" }
func (*Ikev2SetIkeTransforms) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2SetIkeTransformsReply represents VPP binary API message 'ikev2_set_ike_transforms_reply'.
type Ikev2SetIkeTransformsReply struct {
	Retval int32
}

func (m *Ikev2SetIkeTransformsReply) Reset()                        { *m = Ikev2SetIkeTransformsReply{} }
func (*Ikev2SetIkeTransformsReply) GetMessageName() string          { return "ikev2_set_ike_transforms_reply" }
func (*Ikev2SetIkeTransformsReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2SetIkeTransformsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2SetLocalKey represents VPP binary API message 'ikev2_set_local_key'.
type Ikev2SetLocalKey struct {
	KeyFile []byte `struc:"[256]byte"`
}

func (m *Ikev2SetLocalKey) Reset()                        { *m = Ikev2SetLocalKey{} }
func (*Ikev2SetLocalKey) GetMessageName() string          { return "ikev2_set_local_key" }
func (*Ikev2SetLocalKey) GetCrcString() string            { return "e4996cd5" }
func (*Ikev2SetLocalKey) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2SetLocalKeyReply represents VPP binary API message 'ikev2_set_local_key_reply'.
type Ikev2SetLocalKeyReply struct {
	Retval int32
}

func (m *Ikev2SetLocalKeyReply) Reset()                        { *m = Ikev2SetLocalKeyReply{} }
func (*Ikev2SetLocalKeyReply) GetMessageName() string          { return "ikev2_set_local_key_reply" }
func (*Ikev2SetLocalKeyReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2SetLocalKeyReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2SetResponder represents VPP binary API message 'ikev2_set_responder'.
type Ikev2SetResponder struct {
	Name      string `struc:"[64]byte"`
	SwIfIndex InterfaceIndex
	Address   IP4Address
}

func (m *Ikev2SetResponder) Reset()                        { *m = Ikev2SetResponder{} }
func (*Ikev2SetResponder) GetMessageName() string          { return "ikev2_set_responder" }
func (*Ikev2SetResponder) GetCrcString() string            { return "f0d3dc80" }
func (*Ikev2SetResponder) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2SetResponderReply represents VPP binary API message 'ikev2_set_responder_reply'.
type Ikev2SetResponderReply struct {
	Retval int32
}

func (m *Ikev2SetResponderReply) Reset()                        { *m = Ikev2SetResponderReply{} }
func (*Ikev2SetResponderReply) GetMessageName() string          { return "ikev2_set_responder_reply" }
func (*Ikev2SetResponderReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2SetResponderReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Ikev2SetSaLifetime represents VPP binary API message 'ikev2_set_sa_lifetime'.
type Ikev2SetSaLifetime struct {
	Name            string `struc:"[64]byte"`
	Lifetime        uint64
	LifetimeJitter  uint32
	Handover        uint32
	LifetimeMaxdata uint64
}

func (m *Ikev2SetSaLifetime) Reset()                        { *m = Ikev2SetSaLifetime{} }
func (*Ikev2SetSaLifetime) GetMessageName() string          { return "ikev2_set_sa_lifetime" }
func (*Ikev2SetSaLifetime) GetCrcString() string            { return "7039feaa" }
func (*Ikev2SetSaLifetime) GetMessageType() api.MessageType { return api.RequestMessage }

// Ikev2SetSaLifetimeReply represents VPP binary API message 'ikev2_set_sa_lifetime_reply'.
type Ikev2SetSaLifetimeReply struct {
	Retval int32
}

func (m *Ikev2SetSaLifetimeReply) Reset()                        { *m = Ikev2SetSaLifetimeReply{} }
func (*Ikev2SetSaLifetimeReply) GetMessageName() string          { return "ikev2_set_sa_lifetime_reply" }
func (*Ikev2SetSaLifetimeReply) GetCrcString() string            { return "e8d4e804" }
func (*Ikev2SetSaLifetimeReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*Ikev2InitiateDelChildSa)(nil), "ikev2.Ikev2InitiateDelChildSa")
	api.RegisterMessage((*Ikev2InitiateDelChildSaReply)(nil), "ikev2.Ikev2InitiateDelChildSaReply")
	api.RegisterMessage((*Ikev2InitiateDelIkeSa)(nil), "ikev2.Ikev2InitiateDelIkeSa")
	api.RegisterMessage((*Ikev2InitiateDelIkeSaReply)(nil), "ikev2.Ikev2InitiateDelIkeSaReply")
	api.RegisterMessage((*Ikev2InitiateRekeyChildSa)(nil), "ikev2.Ikev2InitiateRekeyChildSa")
	api.RegisterMessage((*Ikev2InitiateRekeyChildSaReply)(nil), "ikev2.Ikev2InitiateRekeyChildSaReply")
	api.RegisterMessage((*Ikev2InitiateSaInit)(nil), "ikev2.Ikev2InitiateSaInit")
	api.RegisterMessage((*Ikev2InitiateSaInitReply)(nil), "ikev2.Ikev2InitiateSaInitReply")
	api.RegisterMessage((*Ikev2PluginGetVersion)(nil), "ikev2.Ikev2PluginGetVersion")
	api.RegisterMessage((*Ikev2PluginGetVersionReply)(nil), "ikev2.Ikev2PluginGetVersionReply")
	api.RegisterMessage((*Ikev2ProfileAddDel)(nil), "ikev2.Ikev2ProfileAddDel")
	api.RegisterMessage((*Ikev2ProfileAddDelReply)(nil), "ikev2.Ikev2ProfileAddDelReply")
	api.RegisterMessage((*Ikev2ProfileSetAuth)(nil), "ikev2.Ikev2ProfileSetAuth")
	api.RegisterMessage((*Ikev2ProfileSetAuthReply)(nil), "ikev2.Ikev2ProfileSetAuthReply")
	api.RegisterMessage((*Ikev2ProfileSetID)(nil), "ikev2.Ikev2ProfileSetID")
	api.RegisterMessage((*Ikev2ProfileSetIDReply)(nil), "ikev2.Ikev2ProfileSetIDReply")
	api.RegisterMessage((*Ikev2ProfileSetTs)(nil), "ikev2.Ikev2ProfileSetTs")
	api.RegisterMessage((*Ikev2ProfileSetTsReply)(nil), "ikev2.Ikev2ProfileSetTsReply")
	api.RegisterMessage((*Ikev2SetEspTransforms)(nil), "ikev2.Ikev2SetEspTransforms")
	api.RegisterMessage((*Ikev2SetEspTransformsReply)(nil), "ikev2.Ikev2SetEspTransformsReply")
	api.RegisterMessage((*Ikev2SetIkeTransforms)(nil), "ikev2.Ikev2SetIkeTransforms")
	api.RegisterMessage((*Ikev2SetIkeTransformsReply)(nil), "ikev2.Ikev2SetIkeTransformsReply")
	api.RegisterMessage((*Ikev2SetLocalKey)(nil), "ikev2.Ikev2SetLocalKey")
	api.RegisterMessage((*Ikev2SetLocalKeyReply)(nil), "ikev2.Ikev2SetLocalKeyReply")
	api.RegisterMessage((*Ikev2SetResponder)(nil), "ikev2.Ikev2SetResponder")
	api.RegisterMessage((*Ikev2SetResponderReply)(nil), "ikev2.Ikev2SetResponderReply")
	api.RegisterMessage((*Ikev2SetSaLifetime)(nil), "ikev2.Ikev2SetSaLifetime")
	api.RegisterMessage((*Ikev2SetSaLifetimeReply)(nil), "ikev2.Ikev2SetSaLifetimeReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Ikev2InitiateDelChildSa)(nil),
		(*Ikev2InitiateDelChildSaReply)(nil),
		(*Ikev2InitiateDelIkeSa)(nil),
		(*Ikev2InitiateDelIkeSaReply)(nil),
		(*Ikev2InitiateRekeyChildSa)(nil),
		(*Ikev2InitiateRekeyChildSaReply)(nil),
		(*Ikev2InitiateSaInit)(nil),
		(*Ikev2InitiateSaInitReply)(nil),
		(*Ikev2PluginGetVersion)(nil),
		(*Ikev2PluginGetVersionReply)(nil),
		(*Ikev2ProfileAddDel)(nil),
		(*Ikev2ProfileAddDelReply)(nil),
		(*Ikev2ProfileSetAuth)(nil),
		(*Ikev2ProfileSetAuthReply)(nil),
		(*Ikev2ProfileSetID)(nil),
		(*Ikev2ProfileSetIDReply)(nil),
		(*Ikev2ProfileSetTs)(nil),
		(*Ikev2ProfileSetTsReply)(nil),
		(*Ikev2SetEspTransforms)(nil),
		(*Ikev2SetEspTransformsReply)(nil),
		(*Ikev2SetIkeTransforms)(nil),
		(*Ikev2SetIkeTransformsReply)(nil),
		(*Ikev2SetLocalKey)(nil),
		(*Ikev2SetLocalKeyReply)(nil),
		(*Ikev2SetResponder)(nil),
		(*Ikev2SetResponderReply)(nil),
		(*Ikev2SetSaLifetime)(nil),
		(*Ikev2SetSaLifetimeReply)(nil),
	}
}

// RPCService represents RPC service API for ikev2 module.
type RPCService interface {
	Ikev2InitiateDelChildSa(ctx context.Context, in *Ikev2InitiateDelChildSa) (*Ikev2InitiateDelChildSaReply, error)
	Ikev2InitiateDelIkeSa(ctx context.Context, in *Ikev2InitiateDelIkeSa) (*Ikev2InitiateDelIkeSaReply, error)
	Ikev2InitiateRekeyChildSa(ctx context.Context, in *Ikev2InitiateRekeyChildSa) (*Ikev2InitiateRekeyChildSaReply, error)
	Ikev2InitiateSaInit(ctx context.Context, in *Ikev2InitiateSaInit) (*Ikev2InitiateSaInitReply, error)
	Ikev2PluginGetVersion(ctx context.Context, in *Ikev2PluginGetVersion) (*Ikev2PluginGetVersionReply, error)
	Ikev2ProfileAddDel(ctx context.Context, in *Ikev2ProfileAddDel) (*Ikev2ProfileAddDelReply, error)
	Ikev2ProfileSetAuth(ctx context.Context, in *Ikev2ProfileSetAuth) (*Ikev2ProfileSetAuthReply, error)
	Ikev2ProfileSetID(ctx context.Context, in *Ikev2ProfileSetID) (*Ikev2ProfileSetIDReply, error)
	Ikev2ProfileSetTs(ctx context.Context, in *Ikev2ProfileSetTs) (*Ikev2ProfileSetTsReply, error)
	Ikev2SetEspTransforms(ctx context.Context, in *Ikev2SetEspTransforms) (*Ikev2SetEspTransformsReply, error)
	Ikev2SetIkeTransforms(ctx context.Context, in *Ikev2SetIkeTransforms) (*Ikev2SetIkeTransformsReply, error)
	Ikev2SetLocalKey(ctx context.Context, in *Ikev2SetLocalKey) (*Ikev2SetLocalKeyReply, error)
	Ikev2SetResponder(ctx context.Context, in *Ikev2SetResponder) (*Ikev2SetResponderReply, error)
	Ikev2SetSaLifetime(ctx context.Context, in *Ikev2SetSaLifetime) (*Ikev2SetSaLifetimeReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) Ikev2InitiateDelChildSa(ctx context.Context, in *Ikev2InitiateDelChildSa) (*Ikev2InitiateDelChildSaReply, error) {
	out := new(Ikev2InitiateDelChildSaReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2InitiateDelIkeSa(ctx context.Context, in *Ikev2InitiateDelIkeSa) (*Ikev2InitiateDelIkeSaReply, error) {
	out := new(Ikev2InitiateDelIkeSaReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2InitiateRekeyChildSa(ctx context.Context, in *Ikev2InitiateRekeyChildSa) (*Ikev2InitiateRekeyChildSaReply, error) {
	out := new(Ikev2InitiateRekeyChildSaReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2InitiateSaInit(ctx context.Context, in *Ikev2InitiateSaInit) (*Ikev2InitiateSaInitReply, error) {
	out := new(Ikev2InitiateSaInitReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2PluginGetVersion(ctx context.Context, in *Ikev2PluginGetVersion) (*Ikev2PluginGetVersionReply, error) {
	out := new(Ikev2PluginGetVersionReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2ProfileAddDel(ctx context.Context, in *Ikev2ProfileAddDel) (*Ikev2ProfileAddDelReply, error) {
	out := new(Ikev2ProfileAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2ProfileSetAuth(ctx context.Context, in *Ikev2ProfileSetAuth) (*Ikev2ProfileSetAuthReply, error) {
	out := new(Ikev2ProfileSetAuthReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2ProfileSetID(ctx context.Context, in *Ikev2ProfileSetID) (*Ikev2ProfileSetIDReply, error) {
	out := new(Ikev2ProfileSetIDReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2ProfileSetTs(ctx context.Context, in *Ikev2ProfileSetTs) (*Ikev2ProfileSetTsReply, error) {
	out := new(Ikev2ProfileSetTsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2SetEspTransforms(ctx context.Context, in *Ikev2SetEspTransforms) (*Ikev2SetEspTransformsReply, error) {
	out := new(Ikev2SetEspTransformsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2SetIkeTransforms(ctx context.Context, in *Ikev2SetIkeTransforms) (*Ikev2SetIkeTransformsReply, error) {
	out := new(Ikev2SetIkeTransformsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2SetLocalKey(ctx context.Context, in *Ikev2SetLocalKey) (*Ikev2SetLocalKeyReply, error) {
	out := new(Ikev2SetLocalKeyReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2SetResponder(ctx context.Context, in *Ikev2SetResponder) (*Ikev2SetResponderReply, error) {
	out := new(Ikev2SetResponderReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2SetSaLifetime(ctx context.Context, in *Ikev2SetSaLifetime) (*Ikev2SetSaLifetimeReply, error) {
	out := new(Ikev2SetSaLifetimeReply)
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
