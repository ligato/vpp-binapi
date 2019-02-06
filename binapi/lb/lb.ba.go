// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/lb.api.json

/*
 Package lb is a generated from VPP binary API module 'lb'.

 It contains following objects:
	  8 messages
	  4 services

*/
package lb

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
//	    "lb_conf": {
//	        "reply": "lb_conf_reply"
//	    },
//	    "lb_add_del_vip": {
//	        "reply": "lb_add_del_vip_reply"
//	    },
//	    "lb_add_del_as": {
//	        "reply": "lb_add_del_as_reply"
//	    },
//	    "lb_flush_vip": {
//	        "reply": "lb_flush_vip_reply"
//	    }
//	},
//
type Services interface {
	LbAddDelAs(*LbAddDelAs) (*LbAddDelAsReply, error)
	LbAddDelVip(*LbAddDelVip) (*LbAddDelVipReply, error)
	LbConf(*LbConf) (*LbConfReply, error)
	LbFlushVip(*LbFlushVip) (*LbFlushVipReply, error)
}

/* Messages */

// LbConf represents VPP binary API message 'lb_conf':
//
//	"lb_conf",
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
//	    "ip4_src_address"
//	],
//	[
//	    "u8",
//	    "ip6_src_address",
//	    16
//	],
//	[
//	    "u32",
//	    "sticky_buckets_per_core"
//	],
//	[
//	    "u32",
//	    "flow_timeout"
//	],
//	{
//	    "crc": "0x4ae4f864"
//	}
//
type LbConf struct {
	IP4SrcAddress        uint32
	IP6SrcAddress        []byte `struc:"[16]byte"`
	StickyBucketsPerCore uint32
	FlowTimeout          uint32
}

func (*LbConf) GetMessageName() string {
	return "lb_conf"
}
func (*LbConf) GetCrcString() string {
	return "4ae4f864"
}
func (*LbConf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// LbConfReply represents VPP binary API message 'lb_conf_reply':
//
//	"lb_conf_reply",
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
type LbConfReply struct {
	Retval int32
}

func (*LbConfReply) GetMessageName() string {
	return "lb_conf_reply"
}
func (*LbConfReply) GetCrcString() string {
	return "e8d4e804"
}
func (*LbConfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// LbAddDelVip represents VPP binary API message 'lb_add_del_vip':
//
//	"lb_add_del_vip",
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
//	    "ip_prefix",
//	    16
//	],
//	[
//	    "u8",
//	    "prefix_length"
//	],
//	[
//	    "u8",
//	    "protocol"
//	],
//	[
//	    "u16",
//	    "port"
//	],
//	[
//	    "u8",
//	    "encap"
//	],
//	[
//	    "u8",
//	    "dscp"
//	],
//	[
//	    "u8",
//	    "type"
//	],
//	[
//	    "u16",
//	    "target_port"
//	],
//	[
//	    "u16",
//	    "node_port"
//	],
//	[
//	    "u32",
//	    "new_flows_table_length"
//	],
//	[
//	    "u8",
//	    "is_del"
//	],
//	{
//	    "crc": "0xd67d5a49"
//	}
//
type LbAddDelVip struct {
	IPPrefix            []byte `struc:"[16]byte"`
	PrefixLength        uint8
	Protocol            uint8
	Port                uint16
	Encap               uint8
	Dscp                uint8
	Type                uint8
	TargetPort          uint16
	NodePort            uint16
	NewFlowsTableLength uint32
	IsDel               uint8
}

func (*LbAddDelVip) GetMessageName() string {
	return "lb_add_del_vip"
}
func (*LbAddDelVip) GetCrcString() string {
	return "d67d5a49"
}
func (*LbAddDelVip) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// LbAddDelVipReply represents VPP binary API message 'lb_add_del_vip_reply':
//
//	"lb_add_del_vip_reply",
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
type LbAddDelVipReply struct {
	Retval int32
}

func (*LbAddDelVipReply) GetMessageName() string {
	return "lb_add_del_vip_reply"
}
func (*LbAddDelVipReply) GetCrcString() string {
	return "e8d4e804"
}
func (*LbAddDelVipReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// LbAddDelAs represents VPP binary API message 'lb_add_del_as':
//
//	"lb_add_del_as",
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
//	    "vip_ip_prefix",
//	    16
//	],
//	[
//	    "u8",
//	    "vip_prefix_length"
//	],
//	[
//	    "u8",
//	    "protocol"
//	],
//	[
//	    "u16",
//	    "port"
//	],
//	[
//	    "u8",
//	    "as_address",
//	    16
//	],
//	[
//	    "u8",
//	    "is_del"
//	],
//	[
//	    "u8",
//	    "is_flush"
//	],
//	{
//	    "crc": "0xb2252622"
//	}
//
type LbAddDelAs struct {
	VipIPPrefix     []byte `struc:"[16]byte"`
	VipPrefixLength uint8
	Protocol        uint8
	Port            uint16
	AsAddress       []byte `struc:"[16]byte"`
	IsDel           uint8
	IsFlush         uint8
}

func (*LbAddDelAs) GetMessageName() string {
	return "lb_add_del_as"
}
func (*LbAddDelAs) GetCrcString() string {
	return "b2252622"
}
func (*LbAddDelAs) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// LbAddDelAsReply represents VPP binary API message 'lb_add_del_as_reply':
//
//	"lb_add_del_as_reply",
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
type LbAddDelAsReply struct {
	Retval int32
}

func (*LbAddDelAsReply) GetMessageName() string {
	return "lb_add_del_as_reply"
}
func (*LbAddDelAsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*LbAddDelAsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// LbFlushVip represents VPP binary API message 'lb_flush_vip':
//
//	"lb_flush_vip",
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
//	    "ip_prefix",
//	    16
//	],
//	[
//	    "u8",
//	    "prefix_length"
//	],
//	[
//	    "u8",
//	    "protocol"
//	],
//	[
//	    "u16",
//	    "port"
//	],
//	{
//	    "crc": "0x21d6df07"
//	}
//
type LbFlushVip struct {
	IPPrefix     []byte `struc:"[16]byte"`
	PrefixLength uint8
	Protocol     uint8
	Port         uint16
}

func (*LbFlushVip) GetMessageName() string {
	return "lb_flush_vip"
}
func (*LbFlushVip) GetCrcString() string {
	return "21d6df07"
}
func (*LbFlushVip) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// LbFlushVipReply represents VPP binary API message 'lb_flush_vip_reply':
//
//	"lb_flush_vip_reply",
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
type LbFlushVipReply struct {
	Retval int32
}

func (*LbFlushVipReply) GetMessageName() string {
	return "lb_flush_vip_reply"
}
func (*LbFlushVipReply) GetCrcString() string {
	return "e8d4e804"
}
func (*LbFlushVipReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*LbConf)(nil), "lb.LbConf")
	api.RegisterMessage((*LbConfReply)(nil), "lb.LbConfReply")
	api.RegisterMessage((*LbAddDelVip)(nil), "lb.LbAddDelVip")
	api.RegisterMessage((*LbAddDelVipReply)(nil), "lb.LbAddDelVipReply")
	api.RegisterMessage((*LbAddDelAs)(nil), "lb.LbAddDelAs")
	api.RegisterMessage((*LbAddDelAsReply)(nil), "lb.LbAddDelAsReply")
	api.RegisterMessage((*LbFlushVip)(nil), "lb.LbFlushVip")
	api.RegisterMessage((*LbFlushVipReply)(nil), "lb.LbFlushVipReply")
}
