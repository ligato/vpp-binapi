// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/dhcp6_ia_na_client_cp.api.json

/*
 Package dhcp6_ia_na_client_cp is a generated from VPP binary API module 'dhcp6_ia_na_client_cp'.

 It contains following objects:
	  2 messages
	  1 service

*/
package dhcp6_ia_na_client_cp

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x3fc84d7c

// Services represents VPP binary API services:
//
//	"services": {
//	    "dhcp6_client_enable_disable": {
//	        "reply": "dhcp6_client_enable_disable_reply"
//	    }
//	},
//
type Services interface {
	DHCP6ClientEnableDisable(*DHCP6ClientEnableDisable) (*DHCP6ClientEnableDisableReply, error)
}

/* Messages */

// DHCP6ClientEnableDisable represents VPP binary API message 'dhcp6_client_enable_disable':
//
//	"dhcp6_client_enable_disable",
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
//	    "enable"
//	],
//	{
//	    "crc": "0xa36fadc0"
//	}
//
type DHCP6ClientEnableDisable struct {
	SwIfIndex uint32
	Enable    uint8
}

func (*DHCP6ClientEnableDisable) GetMessageName() string {
	return "dhcp6_client_enable_disable"
}
func (*DHCP6ClientEnableDisable) GetCrcString() string {
	return "a36fadc0"
}
func (*DHCP6ClientEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6ClientEnableDisableReply represents VPP binary API message 'dhcp6_client_enable_disable_reply':
//
//	"dhcp6_client_enable_disable_reply",
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
type DHCP6ClientEnableDisableReply struct {
	Retval int32
}

func (*DHCP6ClientEnableDisableReply) GetMessageName() string {
	return "dhcp6_client_enable_disable_reply"
}
func (*DHCP6ClientEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6ClientEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*DHCP6ClientEnableDisable)(nil), "dhcp6_ia_na_client_cp.DHCP6ClientEnableDisable")
	api.RegisterMessage((*DHCP6ClientEnableDisableReply)(nil), "dhcp6_ia_na_client_cp.DHCP6ClientEnableDisableReply")
}
