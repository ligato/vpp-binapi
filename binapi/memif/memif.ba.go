// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/memif.api.json

/*
 Package memif is a generated from VPP binary API module 'memif'.

 It contains following objects:
	 10 messages
	  5 services

*/
package memif

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
//	    "memif_socket_filename_add_del": {
//	        "reply": "memif_socket_filename_add_del_reply"
//	    },
//	    "memif_create": {
//	        "reply": "memif_create_reply"
//	    },
//	    "memif_delete": {
//	        "reply": "memif_delete_reply"
//	    },
//	    "memif_socket_filename_dump": {
//	        "reply": "memif_socket_filename_details",
//	        "stream": true
//	    },
//	    "memif_dump": {
//	        "reply": "memif_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpMemif(*MemifDump) ([]*MemifDetails, error)
	DumpMemifSocketFilename(*MemifSocketFilenameDump) ([]*MemifSocketFilenameDetails, error)
	MemifCreate(*MemifCreate) (*MemifCreateReply, error)
	MemifDelete(*MemifDelete) (*MemifDeleteReply, error)
	MemifSocketFilenameAddDel(*MemifSocketFilenameAddDel) (*MemifSocketFilenameAddDelReply, error)
}

/* Messages */

// MemifSocketFilenameAddDel represents VPP binary API message 'memif_socket_filename_add_del':
//
//	"memif_socket_filename_add_del",
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
//	    "u32",
//	    "socket_id"
//	],
//	[
//	    "u8",
//	    "socket_filename",
//	    128
//	],
//	{
//	    "crc": "0x30e3929d"
//	}
//
type MemifSocketFilenameAddDel struct {
	IsAdd          uint8
	SocketID       uint32
	SocketFilename []byte `struc:"[128]byte"`
}

func (*MemifSocketFilenameAddDel) GetMessageName() string {
	return "memif_socket_filename_add_del"
}
func (*MemifSocketFilenameAddDel) GetCrcString() string {
	return "30e3929d"
}
func (*MemifSocketFilenameAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifSocketFilenameAddDelReply represents VPP binary API message 'memif_socket_filename_add_del_reply':
//
//	"memif_socket_filename_add_del_reply",
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
type MemifSocketFilenameAddDelReply struct {
	Retval int32
}

func (*MemifSocketFilenameAddDelReply) GetMessageName() string {
	return "memif_socket_filename_add_del_reply"
}
func (*MemifSocketFilenameAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemifSocketFilenameAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifCreate represents VPP binary API message 'memif_create':
//
//	"memif_create",
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
//	    "role"
//	],
//	[
//	    "u8",
//	    "mode"
//	],
//	[
//	    "u8",
//	    "rx_queues"
//	],
//	[
//	    "u8",
//	    "tx_queues"
//	],
//	[
//	    "u32",
//	    "id"
//	],
//	[
//	    "u32",
//	    "socket_id"
//	],
//	[
//	    "u8",
//	    "secret",
//	    24
//	],
//	[
//	    "u32",
//	    "ring_size"
//	],
//	[
//	    "u16",
//	    "buffer_size"
//	],
//	[
//	    "u8",
//	    "hw_addr",
//	    6
//	],
//	{
//	    "crc": "0x6597cdb2"
//	}
//
type MemifCreate struct {
	Role       uint8
	Mode       uint8
	RxQueues   uint8
	TxQueues   uint8
	ID         uint32
	SocketID   uint32
	Secret     []byte `struc:"[24]byte"`
	RingSize   uint32
	BufferSize uint16
	HwAddr     []byte `struc:"[6]byte"`
}

func (*MemifCreate) GetMessageName() string {
	return "memif_create"
}
func (*MemifCreate) GetCrcString() string {
	return "6597cdb2"
}
func (*MemifCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifCreateReply represents VPP binary API message 'memif_create_reply':
//
//	"memif_create_reply",
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
type MemifCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*MemifCreateReply) GetMessageName() string {
	return "memif_create_reply"
}
func (*MemifCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*MemifCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifDelete represents VPP binary API message 'memif_delete':
//
//	"memif_delete",
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
type MemifDelete struct {
	SwIfIndex uint32
}

func (*MemifDelete) GetMessageName() string {
	return "memif_delete"
}
func (*MemifDelete) GetCrcString() string {
	return "529cb13f"
}
func (*MemifDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifDeleteReply represents VPP binary API message 'memif_delete_reply':
//
//	"memif_delete_reply",
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
type MemifDeleteReply struct {
	Retval int32
}

func (*MemifDeleteReply) GetMessageName() string {
	return "memif_delete_reply"
}
func (*MemifDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemifDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifSocketFilenameDetails represents VPP binary API message 'memif_socket_filename_details':
//
//	"memif_socket_filename_details",
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
//	    "socket_id"
//	],
//	[
//	    "u8",
//	    "socket_filename",
//	    128
//	],
//	{
//	    "crc": "0xe347e32f"
//	}
//
type MemifSocketFilenameDetails struct {
	SocketID       uint32
	SocketFilename []byte `struc:"[128]byte"`
}

func (*MemifSocketFilenameDetails) GetMessageName() string {
	return "memif_socket_filename_details"
}
func (*MemifSocketFilenameDetails) GetCrcString() string {
	return "e347e32f"
}
func (*MemifSocketFilenameDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifSocketFilenameDump represents VPP binary API message 'memif_socket_filename_dump':
//
//	"memif_socket_filename_dump",
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
type MemifSocketFilenameDump struct{}

func (*MemifSocketFilenameDump) GetMessageName() string {
	return "memif_socket_filename_dump"
}
func (*MemifSocketFilenameDump) GetCrcString() string {
	return "51077d14"
}
func (*MemifSocketFilenameDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifDetails represents VPP binary API message 'memif_details':
//
//	"memif_details",
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
//	    "if_name",
//	    64
//	],
//	[
//	    "u8",
//	    "hw_addr",
//	    6
//	],
//	[
//	    "u32",
//	    "id"
//	],
//	[
//	    "u8",
//	    "role"
//	],
//	[
//	    "u8",
//	    "mode"
//	],
//	[
//	    "u32",
//	    "socket_id"
//	],
//	[
//	    "u32",
//	    "ring_size"
//	],
//	[
//	    "u16",
//	    "buffer_size"
//	],
//	[
//	    "u8",
//	    "admin_up_down"
//	],
//	[
//	    "u8",
//	    "link_up_down"
//	],
//	{
//	    "crc": "0x4f5a3397"
//	}
//
type MemifDetails struct {
	SwIfIndex   uint32
	IfName      []byte `struc:"[64]byte"`
	HwAddr      []byte `struc:"[6]byte"`
	ID          uint32
	Role        uint8
	Mode        uint8
	SocketID    uint32
	RingSize    uint32
	BufferSize  uint16
	AdminUpDown uint8
	LinkUpDown  uint8
}

func (*MemifDetails) GetMessageName() string {
	return "memif_details"
}
func (*MemifDetails) GetCrcString() string {
	return "4f5a3397"
}
func (*MemifDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifDump represents VPP binary API message 'memif_dump':
//
//	"memif_dump",
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
type MemifDump struct{}

func (*MemifDump) GetMessageName() string {
	return "memif_dump"
}
func (*MemifDump) GetCrcString() string {
	return "51077d14"
}
func (*MemifDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*MemifSocketFilenameAddDel)(nil), "memif.MemifSocketFilenameAddDel")
	api.RegisterMessage((*MemifSocketFilenameAddDelReply)(nil), "memif.MemifSocketFilenameAddDelReply")
	api.RegisterMessage((*MemifCreate)(nil), "memif.MemifCreate")
	api.RegisterMessage((*MemifCreateReply)(nil), "memif.MemifCreateReply")
	api.RegisterMessage((*MemifDelete)(nil), "memif.MemifDelete")
	api.RegisterMessage((*MemifDeleteReply)(nil), "memif.MemifDeleteReply")
	api.RegisterMessage((*MemifSocketFilenameDetails)(nil), "memif.MemifSocketFilenameDetails")
	api.RegisterMessage((*MemifSocketFilenameDump)(nil), "memif.MemifSocketFilenameDump")
	api.RegisterMessage((*MemifDetails)(nil), "memif.MemifDetails")
	api.RegisterMessage((*MemifDump)(nil), "memif.MemifDump")
}
