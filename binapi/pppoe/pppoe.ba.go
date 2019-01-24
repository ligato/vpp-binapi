// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/pppoe.api.json

/*
 Package pppoe is a generated from VPP binary API module 'pppoe'.

 It contains following objects:
	  4 messages
	  2 services

*/
package pppoe

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x4def67c4

// Services represents VPP binary API services:
//
//	"services": {
//	    "pppoe_session_dump": {
//	        "reply": "pppoe_session_details",
//	        "stream": true
//	    },
//	    "pppoe_add_del_session": {
//	        "reply": "pppoe_add_del_session_reply"
//	    }
//	},
//
type Services interface {
	DumpPppoeSession(*PppoeSessionDump) ([]*PppoeSessionDetails, error)
	PppoeAddDelSession(*PppoeAddDelSession) (*PppoeAddDelSessionReply, error)
}

/* Messages */

// PppoeAddDelSession represents VPP binary API message 'pppoe_add_del_session':
//
//	"pppoe_add_del_session",
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
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u16",
//	    "session_id"
//	],
//	[
//	    "u8",
//	    "client_ip",
//	    16
//	],
//	[
//	    "u32",
//	    "decap_vrf_id"
//	],
//	[
//	    "u8",
//	    "client_mac",
//	    6
//	],
//	{
//	    "crc": "0x766cbfeb"
//	}
//
type PppoeAddDelSession struct {
	IsAdd      uint8
	IsIPv6     uint8
	SessionID  uint16
	ClientIP   []byte `struc:"[16]byte"`
	DecapVrfID uint32
	ClientMac  []byte `struc:"[6]byte"`
}

func (*PppoeAddDelSession) GetMessageName() string {
	return "pppoe_add_del_session"
}
func (*PppoeAddDelSession) GetCrcString() string {
	return "766cbfeb"
}
func (*PppoeAddDelSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PppoeAddDelSessionReply represents VPP binary API message 'pppoe_add_del_session_reply':
//
//	"pppoe_add_del_session_reply",
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
type PppoeAddDelSessionReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*PppoeAddDelSessionReply) GetMessageName() string {
	return "pppoe_add_del_session_reply"
}
func (*PppoeAddDelSessionReply) GetCrcString() string {
	return "fda5941f"
}
func (*PppoeAddDelSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PppoeSessionDump represents VPP binary API message 'pppoe_session_dump':
//
//	"pppoe_session_dump",
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
type PppoeSessionDump struct {
	SwIfIndex uint32
}

func (*PppoeSessionDump) GetMessageName() string {
	return "pppoe_session_dump"
}
func (*PppoeSessionDump) GetCrcString() string {
	return "529cb13f"
}
func (*PppoeSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PppoeSessionDetails represents VPP binary API message 'pppoe_session_details':
//
//	"pppoe_session_details",
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
//	    "is_ipv6"
//	],
//	[
//	    "u16",
//	    "session_id"
//	],
//	[
//	    "u8",
//	    "client_ip",
//	    16
//	],
//	[
//	    "u32",
//	    "encap_if_index"
//	],
//	[
//	    "u32",
//	    "decap_vrf_id"
//	],
//	[
//	    "u8",
//	    "local_mac",
//	    6
//	],
//	[
//	    "u8",
//	    "client_mac",
//	    6
//	],
//	{
//	    "crc": "0x358fc7a8"
//	}
//
type PppoeSessionDetails struct {
	SwIfIndex    uint32
	IsIPv6       uint8
	SessionID    uint16
	ClientIP     []byte `struc:"[16]byte"`
	EncapIfIndex uint32
	DecapVrfID   uint32
	LocalMac     []byte `struc:"[6]byte"`
	ClientMac    []byte `struc:"[6]byte"`
}

func (*PppoeSessionDetails) GetMessageName() string {
	return "pppoe_session_details"
}
func (*PppoeSessionDetails) GetCrcString() string {
	return "358fc7a8"
}
func (*PppoeSessionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*PppoeAddDelSession)(nil), "pppoe.PppoeAddDelSession")
	api.RegisterMessage((*PppoeAddDelSessionReply)(nil), "pppoe.PppoeAddDelSessionReply")
	api.RegisterMessage((*PppoeSessionDump)(nil), "pppoe.PppoeSessionDump")
	api.RegisterMessage((*PppoeSessionDetails)(nil), "pppoe.PppoeSessionDetails")
}
