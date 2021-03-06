// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/plugins/map.api.json

/*
Package maps is a generated VPP binary API for 'map' module.

It consists of:
	 10 enums
	  6 aliases
	  6 types
	  1 union
	 30 messages
	 15 services
*/
package maps

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
	ModuleName = "map"
	// APIVersion is the API version of this module.
	APIVersion = "4.1.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x82b79829
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
	IF_API_TYPE_HARDWARE IfType = 0
	IF_API_TYPE_SUB      IfType = 1
	IF_API_TYPE_P2P      IfType = 2
	IF_API_TYPE_PIPE     IfType = 3
)

var IfType_name = map[uint32]string{
	0: "IF_API_TYPE_HARDWARE",
	1: "IF_API_TYPE_SUB",
	2: "IF_API_TYPE_P2P",
	3: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 0,
	"IF_API_TYPE_SUB":      1,
	"IF_API_TYPE_P2P":      2,
	"IF_API_TYPE_PIPE":     3,
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
	MTU_PROTO_API_L3   MtuProto = 0
	MTU_PROTO_API_IP4  MtuProto = 1
	MTU_PROTO_API_IP6  MtuProto = 2
	MTU_PROTO_API_MPLS MtuProto = 3
)

var MtuProto_name = map[uint32]string{
	0: "MTU_PROTO_API_L3",
	1: "MTU_PROTO_API_IP4",
	2: "MTU_PROTO_API_IP6",
	3: "MTU_PROTO_API_MPLS",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   0,
	"MTU_PROTO_API_IP4":  1,
	"MTU_PROTO_API_IP6":  2,
	"MTU_PROTO_API_MPLS": 3,
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

// MapAddDelRule represents VPP binary API message 'map_add_del_rule'.
type MapAddDelRule struct {
	Index  uint32
	IsAdd  bool
	IP6Dst IP6Address
	Psid   uint16
}

func (m *MapAddDelRule) Reset()                        { *m = MapAddDelRule{} }
func (*MapAddDelRule) GetMessageName() string          { return "map_add_del_rule" }
func (*MapAddDelRule) GetCrcString() string            { return "c65b32f7" }
func (*MapAddDelRule) GetMessageType() api.MessageType { return api.RequestMessage }

// MapAddDelRuleReply represents VPP binary API message 'map_add_del_rule_reply'.
type MapAddDelRuleReply struct {
	Retval int32
}

func (m *MapAddDelRuleReply) Reset()                        { *m = MapAddDelRuleReply{} }
func (*MapAddDelRuleReply) GetMessageName() string          { return "map_add_del_rule_reply" }
func (*MapAddDelRuleReply) GetCrcString() string            { return "e8d4e804" }
func (*MapAddDelRuleReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapAddDomain represents VPP binary API message 'map_add_domain'.
type MapAddDomain struct {
	IP6Prefix  IP6Prefix
	IP4Prefix  IP4Prefix
	IP6Src     IP6Prefix
	EaBitsLen  uint8
	PsidOffset uint8
	PsidLength uint8
	Mtu        uint16
	Tag        string `struc:"[64]byte"`
}

func (m *MapAddDomain) Reset()                        { *m = MapAddDomain{} }
func (*MapAddDomain) GetMessageName() string          { return "map_add_domain" }
func (*MapAddDomain) GetCrcString() string            { return "7a5a18c9" }
func (*MapAddDomain) GetMessageType() api.MessageType { return api.RequestMessage }

// MapAddDomainReply represents VPP binary API message 'map_add_domain_reply'.
type MapAddDomainReply struct {
	Index  uint32
	Retval int32
}

func (m *MapAddDomainReply) Reset()                        { *m = MapAddDomainReply{} }
func (*MapAddDomainReply) GetMessageName() string          { return "map_add_domain_reply" }
func (*MapAddDomainReply) GetCrcString() string            { return "3e6d4e2c" }
func (*MapAddDomainReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapDelDomain represents VPP binary API message 'map_del_domain'.
type MapDelDomain struct {
	Index uint32
}

func (m *MapDelDomain) Reset()                        { *m = MapDelDomain{} }
func (*MapDelDomain) GetMessageName() string          { return "map_del_domain" }
func (*MapDelDomain) GetCrcString() string            { return "8ac76db6" }
func (*MapDelDomain) GetMessageType() api.MessageType { return api.RequestMessage }

// MapDelDomainReply represents VPP binary API message 'map_del_domain_reply'.
type MapDelDomainReply struct {
	Retval int32
}

func (m *MapDelDomainReply) Reset()                        { *m = MapDelDomainReply{} }
func (*MapDelDomainReply) GetMessageName() string          { return "map_del_domain_reply" }
func (*MapDelDomainReply) GetCrcString() string            { return "e8d4e804" }
func (*MapDelDomainReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapDomainDetails represents VPP binary API message 'map_domain_details'.
type MapDomainDetails struct {
	DomainIndex uint32
	IP6Prefix   IP6Prefix
	IP4Prefix   IP4Prefix
	IP6Src      IP6Prefix
	EaBitsLen   uint8
	PsidOffset  uint8
	PsidLength  uint8
	Flags       uint8
	Mtu         uint16
	Tag         string `struc:"[64]byte"`
}

func (m *MapDomainDetails) Reset()                        { *m = MapDomainDetails{} }
func (*MapDomainDetails) GetMessageName() string          { return "map_domain_details" }
func (*MapDomainDetails) GetCrcString() string            { return "fc1859dd" }
func (*MapDomainDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapDomainDump represents VPP binary API message 'map_domain_dump'.
type MapDomainDump struct{}

func (m *MapDomainDump) Reset()                        { *m = MapDomainDump{} }
func (*MapDomainDump) GetMessageName() string          { return "map_domain_dump" }
func (*MapDomainDump) GetCrcString() string            { return "51077d14" }
func (*MapDomainDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MapIfEnableDisable represents VPP binary API message 'map_if_enable_disable'.
type MapIfEnableDisable struct {
	SwIfIndex     InterfaceIndex
	IsEnable      bool
	IsTranslation bool
}

func (m *MapIfEnableDisable) Reset()                        { *m = MapIfEnableDisable{} }
func (*MapIfEnableDisable) GetMessageName() string          { return "map_if_enable_disable" }
func (*MapIfEnableDisable) GetCrcString() string            { return "59bb32f4" }
func (*MapIfEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// MapIfEnableDisableReply represents VPP binary API message 'map_if_enable_disable_reply'.
type MapIfEnableDisableReply struct {
	Retval int32
}

func (m *MapIfEnableDisableReply) Reset()                        { *m = MapIfEnableDisableReply{} }
func (*MapIfEnableDisableReply) GetMessageName() string          { return "map_if_enable_disable_reply" }
func (*MapIfEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*MapIfEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamAddDelPreResolve represents VPP binary API message 'map_param_add_del_pre_resolve'.
type MapParamAddDelPreResolve struct {
	IsAdd        bool
	IP4NhAddress IP4Address
	IP6NhAddress IP6Address
}

func (m *MapParamAddDelPreResolve) Reset()                        { *m = MapParamAddDelPreResolve{} }
func (*MapParamAddDelPreResolve) GetMessageName() string          { return "map_param_add_del_pre_resolve" }
func (*MapParamAddDelPreResolve) GetCrcString() string            { return "17008c66" }
func (*MapParamAddDelPreResolve) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamAddDelPreResolveReply represents VPP binary API message 'map_param_add_del_pre_resolve_reply'.
type MapParamAddDelPreResolveReply struct {
	Retval int32
}

func (m *MapParamAddDelPreResolveReply) Reset() { *m = MapParamAddDelPreResolveReply{} }
func (*MapParamAddDelPreResolveReply) GetMessageName() string {
	return "map_param_add_del_pre_resolve_reply"
}
func (*MapParamAddDelPreResolveReply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamAddDelPreResolveReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamGet represents VPP binary API message 'map_param_get'.
type MapParamGet struct{}

func (m *MapParamGet) Reset()                        { *m = MapParamGet{} }
func (*MapParamGet) GetMessageName() string          { return "map_param_get" }
func (*MapParamGet) GetCrcString() string            { return "51077d14" }
func (*MapParamGet) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamGetReply represents VPP binary API message 'map_param_get_reply'.
type MapParamGetReply struct {
	Retval                 int32
	FragInner              uint8
	FragIgnoreDf           uint8
	ICMPIP4ErrRelaySrc     IP4Address
	ICMP6EnableUnreachable bool
	IP4NhAddress           IP4Address
	IP6NhAddress           IP6Address
	IP4LifetimeMs          uint16
	IP4PoolSize            uint16
	IP4Buffers             uint32
	IP4HtRatio             float64
	SecCheckEnable         bool
	SecCheckFragments      bool
	TcCopy                 bool
	TcClass                uint8
}

func (m *MapParamGetReply) Reset()                        { *m = MapParamGetReply{} }
func (*MapParamGetReply) GetMessageName() string          { return "map_param_get_reply" }
func (*MapParamGetReply) GetCrcString() string            { return "28092156" }
func (*MapParamGetReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamSetFragmentation represents VPP binary API message 'map_param_set_fragmentation'.
type MapParamSetFragmentation struct {
	Inner    bool
	IgnoreDf bool
}

func (m *MapParamSetFragmentation) Reset()                        { *m = MapParamSetFragmentation{} }
func (*MapParamSetFragmentation) GetMessageName() string          { return "map_param_set_fragmentation" }
func (*MapParamSetFragmentation) GetCrcString() string            { return "9ff54d90" }
func (*MapParamSetFragmentation) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamSetFragmentationReply represents VPP binary API message 'map_param_set_fragmentation_reply'.
type MapParamSetFragmentationReply struct {
	Retval int32
}

func (m *MapParamSetFragmentationReply) Reset() { *m = MapParamSetFragmentationReply{} }
func (*MapParamSetFragmentationReply) GetMessageName() string {
	return "map_param_set_fragmentation_reply"
}
func (*MapParamSetFragmentationReply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamSetFragmentationReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamSetICMP represents VPP binary API message 'map_param_set_icmp'.
type MapParamSetICMP struct {
	IP4ErrRelaySrc IP4Address
}

func (m *MapParamSetICMP) Reset()                        { *m = MapParamSetICMP{} }
func (*MapParamSetICMP) GetMessageName() string          { return "map_param_set_icmp" }
func (*MapParamSetICMP) GetCrcString() string            { return "58210cbf" }
func (*MapParamSetICMP) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamSetICMP6 represents VPP binary API message 'map_param_set_icmp6'.
type MapParamSetICMP6 struct {
	EnableUnreachable bool
}

func (m *MapParamSetICMP6) Reset()                        { *m = MapParamSetICMP6{} }
func (*MapParamSetICMP6) GetMessageName() string          { return "map_param_set_icmp6" }
func (*MapParamSetICMP6) GetCrcString() string            { return "5d01f8c1" }
func (*MapParamSetICMP6) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamSetICMP6Reply represents VPP binary API message 'map_param_set_icmp6_reply'.
type MapParamSetICMP6Reply struct {
	Retval int32
}

func (m *MapParamSetICMP6Reply) Reset()                        { *m = MapParamSetICMP6Reply{} }
func (*MapParamSetICMP6Reply) GetMessageName() string          { return "map_param_set_icmp6_reply" }
func (*MapParamSetICMP6Reply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamSetICMP6Reply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamSetICMPReply represents VPP binary API message 'map_param_set_icmp_reply'.
type MapParamSetICMPReply struct {
	Retval int32
}

func (m *MapParamSetICMPReply) Reset()                        { *m = MapParamSetICMPReply{} }
func (*MapParamSetICMPReply) GetMessageName() string          { return "map_param_set_icmp_reply" }
func (*MapParamSetICMPReply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamSetICMPReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamSetSecurityCheck represents VPP binary API message 'map_param_set_security_check'.
type MapParamSetSecurityCheck struct {
	Enable    bool
	Fragments bool
}

func (m *MapParamSetSecurityCheck) Reset()                        { *m = MapParamSetSecurityCheck{} }
func (*MapParamSetSecurityCheck) GetMessageName() string          { return "map_param_set_security_check" }
func (*MapParamSetSecurityCheck) GetCrcString() string            { return "6abe9836" }
func (*MapParamSetSecurityCheck) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamSetSecurityCheckReply represents VPP binary API message 'map_param_set_security_check_reply'.
type MapParamSetSecurityCheckReply struct {
	Retval int32
}

func (m *MapParamSetSecurityCheckReply) Reset() { *m = MapParamSetSecurityCheckReply{} }
func (*MapParamSetSecurityCheckReply) GetMessageName() string {
	return "map_param_set_security_check_reply"
}
func (*MapParamSetSecurityCheckReply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamSetSecurityCheckReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamSetTCP represents VPP binary API message 'map_param_set_tcp'.
type MapParamSetTCP struct {
	TCPMss uint16
}

func (m *MapParamSetTCP) Reset()                        { *m = MapParamSetTCP{} }
func (*MapParamSetTCP) GetMessageName() string          { return "map_param_set_tcp" }
func (*MapParamSetTCP) GetCrcString() string            { return "87a825d9" }
func (*MapParamSetTCP) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamSetTCPReply represents VPP binary API message 'map_param_set_tcp_reply'.
type MapParamSetTCPReply struct {
	Retval int32
}

func (m *MapParamSetTCPReply) Reset()                        { *m = MapParamSetTCPReply{} }
func (*MapParamSetTCPReply) GetMessageName() string          { return "map_param_set_tcp_reply" }
func (*MapParamSetTCPReply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamSetTCPReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapParamSetTrafficClass represents VPP binary API message 'map_param_set_traffic_class'.
type MapParamSetTrafficClass struct {
	Copy    bool
	TcClass uint8
}

func (m *MapParamSetTrafficClass) Reset()                        { *m = MapParamSetTrafficClass{} }
func (*MapParamSetTrafficClass) GetMessageName() string          { return "map_param_set_traffic_class" }
func (*MapParamSetTrafficClass) GetCrcString() string            { return "9cac455c" }
func (*MapParamSetTrafficClass) GetMessageType() api.MessageType { return api.RequestMessage }

// MapParamSetTrafficClassReply represents VPP binary API message 'map_param_set_traffic_class_reply'.
type MapParamSetTrafficClassReply struct {
	Retval int32
}

func (m *MapParamSetTrafficClassReply) Reset() { *m = MapParamSetTrafficClassReply{} }
func (*MapParamSetTrafficClassReply) GetMessageName() string {
	return "map_param_set_traffic_class_reply"
}
func (*MapParamSetTrafficClassReply) GetCrcString() string            { return "e8d4e804" }
func (*MapParamSetTrafficClassReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapRuleDetails represents VPP binary API message 'map_rule_details'.
type MapRuleDetails struct {
	IP6Dst IP6Address
	Psid   uint16
}

func (m *MapRuleDetails) Reset()                        { *m = MapRuleDetails{} }
func (*MapRuleDetails) GetMessageName() string          { return "map_rule_details" }
func (*MapRuleDetails) GetCrcString() string            { return "c7cbeea5" }
func (*MapRuleDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MapRuleDump represents VPP binary API message 'map_rule_dump'.
type MapRuleDump struct {
	DomainIndex uint32
}

func (m *MapRuleDump) Reset()                        { *m = MapRuleDump{} }
func (*MapRuleDump) GetMessageName() string          { return "map_rule_dump" }
func (*MapRuleDump) GetCrcString() string            { return "e43e6ff6" }
func (*MapRuleDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MapSummaryStats represents VPP binary API message 'map_summary_stats'.
type MapSummaryStats struct{}

func (m *MapSummaryStats) Reset()                        { *m = MapSummaryStats{} }
func (*MapSummaryStats) GetMessageName() string          { return "map_summary_stats" }
func (*MapSummaryStats) GetCrcString() string            { return "51077d14" }
func (*MapSummaryStats) GetMessageType() api.MessageType { return api.RequestMessage }

// MapSummaryStatsReply represents VPP binary API message 'map_summary_stats_reply'.
type MapSummaryStatsReply struct {
	Retval             int32
	TotalBindings      uint64
	TotalPkts          []uint64 `struc:"[2]uint64"`
	TotalBytes         []uint64 `struc:"[2]uint64"`
	TotalIP4Fragments  uint64
	TotalSecurityCheck []uint64 `struc:"[2]uint64"`
}

func (m *MapSummaryStatsReply) Reset()                        { *m = MapSummaryStatsReply{} }
func (*MapSummaryStatsReply) GetMessageName() string          { return "map_summary_stats_reply" }
func (*MapSummaryStatsReply) GetCrcString() string            { return "0e4ace0e" }
func (*MapSummaryStatsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*MapAddDelRule)(nil), "map.MapAddDelRule")
	api.RegisterMessage((*MapAddDelRuleReply)(nil), "map.MapAddDelRuleReply")
	api.RegisterMessage((*MapAddDomain)(nil), "map.MapAddDomain")
	api.RegisterMessage((*MapAddDomainReply)(nil), "map.MapAddDomainReply")
	api.RegisterMessage((*MapDelDomain)(nil), "map.MapDelDomain")
	api.RegisterMessage((*MapDelDomainReply)(nil), "map.MapDelDomainReply")
	api.RegisterMessage((*MapDomainDetails)(nil), "map.MapDomainDetails")
	api.RegisterMessage((*MapDomainDump)(nil), "map.MapDomainDump")
	api.RegisterMessage((*MapIfEnableDisable)(nil), "map.MapIfEnableDisable")
	api.RegisterMessage((*MapIfEnableDisableReply)(nil), "map.MapIfEnableDisableReply")
	api.RegisterMessage((*MapParamAddDelPreResolve)(nil), "map.MapParamAddDelPreResolve")
	api.RegisterMessage((*MapParamAddDelPreResolveReply)(nil), "map.MapParamAddDelPreResolveReply")
	api.RegisterMessage((*MapParamGet)(nil), "map.MapParamGet")
	api.RegisterMessage((*MapParamGetReply)(nil), "map.MapParamGetReply")
	api.RegisterMessage((*MapParamSetFragmentation)(nil), "map.MapParamSetFragmentation")
	api.RegisterMessage((*MapParamSetFragmentationReply)(nil), "map.MapParamSetFragmentationReply")
	api.RegisterMessage((*MapParamSetICMP)(nil), "map.MapParamSetICMP")
	api.RegisterMessage((*MapParamSetICMP6)(nil), "map.MapParamSetICMP6")
	api.RegisterMessage((*MapParamSetICMP6Reply)(nil), "map.MapParamSetICMP6Reply")
	api.RegisterMessage((*MapParamSetICMPReply)(nil), "map.MapParamSetICMPReply")
	api.RegisterMessage((*MapParamSetSecurityCheck)(nil), "map.MapParamSetSecurityCheck")
	api.RegisterMessage((*MapParamSetSecurityCheckReply)(nil), "map.MapParamSetSecurityCheckReply")
	api.RegisterMessage((*MapParamSetTCP)(nil), "map.MapParamSetTCP")
	api.RegisterMessage((*MapParamSetTCPReply)(nil), "map.MapParamSetTCPReply")
	api.RegisterMessage((*MapParamSetTrafficClass)(nil), "map.MapParamSetTrafficClass")
	api.RegisterMessage((*MapParamSetTrafficClassReply)(nil), "map.MapParamSetTrafficClassReply")
	api.RegisterMessage((*MapRuleDetails)(nil), "map.MapRuleDetails")
	api.RegisterMessage((*MapRuleDump)(nil), "map.MapRuleDump")
	api.RegisterMessage((*MapSummaryStats)(nil), "map.MapSummaryStats")
	api.RegisterMessage((*MapSummaryStatsReply)(nil), "map.MapSummaryStatsReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MapAddDelRule)(nil),
		(*MapAddDelRuleReply)(nil),
		(*MapAddDomain)(nil),
		(*MapAddDomainReply)(nil),
		(*MapDelDomain)(nil),
		(*MapDelDomainReply)(nil),
		(*MapDomainDetails)(nil),
		(*MapDomainDump)(nil),
		(*MapIfEnableDisable)(nil),
		(*MapIfEnableDisableReply)(nil),
		(*MapParamAddDelPreResolve)(nil),
		(*MapParamAddDelPreResolveReply)(nil),
		(*MapParamGet)(nil),
		(*MapParamGetReply)(nil),
		(*MapParamSetFragmentation)(nil),
		(*MapParamSetFragmentationReply)(nil),
		(*MapParamSetICMP)(nil),
		(*MapParamSetICMP6)(nil),
		(*MapParamSetICMP6Reply)(nil),
		(*MapParamSetICMPReply)(nil),
		(*MapParamSetSecurityCheck)(nil),
		(*MapParamSetSecurityCheckReply)(nil),
		(*MapParamSetTCP)(nil),
		(*MapParamSetTCPReply)(nil),
		(*MapParamSetTrafficClass)(nil),
		(*MapParamSetTrafficClassReply)(nil),
		(*MapRuleDetails)(nil),
		(*MapRuleDump)(nil),
		(*MapSummaryStats)(nil),
		(*MapSummaryStatsReply)(nil),
	}
}

// RPCService represents RPC service API for map module.
type RPCService interface {
	DumpMapDomain(ctx context.Context, in *MapDomainDump) (RPCService_DumpMapDomainClient, error)
	DumpMapRule(ctx context.Context, in *MapRuleDump) (RPCService_DumpMapRuleClient, error)
	MapAddDelRule(ctx context.Context, in *MapAddDelRule) (*MapAddDelRuleReply, error)
	MapAddDomain(ctx context.Context, in *MapAddDomain) (*MapAddDomainReply, error)
	MapDelDomain(ctx context.Context, in *MapDelDomain) (*MapDelDomainReply, error)
	MapIfEnableDisable(ctx context.Context, in *MapIfEnableDisable) (*MapIfEnableDisableReply, error)
	MapParamAddDelPreResolve(ctx context.Context, in *MapParamAddDelPreResolve) (*MapParamAddDelPreResolveReply, error)
	MapParamGet(ctx context.Context, in *MapParamGet) (*MapParamGetReply, error)
	MapParamSetFragmentation(ctx context.Context, in *MapParamSetFragmentation) (*MapParamSetFragmentationReply, error)
	MapParamSetICMP(ctx context.Context, in *MapParamSetICMP) (*MapParamSetICMPReply, error)
	MapParamSetICMP6(ctx context.Context, in *MapParamSetICMP6) (*MapParamSetICMP6Reply, error)
	MapParamSetSecurityCheck(ctx context.Context, in *MapParamSetSecurityCheck) (*MapParamSetSecurityCheckReply, error)
	MapParamSetTCP(ctx context.Context, in *MapParamSetTCP) (*MapParamSetTCPReply, error)
	MapParamSetTrafficClass(ctx context.Context, in *MapParamSetTrafficClass) (*MapParamSetTrafficClassReply, error)
	MapSummaryStats(ctx context.Context, in *MapSummaryStats) (*MapSummaryStatsReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMapDomain(ctx context.Context, in *MapDomainDump) (RPCService_DumpMapDomainClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMapDomainClient{stream}
	return x, nil
}

type RPCService_DumpMapDomainClient interface {
	Recv() (*MapDomainDetails, error)
}

type serviceClient_DumpMapDomainClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMapDomainClient) Recv() (*MapDomainDetails, error) {
	m := new(MapDomainDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMapRule(ctx context.Context, in *MapRuleDump) (RPCService_DumpMapRuleClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMapRuleClient{stream}
	return x, nil
}

type RPCService_DumpMapRuleClient interface {
	Recv() (*MapRuleDetails, error)
}

type serviceClient_DumpMapRuleClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMapRuleClient) Recv() (*MapRuleDetails, error) {
	m := new(MapRuleDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MapAddDelRule(ctx context.Context, in *MapAddDelRule) (*MapAddDelRuleReply, error) {
	out := new(MapAddDelRuleReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapAddDomain(ctx context.Context, in *MapAddDomain) (*MapAddDomainReply, error) {
	out := new(MapAddDomainReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapDelDomain(ctx context.Context, in *MapDelDomain) (*MapDelDomainReply, error) {
	out := new(MapDelDomainReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapIfEnableDisable(ctx context.Context, in *MapIfEnableDisable) (*MapIfEnableDisableReply, error) {
	out := new(MapIfEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamAddDelPreResolve(ctx context.Context, in *MapParamAddDelPreResolve) (*MapParamAddDelPreResolveReply, error) {
	out := new(MapParamAddDelPreResolveReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamGet(ctx context.Context, in *MapParamGet) (*MapParamGetReply, error) {
	out := new(MapParamGetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamSetFragmentation(ctx context.Context, in *MapParamSetFragmentation) (*MapParamSetFragmentationReply, error) {
	out := new(MapParamSetFragmentationReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamSetICMP(ctx context.Context, in *MapParamSetICMP) (*MapParamSetICMPReply, error) {
	out := new(MapParamSetICMPReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamSetICMP6(ctx context.Context, in *MapParamSetICMP6) (*MapParamSetICMP6Reply, error) {
	out := new(MapParamSetICMP6Reply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamSetSecurityCheck(ctx context.Context, in *MapParamSetSecurityCheck) (*MapParamSetSecurityCheckReply, error) {
	out := new(MapParamSetSecurityCheckReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamSetTCP(ctx context.Context, in *MapParamSetTCP) (*MapParamSetTCPReply, error) {
	out := new(MapParamSetTCPReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapParamSetTrafficClass(ctx context.Context, in *MapParamSetTrafficClass) (*MapParamSetTrafficClassReply, error) {
	out := new(MapParamSetTrafficClassReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MapSummaryStats(ctx context.Context, in *MapSummaryStats) (*MapSummaryStatsReply, error) {
	out := new(MapSummaryStatsReply)
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
