// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/mpls.api.json

/*
Package mpls is a generated VPP binary API for 'mpls' module.

It consists of:
	  5 enums
	  2 aliases
	 11 types
	  1 union
	 16 messages
	  8 services
*/
package mpls

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
	ModuleName = "mpls"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x74ef3e0f
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

// IPProto represents VPP binary API enum 'ip_proto'.
type IPProto uint32

const (
	IP_API_PROTO_TCP   IPProto = 6
	IP_API_PROTO_UDP   IPProto = 17
	IP_API_PROTO_EIGRP IPProto = 88
	IP_API_PROTO_OSPF  IPProto = 89
)

var IPProto_name = map[uint32]string{
	6:  "IP_API_PROTO_TCP",
	17: "IP_API_PROTO_UDP",
	88: "IP_API_PROTO_EIGRP",
	89: "IP_API_PROTO_OSPF",
}

var IPProto_value = map[string]uint32{
	"IP_API_PROTO_TCP":   6,
	"IP_API_PROTO_UDP":   17,
	"IP_API_PROTO_EIGRP": 88,
	"IP_API_PROTO_OSPF":  89,
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

// MplsRoute represents VPP binary API type 'mpls_route'.
type MplsRoute struct {
	MrTableID     uint32
	MrLabel       uint32
	MrEos         uint8
	MrEosProto    uint8
	MrIsMulticast uint8
	MrNPaths      uint8 `struc:"sizeof=MrPaths"`
	MrPaths       []FibPath
}

func (*MplsRoute) GetTypeName() string {
	return "mpls_route"
}

// MplsTable represents VPP binary API type 'mpls_table'.
type MplsTable struct {
	MtTableID uint32
	MtName    []byte `struc:"[64]byte"`
}

func (*MplsTable) GetTypeName() string {
	return "mpls_table"
}

// MplsTunnel represents VPP binary API type 'mpls_tunnel'.
type MplsTunnel struct {
	MtSwIfIndex   uint32
	MtTunnelIndex uint32
	MtL2Only      uint8
	MtIsMulticast uint8
	MtNPaths      uint8 `struc:"sizeof=MtPaths"`
	MtPaths       []FibPath
}

func (*MplsTunnel) GetTypeName() string {
	return "mpls_tunnel"
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

// MplsIPBindUnbind represents VPP binary API message 'mpls_ip_bind_unbind'.
type MplsIPBindUnbind struct {
	MbMplsTableID uint32
	MbLabel       uint32
	MbIPTableID   uint32
	MbIsBind      uint8
	MbPrefix      Prefix
}

func (*MplsIPBindUnbind) GetMessageName() string {
	return "mpls_ip_bind_unbind"
}
func (*MplsIPBindUnbind) GetCrcString() string {
	return "76e7ae51"
}
func (*MplsIPBindUnbind) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MplsIPBindUnbindReply represents VPP binary API message 'mpls_ip_bind_unbind_reply'.
type MplsIPBindUnbindReply struct {
	Retval int32
}

func (*MplsIPBindUnbindReply) GetMessageName() string {
	return "mpls_ip_bind_unbind_reply"
}
func (*MplsIPBindUnbindReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MplsIPBindUnbindReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsRouteAddDel represents VPP binary API message 'mpls_route_add_del'.
type MplsRouteAddDel struct {
	MrIsAdd       uint8
	MrIsMultipath uint8
	MrRoute       MplsRoute
}

func (*MplsRouteAddDel) GetMessageName() string {
	return "mpls_route_add_del"
}
func (*MplsRouteAddDel) GetCrcString() string {
	return "3d7ca673"
}
func (*MplsRouteAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MplsRouteAddDelReply represents VPP binary API message 'mpls_route_add_del_reply'.
type MplsRouteAddDelReply struct {
	Retval     int32
	StatsIndex uint32
}

func (*MplsRouteAddDelReply) GetMessageName() string {
	return "mpls_route_add_del_reply"
}
func (*MplsRouteAddDelReply) GetCrcString() string {
	return "1992deab"
}
func (*MplsRouteAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsRouteDetails represents VPP binary API message 'mpls_route_details'.
type MplsRouteDetails struct {
	MrRoute MplsRoute
}

func (*MplsRouteDetails) GetMessageName() string {
	return "mpls_route_details"
}
func (*MplsRouteDetails) GetCrcString() string {
	return "463be12f"
}
func (*MplsRouteDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsRouteDump represents VPP binary API message 'mpls_route_dump'.
type MplsRouteDump struct {
	Table MplsTable
}

func (*MplsRouteDump) GetMessageName() string {
	return "mpls_route_dump"
}
func (*MplsRouteDump) GetCrcString() string {
	return "fa14d170"
}
func (*MplsRouteDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MplsTableAddDel represents VPP binary API message 'mpls_table_add_del'.
type MplsTableAddDel struct {
	MtIsAdd uint8
	MtTable MplsTable
}

func (*MplsTableAddDel) GetMessageName() string {
	return "mpls_table_add_del"
}
func (*MplsTableAddDel) GetCrcString() string {
	return "b4ee1c03"
}
func (*MplsTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MplsTableAddDelReply represents VPP binary API message 'mpls_table_add_del_reply'.
type MplsTableAddDelReply struct {
	Retval int32
}

func (*MplsTableAddDelReply) GetMessageName() string {
	return "mpls_table_add_del_reply"
}
func (*MplsTableAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MplsTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsTableDetails represents VPP binary API message 'mpls_table_details'.
type MplsTableDetails struct {
	MtTable MplsTable
}

func (*MplsTableDetails) GetMessageName() string {
	return "mpls_table_details"
}
func (*MplsTableDetails) GetCrcString() string {
	return "5624404a"
}
func (*MplsTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsTableDump represents VPP binary API message 'mpls_table_dump'.
type MplsTableDump struct{}

func (*MplsTableDump) GetMessageName() string {
	return "mpls_table_dump"
}
func (*MplsTableDump) GetCrcString() string {
	return "51077d14"
}
func (*MplsTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MplsTunnelAddDel represents VPP binary API message 'mpls_tunnel_add_del'.
type MplsTunnelAddDel struct {
	MtIsAdd  uint8
	MtTunnel MplsTunnel
}

func (*MplsTunnelAddDel) GetMessageName() string {
	return "mpls_tunnel_add_del"
}
func (*MplsTunnelAddDel) GetCrcString() string {
	return "81025192"
}
func (*MplsTunnelAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MplsTunnelAddDelReply represents VPP binary API message 'mpls_tunnel_add_del_reply'.
type MplsTunnelAddDelReply struct {
	Retval      int32
	SwIfIndex   uint32
	TunnelIndex uint32
}

func (*MplsTunnelAddDelReply) GetMessageName() string {
	return "mpls_tunnel_add_del_reply"
}
func (*MplsTunnelAddDelReply) GetCrcString() string {
	return "cc62a1ce"
}
func (*MplsTunnelAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsTunnelDetails represents VPP binary API message 'mpls_tunnel_details'.
type MplsTunnelDetails struct {
	MtTunnel MplsTunnel
}

func (*MplsTunnelDetails) GetMessageName() string {
	return "mpls_tunnel_details"
}
func (*MplsTunnelDetails) GetCrcString() string {
	return "8785712e"
}
func (*MplsTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MplsTunnelDump represents VPP binary API message 'mpls_tunnel_dump'.
type MplsTunnelDump struct {
	SwIfIndex uint32
}

func (*MplsTunnelDump) GetMessageName() string {
	return "mpls_tunnel_dump"
}
func (*MplsTunnelDump) GetCrcString() string {
	return "529cb13f"
}
func (*MplsTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetMplsEnable represents VPP binary API message 'sw_interface_set_mpls_enable'.
type SwInterfaceSetMplsEnable struct {
	SwIfIndex uint32
	Enable    uint8
}

func (*SwInterfaceSetMplsEnable) GetMessageName() string {
	return "sw_interface_set_mpls_enable"
}
func (*SwInterfaceSetMplsEnable) GetCrcString() string {
	return "a36fadc0"
}
func (*SwInterfaceSetMplsEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetMplsEnableReply represents VPP binary API message 'sw_interface_set_mpls_enable_reply'.
type SwInterfaceSetMplsEnableReply struct {
	Retval int32
}

func (*SwInterfaceSetMplsEnableReply) GetMessageName() string {
	return "sw_interface_set_mpls_enable_reply"
}
func (*SwInterfaceSetMplsEnableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetMplsEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

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
