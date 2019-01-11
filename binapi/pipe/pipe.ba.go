// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/pipe.api.json

/*
 Package pipe is a generated from VPP binary API module 'pipe'.

 It contains following objects:
	  6 messages
	  3 services

*/
package pipe

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
//
//	"services": {
//	    "pipe_dump": {
//	        "reply": "pipe_details",
//	        "stream": true
//	    },
//	    "pipe_delete": {
//	        "reply": "pipe_delete_reply"
//	    },
//	    "pipe_create": {
//	        "reply": "pipe_create_reply"
//	    }
//	},
//
type Services interface {
	DumpPipe(*PipeDump) ([]*PipeDetails, error)
	PipeCreate(*PipeCreate) (*PipeCreateReply, error)
	PipeDelete(*PipeDelete) (*PipeDeleteReply, error)
}

/* Messages */

// PipeCreate represents VPP binary API message 'pipe_create':
//
//	"pipe_create",
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
//	    "is_specified"
//	],
//	[
//	    "u32",
//	    "user_instance"
//	],
//	{
//	    "crc": "0xdb413409"
//	}
//
type PipeCreate struct {
	IsSpecified  uint8
	UserInstance uint32
}

func (*PipeCreate) GetMessageName() string {
	return "pipe_create"
}
func (*PipeCreate) GetCrcString() string {
	return "db413409"
}
func (*PipeCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PipeCreateReply represents VPP binary API message 'pipe_create_reply':
//
//	"pipe_create_reply",
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
//	[
//	    "u32",
//	    "pipe_sw_if_index",
//	    2
//	],
//	{
//	    "crc": "0x9f0eef14"
//	}
//
type PipeCreateReply struct {
	Retval        int32
	SwIfIndex     uint32
	PipeSwIfIndex []uint32 `struc:"[2]uint32"`
}

func (*PipeCreateReply) GetMessageName() string {
	return "pipe_create_reply"
}
func (*PipeCreateReply) GetCrcString() string {
	return "9f0eef14"
}
func (*PipeCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PipeDelete represents VPP binary API message 'pipe_delete':
//
//	"pipe_delete",
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
type PipeDelete struct {
	SwIfIndex uint32
}

func (*PipeDelete) GetMessageName() string {
	return "pipe_delete"
}
func (*PipeDelete) GetCrcString() string {
	return "529cb13f"
}
func (*PipeDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PipeDeleteReply represents VPP binary API message 'pipe_delete_reply':
//
//	"pipe_delete_reply",
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
type PipeDeleteReply struct {
	Retval int32
}

func (*PipeDeleteReply) GetMessageName() string {
	return "pipe_delete_reply"
}
func (*PipeDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PipeDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PipeDump represents VPP binary API message 'pipe_dump':
//
//	"pipe_dump",
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
//	{
//	    "crc": "0x51077d14"
//	}
//
type PipeDump struct{}

func (*PipeDump) GetMessageName() string {
	return "pipe_dump"
}
func (*PipeDump) GetCrcString() string {
	return "51077d14"
}
func (*PipeDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PipeDetails represents VPP binary API message 'pipe_details':
//
//	"pipe_details",
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
//	    "u32",
//	    "pipe_sw_if_index",
//	    2
//	],
//	[
//	    "u32",
//	    "instance"
//	],
//	{
//	    "crc": "0x91286b09"
//	}
//
type PipeDetails struct {
	SwIfIndex     uint32
	PipeSwIfIndex []uint32 `struc:"[2]uint32"`
	Instance      uint32
}

func (*PipeDetails) GetMessageName() string {
	return "pipe_details"
}
func (*PipeDetails) GetCrcString() string {
	return "91286b09"
}
func (*PipeDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*PipeCreate)(nil), "pipe.PipeCreate")
	api.RegisterMessage((*PipeCreateReply)(nil), "pipe.PipeCreateReply")
	api.RegisterMessage((*PipeDelete)(nil), "pipe.PipeDelete")
	api.RegisterMessage((*PipeDeleteReply)(nil), "pipe.PipeDeleteReply")
	api.RegisterMessage((*PipeDump)(nil), "pipe.PipeDump")
	api.RegisterMessage((*PipeDetails)(nil), "pipe.PipeDetails")
}