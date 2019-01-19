// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/acl.api.json

/*
 Package acl is a generated from VPP binary API module 'acl'.

 It contains following objects:
	 36 messages
	  2 types
	 18 services

*/
package acl

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x8ed22cb9

// Services represents VPP binary API services:
//
//	"services": {
//	    "acl_plugin_get_version": {
//	        "reply": "acl_plugin_get_version_reply"
//	    },
//	    "acl_dump": {
//	        "reply": "acl_details",
//	        "stream": true
//	    },
//	    "acl_interface_add_del": {
//	        "reply": "acl_interface_add_del_reply"
//	    },
//	    "acl_del": {
//	        "reply": "acl_del_reply"
//	    },
//	    "macip_acl_del": {
//	        "reply": "macip_acl_del_reply"
//	    },
//	    "acl_plugin_control_ping": {
//	        "reply": "acl_plugin_control_ping_reply"
//	    },
//	    "macip_acl_interface_get": {
//	        "reply": "macip_acl_interface_get_reply"
//	    },
//	    "acl_interface_etype_whitelist_dump": {
//	        "reply": "acl_interface_etype_whitelist_details",
//	        "stream": true
//	    },
//	    "macip_acl_interface_add_del": {
//	        "reply": "macip_acl_interface_add_del_reply"
//	    },
//	    "acl_add_replace": {
//	        "reply": "acl_add_replace_reply"
//	    },
//	    "acl_plugin_get_conn_table_max_entries": {
//	        "reply": "acl_plugin_get_conn_table_max_entries_reply"
//	    },
//	    "acl_interface_list_dump": {
//	        "reply": "acl_interface_list_details",
//	        "stream": true
//	    },
//	    "acl_interface_set_acl_list": {
//	        "reply": "acl_interface_set_acl_list_reply"
//	    },
//	    "macip_acl_add": {
//	        "reply": "macip_acl_add_reply"
//	    },
//	    "acl_interface_set_etype_whitelist": {
//	        "reply": "acl_interface_set_etype_whitelist_reply"
//	    },
//	    "macip_acl_add_replace": {
//	        "reply": "macip_acl_add_replace_reply"
//	    },
//	    "macip_acl_dump": {
//	        "reply": "macip_acl_details",
//	        "stream": true
//	    },
//	    "macip_acl_interface_list_dump": {
//	        "reply": "macip_acl_interface_list_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpACL(*ACLDump) ([]*ACLDetails, error)
	DumpACLInterfaceEtypeWhitelist(*ACLInterfaceEtypeWhitelistDump) ([]*ACLInterfaceEtypeWhitelistDetails, error)
	DumpACLInterfaceList(*ACLInterfaceListDump) ([]*ACLInterfaceListDetails, error)
	DumpMacipACL(*MacipACLDump) ([]*MacipACLDetails, error)
	DumpMacipACLInterfaceList(*MacipACLInterfaceListDump) ([]*MacipACLInterfaceListDetails, error)
	ACLAddReplace(*ACLAddReplace) (*ACLAddReplaceReply, error)
	ACLDel(*ACLDel) (*ACLDelReply, error)
	ACLInterfaceAddDel(*ACLInterfaceAddDel) (*ACLInterfaceAddDelReply, error)
	ACLInterfaceSetACLList(*ACLInterfaceSetACLList) (*ACLInterfaceSetACLListReply, error)
	ACLInterfaceSetEtypeWhitelist(*ACLInterfaceSetEtypeWhitelist) (*ACLInterfaceSetEtypeWhitelistReply, error)
	ACLPluginControlPing(*ACLPluginControlPing) (*ACLPluginControlPingReply, error)
	ACLPluginGetConnTableMaxEntries(*ACLPluginGetConnTableMaxEntries) (*ACLPluginGetConnTableMaxEntriesReply, error)
	ACLPluginGetVersion(*ACLPluginGetVersion) (*ACLPluginGetVersionReply, error)
	MacipACLAdd(*MacipACLAdd) (*MacipACLAddReply, error)
	MacipACLAddReplace(*MacipACLAddReplace) (*MacipACLAddReplaceReply, error)
	MacipACLDel(*MacipACLDel) (*MacipACLDelReply, error)
	MacipACLInterfaceAddDel(*MacipACLInterfaceAddDel) (*MacipACLInterfaceAddDelReply, error)
	MacipACLInterfaceGet(*MacipACLInterfaceGet) (*MacipACLInterfaceGetReply, error)
}

/* Types */

// ACLRule represents VPP binary API type 'acl_rule':
//
//	"acl_rule",
//	[
//	    "u8",
//	    "is_permit"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "src_ip_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "src_ip_prefix_len"
//	],
//	[
//	    "u8",
//	    "dst_ip_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "dst_ip_prefix_len"
//	],
//	[
//	    "u8",
//	    "proto"
//	],
//	[
//	    "u16",
//	    "srcport_or_icmptype_first"
//	],
//	[
//	    "u16",
//	    "srcport_or_icmptype_last"
//	],
//	[
//	    "u16",
//	    "dstport_or_icmpcode_first"
//	],
//	[
//	    "u16",
//	    "dstport_or_icmpcode_last"
//	],
//	[
//	    "u8",
//	    "tcp_flags_mask"
//	],
//	[
//	    "u8",
//	    "tcp_flags_value"
//	],
//	{
//	    "crc": "0x6f99bf4d"
//	}
//
type ACLRule struct {
	IsPermit               uint8
	IsIPv6                 uint8
	SrcIPAddr              []byte `struc:"[16]byte"`
	SrcIPPrefixLen         uint8
	DstIPAddr              []byte `struc:"[16]byte"`
	DstIPPrefixLen         uint8
	Proto                  uint8
	SrcportOrIcmptypeFirst uint16
	SrcportOrIcmptypeLast  uint16
	DstportOrIcmpcodeFirst uint16
	DstportOrIcmpcodeLast  uint16
	TCPFlagsMask           uint8
	TCPFlagsValue          uint8
}

func (*ACLRule) GetTypeName() string {
	return "acl_rule"
}
func (*ACLRule) GetCrcString() string {
	return "6f99bf4d"
}

// MacipACLRule represents VPP binary API type 'macip_acl_rule':
//
//	"macip_acl_rule",
//	[
//	    "u8",
//	    "is_permit"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "src_mac",
//	    6
//	],
//	[
//	    "u8",
//	    "src_mac_mask",
//	    6
//	],
//	[
//	    "u8",
//	    "src_ip_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "src_ip_prefix_len"
//	],
//	{
//	    "crc": "0x70589f1e"
//	}
//
type MacipACLRule struct {
	IsPermit       uint8
	IsIPv6         uint8
	SrcMac         []byte `struc:"[6]byte"`
	SrcMacMask     []byte `struc:"[6]byte"`
	SrcIPAddr      []byte `struc:"[16]byte"`
	SrcIPPrefixLen uint8
}

func (*MacipACLRule) GetTypeName() string {
	return "macip_acl_rule"
}
func (*MacipACLRule) GetCrcString() string {
	return "70589f1e"
}

/* Messages */

// ACLPluginGetVersion represents VPP binary API message 'acl_plugin_get_version':
//
//	"acl_plugin_get_version",
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
type ACLPluginGetVersion struct{}

func (*ACLPluginGetVersion) GetMessageName() string {
	return "acl_plugin_get_version"
}
func (*ACLPluginGetVersion) GetCrcString() string {
	return "51077d14"
}
func (*ACLPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLPluginGetVersionReply represents VPP binary API message 'acl_plugin_get_version_reply':
//
//	"acl_plugin_get_version_reply",
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
//	    "major"
//	],
//	[
//	    "u32",
//	    "minor"
//	],
//	{
//	    "crc": "0x9b32cf86"
//	}
//
type ACLPluginGetVersionReply struct {
	Major uint32
	Minor uint32
}

func (*ACLPluginGetVersionReply) GetMessageName() string {
	return "acl_plugin_get_version_reply"
}
func (*ACLPluginGetVersionReply) GetCrcString() string {
	return "9b32cf86"
}
func (*ACLPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLPluginControlPing represents VPP binary API message 'acl_plugin_control_ping':
//
//	"acl_plugin_control_ping",
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
type ACLPluginControlPing struct{}

func (*ACLPluginControlPing) GetMessageName() string {
	return "acl_plugin_control_ping"
}
func (*ACLPluginControlPing) GetCrcString() string {
	return "51077d14"
}
func (*ACLPluginControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLPluginControlPingReply represents VPP binary API message 'acl_plugin_control_ping_reply':
//
//	"acl_plugin_control_ping_reply",
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
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "vpe_pid"
//	],
//	{
//	    "crc": "0xf6b0b8ca"
//	}
//
type ACLPluginControlPingReply struct {
	Retval      int32
	ClientIndex uint32
	VpePID      uint32
}

func (*ACLPluginControlPingReply) GetMessageName() string {
	return "acl_plugin_control_ping_reply"
}
func (*ACLPluginControlPingReply) GetCrcString() string {
	return "f6b0b8ca"
}
func (*ACLPluginControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLPluginGetConnTableMaxEntries represents VPP binary API message 'acl_plugin_get_conn_table_max_entries':
//
//	"acl_plugin_get_conn_table_max_entries",
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
type ACLPluginGetConnTableMaxEntries struct{}

func (*ACLPluginGetConnTableMaxEntries) GetMessageName() string {
	return "acl_plugin_get_conn_table_max_entries"
}
func (*ACLPluginGetConnTableMaxEntries) GetCrcString() string {
	return "51077d14"
}
func (*ACLPluginGetConnTableMaxEntries) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLPluginGetConnTableMaxEntriesReply represents VPP binary API message 'acl_plugin_get_conn_table_max_entries_reply':
//
//	"acl_plugin_get_conn_table_max_entries_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u64",
//	    "conn_table_max_entries"
//	],
//	{
//	    "crc": "0x7a096d3d"
//	}
//
type ACLPluginGetConnTableMaxEntriesReply struct {
	ConnTableMaxEntries uint64
}

func (*ACLPluginGetConnTableMaxEntriesReply) GetMessageName() string {
	return "acl_plugin_get_conn_table_max_entries_reply"
}
func (*ACLPluginGetConnTableMaxEntriesReply) GetCrcString() string {
	return "7a096d3d"
}
func (*ACLPluginGetConnTableMaxEntriesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLAddReplace represents VPP binary API message 'acl_add_replace':
//
//	"acl_add_replace",
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
//	    "acl_index"
//	],
//	[
//	    "u8",
//	    "tag",
//	    64
//	],
//	[
//	    "u32",
//	    "count"
//	],
//	[
//	    "vl_api_acl_rule_t",
//	    "r",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xe839997e"
//	}
//
type ACLAddReplace struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []ACLRule
}

func (*ACLAddReplace) GetMessageName() string {
	return "acl_add_replace"
}
func (*ACLAddReplace) GetCrcString() string {
	return "e839997e"
}
func (*ACLAddReplace) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLAddReplaceReply represents VPP binary API message 'acl_add_replace_reply':
//
//	"acl_add_replace_reply",
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
//	    "acl_index"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xac407b0c"
//	}
//
type ACLAddReplaceReply struct {
	ACLIndex uint32
	Retval   int32
}

func (*ACLAddReplaceReply) GetMessageName() string {
	return "acl_add_replace_reply"
}
func (*ACLAddReplaceReply) GetCrcString() string {
	return "ac407b0c"
}
func (*ACLAddReplaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLDel represents VPP binary API message 'acl_del':
//
//	"acl_del",
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
//	    "acl_index"
//	],
//	{
//	    "crc": "0xef34fea4"
//	}
//
type ACLDel struct {
	ACLIndex uint32
}

func (*ACLDel) GetMessageName() string {
	return "acl_del"
}
func (*ACLDel) GetCrcString() string {
	return "ef34fea4"
}
func (*ACLDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLDelReply represents VPP binary API message 'acl_del_reply':
//
//	"acl_del_reply",
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
type ACLDelReply struct {
	Retval int32
}

func (*ACLDelReply) GetMessageName() string {
	return "acl_del_reply"
}
func (*ACLDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ACLDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLInterfaceAddDel represents VPP binary API message 'acl_interface_add_del':
//
//	"acl_interface_add_del",
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
//	    "is_input"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u32",
//	    "acl_index"
//	],
//	{
//	    "crc": "0x0b2aedd1"
//	}
//
type ACLInterfaceAddDel struct {
	IsAdd     uint8
	IsInput   uint8
	SwIfIndex uint32
	ACLIndex  uint32
}

func (*ACLInterfaceAddDel) GetMessageName() string {
	return "acl_interface_add_del"
}
func (*ACLInterfaceAddDel) GetCrcString() string {
	return "0b2aedd1"
}
func (*ACLInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLInterfaceAddDelReply represents VPP binary API message 'acl_interface_add_del_reply':
//
//	"acl_interface_add_del_reply",
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
type ACLInterfaceAddDelReply struct {
	Retval int32
}

func (*ACLInterfaceAddDelReply) GetMessageName() string {
	return "acl_interface_add_del_reply"
}
func (*ACLInterfaceAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ACLInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLInterfaceSetACLList represents VPP binary API message 'acl_interface_set_acl_list':
//
//	"acl_interface_set_acl_list",
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
//	    "count"
//	],
//	[
//	    "u8",
//	    "n_input"
//	],
//	[
//	    "u32",
//	    "acls",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0x8baece38"
//	}
//
type ACLInterfaceSetACLList struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Acls"`
	NInput    uint8
	Acls      []uint32
}

func (*ACLInterfaceSetACLList) GetMessageName() string {
	return "acl_interface_set_acl_list"
}
func (*ACLInterfaceSetACLList) GetCrcString() string {
	return "8baece38"
}
func (*ACLInterfaceSetACLList) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLInterfaceSetACLListReply represents VPP binary API message 'acl_interface_set_acl_list_reply':
//
//	"acl_interface_set_acl_list_reply",
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
type ACLInterfaceSetACLListReply struct {
	Retval int32
}

func (*ACLInterfaceSetACLListReply) GetMessageName() string {
	return "acl_interface_set_acl_list_reply"
}
func (*ACLInterfaceSetACLListReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ACLInterfaceSetACLListReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLDump represents VPP binary API message 'acl_dump':
//
//	"acl_dump",
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
//	    "acl_index"
//	],
//	{
//	    "crc": "0xef34fea4"
//	}
//
type ACLDump struct {
	ACLIndex uint32
}

func (*ACLDump) GetMessageName() string {
	return "acl_dump"
}
func (*ACLDump) GetCrcString() string {
	return "ef34fea4"
}
func (*ACLDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLDetails represents VPP binary API message 'acl_details':
//
//	"acl_details",
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
//	    "acl_index"
//	],
//	[
//	    "u8",
//	    "tag",
//	    64
//	],
//	[
//	    "u32",
//	    "count"
//	],
//	[
//	    "vl_api_acl_rule_t",
//	    "r",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0x5bd895be"
//	}
//
type ACLDetails struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []ACLRule
}

func (*ACLDetails) GetMessageName() string {
	return "acl_details"
}
func (*ACLDetails) GetCrcString() string {
	return "5bd895be"
}
func (*ACLDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLInterfaceListDump represents VPP binary API message 'acl_interface_list_dump':
//
//	"acl_interface_list_dump",
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
type ACLInterfaceListDump struct {
	SwIfIndex uint32
}

func (*ACLInterfaceListDump) GetMessageName() string {
	return "acl_interface_list_dump"
}
func (*ACLInterfaceListDump) GetCrcString() string {
	return "529cb13f"
}
func (*ACLInterfaceListDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLInterfaceListDetails represents VPP binary API message 'acl_interface_list_details':
//
//	"acl_interface_list_details",
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
//	    "count"
//	],
//	[
//	    "u8",
//	    "n_input"
//	],
//	[
//	    "u32",
//	    "acls",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xd5e80809"
//	}
//
type ACLInterfaceListDetails struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Acls"`
	NInput    uint8
	Acls      []uint32
}

func (*ACLInterfaceListDetails) GetMessageName() string {
	return "acl_interface_list_details"
}
func (*ACLInterfaceListDetails) GetCrcString() string {
	return "d5e80809"
}
func (*ACLInterfaceListDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLAdd represents VPP binary API message 'macip_acl_add':
//
//	"macip_acl_add",
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
//	    "tag",
//	    64
//	],
//	[
//	    "u32",
//	    "count"
//	],
//	[
//	    "vl_api_macip_acl_rule_t",
//	    "r",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xb3d3d65a"
//	}
//
type MacipACLAdd struct {
	Tag   []byte `struc:"[64]byte"`
	Count uint32 `struc:"sizeof=R"`
	R     []MacipACLRule
}

func (*MacipACLAdd) GetMessageName() string {
	return "macip_acl_add"
}
func (*MacipACLAdd) GetCrcString() string {
	return "b3d3d65a"
}
func (*MacipACLAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLAddReply represents VPP binary API message 'macip_acl_add_reply':
//
//	"macip_acl_add_reply",
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
//	    "acl_index"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xac407b0c"
//	}
//
type MacipACLAddReply struct {
	ACLIndex uint32
	Retval   int32
}

func (*MacipACLAddReply) GetMessageName() string {
	return "macip_acl_add_reply"
}
func (*MacipACLAddReply) GetCrcString() string {
	return "ac407b0c"
}
func (*MacipACLAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLAddReplace represents VPP binary API message 'macip_acl_add_replace':
//
//	"macip_acl_add_replace",
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
//	    "acl_index"
//	],
//	[
//	    "u8",
//	    "tag",
//	    64
//	],
//	[
//	    "u32",
//	    "count"
//	],
//	[
//	    "vl_api_macip_acl_rule_t",
//	    "r",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xa0e8c01b"
//	}
//
type MacipACLAddReplace struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []MacipACLRule
}

func (*MacipACLAddReplace) GetMessageName() string {
	return "macip_acl_add_replace"
}
func (*MacipACLAddReplace) GetCrcString() string {
	return "a0e8c01b"
}
func (*MacipACLAddReplace) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLAddReplaceReply represents VPP binary API message 'macip_acl_add_replace_reply':
//
//	"macip_acl_add_replace_reply",
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
//	    "acl_index"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xac407b0c"
//	}
//
type MacipACLAddReplaceReply struct {
	ACLIndex uint32
	Retval   int32
}

func (*MacipACLAddReplaceReply) GetMessageName() string {
	return "macip_acl_add_replace_reply"
}
func (*MacipACLAddReplaceReply) GetCrcString() string {
	return "ac407b0c"
}
func (*MacipACLAddReplaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLDel represents VPP binary API message 'macip_acl_del':
//
//	"macip_acl_del",
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
//	    "acl_index"
//	],
//	{
//	    "crc": "0xef34fea4"
//	}
//
type MacipACLDel struct {
	ACLIndex uint32
}

func (*MacipACLDel) GetMessageName() string {
	return "macip_acl_del"
}
func (*MacipACLDel) GetCrcString() string {
	return "ef34fea4"
}
func (*MacipACLDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLDelReply represents VPP binary API message 'macip_acl_del_reply':
//
//	"macip_acl_del_reply",
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
type MacipACLDelReply struct {
	Retval int32
}

func (*MacipACLDelReply) GetMessageName() string {
	return "macip_acl_del_reply"
}
func (*MacipACLDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MacipACLDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLInterfaceAddDel represents VPP binary API message 'macip_acl_interface_add_del':
//
//	"macip_acl_interface_add_del",
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
//	    "sw_if_index"
//	],
//	[
//	    "u32",
//	    "acl_index"
//	],
//	{
//	    "crc": "0x6a6be97c"
//	}
//
type MacipACLInterfaceAddDel struct {
	IsAdd     uint8
	SwIfIndex uint32
	ACLIndex  uint32
}

func (*MacipACLInterfaceAddDel) GetMessageName() string {
	return "macip_acl_interface_add_del"
}
func (*MacipACLInterfaceAddDel) GetCrcString() string {
	return "6a6be97c"
}
func (*MacipACLInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLInterfaceAddDelReply represents VPP binary API message 'macip_acl_interface_add_del_reply':
//
//	"macip_acl_interface_add_del_reply",
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
type MacipACLInterfaceAddDelReply struct {
	Retval int32
}

func (*MacipACLInterfaceAddDelReply) GetMessageName() string {
	return "macip_acl_interface_add_del_reply"
}
func (*MacipACLInterfaceAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MacipACLInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLDump represents VPP binary API message 'macip_acl_dump':
//
//	"macip_acl_dump",
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
//	    "acl_index"
//	],
//	{
//	    "crc": "0xef34fea4"
//	}
//
type MacipACLDump struct {
	ACLIndex uint32
}

func (*MacipACLDump) GetMessageName() string {
	return "macip_acl_dump"
}
func (*MacipACLDump) GetCrcString() string {
	return "ef34fea4"
}
func (*MacipACLDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLDetails represents VPP binary API message 'macip_acl_details':
//
//	"macip_acl_details",
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
//	    "acl_index"
//	],
//	[
//	    "u8",
//	    "tag",
//	    64
//	],
//	[
//	    "u32",
//	    "count"
//	],
//	[
//	    "vl_api_macip_acl_rule_t",
//	    "r",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xdd2b55ba"
//	}
//
type MacipACLDetails struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []MacipACLRule
}

func (*MacipACLDetails) GetMessageName() string {
	return "macip_acl_details"
}
func (*MacipACLDetails) GetCrcString() string {
	return "dd2b55ba"
}
func (*MacipACLDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLInterfaceGet represents VPP binary API message 'macip_acl_interface_get':
//
//	"macip_acl_interface_get",
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
type MacipACLInterfaceGet struct{}

func (*MacipACLInterfaceGet) GetMessageName() string {
	return "macip_acl_interface_get"
}
func (*MacipACLInterfaceGet) GetCrcString() string {
	return "51077d14"
}
func (*MacipACLInterfaceGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLInterfaceGetReply represents VPP binary API message 'macip_acl_interface_get_reply':
//
//	"macip_acl_interface_get_reply",
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
//	    "count"
//	],
//	[
//	    "u32",
//	    "acls",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xaccf9b05"
//	}
//
type MacipACLInterfaceGetReply struct {
	Count uint32 `struc:"sizeof=Acls"`
	Acls  []uint32
}

func (*MacipACLInterfaceGetReply) GetMessageName() string {
	return "macip_acl_interface_get_reply"
}
func (*MacipACLInterfaceGetReply) GetCrcString() string {
	return "accf9b05"
}
func (*MacipACLInterfaceGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MacipACLInterfaceListDump represents VPP binary API message 'macip_acl_interface_list_dump':
//
//	"macip_acl_interface_list_dump",
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
type MacipACLInterfaceListDump struct {
	SwIfIndex uint32
}

func (*MacipACLInterfaceListDump) GetMessageName() string {
	return "macip_acl_interface_list_dump"
}
func (*MacipACLInterfaceListDump) GetCrcString() string {
	return "529cb13f"
}
func (*MacipACLInterfaceListDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MacipACLInterfaceListDetails represents VPP binary API message 'macip_acl_interface_list_details':
//
//	"macip_acl_interface_list_details",
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
//	    "count"
//	],
//	[
//	    "u32",
//	    "acls",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0x29783fa0"
//	}
//
type MacipACLInterfaceListDetails struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Acls"`
	Acls      []uint32
}

func (*MacipACLInterfaceListDetails) GetMessageName() string {
	return "macip_acl_interface_list_details"
}
func (*MacipACLInterfaceListDetails) GetCrcString() string {
	return "29783fa0"
}
func (*MacipACLInterfaceListDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLInterfaceSetEtypeWhitelist represents VPP binary API message 'acl_interface_set_etype_whitelist':
//
//	"acl_interface_set_etype_whitelist",
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
//	    "count"
//	],
//	[
//	    "u8",
//	    "n_input"
//	],
//	[
//	    "u16",
//	    "whitelist",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xf515efc5"
//	}
//
type ACLInterfaceSetEtypeWhitelist struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Whitelist"`
	NInput    uint8
	Whitelist []uint16
}

func (*ACLInterfaceSetEtypeWhitelist) GetMessageName() string {
	return "acl_interface_set_etype_whitelist"
}
func (*ACLInterfaceSetEtypeWhitelist) GetCrcString() string {
	return "f515efc5"
}
func (*ACLInterfaceSetEtypeWhitelist) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLInterfaceSetEtypeWhitelistReply represents VPP binary API message 'acl_interface_set_etype_whitelist_reply':
//
//	"acl_interface_set_etype_whitelist_reply",
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
type ACLInterfaceSetEtypeWhitelistReply struct {
	Retval int32
}

func (*ACLInterfaceSetEtypeWhitelistReply) GetMessageName() string {
	return "acl_interface_set_etype_whitelist_reply"
}
func (*ACLInterfaceSetEtypeWhitelistReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ACLInterfaceSetEtypeWhitelistReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ACLInterfaceEtypeWhitelistDump represents VPP binary API message 'acl_interface_etype_whitelist_dump':
//
//	"acl_interface_etype_whitelist_dump",
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
type ACLInterfaceEtypeWhitelistDump struct {
	SwIfIndex uint32
}

func (*ACLInterfaceEtypeWhitelistDump) GetMessageName() string {
	return "acl_interface_etype_whitelist_dump"
}
func (*ACLInterfaceEtypeWhitelistDump) GetCrcString() string {
	return "529cb13f"
}
func (*ACLInterfaceEtypeWhitelistDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ACLInterfaceEtypeWhitelistDetails represents VPP binary API message 'acl_interface_etype_whitelist_details':
//
//	"acl_interface_etype_whitelist_details",
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
//	    "count"
//	],
//	[
//	    "u8",
//	    "n_input"
//	],
//	[
//	    "u16",
//	    "whitelist",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0x6a5d4e81"
//	}
//
type ACLInterfaceEtypeWhitelistDetails struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Whitelist"`
	NInput    uint8
	Whitelist []uint16
}

func (*ACLInterfaceEtypeWhitelistDetails) GetMessageName() string {
	return "acl_interface_etype_whitelist_details"
}
func (*ACLInterfaceEtypeWhitelistDetails) GetCrcString() string {
	return "6a5d4e81"
}
func (*ACLInterfaceEtypeWhitelistDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*ACLPluginGetVersion)(nil), "acl.ACLPluginGetVersion")
	api.RegisterMessage((*ACLPluginGetVersionReply)(nil), "acl.ACLPluginGetVersionReply")
	api.RegisterMessage((*ACLPluginControlPing)(nil), "acl.ACLPluginControlPing")
	api.RegisterMessage((*ACLPluginControlPingReply)(nil), "acl.ACLPluginControlPingReply")
	api.RegisterMessage((*ACLPluginGetConnTableMaxEntries)(nil), "acl.ACLPluginGetConnTableMaxEntries")
	api.RegisterMessage((*ACLPluginGetConnTableMaxEntriesReply)(nil), "acl.ACLPluginGetConnTableMaxEntriesReply")
	api.RegisterMessage((*ACLAddReplace)(nil), "acl.ACLAddReplace")
	api.RegisterMessage((*ACLAddReplaceReply)(nil), "acl.ACLAddReplaceReply")
	api.RegisterMessage((*ACLDel)(nil), "acl.ACLDel")
	api.RegisterMessage((*ACLDelReply)(nil), "acl.ACLDelReply")
	api.RegisterMessage((*ACLInterfaceAddDel)(nil), "acl.ACLInterfaceAddDel")
	api.RegisterMessage((*ACLInterfaceAddDelReply)(nil), "acl.ACLInterfaceAddDelReply")
	api.RegisterMessage((*ACLInterfaceSetACLList)(nil), "acl.ACLInterfaceSetACLList")
	api.RegisterMessage((*ACLInterfaceSetACLListReply)(nil), "acl.ACLInterfaceSetACLListReply")
	api.RegisterMessage((*ACLDump)(nil), "acl.ACLDump")
	api.RegisterMessage((*ACLDetails)(nil), "acl.ACLDetails")
	api.RegisterMessage((*ACLInterfaceListDump)(nil), "acl.ACLInterfaceListDump")
	api.RegisterMessage((*ACLInterfaceListDetails)(nil), "acl.ACLInterfaceListDetails")
	api.RegisterMessage((*MacipACLAdd)(nil), "acl.MacipACLAdd")
	api.RegisterMessage((*MacipACLAddReply)(nil), "acl.MacipACLAddReply")
	api.RegisterMessage((*MacipACLAddReplace)(nil), "acl.MacipACLAddReplace")
	api.RegisterMessage((*MacipACLAddReplaceReply)(nil), "acl.MacipACLAddReplaceReply")
	api.RegisterMessage((*MacipACLDel)(nil), "acl.MacipACLDel")
	api.RegisterMessage((*MacipACLDelReply)(nil), "acl.MacipACLDelReply")
	api.RegisterMessage((*MacipACLInterfaceAddDel)(nil), "acl.MacipACLInterfaceAddDel")
	api.RegisterMessage((*MacipACLInterfaceAddDelReply)(nil), "acl.MacipACLInterfaceAddDelReply")
	api.RegisterMessage((*MacipACLDump)(nil), "acl.MacipACLDump")
	api.RegisterMessage((*MacipACLDetails)(nil), "acl.MacipACLDetails")
	api.RegisterMessage((*MacipACLInterfaceGet)(nil), "acl.MacipACLInterfaceGet")
	api.RegisterMessage((*MacipACLInterfaceGetReply)(nil), "acl.MacipACLInterfaceGetReply")
	api.RegisterMessage((*MacipACLInterfaceListDump)(nil), "acl.MacipACLInterfaceListDump")
	api.RegisterMessage((*MacipACLInterfaceListDetails)(nil), "acl.MacipACLInterfaceListDetails")
	api.RegisterMessage((*ACLInterfaceSetEtypeWhitelist)(nil), "acl.ACLInterfaceSetEtypeWhitelist")
	api.RegisterMessage((*ACLInterfaceSetEtypeWhitelistReply)(nil), "acl.ACLInterfaceSetEtypeWhitelistReply")
	api.RegisterMessage((*ACLInterfaceEtypeWhitelistDump)(nil), "acl.ACLInterfaceEtypeWhitelistDump")
	api.RegisterMessage((*ACLInterfaceEtypeWhitelistDetails)(nil), "acl.ACLInterfaceEtypeWhitelistDetails")
}
