// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/punt.api.json

/*
 Package punt is a generated from VPP binary API module 'punt'.

 It contains following objects:
	 10 messages
	  1 type
	  5 services

*/
package punt

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x1c487f2b

// Services represents VPP binary API services:
//
//	"services": {
//	    "set_punt": {
//	        "reply": "set_punt_reply"
//	    },
//	    "punt_socket_dump": {
//	        "reply": "punt_socket_details",
//	        "stream": true
//	    },
//	    "punt_socket_deregister": {
//	        "reply": "punt_socket_deregister_reply"
//	    },
//	    "punt_dump": {
//	        "reply": "punt_details",
//	        "stream": true
//	    },
//	    "punt_socket_register": {
//	        "reply": "punt_socket_register_reply"
//	    }
//	},
//
type Services interface {
	DumpPunt(*PuntDump) ([]*PuntDetails, error)
	DumpPuntSocket(*PuntSocketDump) ([]*PuntSocketDetails, error)
	PuntSocketDeregister(*PuntSocketDeregister) (*PuntSocketDeregisterReply, error)
	PuntSocketRegister(*PuntSocketRegister) (*PuntSocketRegisterReply, error)
	SetPunt(*SetPunt) (*SetPuntReply, error)
}

/* Types */

// Punt represents VPP binary API type 'punt':
//
//	"punt",
//	[
//	    "u8",
//	    "ipv"
//	],
//	[
//	    "u8",
//	    "l4_protocol"
//	],
//	[
//	    "u16",
//	    "l4_port"
//	],
//	{
//	    "crc": "0xfe4f98ac"
//	}
//
type Punt struct {
	IPv        uint8
	L4Protocol uint8
	L4Port     uint16
}

func (*Punt) GetTypeName() string {
	return "punt"
}
func (*Punt) GetCrcString() string {
	return "fe4f98ac"
}

/* Messages */

// SetPunt represents VPP binary API message 'set_punt':
//
//	"set_punt",
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
//	    "vl_api_punt_t",
//	    "punt"
//	],
//	{
//	    "crc": "0x332d88dc"
//	}
//
type SetPunt struct {
	IsAdd uint8
	Punt  Punt
}

func (*SetPunt) GetMessageName() string {
	return "set_punt"
}
func (*SetPunt) GetCrcString() string {
	return "332d88dc"
}
func (*SetPunt) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SetPuntReply represents VPP binary API message 'set_punt_reply':
//
//	"set_punt_reply",
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
type SetPuntReply struct {
	Retval int32
}

func (*SetPuntReply) GetMessageName() string {
	return "set_punt_reply"
}
func (*SetPuntReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SetPuntReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PuntDump represents VPP binary API message 'punt_dump':
//
//	"punt_dump",
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
//	{
//	    "crc": "0xde883da4"
//	}
//
type PuntDump struct {
	IsIPv6 uint8
}

func (*PuntDump) GetMessageName() string {
	return "punt_dump"
}
func (*PuntDump) GetCrcString() string {
	return "de883da4"
}
func (*PuntDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PuntDetails represents VPP binary API message 'punt_details':
//
//	"punt_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_punt_t",
//	    "punt"
//	],
//	{
//	    "crc": "0xe905318e"
//	}
//
type PuntDetails struct {
	Punt Punt
}

func (*PuntDetails) GetMessageName() string {
	return "punt_details"
}
func (*PuntDetails) GetCrcString() string {
	return "e905318e"
}
func (*PuntDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PuntSocketRegister represents VPP binary API message 'punt_socket_register':
//
//	"punt_socket_register",
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
//	    "header_version"
//	],
//	[
//	    "vl_api_punt_t",
//	    "punt"
//	],
//	[
//	    "u8",
//	    "pathname",
//	    108
//	],
//	{
//	    "crc": "0x9f3e2877"
//	}
//
type PuntSocketRegister struct {
	HeaderVersion uint32
	Punt          Punt
	Pathname      []byte `struc:"[108]byte"`
}

func (*PuntSocketRegister) GetMessageName() string {
	return "punt_socket_register"
}
func (*PuntSocketRegister) GetCrcString() string {
	return "9f3e2877"
}
func (*PuntSocketRegister) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PuntSocketRegisterReply represents VPP binary API message 'punt_socket_register_reply':
//
//	"punt_socket_register_reply",
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
//	    "u8",
//	    "pathname",
//	    64
//	],
//	{
//	    "crc": "0x42dc0ee6"
//	}
//
type PuntSocketRegisterReply struct {
	Retval   int32
	Pathname []byte `struc:"[64]byte"`
}

func (*PuntSocketRegisterReply) GetMessageName() string {
	return "punt_socket_register_reply"
}
func (*PuntSocketRegisterReply) GetCrcString() string {
	return "42dc0ee6"
}
func (*PuntSocketRegisterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PuntSocketDump represents VPP binary API message 'punt_socket_dump':
//
//	"punt_socket_dump",
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
//	{
//	    "crc": "0xde883da4"
//	}
//
type PuntSocketDump struct {
	IsIPv6 uint8
}

func (*PuntSocketDump) GetMessageName() string {
	return "punt_socket_dump"
}
func (*PuntSocketDump) GetCrcString() string {
	return "de883da4"
}
func (*PuntSocketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PuntSocketDetails represents VPP binary API message 'punt_socket_details':
//
//	"punt_socket_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_punt_t",
//	    "punt"
//	],
//	[
//	    "u8",
//	    "pathname",
//	    108
//	],
//	{
//	    "crc": "0x8911c6c5"
//	}
//
type PuntSocketDetails struct {
	Punt     Punt
	Pathname []byte `struc:"[108]byte"`
}

func (*PuntSocketDetails) GetMessageName() string {
	return "punt_socket_details"
}
func (*PuntSocketDetails) GetCrcString() string {
	return "8911c6c5"
}
func (*PuntSocketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PuntSocketDeregister represents VPP binary API message 'punt_socket_deregister':
//
//	"punt_socket_deregister",
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
//	    "vl_api_punt_t",
//	    "punt"
//	],
//	{
//	    "crc": "0x0603ba46"
//	}
//
type PuntSocketDeregister struct {
	Punt Punt
}

func (*PuntSocketDeregister) GetMessageName() string {
	return "punt_socket_deregister"
}
func (*PuntSocketDeregister) GetCrcString() string {
	return "0603ba46"
}
func (*PuntSocketDeregister) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PuntSocketDeregisterReply represents VPP binary API message 'punt_socket_deregister_reply':
//
//	"punt_socket_deregister_reply",
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
type PuntSocketDeregisterReply struct {
	Retval int32
}

func (*PuntSocketDeregisterReply) GetMessageName() string {
	return "punt_socket_deregister_reply"
}
func (*PuntSocketDeregisterReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PuntSocketDeregisterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SetPunt)(nil), "punt.SetPunt")
	api.RegisterMessage((*SetPuntReply)(nil), "punt.SetPuntReply")
	api.RegisterMessage((*PuntDump)(nil), "punt.PuntDump")
	api.RegisterMessage((*PuntDetails)(nil), "punt.PuntDetails")
	api.RegisterMessage((*PuntSocketRegister)(nil), "punt.PuntSocketRegister")
	api.RegisterMessage((*PuntSocketRegisterReply)(nil), "punt.PuntSocketRegisterReply")
	api.RegisterMessage((*PuntSocketDump)(nil), "punt.PuntSocketDump")
	api.RegisterMessage((*PuntSocketDetails)(nil), "punt.PuntSocketDetails")
	api.RegisterMessage((*PuntSocketDeregister)(nil), "punt.PuntSocketDeregister")
	api.RegisterMessage((*PuntSocketDeregisterReply)(nil), "punt.PuntSocketDeregisterReply")
}
