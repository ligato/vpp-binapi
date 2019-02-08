// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/vxlan_gpe.api.json

/*
 Package vxlan_gpe is a generated from VPP binary API module 'vxlan_gpe'.

 It contains following objects:
	  3 services
	  6 messages
*/
package vxlan_gpe

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x7849c91

// Services represents VPP binary API services:
//
//	"services": {
//	    "vxlan_gpe_add_del_tunnel": {
//	        "reply": "vxlan_gpe_add_del_tunnel_reply"
//	    },
//	    "sw_interface_set_vxlan_gpe_bypass": {
//	        "reply": "sw_interface_set_vxlan_gpe_bypass_reply"
//	    },
//	    "vxlan_gpe_tunnel_dump": {
//	        "reply": "vxlan_gpe_tunnel_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpVxlanGpeTunnel(*VxlanGpeTunnelDump) ([]*VxlanGpeTunnelDetails, error)
	SwInterfaceSetVxlanGpeBypass(*SwInterfaceSetVxlanGpeBypass) (*SwInterfaceSetVxlanGpeBypassReply, error)
	VxlanGpeAddDelTunnel(*VxlanGpeAddDelTunnel) (*VxlanGpeAddDelTunnelReply, error)
}

/* Messages */

// SwInterfaceSetVxlanGpeBypass represents VPP binary API message 'sw_interface_set_vxlan_gpe_bypass':
//
//	"sw_interface_set_vxlan_gpe_bypass",
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
type SwInterfaceSetVxlanGpeBypass struct {
	SwIfIndex uint32
	IsIPv6    uint8
	Enable    uint8
}

func (*SwInterfaceSetVxlanGpeBypass) GetMessageName() string {
	return "sw_interface_set_vxlan_gpe_bypass"
}
func (*SwInterfaceSetVxlanGpeBypass) GetCrcString() string {
	return "e74ca095"
}
func (*SwInterfaceSetVxlanGpeBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetVxlanGpeBypassReply represents VPP binary API message 'sw_interface_set_vxlan_gpe_bypass_reply':
//
//	"sw_interface_set_vxlan_gpe_bypass_reply",
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
type SwInterfaceSetVxlanGpeBypassReply struct {
	Retval int32
}

func (*SwInterfaceSetVxlanGpeBypassReply) GetMessageName() string {
	return "sw_interface_set_vxlan_gpe_bypass_reply"
}
func (*SwInterfaceSetVxlanGpeBypassReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetVxlanGpeBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// VxlanGpeAddDelTunnel represents VPP binary API message 'vxlan_gpe_add_del_tunnel':
//
//	"vxlan_gpe_add_del_tunnel",
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
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "local",
//	    16
//	],
//	[
//	    "u8",
//	    "remote",
//	    16
//	],
//	[
//	    "u32",
//	    "mcast_sw_if_index"
//	],
//	[
//	    "u32",
//	    "encap_vrf_id"
//	],
//	[
//	    "u32",
//	    "decap_vrf_id"
//	],
//	[
//	    "u8",
//	    "protocol"
//	],
//	[
//	    "u32",
//	    "vni"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	{
//	    "crc": "0xd15850ba"
//	}
//
type VxlanGpeAddDelTunnel struct {
	IsIPv6         uint8
	Local          []byte `struc:"[16]byte"`
	Remote         []byte `struc:"[16]byte"`
	McastSwIfIndex uint32
	EncapVrfID     uint32
	DecapVrfID     uint32
	Protocol       uint8
	Vni            uint32
	IsAdd          uint8
}

func (*VxlanGpeAddDelTunnel) GetMessageName() string {
	return "vxlan_gpe_add_del_tunnel"
}
func (*VxlanGpeAddDelTunnel) GetCrcString() string {
	return "d15850ba"
}
func (*VxlanGpeAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// VxlanGpeAddDelTunnelReply represents VPP binary API message 'vxlan_gpe_add_del_tunnel_reply':
//
//	"vxlan_gpe_add_del_tunnel_reply",
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
type VxlanGpeAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*VxlanGpeAddDelTunnelReply) GetMessageName() string {
	return "vxlan_gpe_add_del_tunnel_reply"
}
func (*VxlanGpeAddDelTunnelReply) GetCrcString() string {
	return "fda5941f"
}
func (*VxlanGpeAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// VxlanGpeTunnelDetails represents VPP binary API message 'vxlan_gpe_tunnel_details':
//
//	"vxlan_gpe_tunnel_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
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
//	    "local",
//	    16
//	],
//	[
//	    "u8",
//	    "remote",
//	    16
//	],
//	[
//	    "u32",
//	    "vni"
//	],
//	[
//	    "u8",
//	    "protocol"
//	],
//	[
//	    "u32",
//	    "mcast_sw_if_index"
//	],
//	[
//	    "u32",
//	    "encap_vrf_id"
//	],
//	[
//	    "u32",
//	    "decap_vrf_id"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	{
//	    "crc": "0x2673fbfa"
//	}
//
type VxlanGpeTunnelDetails struct {
	SwIfIndex      uint32
	Local          []byte `struc:"[16]byte"`
	Remote         []byte `struc:"[16]byte"`
	Vni            uint32
	Protocol       uint8
	McastSwIfIndex uint32
	EncapVrfID     uint32
	DecapVrfID     uint32
	IsIPv6         uint8
}

func (*VxlanGpeTunnelDetails) GetMessageName() string {
	return "vxlan_gpe_tunnel_details"
}
func (*VxlanGpeTunnelDetails) GetCrcString() string {
	return "2673fbfa"
}
func (*VxlanGpeTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// VxlanGpeTunnelDump represents VPP binary API message 'vxlan_gpe_tunnel_dump':
//
//	"vxlan_gpe_tunnel_dump",
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
type VxlanGpeTunnelDump struct {
	SwIfIndex uint32
}

func (*VxlanGpeTunnelDump) GetMessageName() string {
	return "vxlan_gpe_tunnel_dump"
}
func (*VxlanGpeTunnelDump) GetCrcString() string {
	return "529cb13f"
}
func (*VxlanGpeTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*SwInterfaceSetVxlanGpeBypass)(nil), "vxlan_gpe.SwInterfaceSetVxlanGpeBypass")
	api.RegisterMessage((*SwInterfaceSetVxlanGpeBypassReply)(nil), "vxlan_gpe.SwInterfaceSetVxlanGpeBypassReply")
	api.RegisterMessage((*VxlanGpeAddDelTunnel)(nil), "vxlan_gpe.VxlanGpeAddDelTunnel")
	api.RegisterMessage((*VxlanGpeAddDelTunnelReply)(nil), "vxlan_gpe.VxlanGpeAddDelTunnelReply")
	api.RegisterMessage((*VxlanGpeTunnelDetails)(nil), "vxlan_gpe.VxlanGpeTunnelDetails")
	api.RegisterMessage((*VxlanGpeTunnelDump)(nil), "vxlan_gpe.VxlanGpeTunnelDump")
}
