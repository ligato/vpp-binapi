// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/ipsec_gre.api.json

/*
Package ipsec_gre is a generated VPP binary API for 'ipsec_gre' module.

It consists of:
	  1 enum
	  2 aliases
	  6 types
	  1 union
	  4 messages
	  2 services
*/
package ipsec_gre

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
	ModuleName = "ipsec_gre"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xa52c0f0a
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

// IpsecGreTunnel represents VPP binary API type 'ipsec_gre_tunnel'.
type IpsecGreTunnel struct {
	ClientIndex uint32
	Context     uint32
	LocalSaID   uint32
	RemoteSaID  uint32
	IsAdd       uint8
	SwIfIndex   uint32
	Src         IP4Address
	Dst         IP4Address
}

func (*IpsecGreTunnel) GetTypeName() string  { return "ipsec_gre_tunnel" }
func (*IpsecGreTunnel) GetCrcString() string { return "f9425226" }

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

// IpsecGreTunnelAddDel represents VPP binary API message 'ipsec_gre_tunnel_add_del'.
type IpsecGreTunnelAddDel struct {
	IsAdd  uint8
	Tunnel IpsecGreTunnel
}

func (m *IpsecGreTunnelAddDel) Reset()                        { *m = IpsecGreTunnelAddDel{} }
func (*IpsecGreTunnelAddDel) GetMessageName() string          { return "ipsec_gre_tunnel_add_del" }
func (*IpsecGreTunnelAddDel) GetCrcString() string            { return "17985676" }
func (*IpsecGreTunnelAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// IpsecGreTunnelAddDelReply represents VPP binary API message 'ipsec_gre_tunnel_add_del_reply'.
type IpsecGreTunnelAddDelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (m *IpsecGreTunnelAddDelReply) Reset()                        { *m = IpsecGreTunnelAddDelReply{} }
func (*IpsecGreTunnelAddDelReply) GetMessageName() string          { return "ipsec_gre_tunnel_add_del_reply" }
func (*IpsecGreTunnelAddDelReply) GetCrcString() string            { return "fda5941f" }
func (*IpsecGreTunnelAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// IpsecGreTunnelDetails represents VPP binary API message 'ipsec_gre_tunnel_details'.
type IpsecGreTunnelDetails struct {
	Tunnel IpsecGreTunnel
}

func (m *IpsecGreTunnelDetails) Reset()                        { *m = IpsecGreTunnelDetails{} }
func (*IpsecGreTunnelDetails) GetMessageName() string          { return "ipsec_gre_tunnel_details" }
func (*IpsecGreTunnelDetails) GetCrcString() string            { return "02b686af" }
func (*IpsecGreTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// IpsecGreTunnelDump represents VPP binary API message 'ipsec_gre_tunnel_dump'.
type IpsecGreTunnelDump struct {
	SwIfIndex uint32
}

func (m *IpsecGreTunnelDump) Reset()                        { *m = IpsecGreTunnelDump{} }
func (*IpsecGreTunnelDump) GetMessageName() string          { return "ipsec_gre_tunnel_dump" }
func (*IpsecGreTunnelDump) GetCrcString() string            { return "529cb13f" }
func (*IpsecGreTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*IpsecGreTunnelAddDel)(nil), "ipsec_gre.IpsecGreTunnelAddDel")
	api.RegisterMessage((*IpsecGreTunnelAddDelReply)(nil), "ipsec_gre.IpsecGreTunnelAddDelReply")
	api.RegisterMessage((*IpsecGreTunnelDetails)(nil), "ipsec_gre.IpsecGreTunnelDetails")
	api.RegisterMessage((*IpsecGreTunnelDump)(nil), "ipsec_gre.IpsecGreTunnelDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IpsecGreTunnelAddDel)(nil),
		(*IpsecGreTunnelAddDelReply)(nil),
		(*IpsecGreTunnelDetails)(nil),
		(*IpsecGreTunnelDump)(nil),
	}
}

// RPCService represents RPC service API for ipsec_gre module.
type RPCService interface {
	DumpIpsecGreTunnel(ctx context.Context, in *IpsecGreTunnelDump) (RPCService_DumpIpsecGreTunnelClient, error)
	IpsecGreTunnelAddDel(ctx context.Context, in *IpsecGreTunnelAddDel) (*IpsecGreTunnelAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpIpsecGreTunnel(ctx context.Context, in *IpsecGreTunnelDump) (RPCService_DumpIpsecGreTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIpsecGreTunnelClient{stream}
	return x, nil
}

type RPCService_DumpIpsecGreTunnelClient interface {
	Recv() (*IpsecGreTunnelDetails, error)
}

type serviceClient_DumpIpsecGreTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIpsecGreTunnelClient) Recv() (*IpsecGreTunnelDetails, error) {
	m := new(IpsecGreTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) IpsecGreTunnelAddDel(ctx context.Context, in *IpsecGreTunnelAddDel) (*IpsecGreTunnelAddDelReply, error) {
	out := new(IpsecGreTunnelAddDelReply)
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
