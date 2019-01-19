// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/vxlan_gbp.api.json

/*
 Package vxlan_gbp is a generated from VPP binary API module 'vxlan_gbp'.

 It contains following objects:
	  6 messages
	  6 types
	  2 aliases
	  1 enum
	  1 union
	  3 services

*/
package vxlan_gbp

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0xff1fa142

// Services represents VPP binary API services:
//
//	"services": {
//	    "vxlan_gbp_tunnel_dump": {
//	        "reply": "vxlan_gbp_tunnel_details",
//	        "stream": true
//	    },
//	    "sw_interface_set_vxlan_gbp_bypass": {
//	        "reply": "sw_interface_set_vxlan_gbp_bypass_reply"
//	    },
//	    "vxlan_gbp_tunnel_add_del": {
//	        "reply": "vxlan_gbp_tunnel_add_del_reply"
//	    }
//	},
//
type Services interface {
	DumpVxlanGbpTunnel(*VxlanGbpTunnelDump) ([]*VxlanGbpTunnelDetails, error)
	SwInterfaceSetVxlanGbpBypass(*SwInterfaceSetVxlanGbpBypass) (*SwInterfaceSetVxlanGbpBypassReply, error)
	VxlanGbpTunnelAddDel(*VxlanGbpTunnelAddDel) (*VxlanGbpTunnelAddDelReply, error)
}

/* Enums */

// AddressFamily represents VPP binary API enum 'address_family':
//
//	"address_family",
//	[
//	    "ADDRESS_IP4",
//	    0
//	],
//	[
//	    "ADDRESS_IP6",
//	    1
//	],
//	{
//	    "enumtype": "u32"
//	}
//
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

/* Aliases */

// IP4Address represents VPP binary API alias 'ip4_address':
//
//	"ip4_address": {
//	    "length": 4,
//	    "type": "u8"
//	}
//
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address':
//
//	"ip6_address": {
//	    "length": 16,
//	    "type": "u8"
//	},
//
type IP6Address [16]uint8

/* Types */

// Address represents VPP binary API type 'address':
//
//	"address",
//	[
//	    "vl_api_address_family_t",
//	    "af"
//	],
//	[
//	    "vl_api_address_union_t",
//	    "un"
//	],
//	{
//	    "crc": "0x09f11671"
//	}
//
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}
func (*Address) GetCrcString() string {
	return "09f11671"
}

// Prefix represents VPP binary API type 'prefix':
//
//	"prefix",
//	[
//	    "vl_api_address_t",
//	    "address"
//	],
//	[
//	    "u8",
//	    "address_length"
//	],
//	{
//	    "crc": "0x0403aebc"
//	}
//
type Prefix struct {
	Address       Address
	AddressLength uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}
func (*Prefix) GetCrcString() string {
	return "0403aebc"
}

// Mprefix represents VPP binary API type 'mprefix':
//
//	"mprefix",
//	[
//	    "vl_api_address_family_t",
//	    "af"
//	],
//	[
//	    "u16",
//	    "grp_address_length"
//	],
//	[
//	    "vl_api_address_union_t",
//	    "grp_address"
//	],
//	[
//	    "vl_api_address_union_t",
//	    "src_address"
//	],
//	{
//	    "crc": "0x1c4cba05"
//	}
//
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}
func (*Mprefix) GetCrcString() string {
	return "1c4cba05"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix':
//
//	"ip6_prefix",
//	[
//	    "vl_api_ip6_address_t",
//	    "prefix"
//	],
//	[
//	    "u8",
//	    "len"
//	],
//	{
//	    "crc": "0x779fd64f"
//	}
//
type IP6Prefix struct {
	Prefix IP6Address
	Len    uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}
func (*IP6Prefix) GetCrcString() string {
	return "779fd64f"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix':
//
//	"ip4_prefix",
//	[
//	    "vl_api_ip4_address_t",
//	    "prefix"
//	],
//	[
//	    "u8",
//	    "len"
//	],
//	{
//	    "crc": "0xea8dc11d"
//	}
//
type IP4Prefix struct {
	Prefix IP4Address
	Len    uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}
func (*IP4Prefix) GetCrcString() string {
	return "ea8dc11d"
}

// VxlanGbpTunnel represents VPP binary API type 'vxlan_gbp_tunnel':
//
//	"vxlan_gbp_tunnel",
//	[
//	    "u32",
//	    "instance"
//	],
//	[
//	    "vl_api_address_t",
//	    "src"
//	],
//	[
//	    "vl_api_address_t",
//	    "dst"
//	],
//	[
//	    "u32",
//	    "mcast_sw_if_index"
//	],
//	[
//	    "u32",
//	    "encap_table_id"
//	],
//	[
//	    "u32",
//	    "vni"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x2bb00ea0"
//	}
//
type VxlanGbpTunnel struct {
	Instance       uint32
	Src            Address
	Dst            Address
	McastSwIfIndex uint32
	EncapTableID   uint32
	Vni            uint32
	SwIfIndex      uint32
}

func (*VxlanGbpTunnel) GetTypeName() string {
	return "vxlan_gbp_tunnel"
}
func (*VxlanGbpTunnel) GetCrcString() string {
	return "2bb00ea0"
}

/* Unions */

// AddressUnion represents VPP binary API union 'address_union':
//
//	"address_union",
//	[
//	    "vl_api_ip4_address_t",
//	    "ip4"
//	],
//	[
//	    "vl_api_ip6_address_t",
//	    "ip6"
//	],
//	{
//	    "crc": "0xd68a2fb4"
//	}
//
type AddressUnion struct {
	Union_data [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}
func (*AddressUnion) GetCrcString() string {
	return "d68a2fb4"
}

func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

/* Messages */

// VxlanGbpTunnelAddDel represents VPP binary API message 'vxlan_gbp_tunnel_add_del':
//
//	"vxlan_gbp_tunnel_add_del",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	[
//	    "vl_api_vxlan_gbp_tunnel_t",
//	    "tunnel"
//	],
//	{
//	    "crc": "0x6ac4b80e"
//	}
//
type VxlanGbpTunnelAddDel struct {
	IsAdd  uint8
	Tunnel VxlanGbpTunnel
}

func (*VxlanGbpTunnelAddDel) GetMessageName() string {
	return "vxlan_gbp_tunnel_add_del"
}
func (*VxlanGbpTunnelAddDel) GetCrcString() string {
	return "6ac4b80e"
}
func (*VxlanGbpTunnelAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// VxlanGbpTunnelAddDelReply represents VPP binary API message 'vxlan_gbp_tunnel_add_del_reply':
//
//	"vxlan_gbp_tunnel_add_del_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0xfda5941f"
//	}
//
type VxlanGbpTunnelAddDelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*VxlanGbpTunnelAddDelReply) GetMessageName() string {
	return "vxlan_gbp_tunnel_add_del_reply"
}
func (*VxlanGbpTunnelAddDelReply) GetCrcString() string {
	return "fda5941f"
}
func (*VxlanGbpTunnelAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// VxlanGbpTunnelDump represents VPP binary API message 'vxlan_gbp_tunnel_dump':
//
//	"vxlan_gbp_tunnel_dump",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x529cb13f"
//	}
//
type VxlanGbpTunnelDump struct {
	SwIfIndex uint32
}

func (*VxlanGbpTunnelDump) GetMessageName() string {
	return "vxlan_gbp_tunnel_dump"
}
func (*VxlanGbpTunnelDump) GetCrcString() string {
	return "529cb13f"
}
func (*VxlanGbpTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// VxlanGbpTunnelDetails represents VPP binary API message 'vxlan_gbp_tunnel_details':
//
//	"vxlan_gbp_tunnel_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_vxlan_gbp_tunnel_t",
//	    "tunnel"
//	],
//	{
//	    "crc": "0x7fea68d7"
//	}
//
type VxlanGbpTunnelDetails struct {
	Tunnel VxlanGbpTunnel
}

func (*VxlanGbpTunnelDetails) GetMessageName() string {
	return "vxlan_gbp_tunnel_details"
}
func (*VxlanGbpTunnelDetails) GetCrcString() string {
	return "7fea68d7"
}
func (*VxlanGbpTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSetVxlanGbpBypass represents VPP binary API message 'sw_interface_set_vxlan_gbp_bypass':
//
//	"sw_interface_set_vxlan_gbp_bypass",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0xe74ca095"
//	}
//
type SwInterfaceSetVxlanGbpBypass struct {
	SwIfIndex uint32
	IsIPv6    uint8
	Enable    uint8
}

func (*SwInterfaceSetVxlanGbpBypass) GetMessageName() string {
	return "sw_interface_set_vxlan_gbp_bypass"
}
func (*SwInterfaceSetVxlanGbpBypass) GetCrcString() string {
	return "e74ca095"
}
func (*SwInterfaceSetVxlanGbpBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetVxlanGbpBypassReply represents VPP binary API message 'sw_interface_set_vxlan_gbp_bypass_reply':
//
//	"sw_interface_set_vxlan_gbp_bypass_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type SwInterfaceSetVxlanGbpBypassReply struct {
	Retval int32
}

func (*SwInterfaceSetVxlanGbpBypassReply) GetMessageName() string {
	return "sw_interface_set_vxlan_gbp_bypass_reply"
}
func (*SwInterfaceSetVxlanGbpBypassReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetVxlanGbpBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*VxlanGbpTunnelAddDel)(nil), "vxlan_gbp.VxlanGbpTunnelAddDel")
	api.RegisterMessage((*VxlanGbpTunnelAddDelReply)(nil), "vxlan_gbp.VxlanGbpTunnelAddDelReply")
	api.RegisterMessage((*VxlanGbpTunnelDump)(nil), "vxlan_gbp.VxlanGbpTunnelDump")
	api.RegisterMessage((*VxlanGbpTunnelDetails)(nil), "vxlan_gbp.VxlanGbpTunnelDetails")
	api.RegisterMessage((*SwInterfaceSetVxlanGbpBypass)(nil), "vxlan_gbp.SwInterfaceSetVxlanGbpBypass")
	api.RegisterMessage((*SwInterfaceSetVxlanGbpBypassReply)(nil), "vxlan_gbp.SwInterfaceSetVxlanGbpBypassReply")
}
