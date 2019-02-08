// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/ioam_export.api.json

/*
 Package ioam_export is a generated from VPP binary API module 'ioam_export'.

 It contains following objects:
	  1 service
	  2 messages
*/
package ioam_export

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x22132a2c

// Services represents VPP binary API services:
//
//	"services": {
//	    "ioam_export_ip6_enable_disable": {
//	        "reply": "ioam_export_ip6_enable_disable_reply"
//	    }
//	},
//
type Services interface {
	IoamExportIP6EnableDisable(*IoamExportIP6EnableDisable) (*IoamExportIP6EnableDisableReply, error)
}

/* Messages */

// IoamExportIP6EnableDisable represents VPP binary API message 'ioam_export_ip6_enable_disable':
//
//	"ioam_export_ip6_enable_disable",
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
//	    "is_disable"
//	],
//	[
//	    "u8",
//	    "collector_address",
//	    4
//	],
//	[
//	    "u8",
//	    "src_address",
//	    4
//	],
//	{
//	    "crc": "0x148b82a4"
//	}
//
type IoamExportIP6EnableDisable struct {
	IsDisable        uint8
	CollectorAddress []byte `struc:"[4]byte"`
	SrcAddress       []byte `struc:"[4]byte"`
}

func (*IoamExportIP6EnableDisable) GetMessageName() string {
	return "ioam_export_ip6_enable_disable"
}
func (*IoamExportIP6EnableDisable) GetCrcString() string {
	return "148b82a4"
}
func (*IoamExportIP6EnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IoamExportIP6EnableDisableReply represents VPP binary API message 'ioam_export_ip6_enable_disable_reply':
//
//	"ioam_export_ip6_enable_disable_reply",
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
type IoamExportIP6EnableDisableReply struct {
	Retval int32
}

func (*IoamExportIP6EnableDisableReply) GetMessageName() string {
	return "ioam_export_ip6_enable_disable_reply"
}
func (*IoamExportIP6EnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IoamExportIP6EnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*IoamExportIP6EnableDisable)(nil), "ioam_export.IoamExportIP6EnableDisable")
	api.RegisterMessage((*IoamExportIP6EnableDisableReply)(nil), "ioam_export.IoamExportIP6EnableDisableReply")
}
