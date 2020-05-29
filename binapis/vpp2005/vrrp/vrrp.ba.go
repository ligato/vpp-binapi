// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/plugins/vrrp.api.json

/*
Package vrrp is a generated VPP binary API for 'vrrp' module.

It consists of:
	 12 enums
	  7 aliases
	 11 types
	  1 union
	 14 messages
	  7 services
*/
package vrrp

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
	ModuleName = "vrrp"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x1903f1f1
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

// VrrpVrFlags represents VPP binary API enum 'vrrp_vr_flags'.
type VrrpVrFlags uint32

const (
	VRRP_API_VR_PREEMPT VrrpVrFlags = 1
	VRRP_API_VR_ACCEPT  VrrpVrFlags = 2
	VRRP_API_VR_UNICAST VrrpVrFlags = 4
	VRRP_API_VR_IPV6    VrrpVrFlags = 8
)

var VrrpVrFlags_name = map[uint32]string{
	1: "VRRP_API_VR_PREEMPT",
	2: "VRRP_API_VR_ACCEPT",
	4: "VRRP_API_VR_UNICAST",
	8: "VRRP_API_VR_IPV6",
}

var VrrpVrFlags_value = map[string]uint32{
	"VRRP_API_VR_PREEMPT": 1,
	"VRRP_API_VR_ACCEPT":  2,
	"VRRP_API_VR_UNICAST": 4,
	"VRRP_API_VR_IPV6":    8,
}

func (x VrrpVrFlags) String() string {
	s, ok := VrrpVrFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// VrrpVrState represents VPP binary API enum 'vrrp_vr_state'.
type VrrpVrState uint32

const (
	VRRP_API_VR_STATE_INIT      VrrpVrState = 0
	VRRP_API_VR_STATE_BACKUP    VrrpVrState = 1
	VRRP_API_VR_STATE_MASTER    VrrpVrState = 2
	VRRP_API_VR_STATE_INTF_DOWN VrrpVrState = 3
)

var VrrpVrState_name = map[uint32]string{
	0: "VRRP_API_VR_STATE_INIT",
	1: "VRRP_API_VR_STATE_BACKUP",
	2: "VRRP_API_VR_STATE_MASTER",
	3: "VRRP_API_VR_STATE_INTF_DOWN",
}

var VrrpVrState_value = map[string]uint32{
	"VRRP_API_VR_STATE_INIT":      0,
	"VRRP_API_VR_STATE_BACKUP":    1,
	"VRRP_API_VR_STATE_MASTER":    2,
	"VRRP_API_VR_STATE_INTF_DOWN": 3,
}

func (x VrrpVrState) String() string {
	s, ok := VrrpVrState_name[uint32(x)]
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

// VrrpVrConf represents VPP binary API type 'vrrp_vr_conf'.
type VrrpVrConf struct {
	SwIfIndex InterfaceIndex
	VrID      uint8
	Priority  uint8
	Interval  uint16
	Flags     VrrpVrFlags
}

func (*VrrpVrConf) GetTypeName() string { return "vrrp_vr_conf" }

// VrrpVrKey represents VPP binary API type 'vrrp_vr_key'.
type VrrpVrKey struct {
	SwIfIndex InterfaceIndex
	VrID      uint8
	IsIPv6    uint8
}

func (*VrrpVrKey) GetTypeName() string { return "vrrp_vr_key" }

// VrrpVrRuntime represents VPP binary API type 'vrrp_vr_runtime'.
type VrrpVrRuntime struct {
	State         VrrpVrState
	MasterAdvInt  uint16
	Skew          uint16
	MasterDownInt uint16
	Mac           MacAddress
	Tracking      VrrpVrTracking
}

func (*VrrpVrRuntime) GetTypeName() string { return "vrrp_vr_runtime" }

// VrrpVrTrackIf represents VPP binary API type 'vrrp_vr_track_if'.
type VrrpVrTrackIf struct {
	SwIfIndex InterfaceIndex
	Priority  uint8
}

func (*VrrpVrTrackIf) GetTypeName() string { return "vrrp_vr_track_if" }

// VrrpVrTracking represents VPP binary API type 'vrrp_vr_tracking'.
type VrrpVrTracking struct {
	InterfacesDec uint32
	Priority      uint8
}

func (*VrrpVrTracking) GetTypeName() string { return "vrrp_vr_tracking" }

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

// VrrpVrAddDel represents VPP binary API message 'vrrp_vr_add_del'.
type VrrpVrAddDel struct {
	IsAdd     uint8
	SwIfIndex InterfaceIndex
	VrID      uint8
	Priority  uint8
	Interval  uint16
	Flags     VrrpVrFlags
	NAddrs    uint8 `struc:"sizeof=Addrs"`
	Addrs     []Address
}

func (m *VrrpVrAddDel) Reset()                        { *m = VrrpVrAddDel{} }
func (*VrrpVrAddDel) GetMessageName() string          { return "vrrp_vr_add_del" }
func (*VrrpVrAddDel) GetCrcString() string            { return "6dc4b881" }
func (*VrrpVrAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrAddDelReply represents VPP binary API message 'vrrp_vr_add_del_reply'.
type VrrpVrAddDelReply struct {
	Retval int32
}

func (m *VrrpVrAddDelReply) Reset()                        { *m = VrrpVrAddDelReply{} }
func (*VrrpVrAddDelReply) GetMessageName() string          { return "vrrp_vr_add_del_reply" }
func (*VrrpVrAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*VrrpVrAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VrrpVrDetails represents VPP binary API message 'vrrp_vr_details'.
type VrrpVrDetails struct {
	Config  VrrpVrConf
	Runtime VrrpVrRuntime
	NAddrs  uint8 `struc:"sizeof=Addrs"`
	Addrs   []Address
}

func (m *VrrpVrDetails) Reset()                        { *m = VrrpVrDetails{} }
func (*VrrpVrDetails) GetMessageName() string          { return "vrrp_vr_details" }
func (*VrrpVrDetails) GetCrcString() string            { return "0412fa71" }
func (*VrrpVrDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// VrrpVrDump represents VPP binary API message 'vrrp_vr_dump'.
type VrrpVrDump struct {
	SwIfIndex InterfaceIndex
}

func (m *VrrpVrDump) Reset()                        { *m = VrrpVrDump{} }
func (*VrrpVrDump) GetMessageName() string          { return "vrrp_vr_dump" }
func (*VrrpVrDump) GetCrcString() string            { return "f9e6675e" }
func (*VrrpVrDump) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrPeerDetails represents VPP binary API message 'vrrp_vr_peer_details'.
type VrrpVrPeerDetails struct {
	SwIfIndex  InterfaceIndex
	VrID       uint8
	IsIPv6     uint8
	NPeerAddrs uint8 `struc:"sizeof=PeerAddrs"`
	PeerAddrs  []Address
}

func (m *VrrpVrPeerDetails) Reset()                        { *m = VrrpVrPeerDetails{} }
func (*VrrpVrPeerDetails) GetMessageName() string          { return "vrrp_vr_peer_details" }
func (*VrrpVrPeerDetails) GetCrcString() string            { return "abd9145e" }
func (*VrrpVrPeerDetails) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrPeerDump represents VPP binary API message 'vrrp_vr_peer_dump'.
type VrrpVrPeerDump struct {
	SwIfIndex InterfaceIndex
	IsIPv6    uint8
	VrID      uint8
}

func (m *VrrpVrPeerDump) Reset()                        { *m = VrrpVrPeerDump{} }
func (*VrrpVrPeerDump) GetMessageName() string          { return "vrrp_vr_peer_dump" }
func (*VrrpVrPeerDump) GetCrcString() string            { return "6fa3f7c4" }
func (*VrrpVrPeerDump) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrSetPeers represents VPP binary API message 'vrrp_vr_set_peers'.
type VrrpVrSetPeers struct {
	SwIfIndex InterfaceIndex
	VrID      uint8
	IsIPv6    uint8
	NAddrs    uint8 `struc:"sizeof=Addrs"`
	Addrs     []Address
}

func (m *VrrpVrSetPeers) Reset()                        { *m = VrrpVrSetPeers{} }
func (*VrrpVrSetPeers) GetMessageName() string          { return "vrrp_vr_set_peers" }
func (*VrrpVrSetPeers) GetCrcString() string            { return "baa2e52b" }
func (*VrrpVrSetPeers) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrSetPeersReply represents VPP binary API message 'vrrp_vr_set_peers_reply'.
type VrrpVrSetPeersReply struct {
	Retval int32
}

func (m *VrrpVrSetPeersReply) Reset()                        { *m = VrrpVrSetPeersReply{} }
func (*VrrpVrSetPeersReply) GetMessageName() string          { return "vrrp_vr_set_peers_reply" }
func (*VrrpVrSetPeersReply) GetCrcString() string            { return "e8d4e804" }
func (*VrrpVrSetPeersReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VrrpVrStartStop represents VPP binary API message 'vrrp_vr_start_stop'.
type VrrpVrStartStop struct {
	SwIfIndex InterfaceIndex
	VrID      uint8
	IsIPv6    uint8
	IsStart   uint8
}

func (m *VrrpVrStartStop) Reset()                        { *m = VrrpVrStartStop{} }
func (*VrrpVrStartStop) GetMessageName() string          { return "vrrp_vr_start_stop" }
func (*VrrpVrStartStop) GetCrcString() string            { return "0662a3b7" }
func (*VrrpVrStartStop) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrStartStopReply represents VPP binary API message 'vrrp_vr_start_stop_reply'.
type VrrpVrStartStopReply struct {
	Retval int32
}

func (m *VrrpVrStartStopReply) Reset()                        { *m = VrrpVrStartStopReply{} }
func (*VrrpVrStartStopReply) GetMessageName() string          { return "vrrp_vr_start_stop_reply" }
func (*VrrpVrStartStopReply) GetCrcString() string            { return "e8d4e804" }
func (*VrrpVrStartStopReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VrrpVrTrackIfAddDel represents VPP binary API message 'vrrp_vr_track_if_add_del'.
type VrrpVrTrackIfAddDel struct {
	SwIfIndex InterfaceIndex
	IsIPv6    uint8
	VrID      uint8
	IsAdd     uint8
	NIfs      uint8 `struc:"sizeof=Ifs"`
	Ifs       []VrrpVrTrackIf
}

func (m *VrrpVrTrackIfAddDel) Reset()                        { *m = VrrpVrTrackIfAddDel{} }
func (*VrrpVrTrackIfAddDel) GetMessageName() string          { return "vrrp_vr_track_if_add_del" }
func (*VrrpVrTrackIfAddDel) GetCrcString() string            { return "337f4ba4" }
func (*VrrpVrTrackIfAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrTrackIfAddDelReply represents VPP binary API message 'vrrp_vr_track_if_add_del_reply'.
type VrrpVrTrackIfAddDelReply struct {
	Retval int32
}

func (m *VrrpVrTrackIfAddDelReply) Reset()                        { *m = VrrpVrTrackIfAddDelReply{} }
func (*VrrpVrTrackIfAddDelReply) GetMessageName() string          { return "vrrp_vr_track_if_add_del_reply" }
func (*VrrpVrTrackIfAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*VrrpVrTrackIfAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VrrpVrTrackIfDetails represents VPP binary API message 'vrrp_vr_track_if_details'.
type VrrpVrTrackIfDetails struct {
	SwIfIndex InterfaceIndex
	VrID      uint8
	IsIPv6    uint8
	NIfs      uint8 `struc:"sizeof=Ifs"`
	Ifs       []VrrpVrTrackIf
}

func (m *VrrpVrTrackIfDetails) Reset()                        { *m = VrrpVrTrackIfDetails{} }
func (*VrrpVrTrackIfDetails) GetMessageName() string          { return "vrrp_vr_track_if_details" }
func (*VrrpVrTrackIfDetails) GetCrcString() string            { return "99bcca9c" }
func (*VrrpVrTrackIfDetails) GetMessageType() api.MessageType { return api.RequestMessage }

// VrrpVrTrackIfDump represents VPP binary API message 'vrrp_vr_track_if_dump'.
type VrrpVrTrackIfDump struct {
	SwIfIndex InterfaceIndex
	IsIPv6    uint8
	VrID      uint8
	DumpAll   uint8
}

func (m *VrrpVrTrackIfDump) Reset()                        { *m = VrrpVrTrackIfDump{} }
func (*VrrpVrTrackIfDump) GetMessageName() string          { return "vrrp_vr_track_if_dump" }
func (*VrrpVrTrackIfDump) GetCrcString() string            { return "a34dfc6d" }
func (*VrrpVrTrackIfDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*VrrpVrAddDel)(nil), "vrrp.VrrpVrAddDel")
	api.RegisterMessage((*VrrpVrAddDelReply)(nil), "vrrp.VrrpVrAddDelReply")
	api.RegisterMessage((*VrrpVrDetails)(nil), "vrrp.VrrpVrDetails")
	api.RegisterMessage((*VrrpVrDump)(nil), "vrrp.VrrpVrDump")
	api.RegisterMessage((*VrrpVrPeerDetails)(nil), "vrrp.VrrpVrPeerDetails")
	api.RegisterMessage((*VrrpVrPeerDump)(nil), "vrrp.VrrpVrPeerDump")
	api.RegisterMessage((*VrrpVrSetPeers)(nil), "vrrp.VrrpVrSetPeers")
	api.RegisterMessage((*VrrpVrSetPeersReply)(nil), "vrrp.VrrpVrSetPeersReply")
	api.RegisterMessage((*VrrpVrStartStop)(nil), "vrrp.VrrpVrStartStop")
	api.RegisterMessage((*VrrpVrStartStopReply)(nil), "vrrp.VrrpVrStartStopReply")
	api.RegisterMessage((*VrrpVrTrackIfAddDel)(nil), "vrrp.VrrpVrTrackIfAddDel")
	api.RegisterMessage((*VrrpVrTrackIfAddDelReply)(nil), "vrrp.VrrpVrTrackIfAddDelReply")
	api.RegisterMessage((*VrrpVrTrackIfDetails)(nil), "vrrp.VrrpVrTrackIfDetails")
	api.RegisterMessage((*VrrpVrTrackIfDump)(nil), "vrrp.VrrpVrTrackIfDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*VrrpVrAddDel)(nil),
		(*VrrpVrAddDelReply)(nil),
		(*VrrpVrDetails)(nil),
		(*VrrpVrDump)(nil),
		(*VrrpVrPeerDetails)(nil),
		(*VrrpVrPeerDump)(nil),
		(*VrrpVrSetPeers)(nil),
		(*VrrpVrSetPeersReply)(nil),
		(*VrrpVrStartStop)(nil),
		(*VrrpVrStartStopReply)(nil),
		(*VrrpVrTrackIfAddDel)(nil),
		(*VrrpVrTrackIfAddDelReply)(nil),
		(*VrrpVrTrackIfDetails)(nil),
		(*VrrpVrTrackIfDump)(nil),
	}
}

// RPCService represents RPC service API for vrrp module.
type RPCService interface {
	DumpVrrpVr(ctx context.Context, in *VrrpVrDump) (RPCService_DumpVrrpVrClient, error)
	DumpVrrpVrPeer(ctx context.Context, in *VrrpVrPeerDump) (RPCService_DumpVrrpVrPeerClient, error)
	DumpVrrpVrTrackIf(ctx context.Context, in *VrrpVrTrackIfDump) (RPCService_DumpVrrpVrTrackIfClient, error)
	VrrpVrAddDel(ctx context.Context, in *VrrpVrAddDel) (*VrrpVrAddDelReply, error)
	VrrpVrSetPeers(ctx context.Context, in *VrrpVrSetPeers) (*VrrpVrSetPeersReply, error)
	VrrpVrStartStop(ctx context.Context, in *VrrpVrStartStop) (*VrrpVrStartStopReply, error)
	VrrpVrTrackIfAddDel(ctx context.Context, in *VrrpVrTrackIfAddDel) (*VrrpVrTrackIfAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpVrrpVr(ctx context.Context, in *VrrpVrDump) (RPCService_DumpVrrpVrClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpVrrpVrClient{stream}
	return x, nil
}

type RPCService_DumpVrrpVrClient interface {
	Recv() (*VrrpVrDetails, error)
}

type serviceClient_DumpVrrpVrClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpVrrpVrClient) Recv() (*VrrpVrDetails, error) {
	m := new(VrrpVrDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpVrrpVrPeer(ctx context.Context, in *VrrpVrPeerDump) (RPCService_DumpVrrpVrPeerClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpVrrpVrPeerClient{stream}
	return x, nil
}

type RPCService_DumpVrrpVrPeerClient interface {
	Recv() (*VrrpVrPeerDetails, error)
}

type serviceClient_DumpVrrpVrPeerClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpVrrpVrPeerClient) Recv() (*VrrpVrPeerDetails, error) {
	m := new(VrrpVrPeerDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpVrrpVrTrackIf(ctx context.Context, in *VrrpVrTrackIfDump) (RPCService_DumpVrrpVrTrackIfClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpVrrpVrTrackIfClient{stream}
	return x, nil
}

type RPCService_DumpVrrpVrTrackIfClient interface {
	Recv() (*VrrpVrTrackIfDetails, error)
}

type serviceClient_DumpVrrpVrTrackIfClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpVrrpVrTrackIfClient) Recv() (*VrrpVrTrackIfDetails, error) {
	m := new(VrrpVrTrackIfDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) VrrpVrAddDel(ctx context.Context, in *VrrpVrAddDel) (*VrrpVrAddDelReply, error) {
	out := new(VrrpVrAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrSetPeers(ctx context.Context, in *VrrpVrSetPeers) (*VrrpVrSetPeersReply, error) {
	out := new(VrrpVrSetPeersReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrStartStop(ctx context.Context, in *VrrpVrStartStop) (*VrrpVrStartStopReply, error) {
	out := new(VrrpVrStartStopReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrTrackIfAddDel(ctx context.Context, in *VrrpVrTrackIfAddDel) (*VrrpVrTrackIfAddDelReply, error) {
	out := new(VrrpVrTrackIfAddDelReply)
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
