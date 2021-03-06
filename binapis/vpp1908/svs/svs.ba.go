// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/plugins/svs.api.json

/*
Package svs is a generated VPP binary API for 'svs' module.

It consists of:
	  4 enums
	  5 aliases
	  6 types
	  1 union
	 10 messages
	  5 services
*/
package svs

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
	ModuleName = "svs"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x37c625eb
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

// SvsDetails represents VPP binary API message 'svs_details'.
type SvsDetails struct {
	TableID   uint32
	SwIfIndex uint32
	Af        AddressFamily
}

func (m *SvsDetails) Reset()                        { *m = SvsDetails{} }
func (*SvsDetails) GetMessageName() string          { return "svs_details" }
func (*SvsDetails) GetCrcString() string            { return "a4e17ae5" }
func (*SvsDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SvsDump represents VPP binary API message 'svs_dump'.
type SvsDump struct{}

func (m *SvsDump) Reset()                        { *m = SvsDump{} }
func (*SvsDump) GetMessageName() string          { return "svs_dump" }
func (*SvsDump) GetCrcString() string            { return "51077d14" }
func (*SvsDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SvsEnableDisable represents VPP binary API message 'svs_enable_disable'.
type SvsEnableDisable struct {
	IsEnable  uint8
	Af        AddressFamily
	TableID   uint32
	SwIfIndex uint32
}

func (m *SvsEnableDisable) Reset()                        { *m = SvsEnableDisable{} }
func (*SvsEnableDisable) GetMessageName() string          { return "svs_enable_disable" }
func (*SvsEnableDisable) GetCrcString() string            { return "1ad57a4b" }
func (*SvsEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// SvsEnableDisableReply represents VPP binary API message 'svs_enable_disable_reply'.
type SvsEnableDisableReply struct {
	Retval int32
}

func (m *SvsEnableDisableReply) Reset()                        { *m = SvsEnableDisableReply{} }
func (*SvsEnableDisableReply) GetMessageName() string          { return "svs_enable_disable_reply" }
func (*SvsEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*SvsEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SvsPluginGetVersion represents VPP binary API message 'svs_plugin_get_version'.
type SvsPluginGetVersion struct{}

func (m *SvsPluginGetVersion) Reset()                        { *m = SvsPluginGetVersion{} }
func (*SvsPluginGetVersion) GetMessageName() string          { return "svs_plugin_get_version" }
func (*SvsPluginGetVersion) GetCrcString() string            { return "51077d14" }
func (*SvsPluginGetVersion) GetMessageType() api.MessageType { return api.RequestMessage }

// SvsPluginGetVersionReply represents VPP binary API message 'svs_plugin_get_version_reply'.
type SvsPluginGetVersionReply struct {
	Major uint32
	Minor uint32
}

func (m *SvsPluginGetVersionReply) Reset()                        { *m = SvsPluginGetVersionReply{} }
func (*SvsPluginGetVersionReply) GetMessageName() string          { return "svs_plugin_get_version_reply" }
func (*SvsPluginGetVersionReply) GetCrcString() string            { return "9b32cf86" }
func (*SvsPluginGetVersionReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SvsRouteAddDel represents VPP binary API message 'svs_route_add_del'.
type SvsRouteAddDel struct {
	IsAdd         uint8
	Prefix        Prefix
	TableID       uint32
	SourceTableID uint32
}

func (m *SvsRouteAddDel) Reset()                        { *m = SvsRouteAddDel{} }
func (*SvsRouteAddDel) GetMessageName() string          { return "svs_route_add_del" }
func (*SvsRouteAddDel) GetCrcString() string            { return "53cde50b" }
func (*SvsRouteAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// SvsRouteAddDelReply represents VPP binary API message 'svs_route_add_del_reply'.
type SvsRouteAddDelReply struct {
	Retval int32
}

func (m *SvsRouteAddDelReply) Reset()                        { *m = SvsRouteAddDelReply{} }
func (*SvsRouteAddDelReply) GetMessageName() string          { return "svs_route_add_del_reply" }
func (*SvsRouteAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*SvsRouteAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SvsTableAddDel represents VPP binary API message 'svs_table_add_del'.
type SvsTableAddDel struct {
	IsAdd   uint8
	Af      AddressFamily
	TableID uint32
}

func (m *SvsTableAddDel) Reset()                        { *m = SvsTableAddDel{} }
func (*SvsTableAddDel) GetMessageName() string          { return "svs_table_add_del" }
func (*SvsTableAddDel) GetCrcString() string            { return "b74bff95" }
func (*SvsTableAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// SvsTableAddDelReply represents VPP binary API message 'svs_table_add_del_reply'.
type SvsTableAddDelReply struct {
	Retval int32
}

func (m *SvsTableAddDelReply) Reset()                        { *m = SvsTableAddDelReply{} }
func (*SvsTableAddDelReply) GetMessageName() string          { return "svs_table_add_del_reply" }
func (*SvsTableAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*SvsTableAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*SvsDetails)(nil), "svs.SvsDetails")
	api.RegisterMessage((*SvsDump)(nil), "svs.SvsDump")
	api.RegisterMessage((*SvsEnableDisable)(nil), "svs.SvsEnableDisable")
	api.RegisterMessage((*SvsEnableDisableReply)(nil), "svs.SvsEnableDisableReply")
	api.RegisterMessage((*SvsPluginGetVersion)(nil), "svs.SvsPluginGetVersion")
	api.RegisterMessage((*SvsPluginGetVersionReply)(nil), "svs.SvsPluginGetVersionReply")
	api.RegisterMessage((*SvsRouteAddDel)(nil), "svs.SvsRouteAddDel")
	api.RegisterMessage((*SvsRouteAddDelReply)(nil), "svs.SvsRouteAddDelReply")
	api.RegisterMessage((*SvsTableAddDel)(nil), "svs.SvsTableAddDel")
	api.RegisterMessage((*SvsTableAddDelReply)(nil), "svs.SvsTableAddDelReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SvsDetails)(nil),
		(*SvsDump)(nil),
		(*SvsEnableDisable)(nil),
		(*SvsEnableDisableReply)(nil),
		(*SvsPluginGetVersion)(nil),
		(*SvsPluginGetVersionReply)(nil),
		(*SvsRouteAddDel)(nil),
		(*SvsRouteAddDelReply)(nil),
		(*SvsTableAddDel)(nil),
		(*SvsTableAddDelReply)(nil),
	}
}

// RPCService represents RPC service API for svs module.
type RPCService interface {
	DumpSvs(ctx context.Context, in *SvsDump) (RPCService_DumpSvsClient, error)
	SvsEnableDisable(ctx context.Context, in *SvsEnableDisable) (*SvsEnableDisableReply, error)
	SvsPluginGetVersion(ctx context.Context, in *SvsPluginGetVersion) (*SvsPluginGetVersionReply, error)
	SvsRouteAddDel(ctx context.Context, in *SvsRouteAddDel) (*SvsRouteAddDelReply, error)
	SvsTableAddDel(ctx context.Context, in *SvsTableAddDel) (*SvsTableAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSvs(ctx context.Context, in *SvsDump) (RPCService_DumpSvsClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSvsClient{stream}
	return x, nil
}

type RPCService_DumpSvsClient interface {
	Recv() (*SvsDetails, error)
}

type serviceClient_DumpSvsClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSvsClient) Recv() (*SvsDetails, error) {
	m := new(SvsDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) SvsEnableDisable(ctx context.Context, in *SvsEnableDisable) (*SvsEnableDisableReply, error) {
	out := new(SvsEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SvsPluginGetVersion(ctx context.Context, in *SvsPluginGetVersion) (*SvsPluginGetVersionReply, error) {
	out := new(SvsPluginGetVersionReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SvsRouteAddDel(ctx context.Context, in *SvsRouteAddDel) (*SvsRouteAddDelReply, error) {
	out := new(SvsRouteAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SvsTableAddDel(ctx context.Context, in *SvsTableAddDel) (*SvsTableAddDelReply, error) {
	out := new(SvsTableAddDelReply)
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
