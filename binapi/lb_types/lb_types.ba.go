// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/plugins/lb_types.api.json

/*
Package lb_types is a generated VPP binary API for 'lb_types' module.

It consists of:
	  9 enums
	  2 aliases
	  7 types
	  1 union
*/
package lb_types

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
	ModuleName = "lb_types"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x34b91915
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

// LbEncapType represents VPP binary API enum 'lb_encap_type'.
type LbEncapType uint32

const (
	LB_API_ENCAP_TYPE_GRE4  LbEncapType = 0
	LB_API_ENCAP_TYPE_GRE6  LbEncapType = 1
	LB_API_ENCAP_TYPE_L3DSR LbEncapType = 2
	LB_API_ENCAP_TYPE_NAT4  LbEncapType = 3
	LB_API_ENCAP_TYPE_NAT6  LbEncapType = 4
	LB_API_ENCAP_N_TYPES    LbEncapType = 5
)

var LbEncapType_name = map[uint32]string{
	0: "LB_API_ENCAP_TYPE_GRE4",
	1: "LB_API_ENCAP_TYPE_GRE6",
	2: "LB_API_ENCAP_TYPE_L3DSR",
	3: "LB_API_ENCAP_TYPE_NAT4",
	4: "LB_API_ENCAP_TYPE_NAT6",
	5: "LB_API_ENCAP_N_TYPES",
}

var LbEncapType_value = map[string]uint32{
	"LB_API_ENCAP_TYPE_GRE4":  0,
	"LB_API_ENCAP_TYPE_GRE6":  1,
	"LB_API_ENCAP_TYPE_L3DSR": 2,
	"LB_API_ENCAP_TYPE_NAT4":  3,
	"LB_API_ENCAP_TYPE_NAT6":  4,
	"LB_API_ENCAP_N_TYPES":    5,
}

func (x LbEncapType) String() string {
	s, ok := LbEncapType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbLkpTypeT represents VPP binary API enum 'lb_lkp_type_t'.
type LbLkpTypeT uint32

const (
	LB_API_LKP_SAME_IP_PORT LbLkpTypeT = 0
	LB_API_LKP_DIFF_IP_PORT LbLkpTypeT = 1
	LB_API_LKP_ALL_PORT_IP  LbLkpTypeT = 2
	LB_API_LKP_N_TYPES      LbLkpTypeT = 3
)

var LbLkpTypeT_name = map[uint32]string{
	0: "LB_API_LKP_SAME_IP_PORT",
	1: "LB_API_LKP_DIFF_IP_PORT",
	2: "LB_API_LKP_ALL_PORT_IP",
	3: "LB_API_LKP_N_TYPES",
}

var LbLkpTypeT_value = map[string]uint32{
	"LB_API_LKP_SAME_IP_PORT": 0,
	"LB_API_LKP_DIFF_IP_PORT": 1,
	"LB_API_LKP_ALL_PORT_IP":  2,
	"LB_API_LKP_N_TYPES":      3,
}

func (x LbLkpTypeT) String() string {
	s, ok := LbLkpTypeT_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbNatProtocol represents VPP binary API enum 'lb_nat_protocol'.
type LbNatProtocol uint32

const (
	LB_API_NAT_PROTOCOL_UDP LbNatProtocol = 6
	LB_API_NAT_PROTOCOL_TCP LbNatProtocol = 23
	LB_API_NAT_PROTOCOL_ANY LbNatProtocol = 4.294967295e+09
)

var LbNatProtocol_name = map[uint32]string{
	6:               "LB_API_NAT_PROTOCOL_UDP",
	23:              "LB_API_NAT_PROTOCOL_TCP",
	4.294967295e+09: "LB_API_NAT_PROTOCOL_ANY",
}

var LbNatProtocol_value = map[string]uint32{
	"LB_API_NAT_PROTOCOL_UDP": 6,
	"LB_API_NAT_PROTOCOL_TCP": 23,
	"LB_API_NAT_PROTOCOL_ANY": 4.294967295e+09,
}

func (x LbNatProtocol) String() string {
	s, ok := LbNatProtocol_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbSrvType represents VPP binary API enum 'lb_srv_type'.
type LbSrvType uint32

const (
	LB_API_SRV_TYPE_CLUSTERIP LbSrvType = 0
	LB_API_SRV_TYPE_NODEPORT  LbSrvType = 1
	LB_API_SRV_N_TYPES        LbSrvType = 2
)

var LbSrvType_name = map[uint32]string{
	0: "LB_API_SRV_TYPE_CLUSTERIP",
	1: "LB_API_SRV_TYPE_NODEPORT",
	2: "LB_API_SRV_N_TYPES",
}

var LbSrvType_value = map[string]uint32{
	"LB_API_SRV_TYPE_CLUSTERIP": 0,
	"LB_API_SRV_TYPE_NODEPORT":  1,
	"LB_API_SRV_N_TYPES":        2,
}

func (x LbSrvType) String() string {
	s, ok := LbSrvType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbVipType represents VPP binary API enum 'lb_vip_type'.
type LbVipType uint32

const (
	LB_API_VIP_TYPE_IP6_GRE6  LbVipType = 0
	LB_API_VIP_TYPE_IP6_GRE4  LbVipType = 1
	LB_API_VIP_TYPE_IP4_GRE6  LbVipType = 2
	LB_API_VIP_TYPE_IP4_GRE4  LbVipType = 3
	LB_API_VIP_TYPE_IP4_L3DSR LbVipType = 4
	LB_API_VIP_TYPE_IP4_NAT4  LbVipType = 5
	LB_API_VIP_TYPE_IP6_NAT6  LbVipType = 6
	LB_API_VIP_N_TYPES        LbVipType = 7
)

var LbVipType_name = map[uint32]string{
	0: "LB_API_VIP_TYPE_IP6_GRE6",
	1: "LB_API_VIP_TYPE_IP6_GRE4",
	2: "LB_API_VIP_TYPE_IP4_GRE6",
	3: "LB_API_VIP_TYPE_IP4_GRE4",
	4: "LB_API_VIP_TYPE_IP4_L3DSR",
	5: "LB_API_VIP_TYPE_IP4_NAT4",
	6: "LB_API_VIP_TYPE_IP6_NAT6",
	7: "LB_API_VIP_N_TYPES",
}

var LbVipType_value = map[string]uint32{
	"LB_API_VIP_TYPE_IP6_GRE6":  0,
	"LB_API_VIP_TYPE_IP6_GRE4":  1,
	"LB_API_VIP_TYPE_IP4_GRE6":  2,
	"LB_API_VIP_TYPE_IP4_GRE4":  3,
	"LB_API_VIP_TYPE_IP4_L3DSR": 4,
	"LB_API_VIP_TYPE_IP4_NAT4":  5,
	"LB_API_VIP_TYPE_IP6_NAT6":  6,
	"LB_API_VIP_N_TYPES":        7,
}

func (x LbVipType) String() string {
	s, ok := LbVipType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IP4Address represents VPP binary API alias 'ip4_address'.
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address'.
type IP6Address [16]uint8

// Address represents VPP binary API type 'address'.
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address
	Len     uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address
	Len     uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}

// LbVip represents VPP binary API type 'lb_vip'.
type LbVip struct {
	Pfx      Prefix
	Protocol IPProto
	Port     uint16
}

func (*LbVip) GetTypeName() string {
	return "lb_vip"
}

// Mprefix represents VPP binary API type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}

// Prefix represents VPP binary API type 'prefix'.
type Prefix struct {
	Address Address
	Len     uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}

// PrefixMatcher represents VPP binary API type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8
	Ge uint8
}

func (*PrefixMatcher) GetTypeName() string {
	return "prefix_matcher"
}

// AddressUnion represents VPP binary API union 'address_union'.
type AddressUnion struct {
	XXX_UnionData [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}

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
