// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/bier.api.json

/*
Package bier is a generated VPP binary API for 'bier' module.

It consists of:
	  7 enums
	  2 aliases
	 10 types
	  1 union
	 22 messages
	 11 services
*/
package bier

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
	ModuleName = "bier"
	// APIVersion is the API version of this module.
	APIVersion = "1.2.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x13f02645
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

// FibPathFlags represents VPP binary API enum 'fib_path_flags'.
type FibPathFlags uint32

const (
	FIB_API_PATH_FLAG_NONE                 FibPathFlags = 0
	FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED FibPathFlags = 1
	FIB_API_PATH_FLAG_RESOLVE_VIA_HOST     FibPathFlags = 2
)

var FibPathFlags_name = map[uint32]string{
	0: "FIB_API_PATH_FLAG_NONE",
	1: "FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED",
	2: "FIB_API_PATH_FLAG_RESOLVE_VIA_HOST",
}

var FibPathFlags_value = map[string]uint32{
	"FIB_API_PATH_FLAG_NONE":                 0,
	"FIB_API_PATH_FLAG_RESOLVE_VIA_ATTACHED": 1,
	"FIB_API_PATH_FLAG_RESOLVE_VIA_HOST":     2,
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

// BierRoute represents VPP binary API type 'bier_route'.
type BierRoute struct {
	BrBp     uint32
	BrTblID  BierTableID
	BrNPaths uint8 `struc:"sizeof=BrPaths"`
	BrPaths  []FibPath
}

func (*BierRoute) GetTypeName() string {
	return "bier_route"
}

// BierTableID represents VPP binary API type 'bier_table_id'.
type BierTableID struct {
	BtSet       uint8
	BtSubDomain uint8
	BtHdrLenID  uint8
}

func (*BierTableID) GetTypeName() string {
	return "bier_table_id"
}

// FibMplsLabel represents VPP binary API type 'fib_mpls_label'.
type FibMplsLabel struct {
	IsUniform uint8
	Label     uint32
	TTL       uint8
	Exp       uint8
}

func (*FibMplsLabel) GetTypeName() string {
	return "fib_mpls_label"
}

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

func (*FibPath) GetTypeName() string {
	return "fib_path"
}

// FibPathNh represents VPP binary API type 'fib_path_nh'.
type FibPathNh struct {
	Address            AddressUnion
	ViaLabel           uint32
	ObjID              uint32
	ClassifyTableIndex uint32
}

func (*FibPathNh) GetTypeName() string {
	return "fib_path_nh"
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

// BierDispEntryAddDel represents VPP binary API message 'bier_disp_entry_add_del'.
type BierDispEntryAddDel struct {
	BdeBp           uint16
	BdeTblID        uint32
	BdeIsAdd        uint8
	BdePayloadProto uint8
	BdeNPaths       uint8 `struc:"sizeof=BdePaths"`
	BdePaths        []FibPath
}

func (*BierDispEntryAddDel) GetMessageName() string {
	return "bier_disp_entry_add_del"
}
func (*BierDispEntryAddDel) GetCrcString() string {
	return "e10bd9c4"
}
func (*BierDispEntryAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispEntryAddDelReply represents VPP binary API message 'bier_disp_entry_add_del_reply'.
type BierDispEntryAddDelReply struct {
	Retval int32
}

func (*BierDispEntryAddDelReply) GetMessageName() string {
	return "bier_disp_entry_add_del_reply"
}
func (*BierDispEntryAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierDispEntryAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispEntryDetails represents VPP binary API message 'bier_disp_entry_details'.
type BierDispEntryDetails struct {
	BdeBp           uint16
	BdeTblID        uint32
	BdeIsAdd        uint8
	BdePayloadProto uint8
	BdeNPaths       uint8 `struc:"sizeof=BdePaths"`
	BdePaths        []FibPath
}

func (*BierDispEntryDetails) GetMessageName() string {
	return "bier_disp_entry_details"
}
func (*BierDispEntryDetails) GetCrcString() string {
	return "c942baef"
}
func (*BierDispEntryDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispEntryDump represents VPP binary API message 'bier_disp_entry_dump'.
type BierDispEntryDump struct {
	BdeTblID uint32
}

func (*BierDispEntryDump) GetMessageName() string {
	return "bier_disp_entry_dump"
}
func (*BierDispEntryDump) GetCrcString() string {
	return "b5fa54ad"
}
func (*BierDispEntryDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispTableAddDel represents VPP binary API message 'bier_disp_table_add_del'.
type BierDispTableAddDel struct {
	BdtTblID uint32
	BdtIsAdd uint8
}

func (*BierDispTableAddDel) GetMessageName() string {
	return "bier_disp_table_add_del"
}
func (*BierDispTableAddDel) GetCrcString() string {
	return "7671b2cb"
}
func (*BierDispTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispTableAddDelReply represents VPP binary API message 'bier_disp_table_add_del_reply'.
type BierDispTableAddDelReply struct {
	Retval int32
}

func (*BierDispTableAddDelReply) GetMessageName() string {
	return "bier_disp_table_add_del_reply"
}
func (*BierDispTableAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierDispTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispTableDetails represents VPP binary API message 'bier_disp_table_details'.
type BierDispTableDetails struct {
	BdtTblID uint32
}

func (*BierDispTableDetails) GetMessageName() string {
	return "bier_disp_table_details"
}
func (*BierDispTableDetails) GetCrcString() string {
	return "d27942c0"
}
func (*BierDispTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispTableDump represents VPP binary API message 'bier_disp_table_dump'.
type BierDispTableDump struct{}

func (*BierDispTableDump) GetMessageName() string {
	return "bier_disp_table_dump"
}
func (*BierDispTableDump) GetCrcString() string {
	return "51077d14"
}
func (*BierDispTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierImpAdd represents VPP binary API message 'bier_imp_add'.
type BierImpAdd struct {
	BiTblID  BierTableID
	BiSrc    uint16
	BiNBytes uint8 `struc:"sizeof=BiBytes"`
	BiBytes  []byte
}

func (*BierImpAdd) GetMessageName() string {
	return "bier_imp_add"
}
func (*BierImpAdd) GetCrcString() string {
	return "3856dc3d"
}
func (*BierImpAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierImpAddReply represents VPP binary API message 'bier_imp_add_reply'.
type BierImpAddReply struct {
	Retval  int32
	BiIndex uint32
}

func (*BierImpAddReply) GetMessageName() string {
	return "bier_imp_add_reply"
}
func (*BierImpAddReply) GetCrcString() string {
	return "d49c5793"
}
func (*BierImpAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierImpDel represents VPP binary API message 'bier_imp_del'.
type BierImpDel struct {
	BiIndex uint32
}

func (*BierImpDel) GetMessageName() string {
	return "bier_imp_del"
}
func (*BierImpDel) GetCrcString() string {
	return "7d45edf6"
}
func (*BierImpDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierImpDelReply represents VPP binary API message 'bier_imp_del_reply'.
type BierImpDelReply struct {
	Retval int32
}

func (*BierImpDelReply) GetMessageName() string {
	return "bier_imp_del_reply"
}
func (*BierImpDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierImpDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierImpDetails represents VPP binary API message 'bier_imp_details'.
type BierImpDetails struct {
	BiTblID  BierTableID
	BiSrc    uint16
	BiNBytes uint8 `struc:"sizeof=BiBytes"`
	BiBytes  []byte
}

func (*BierImpDetails) GetMessageName() string {
	return "bier_imp_details"
}
func (*BierImpDetails) GetCrcString() string {
	return "b76192df"
}
func (*BierImpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierImpDump represents VPP binary API message 'bier_imp_dump'.
type BierImpDump struct{}

func (*BierImpDump) GetMessageName() string {
	return "bier_imp_dump"
}
func (*BierImpDump) GetCrcString() string {
	return "51077d14"
}
func (*BierImpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierRouteAddDel represents VPP binary API message 'bier_route_add_del'.
type BierRouteAddDel struct {
	BrIsAdd     uint8
	BrIsReplace uint8
	BrRoute     BierRoute
}

func (*BierRouteAddDel) GetMessageName() string {
	return "bier_route_add_del"
}
func (*BierRouteAddDel) GetCrcString() string {
	return "90121982"
}
func (*BierRouteAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierRouteAddDelReply represents VPP binary API message 'bier_route_add_del_reply'.
type BierRouteAddDelReply struct {
	Retval int32
}

func (*BierRouteAddDelReply) GetMessageName() string {
	return "bier_route_add_del_reply"
}
func (*BierRouteAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierRouteAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierRouteDetails represents VPP binary API message 'bier_route_details'.
type BierRouteDetails struct {
	BrRoute BierRoute
}

func (*BierRouteDetails) GetMessageName() string {
	return "bier_route_details"
}
func (*BierRouteDetails) GetCrcString() string {
	return "39ee6a56"
}
func (*BierRouteDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierRouteDump represents VPP binary API message 'bier_route_dump'.
type BierRouteDump struct {
	BrTblID BierTableID
}

func (*BierRouteDump) GetMessageName() string {
	return "bier_route_dump"
}
func (*BierRouteDump) GetCrcString() string {
	return "38339846"
}
func (*BierRouteDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierTableAddDel represents VPP binary API message 'bier_table_add_del'.
type BierTableAddDel struct {
	BtTblID BierTableID
	BtLabel uint32
	BtIsAdd uint8
}

func (*BierTableAddDel) GetMessageName() string {
	return "bier_table_add_del"
}
func (*BierTableAddDel) GetCrcString() string {
	return "cd3e9483"
}
func (*BierTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierTableAddDelReply represents VPP binary API message 'bier_table_add_del_reply'.
type BierTableAddDelReply struct {
	Retval int32
}

func (*BierTableAddDelReply) GetMessageName() string {
	return "bier_table_add_del_reply"
}
func (*BierTableAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierTableDetails represents VPP binary API message 'bier_table_details'.
type BierTableDetails struct {
	BtLabel uint32
	BtTblID BierTableID
}

func (*BierTableDetails) GetMessageName() string {
	return "bier_table_details"
}
func (*BierTableDetails) GetCrcString() string {
	return "fc44a9dd"
}
func (*BierTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierTableDump represents VPP binary API message 'bier_table_dump'.
type BierTableDump struct{}

func (*BierTableDump) GetMessageName() string {
	return "bier_table_dump"
}
func (*BierTableDump) GetCrcString() string {
	return "51077d14"
}
func (*BierTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*BierDispEntryAddDel)(nil), "bier.BierDispEntryAddDel")
	api.RegisterMessage((*BierDispEntryAddDelReply)(nil), "bier.BierDispEntryAddDelReply")
	api.RegisterMessage((*BierDispEntryDetails)(nil), "bier.BierDispEntryDetails")
	api.RegisterMessage((*BierDispEntryDump)(nil), "bier.BierDispEntryDump")
	api.RegisterMessage((*BierDispTableAddDel)(nil), "bier.BierDispTableAddDel")
	api.RegisterMessage((*BierDispTableAddDelReply)(nil), "bier.BierDispTableAddDelReply")
	api.RegisterMessage((*BierDispTableDetails)(nil), "bier.BierDispTableDetails")
	api.RegisterMessage((*BierDispTableDump)(nil), "bier.BierDispTableDump")
	api.RegisterMessage((*BierImpAdd)(nil), "bier.BierImpAdd")
	api.RegisterMessage((*BierImpAddReply)(nil), "bier.BierImpAddReply")
	api.RegisterMessage((*BierImpDel)(nil), "bier.BierImpDel")
	api.RegisterMessage((*BierImpDelReply)(nil), "bier.BierImpDelReply")
	api.RegisterMessage((*BierImpDetails)(nil), "bier.BierImpDetails")
	api.RegisterMessage((*BierImpDump)(nil), "bier.BierImpDump")
	api.RegisterMessage((*BierRouteAddDel)(nil), "bier.BierRouteAddDel")
	api.RegisterMessage((*BierRouteAddDelReply)(nil), "bier.BierRouteAddDelReply")
	api.RegisterMessage((*BierRouteDetails)(nil), "bier.BierRouteDetails")
	api.RegisterMessage((*BierRouteDump)(nil), "bier.BierRouteDump")
	api.RegisterMessage((*BierTableAddDel)(nil), "bier.BierTableAddDel")
	api.RegisterMessage((*BierTableAddDelReply)(nil), "bier.BierTableAddDelReply")
	api.RegisterMessage((*BierTableDetails)(nil), "bier.BierTableDetails")
	api.RegisterMessage((*BierTableDump)(nil), "bier.BierTableDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BierDispEntryAddDel)(nil),
		(*BierDispEntryAddDelReply)(nil),
		(*BierDispEntryDetails)(nil),
		(*BierDispEntryDump)(nil),
		(*BierDispTableAddDel)(nil),
		(*BierDispTableAddDelReply)(nil),
		(*BierDispTableDetails)(nil),
		(*BierDispTableDump)(nil),
		(*BierImpAdd)(nil),
		(*BierImpAddReply)(nil),
		(*BierImpDel)(nil),
		(*BierImpDelReply)(nil),
		(*BierImpDetails)(nil),
		(*BierImpDump)(nil),
		(*BierRouteAddDel)(nil),
		(*BierRouteAddDelReply)(nil),
		(*BierRouteDetails)(nil),
		(*BierRouteDump)(nil),
		(*BierTableAddDel)(nil),
		(*BierTableAddDelReply)(nil),
		(*BierTableDetails)(nil),
		(*BierTableDump)(nil),
	}
}

// RPCService represents RPC service API for bier module.
type RPCService interface {
	DumpBierDispEntry(ctx context.Context, in *BierDispEntryDump) (RPCService_DumpBierDispEntryClient, error)
	DumpBierDispTable(ctx context.Context, in *BierDispTableDump) (RPCService_DumpBierDispTableClient, error)
	DumpBierImp(ctx context.Context, in *BierImpDump) (RPCService_DumpBierImpClient, error)
	DumpBierRoute(ctx context.Context, in *BierRouteDump) (RPCService_DumpBierRouteClient, error)
	DumpBierTable(ctx context.Context, in *BierTableDump) (RPCService_DumpBierTableClient, error)
	BierDispEntryAddDel(ctx context.Context, in *BierDispEntryAddDel) (*BierDispEntryAddDelReply, error)
	BierDispTableAddDel(ctx context.Context, in *BierDispTableAddDel) (*BierDispTableAddDelReply, error)
	BierImpAdd(ctx context.Context, in *BierImpAdd) (*BierImpAddReply, error)
	BierImpDel(ctx context.Context, in *BierImpDel) (*BierImpDelReply, error)
	BierRouteAddDel(ctx context.Context, in *BierRouteAddDel) (*BierRouteAddDelReply, error)
	BierTableAddDel(ctx context.Context, in *BierTableAddDel) (*BierTableAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpBierDispEntry(ctx context.Context, in *BierDispEntryDump) (RPCService_DumpBierDispEntryClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBierDispEntryClient{stream}
	return x, nil
}

type RPCService_DumpBierDispEntryClient interface {
	Recv() (*BierDispEntryDetails, error)
}

type serviceClient_DumpBierDispEntryClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBierDispEntryClient) Recv() (*BierDispEntryDetails, error) {
	m := new(BierDispEntryDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpBierDispTable(ctx context.Context, in *BierDispTableDump) (RPCService_DumpBierDispTableClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBierDispTableClient{stream}
	return x, nil
}

type RPCService_DumpBierDispTableClient interface {
	Recv() (*BierDispTableDetails, error)
}

type serviceClient_DumpBierDispTableClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBierDispTableClient) Recv() (*BierDispTableDetails, error) {
	m := new(BierDispTableDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpBierImp(ctx context.Context, in *BierImpDump) (RPCService_DumpBierImpClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBierImpClient{stream}
	return x, nil
}

type RPCService_DumpBierImpClient interface {
	Recv() (*BierImpDetails, error)
}

type serviceClient_DumpBierImpClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBierImpClient) Recv() (*BierImpDetails, error) {
	m := new(BierImpDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpBierRoute(ctx context.Context, in *BierRouteDump) (RPCService_DumpBierRouteClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBierRouteClient{stream}
	return x, nil
}

type RPCService_DumpBierRouteClient interface {
	Recv() (*BierRouteDetails, error)
}

type serviceClient_DumpBierRouteClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBierRouteClient) Recv() (*BierRouteDetails, error) {
	m := new(BierRouteDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpBierTable(ctx context.Context, in *BierTableDump) (RPCService_DumpBierTableClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBierTableClient{stream}
	return x, nil
}

type RPCService_DumpBierTableClient interface {
	Recv() (*BierTableDetails, error)
}

type serviceClient_DumpBierTableClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBierTableClient) Recv() (*BierTableDetails, error) {
	m := new(BierTableDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) BierDispEntryAddDel(ctx context.Context, in *BierDispEntryAddDel) (*BierDispEntryAddDelReply, error) {
	out := new(BierDispEntryAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BierDispTableAddDel(ctx context.Context, in *BierDispTableAddDel) (*BierDispTableAddDelReply, error) {
	out := new(BierDispTableAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BierImpAdd(ctx context.Context, in *BierImpAdd) (*BierImpAddReply, error) {
	out := new(BierImpAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BierImpDel(ctx context.Context, in *BierImpDel) (*BierImpDelReply, error) {
	out := new(BierImpDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BierRouteAddDel(ctx context.Context, in *BierRouteAddDel) (*BierRouteAddDelReply, error) {
	out := new(BierRouteAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BierTableAddDel(ctx context.Context, in *BierTableAddDel) (*BierTableAddDelReply, error) {
	out := new(BierTableAddDelReply)
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
