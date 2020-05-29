// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/core/mpls.api.json

/*
Package mpls is a generated VPP binary API for 'mpls' module.

It consists of:
	 13 enums
	  6 aliases
	 12 types
	  1 union
	 16 messages
	  8 services
*/
package mpls

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
	ModuleName = "mpls"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd83a030f
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

// FibPathFlags represents VPP binary API enum 'fib_path_flags'.
type FibPathFlags uint32

const (
	FIB_API_PATH_FLAG_NONE                 FibPathFlags = 0
	FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED FibPathFlags = 1
	FIB_API_PATH_FLAG_RESOLVE_VIA_HOST     FibPathFlags = 2
	FIB_API_PATH_FLAG_POP_PW_CW            FibPathFlags = 4
)

var FibPathFlags_name = map[uint32]string{
	0: "FIB_API_PATH_FLAG_NONE",
	1: "FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED",
	2: "FIB_API_PATH_FLAG_RESOLVE_VIA_HOST",
	4: "FIB_API_PATH_FLAG_POP_PW_CW",
}

var FibPathFlags_value = map[string]uint32{
	"FIB_API_PATH_FLAG_NONE":                 0,
	"FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED": 1,
	"FIB_API_PATH_FLAG_RESOLVE_VIA_HOST":     2,
	"FIB_API_PATH_FLAG_POP_PW_CW":            4,
}

func (x FibPathFlags) String() string {
	s, ok := FibPathFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// FibPathNhProto represents VPP binary API enum 'fib_path_nh_proto'.
type FibPathNhProto uint32

const (
	FIB_API_PATH_NH_PROTO_IP4      FibPathNhProto = 0
	FIB_API_PATH_NH_PROTO_IP6      FibPathNhProto = 1
	FIB_API_PATH_NH_PROTO_MPLS     FibPathNhProto = 2
	FIB_API_PATH_NH_PROTO_ETHERNET FibPathNhProto = 3
	FIB_API_PATH_NH_PROTO_BIER     FibPathNhProto = 4
)

var FibPathNhProto_name = map[uint32]string{
	0: "FIB_API_PATH_NH_PROTO_IP4",
	1: "FIB_API_PATH_NH_PROTO_IP6",
	2: "FIB_API_PATH_NH_PROTO_MPLS",
	3: "FIB_API_PATH_NH_PROTO_ETHERNET",
	4: "FIB_API_PATH_NH_PROTO_BIER",
}

var FibPathNhProto_value = map[string]uint32{
	"FIB_API_PATH_NH_PROTO_IP4":      0,
	"FIB_API_PATH_NH_PROTO_IP6":      1,
	"FIB_API_PATH_NH_PROTO_MPLS":     2,
	"FIB_API_PATH_NH_PROTO_ETHERNET": 3,
	"FIB_API_PATH_NH_PROTO_BIER":     4,
}

func (x FibPathNhProto) String() string {
	s, ok := FibPathNhProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// FibPathType represents VPP binary API enum 'fib_path_type'.
type FibPathType uint32

const (
	FIB_API_PATH_TYPE_NORMAL        FibPathType = 0
	FIB_API_PATH_TYPE_LOCAL         FibPathType = 1
	FIB_API_PATH_TYPE_DROP          FibPathType = 2
	FIB_API_PATH_TYPE_UDP_ENCAP     FibPathType = 3
	FIB_API_PATH_TYPE_BIER_IMP      FibPathType = 4
	FIB_API_PATH_TYPE_ICMP_UNREACH  FibPathType = 5
	FIB_API_PATH_TYPE_ICMP_PROHIBIT FibPathType = 6
	FIB_API_PATH_TYPE_SOURCE_LOOKUP FibPathType = 7
	FIB_API_PATH_TYPE_DVR           FibPathType = 8
	FIB_API_PATH_TYPE_INTERFACE_RX  FibPathType = 9
	FIB_API_PATH_TYPE_CLASSIFY      FibPathType = 10
)

var FibPathType_name = map[uint32]string{
	0:  "FIB_API_PATH_TYPE_NORMAL",
	1:  "FIB_API_PATH_TYPE_LOCAL",
	2:  "FIB_API_PATH_TYPE_DROP",
	3:  "FIB_API_PATH_TYPE_UDP_ENCAP",
	4:  "FIB_API_PATH_TYPE_BIER_IMP",
	5:  "FIB_API_PATH_TYPE_ICMP_UNREACH",
	6:  "FIB_API_PATH_TYPE_ICMP_PROHIBIT",
	7:  "FIB_API_PATH_TYPE_SOURCE_LOOKUP",
	8:  "FIB_API_PATH_TYPE_DVR",
	9:  "FIB_API_PATH_TYPE_INTERFACE_RX",
	10: "FIB_API_PATH_TYPE_CLASSIFY",
}

var FibPathType_value = map[string]uint32{
	"FIB_API_PATH_TYPE_NORMAL":        0,
	"FIB_API_PATH_TYPE_LOCAL":         1,
	"FIB_API_PATH_TYPE_DROP":          2,
	"FIB_API_PATH_TYPE_UDP_ENCAP":     3,
	"FIB_API_PATH_TYPE_BIER_IMP":      4,
	"FIB_API_PATH_TYPE_ICMP_UNREACH":  5,
	"FIB_API_PATH_TYPE_ICMP_PROHIBIT": 6,
	"FIB_API_PATH_TYPE_SOURCE_LOOKUP": 7,
	"FIB_API_PATH_TYPE_DVR":           8,
	"FIB_API_PATH_TYPE_INTERFACE_RX":  9,
	"FIB_API_PATH_TYPE_CLASSIFY":      10,
}

func (x FibPathType) String() string {
	s, ok := FibPathType_name[uint32(x)]
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

// FibMplsLabel represents VPP binary API type 'fib_mpls_label'.
type FibMplsLabel struct {
	IsUniform uint8
	Label     uint32
	TTL       uint8
	Exp       uint8
}

func (*FibMplsLabel) GetTypeName() string { return "fib_mpls_label" }

// FibPath represents VPP binary API type 'fib_path'.
type FibPath struct {
	SwIfIndex  uint32
	TableID    uint32
	RpfID      uint32
	Weight     uint8
	Preference uint8
	Type       FibPathType
	Flags      FibPathFlags
	Proto      FibPathNhProto
	Nh         FibPathNh
	NLabels    uint8
	LabelStack []FibMplsLabel `struc:"[16]FibMplsLabel"`
}

func (*FibPath) GetTypeName() string { return "fib_path" }

// FibPathNh represents VPP binary API type 'fib_path_nh'.
type FibPathNh struct {
	Address            AddressUnion
	ViaLabel           uint32
	ObjID              uint32
	ClassifyTableIndex uint32
}

func (*FibPathNh) GetTypeName() string { return "fib_path_nh" }

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

// MplsRoute represents VPP binary API type 'mpls_route'.
type MplsRoute struct {
	MrTableID     uint32
	MrLabel       uint32
	MrEos         uint8
	MrEosProto    uint8
	MrIsMulticast bool
	MrNPaths      uint8 `struc:"sizeof=MrPaths"`
	MrPaths       []FibPath
}

func (*MplsRoute) GetTypeName() string { return "mpls_route" }

// MplsTable represents VPP binary API type 'mpls_table'.
type MplsTable struct {
	MtTableID uint32
	MtName    string `struc:"[64]byte"`
}

func (*MplsTable) GetTypeName() string { return "mpls_table" }

// MplsTunnel represents VPP binary API type 'mpls_tunnel'.
type MplsTunnel struct {
	MtSwIfIndex   InterfaceIndex
	MtTunnelIndex uint32
	MtL2Only      bool
	MtIsMulticast bool
	MtTag         string `struc:"[64]byte"`
	MtNPaths      uint8  `struc:"sizeof=MtPaths"`
	MtPaths       []FibPath
}

func (*MplsTunnel) GetTypeName() string { return "mpls_tunnel" }

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

// MplsIPBindUnbind represents VPP binary API message 'mpls_ip_bind_unbind'.
type MplsIPBindUnbind struct {
	MbMplsTableID uint32
	MbLabel       uint32
	MbIPTableID   uint32
	MbIsBind      bool
	MbPrefix      Prefix
}

func (m *MplsIPBindUnbind) Reset()                        { *m = MplsIPBindUnbind{} }
func (*MplsIPBindUnbind) GetMessageName() string          { return "mpls_ip_bind_unbind" }
func (*MplsIPBindUnbind) GetCrcString() string            { return "48249a27" }
func (*MplsIPBindUnbind) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsIPBindUnbindReply represents VPP binary API message 'mpls_ip_bind_unbind_reply'.
type MplsIPBindUnbindReply struct {
	Retval int32
}

func (m *MplsIPBindUnbindReply) Reset()                        { *m = MplsIPBindUnbindReply{} }
func (*MplsIPBindUnbindReply) GetMessageName() string          { return "mpls_ip_bind_unbind_reply" }
func (*MplsIPBindUnbindReply) GetCrcString() string            { return "e8d4e804" }
func (*MplsIPBindUnbindReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsRouteAddDel represents VPP binary API message 'mpls_route_add_del'.
type MplsRouteAddDel struct {
	MrIsAdd       bool
	MrIsMultipath bool
	MrRoute       MplsRoute
}

func (m *MplsRouteAddDel) Reset()                        { *m = MplsRouteAddDel{} }
func (*MplsRouteAddDel) GetMessageName() string          { return "mpls_route_add_del" }
func (*MplsRouteAddDel) GetCrcString() string            { return "343cff54" }
func (*MplsRouteAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsRouteAddDelReply represents VPP binary API message 'mpls_route_add_del_reply'.
type MplsRouteAddDelReply struct {
	Retval     int32
	StatsIndex uint32
}

func (m *MplsRouteAddDelReply) Reset()                        { *m = MplsRouteAddDelReply{} }
func (*MplsRouteAddDelReply) GetMessageName() string          { return "mpls_route_add_del_reply" }
func (*MplsRouteAddDelReply) GetCrcString() string            { return "1992deab" }
func (*MplsRouteAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsRouteDetails represents VPP binary API message 'mpls_route_details'.
type MplsRouteDetails struct {
	MrRoute MplsRoute
}

func (m *MplsRouteDetails) Reset()                        { *m = MplsRouteDetails{} }
func (*MplsRouteDetails) GetMessageName() string          { return "mpls_route_details" }
func (*MplsRouteDetails) GetCrcString() string            { return "d0ac384c" }
func (*MplsRouteDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsRouteDump represents VPP binary API message 'mpls_route_dump'.
type MplsRouteDump struct {
	Table MplsTable
}

func (m *MplsRouteDump) Reset()                        { *m = MplsRouteDump{} }
func (*MplsRouteDump) GetMessageName() string          { return "mpls_route_dump" }
func (*MplsRouteDump) GetCrcString() string            { return "935fdefa" }
func (*MplsRouteDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsTableAddDel represents VPP binary API message 'mpls_table_add_del'.
type MplsTableAddDel struct {
	MtIsAdd bool
	MtTable MplsTable
}

func (m *MplsTableAddDel) Reset()                        { *m = MplsTableAddDel{} }
func (*MplsTableAddDel) GetMessageName() string          { return "mpls_table_add_del" }
func (*MplsTableAddDel) GetCrcString() string            { return "57817512" }
func (*MplsTableAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsTableAddDelReply represents VPP binary API message 'mpls_table_add_del_reply'.
type MplsTableAddDelReply struct {
	Retval int32
}

func (m *MplsTableAddDelReply) Reset()                        { *m = MplsTableAddDelReply{} }
func (*MplsTableAddDelReply) GetMessageName() string          { return "mpls_table_add_del_reply" }
func (*MplsTableAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*MplsTableAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTableDetails represents VPP binary API message 'mpls_table_details'.
type MplsTableDetails struct {
	MtTable MplsTable
}

func (m *MplsTableDetails) Reset()                        { *m = MplsTableDetails{} }
func (*MplsTableDetails) GetMessageName() string          { return "mpls_table_details" }
func (*MplsTableDetails) GetCrcString() string            { return "f03ecdc8" }
func (*MplsTableDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTableDump represents VPP binary API message 'mpls_table_dump'.
type MplsTableDump struct{}

func (m *MplsTableDump) Reset()                        { *m = MplsTableDump{} }
func (*MplsTableDump) GetMessageName() string          { return "mpls_table_dump" }
func (*MplsTableDump) GetCrcString() string            { return "51077d14" }
func (*MplsTableDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsTunnelAddDel represents VPP binary API message 'mpls_tunnel_add_del'.
type MplsTunnelAddDel struct {
	MtIsAdd  bool
	MtTunnel MplsTunnel
}

func (m *MplsTunnelAddDel) Reset()                        { *m = MplsTunnelAddDel{} }
func (*MplsTunnelAddDel) GetMessageName() string          { return "mpls_tunnel_add_del" }
func (*MplsTunnelAddDel) GetCrcString() string            { return "e57ce61d" }
func (*MplsTunnelAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsTunnelAddDelReply represents VPP binary API message 'mpls_tunnel_add_del_reply'.
type MplsTunnelAddDelReply struct {
	Retval      int32
	SwIfIndex   InterfaceIndex
	TunnelIndex uint32
}

func (m *MplsTunnelAddDelReply) Reset()                        { *m = MplsTunnelAddDelReply{} }
func (*MplsTunnelAddDelReply) GetMessageName() string          { return "mpls_tunnel_add_del_reply" }
func (*MplsTunnelAddDelReply) GetCrcString() string            { return "afb01472" }
func (*MplsTunnelAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTunnelDetails represents VPP binary API message 'mpls_tunnel_details'.
type MplsTunnelDetails struct {
	MtTunnel MplsTunnel
}

func (m *MplsTunnelDetails) Reset()                        { *m = MplsTunnelDetails{} }
func (*MplsTunnelDetails) GetMessageName() string          { return "mpls_tunnel_details" }
func (*MplsTunnelDetails) GetCrcString() string            { return "f3c0928e" }
func (*MplsTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTunnelDump represents VPP binary API message 'mpls_tunnel_dump'.
type MplsTunnelDump struct {
	SwIfIndex InterfaceIndex
}

func (m *MplsTunnelDump) Reset()                        { *m = MplsTunnelDump{} }
func (*MplsTunnelDump) GetMessageName() string          { return "mpls_tunnel_dump" }
func (*MplsTunnelDump) GetCrcString() string            { return "f9e6675e" }
func (*MplsTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetMplsEnable represents VPP binary API message 'sw_interface_set_mpls_enable'.
type SwInterfaceSetMplsEnable struct {
	SwIfIndex InterfaceIndex
	Enable    bool
}

func (m *SwInterfaceSetMplsEnable) Reset()                        { *m = SwInterfaceSetMplsEnable{} }
func (*SwInterfaceSetMplsEnable) GetMessageName() string          { return "sw_interface_set_mpls_enable" }
func (*SwInterfaceSetMplsEnable) GetCrcString() string            { return "ae6cfcfb" }
func (*SwInterfaceSetMplsEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetMplsEnableReply represents VPP binary API message 'sw_interface_set_mpls_enable_reply'.
type SwInterfaceSetMplsEnableReply struct {
	Retval int32
}

func (m *SwInterfaceSetMplsEnableReply) Reset() { *m = SwInterfaceSetMplsEnableReply{} }
func (*SwInterfaceSetMplsEnableReply) GetMessageName() string {
	return "sw_interface_set_mpls_enable_reply"
}
func (*SwInterfaceSetMplsEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSetMplsEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*MplsIPBindUnbind)(nil), "mpls.MplsIPBindUnbind")
	api.RegisterMessage((*MplsIPBindUnbindReply)(nil), "mpls.MplsIPBindUnbindReply")
	api.RegisterMessage((*MplsRouteAddDel)(nil), "mpls.MplsRouteAddDel")
	api.RegisterMessage((*MplsRouteAddDelReply)(nil), "mpls.MplsRouteAddDelReply")
	api.RegisterMessage((*MplsRouteDetails)(nil), "mpls.MplsRouteDetails")
	api.RegisterMessage((*MplsRouteDump)(nil), "mpls.MplsRouteDump")
	api.RegisterMessage((*MplsTableAddDel)(nil), "mpls.MplsTableAddDel")
	api.RegisterMessage((*MplsTableAddDelReply)(nil), "mpls.MplsTableAddDelReply")
	api.RegisterMessage((*MplsTableDetails)(nil), "mpls.MplsTableDetails")
	api.RegisterMessage((*MplsTableDump)(nil), "mpls.MplsTableDump")
	api.RegisterMessage((*MplsTunnelAddDel)(nil), "mpls.MplsTunnelAddDel")
	api.RegisterMessage((*MplsTunnelAddDelReply)(nil), "mpls.MplsTunnelAddDelReply")
	api.RegisterMessage((*MplsTunnelDetails)(nil), "mpls.MplsTunnelDetails")
	api.RegisterMessage((*MplsTunnelDump)(nil), "mpls.MplsTunnelDump")
	api.RegisterMessage((*SwInterfaceSetMplsEnable)(nil), "mpls.SwInterfaceSetMplsEnable")
	api.RegisterMessage((*SwInterfaceSetMplsEnableReply)(nil), "mpls.SwInterfaceSetMplsEnableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MplsIPBindUnbind)(nil),
		(*MplsIPBindUnbindReply)(nil),
		(*MplsRouteAddDel)(nil),
		(*MplsRouteAddDelReply)(nil),
		(*MplsRouteDetails)(nil),
		(*MplsRouteDump)(nil),
		(*MplsTableAddDel)(nil),
		(*MplsTableAddDelReply)(nil),
		(*MplsTableDetails)(nil),
		(*MplsTableDump)(nil),
		(*MplsTunnelAddDel)(nil),
		(*MplsTunnelAddDelReply)(nil),
		(*MplsTunnelDetails)(nil),
		(*MplsTunnelDump)(nil),
		(*SwInterfaceSetMplsEnable)(nil),
		(*SwInterfaceSetMplsEnableReply)(nil),
	}
}

// RPCService represents RPC service API for mpls module.
type RPCService interface {
	DumpMplsRoute(ctx context.Context, in *MplsRouteDump) (RPCService_DumpMplsRouteClient, error)
	DumpMplsTable(ctx context.Context, in *MplsTableDump) (RPCService_DumpMplsTableClient, error)
	DumpMplsTunnel(ctx context.Context, in *MplsTunnelDump) (RPCService_DumpMplsTunnelClient, error)
	MplsIPBindUnbind(ctx context.Context, in *MplsIPBindUnbind) (*MplsIPBindUnbindReply, error)
	MplsRouteAddDel(ctx context.Context, in *MplsRouteAddDel) (*MplsRouteAddDelReply, error)
	MplsTableAddDel(ctx context.Context, in *MplsTableAddDel) (*MplsTableAddDelReply, error)
	MplsTunnelAddDel(ctx context.Context, in *MplsTunnelAddDel) (*MplsTunnelAddDelReply, error)
	SwInterfaceSetMplsEnable(ctx context.Context, in *SwInterfaceSetMplsEnable) (*SwInterfaceSetMplsEnableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMplsRoute(ctx context.Context, in *MplsRouteDump) (RPCService_DumpMplsRouteClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMplsRouteClient{stream}
	return x, nil
}

type RPCService_DumpMplsRouteClient interface {
	Recv() (*MplsRouteDetails, error)
}

type serviceClient_DumpMplsRouteClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMplsRouteClient) Recv() (*MplsRouteDetails, error) {
	m := new(MplsRouteDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMplsTable(ctx context.Context, in *MplsTableDump) (RPCService_DumpMplsTableClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMplsTableClient{stream}
	return x, nil
}

type RPCService_DumpMplsTableClient interface {
	Recv() (*MplsTableDetails, error)
}

type serviceClient_DumpMplsTableClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMplsTableClient) Recv() (*MplsTableDetails, error) {
	m := new(MplsTableDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMplsTunnel(ctx context.Context, in *MplsTunnelDump) (RPCService_DumpMplsTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMplsTunnelClient{stream}
	return x, nil
}

type RPCService_DumpMplsTunnelClient interface {
	Recv() (*MplsTunnelDetails, error)
}

type serviceClient_DumpMplsTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMplsTunnelClient) Recv() (*MplsTunnelDetails, error) {
	m := new(MplsTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MplsIPBindUnbind(ctx context.Context, in *MplsIPBindUnbind) (*MplsIPBindUnbindReply, error) {
	out := new(MplsIPBindUnbindReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MplsRouteAddDel(ctx context.Context, in *MplsRouteAddDel) (*MplsRouteAddDelReply, error) {
	out := new(MplsRouteAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MplsTableAddDel(ctx context.Context, in *MplsTableAddDel) (*MplsTableAddDelReply, error) {
	out := new(MplsTableAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MplsTunnelAddDel(ctx context.Context, in *MplsTunnelAddDel) (*MplsTunnelAddDelReply, error) {
	out := new(MplsTunnelAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceSetMplsEnable(ctx context.Context, in *SwInterfaceSetMplsEnable) (*SwInterfaceSetMplsEnableReply, error) {
	out := new(SwInterfaceSetMplsEnableReply)
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
