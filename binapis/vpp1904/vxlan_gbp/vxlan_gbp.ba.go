// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/vxlan_gbp.api.json

/*
Package vxlan_gbp is a generated VPP binary API for 'vxlan_gbp' module.

It consists of:
	  2 enums
	  2 aliases
	  6 types
	  1 union
	  6 messages
	  3 services
*/
package vxlan_gbp

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
	ModuleName = "vxlan_gbp"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xa26e715a
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

// VxlanGbpAPITunnelMode represents VPP binary API enum 'vxlan_gbp_api_tunnel_mode'.
type VxlanGbpAPITunnelMode uint32

const (
	VXLAN_GBP_API_TUNNEL_MODE_L2 VxlanGbpAPITunnelMode = 1
	VXLAN_GBP_API_TUNNEL_MODE_L3 VxlanGbpAPITunnelMode = 2
)

var VxlanGbpAPITunnelMode_name = map[uint32]string{
	1: "VXLAN_GBP_API_TUNNEL_MODE_L2",
	2: "VXLAN_GBP_API_TUNNEL_MODE_L3",
}

var VxlanGbpAPITunnelMode_value = map[string]uint32{
	"VXLAN_GBP_API_TUNNEL_MODE_L2": 1,
	"VXLAN_GBP_API_TUNNEL_MODE_L3": 2,
}

func (x VxlanGbpAPITunnelMode) String() string {
	s, ok := VxlanGbpAPITunnelMode_name[uint32(x)]
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

func (*Address) GetTypeName() string  { return "address" }
func (*Address) GetCrcString() string { return "09f11671" }

// IP4Prefix represents VPP binary API type 'ip4_prefix'.
type IP4Prefix struct {
	Prefix IP4Address
	Len    uint8
}

func (*IP4Prefix) GetTypeName() string  { return "ip4_prefix" }
func (*IP4Prefix) GetCrcString() string { return "ea8dc11d" }

// IP6Prefix represents VPP binary API type 'ip6_prefix'.
type IP6Prefix struct {
	Prefix IP6Address
	Len    uint8
}

func (*IP6Prefix) GetTypeName() string  { return "ip6_prefix" }
func (*IP6Prefix) GetCrcString() string { return "779fd64f" }

// Mprefix represents VPP binary API type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string  { return "mprefix" }
func (*Mprefix) GetCrcString() string { return "1c4cba05" }

// Prefix represents VPP binary API type 'prefix'.
type Prefix struct {
	Address       Address
	AddressLength uint8
}

func (*Prefix) GetTypeName() string  { return "prefix" }
func (*Prefix) GetCrcString() string { return "0403aebc" }

// VxlanGbpTunnel represents VPP binary API type 'vxlan_gbp_tunnel'.
type VxlanGbpTunnel struct {
	Instance       uint32
	Src            Address
	Dst            Address
	McastSwIfIndex uint32
	EncapTableID   uint32
	Vni            uint32
	SwIfIndex      uint32
	Mode           VxlanGbpAPITunnelMode
}

func (*VxlanGbpTunnel) GetTypeName() string  { return "vxlan_gbp_tunnel" }
func (*VxlanGbpTunnel) GetCrcString() string { return "31974cf7" }

// AddressUnion represents VPP binary API union 'address_union'.
type AddressUnion struct {
	XXX_UnionData [16]byte
}

func (*AddressUnion) GetTypeName() string  { return "address_union" }
func (*AddressUnion) GetCrcString() string { return "d68a2fb4" }

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

// SwInterfaceSetVxlanGbpBypass represents VPP binary API message 'sw_interface_set_vxlan_gbp_bypass'.
type SwInterfaceSetVxlanGbpBypass struct {
	SwIfIndex uint32
	IsIPv6    uint8
	Enable    uint8
}

func (m *SwInterfaceSetVxlanGbpBypass) Reset() { *m = SwInterfaceSetVxlanGbpBypass{} }
func (*SwInterfaceSetVxlanGbpBypass) GetMessageName() string {
	return "sw_interface_set_vxlan_gbp_bypass"
}
func (*SwInterfaceSetVxlanGbpBypass) GetCrcString() string            { return "e74ca095" }
func (*SwInterfaceSetVxlanGbpBypass) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetVxlanGbpBypassReply represents VPP binary API message 'sw_interface_set_vxlan_gbp_bypass_reply'.
type SwInterfaceSetVxlanGbpBypassReply struct {
	Retval int32
}

func (m *SwInterfaceSetVxlanGbpBypassReply) Reset() { *m = SwInterfaceSetVxlanGbpBypassReply{} }
func (*SwInterfaceSetVxlanGbpBypassReply) GetMessageName() string {
	return "sw_interface_set_vxlan_gbp_bypass_reply"
}
func (*SwInterfaceSetVxlanGbpBypassReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSetVxlanGbpBypassReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGbpTunnelAddDel represents VPP binary API message 'vxlan_gbp_tunnel_add_del'.
type VxlanGbpTunnelAddDel struct {
	IsAdd  uint8
	Tunnel VxlanGbpTunnel
}

func (m *VxlanGbpTunnelAddDel) Reset()                        { *m = VxlanGbpTunnelAddDel{} }
func (*VxlanGbpTunnelAddDel) GetMessageName() string          { return "vxlan_gbp_tunnel_add_del" }
func (*VxlanGbpTunnelAddDel) GetCrcString() string            { return "6ac4b80e" }
func (*VxlanGbpTunnelAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGbpTunnelAddDelReply represents VPP binary API message 'vxlan_gbp_tunnel_add_del_reply'.
type VxlanGbpTunnelAddDelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (m *VxlanGbpTunnelAddDelReply) Reset()                        { *m = VxlanGbpTunnelAddDelReply{} }
func (*VxlanGbpTunnelAddDelReply) GetMessageName() string          { return "vxlan_gbp_tunnel_add_del_reply" }
func (*VxlanGbpTunnelAddDelReply) GetCrcString() string            { return "fda5941f" }
func (*VxlanGbpTunnelAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGbpTunnelDetails represents VPP binary API message 'vxlan_gbp_tunnel_details'.
type VxlanGbpTunnelDetails struct {
	Tunnel VxlanGbpTunnel
}

func (m *VxlanGbpTunnelDetails) Reset()                        { *m = VxlanGbpTunnelDetails{} }
func (*VxlanGbpTunnelDetails) GetMessageName() string          { return "vxlan_gbp_tunnel_details" }
func (*VxlanGbpTunnelDetails) GetCrcString() string            { return "7fea68d7" }
func (*VxlanGbpTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGbpTunnelDump represents VPP binary API message 'vxlan_gbp_tunnel_dump'.
type VxlanGbpTunnelDump struct {
	SwIfIndex uint32
}

func (m *VxlanGbpTunnelDump) Reset()                        { *m = VxlanGbpTunnelDump{} }
func (*VxlanGbpTunnelDump) GetMessageName() string          { return "vxlan_gbp_tunnel_dump" }
func (*VxlanGbpTunnelDump) GetCrcString() string            { return "529cb13f" }
func (*VxlanGbpTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*SwInterfaceSetVxlanGbpBypass)(nil), "vxlan_gbp.SwInterfaceSetVxlanGbpBypass")
	api.RegisterMessage((*SwInterfaceSetVxlanGbpBypassReply)(nil), "vxlan_gbp.SwInterfaceSetVxlanGbpBypassReply")
	api.RegisterMessage((*VxlanGbpTunnelAddDel)(nil), "vxlan_gbp.VxlanGbpTunnelAddDel")
	api.RegisterMessage((*VxlanGbpTunnelAddDelReply)(nil), "vxlan_gbp.VxlanGbpTunnelAddDelReply")
	api.RegisterMessage((*VxlanGbpTunnelDetails)(nil), "vxlan_gbp.VxlanGbpTunnelDetails")
	api.RegisterMessage((*VxlanGbpTunnelDump)(nil), "vxlan_gbp.VxlanGbpTunnelDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceSetVxlanGbpBypass)(nil),
		(*SwInterfaceSetVxlanGbpBypassReply)(nil),
		(*VxlanGbpTunnelAddDel)(nil),
		(*VxlanGbpTunnelAddDelReply)(nil),
		(*VxlanGbpTunnelDetails)(nil),
		(*VxlanGbpTunnelDump)(nil),
	}
}

// RPCService represents RPC service API for vxlan_gbp module.
type RPCService interface {
	DumpVxlanGbpTunnel(ctx context.Context, in *VxlanGbpTunnelDump) (RPCService_DumpVxlanGbpTunnelClient, error)
	SwInterfaceSetVxlanGbpBypass(ctx context.Context, in *SwInterfaceSetVxlanGbpBypass) (*SwInterfaceSetVxlanGbpBypassReply, error)
	VxlanGbpTunnelAddDel(ctx context.Context, in *VxlanGbpTunnelAddDel) (*VxlanGbpTunnelAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpVxlanGbpTunnel(ctx context.Context, in *VxlanGbpTunnelDump) (RPCService_DumpVxlanGbpTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpVxlanGbpTunnelClient{stream}
	return x, nil
}

type RPCService_DumpVxlanGbpTunnelClient interface {
	Recv() (*VxlanGbpTunnelDetails, error)
}

type serviceClient_DumpVxlanGbpTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpVxlanGbpTunnelClient) Recv() (*VxlanGbpTunnelDetails, error) {
	m := new(VxlanGbpTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) SwInterfaceSetVxlanGbpBypass(ctx context.Context, in *SwInterfaceSetVxlanGbpBypass) (*SwInterfaceSetVxlanGbpBypassReply, error) {
	out := new(SwInterfaceSetVxlanGbpBypassReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGbpTunnelAddDel(ctx context.Context, in *VxlanGbpTunnelAddDel) (*VxlanGbpTunnelAddDelReply, error) {
	out := new(VxlanGbpTunnelAddDelReply)
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
