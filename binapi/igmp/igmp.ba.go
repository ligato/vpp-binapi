// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/igmp.api.json

/*
 Package igmp is a generated from VPP binary API module 'igmp'.

 It contains following objects:
	 19 messages
	  7 types
	  2 aliases
	  3 enums
	  1 union
	  9 services

*/
package igmp

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x5c65be9a

// Services represents VPP binary API services:
//
//	"services": {
//	    "want_igmp_events": {
//	        "reply": "want_igmp_events_reply",
//	        "events": [
//	            "igmp_event"
//	        ]
//	    },
//	    "igmp_proxy_device_add_del": {
//	        "reply": "igmp_proxy_device_add_del_reply"
//	    },
//	    "igmp_dump": {
//	        "reply": "igmp_details",
//	        "stream": true
//	    },
//	    "igmp_enable_disable": {
//	        "reply": "igmp_enable_disable_reply"
//	    },
//	    "igmp_proxy_device_add_del_interface": {
//	        "reply": "igmp_proxy_device_add_del_interface_reply"
//	    },
//	    "igmp_group_prefix_set": {
//	        "reply": "igmp_group_prefix_set_reply"
//	    },
//	    "igmp_group_prefix_dump": {
//	        "reply": "igmp_group_prefix_details",
//	        "stream": true
//	    },
//	    "igmp_clear_interface": {
//	        "reply": "igmp_clear_interface_reply"
//	    },
//	    "igmp_listen": {
//	        "reply": "igmp_listen_reply"
//	    }
//	},
//
type Services interface {
	DumpIgmp(*IgmpDump) ([]*IgmpDetails, error)
	DumpIgmpGroupPrefix(*IgmpGroupPrefixDump) ([]*IgmpGroupPrefixDetails, error)
	IgmpClearInterface(*IgmpClearInterface) (*IgmpClearInterfaceReply, error)
	IgmpEnableDisable(*IgmpEnableDisable) (*IgmpEnableDisableReply, error)
	IgmpGroupPrefixSet(*IgmpGroupPrefixSet) (*IgmpGroupPrefixSetReply, error)
	IgmpListen(*IgmpListen) (*IgmpListenReply, error)
	IgmpProxyDeviceAddDel(*IgmpProxyDeviceAddDel) (*IgmpProxyDeviceAddDelReply, error)
	IgmpProxyDeviceAddDelInterface(*IgmpProxyDeviceAddDelInterface) (*IgmpProxyDeviceAddDelInterfaceReply, error)
	WantIgmpEvents(*WantIgmpEvents) (*WantIgmpEventsReply, error)
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

// FilterMode represents VPP binary API enum 'filter_mode':
//
//	"filter_mode",
//	[
//	    "EXCLUDE",
//	    0
//	],
//	[
//	    "INCLUDE",
//	    1
//	],
//	{
//	    "enumtype": "u32"
//	}
//
type FilterMode uint32

const (
	EXCLUDE FilterMode = 0
	INCLUDE FilterMode = 1
)

// GroupPrefixType represents VPP binary API enum 'group_prefix_type':
//
//	"group_prefix_type",
//	[
//	    "ASM",
//	    0
//	],
//	[
//	    "SSM",
//	    1
//	],
//	{
//	    "enumtype": "u32"
//	}
//
type GroupPrefixType uint32

const (
	ASM GroupPrefixType = 0
	SSM GroupPrefixType = 1
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

// IgmpGroup represents VPP binary API type 'igmp_group':
//
//	"igmp_group",
//	[
//	    "vl_api_filter_mode_t",
//	    "filter"
//	],
//	[
//	    "u8",
//	    "n_srcs"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "gaddr"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "saddrs",
//	    0,
//	    "n_srcs"
//	],
//	{
//	    "crc": "0xc0f42621"
//	}
//
type IgmpGroup struct {
	Filter    FilterMode
	NSrcs     uint8 `struc:"sizeof=Saddrs"`
	SwIfIndex uint32
	Gaddr     IP4Address
	Saddrs    []IP4Address
}

func (*IgmpGroup) GetTypeName() string {
	return "igmp_group"
}
func (*IgmpGroup) GetCrcString() string {
	return "c0f42621"
}

// GroupPrefix represents VPP binary API type 'group_prefix':
//
//	"group_prefix",
//	[
//	    "vl_api_group_prefix_type_t",
//	    "type"
//	],
//	[
//	    "vl_api_prefix_t",
//	    "prefix"
//	],
//	{
//	    "crc": "0x5398e813"
//	}
//
type GroupPrefix struct {
	Type   GroupPrefixType
	Prefix Prefix
}

func (*GroupPrefix) GetTypeName() string {
	return "group_prefix"
}
func (*GroupPrefix) GetCrcString() string {
	return "5398e813"
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

// IgmpListen represents VPP binary API message 'igmp_listen':
//
//	"igmp_listen",
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
//	    "vl_api_igmp_group_t",
//	    "group"
//	],
//	{
//	    "crc": "0xa4973479"
//	}
//
type IgmpListen struct {
	Group IgmpGroup
}

func (*IgmpListen) GetMessageName() string {
	return "igmp_listen"
}
func (*IgmpListen) GetCrcString() string {
	return "a4973479"
}
func (*IgmpListen) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpListenReply represents VPP binary API message 'igmp_listen_reply':
//
//	"igmp_listen_reply",
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
type IgmpListenReply struct {
	Retval int32
}

func (*IgmpListenReply) GetMessageName() string {
	return "igmp_listen_reply"
}
func (*IgmpListenReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IgmpListenReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpEnableDisable represents VPP binary API message 'igmp_enable_disable':
//
//	"igmp_enable_disable",
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
//	    "enable"
//	],
//	[
//	    "u8",
//	    "mode"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x5bd2a570"
//	}
//
type IgmpEnableDisable struct {
	Enable    uint8
	Mode      uint8
	SwIfIndex uint32
}

func (*IgmpEnableDisable) GetMessageName() string {
	return "igmp_enable_disable"
}
func (*IgmpEnableDisable) GetCrcString() string {
	return "5bd2a570"
}
func (*IgmpEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpEnableDisableReply represents VPP binary API message 'igmp_enable_disable_reply':
//
//	"igmp_enable_disable_reply",
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
type IgmpEnableDisableReply struct {
	Retval int32
}

func (*IgmpEnableDisableReply) GetMessageName() string {
	return "igmp_enable_disable_reply"
}
func (*IgmpEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IgmpEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpProxyDeviceAddDel represents VPP binary API message 'igmp_proxy_device_add_del':
//
//	"igmp_proxy_device_add_del",
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
//	    "add"
//	],
//	[
//	    "u32",
//	    "vrf_id"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0xf753aab3"
//	}
//
type IgmpProxyDeviceAddDel struct {
	Add       uint8
	VrfID     uint32
	SwIfIndex uint32
}

func (*IgmpProxyDeviceAddDel) GetMessageName() string {
	return "igmp_proxy_device_add_del"
}
func (*IgmpProxyDeviceAddDel) GetCrcString() string {
	return "f753aab3"
}
func (*IgmpProxyDeviceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpProxyDeviceAddDelReply represents VPP binary API message 'igmp_proxy_device_add_del_reply':
//
//	"igmp_proxy_device_add_del_reply",
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
type IgmpProxyDeviceAddDelReply struct {
	Retval int32
}

func (*IgmpProxyDeviceAddDelReply) GetMessageName() string {
	return "igmp_proxy_device_add_del_reply"
}
func (*IgmpProxyDeviceAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IgmpProxyDeviceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpProxyDeviceAddDelInterface represents VPP binary API message 'igmp_proxy_device_add_del_interface':
//
//	"igmp_proxy_device_add_del_interface",
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
//	    "add"
//	],
//	[
//	    "u32",
//	    "vrf_id"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0xf753aab3"
//	}
//
type IgmpProxyDeviceAddDelInterface struct {
	Add       uint8
	VrfID     uint32
	SwIfIndex uint32
}

func (*IgmpProxyDeviceAddDelInterface) GetMessageName() string {
	return "igmp_proxy_device_add_del_interface"
}
func (*IgmpProxyDeviceAddDelInterface) GetCrcString() string {
	return "f753aab3"
}
func (*IgmpProxyDeviceAddDelInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpProxyDeviceAddDelInterfaceReply represents VPP binary API message 'igmp_proxy_device_add_del_interface_reply':
//
//	"igmp_proxy_device_add_del_interface_reply",
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
type IgmpProxyDeviceAddDelInterfaceReply struct {
	Retval int32
}

func (*IgmpProxyDeviceAddDelInterfaceReply) GetMessageName() string {
	return "igmp_proxy_device_add_del_interface_reply"
}
func (*IgmpProxyDeviceAddDelInterfaceReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IgmpProxyDeviceAddDelInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpDump represents VPP binary API message 'igmp_dump':
//
//	"igmp_dump",
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
type IgmpDump struct {
	SwIfIndex uint32
}

func (*IgmpDump) GetMessageName() string {
	return "igmp_dump"
}
func (*IgmpDump) GetCrcString() string {
	return "529cb13f"
}
func (*IgmpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpDetails represents VPP binary API message 'igmp_details':
//
//	"igmp_details",
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
//	    "vl_api_ip4_address_t",
//	    "saddr"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "gaddr"
//	],
//	{
//	    "crc": "0x238a59f3"
//	}
//
type IgmpDetails struct {
	SwIfIndex uint32
	Saddr     IP4Address
	Gaddr     IP4Address
}

func (*IgmpDetails) GetMessageName() string {
	return "igmp_details"
}
func (*IgmpDetails) GetCrcString() string {
	return "238a59f3"
}
func (*IgmpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpClearInterface represents VPP binary API message 'igmp_clear_interface':
//
//	"igmp_clear_interface",
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
type IgmpClearInterface struct {
	SwIfIndex uint32
}

func (*IgmpClearInterface) GetMessageName() string {
	return "igmp_clear_interface"
}
func (*IgmpClearInterface) GetCrcString() string {
	return "529cb13f"
}
func (*IgmpClearInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpClearInterfaceReply represents VPP binary API message 'igmp_clear_interface_reply':
//
//	"igmp_clear_interface_reply",
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
type IgmpClearInterfaceReply struct {
	Retval int32
}

func (*IgmpClearInterfaceReply) GetMessageName() string {
	return "igmp_clear_interface_reply"
}
func (*IgmpClearInterfaceReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IgmpClearInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantIgmpEvents represents VPP binary API message 'want_igmp_events':
//
//	"want_igmp_events",
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
//	    "enable"
//	],
//	[
//	    "u32",
//	    "pid"
//	],
//	{
//	    "crc": "0xcfaccc1f"
//	}
//
type WantIgmpEvents struct {
	Enable uint32
	PID    uint32
}

func (*WantIgmpEvents) GetMessageName() string {
	return "want_igmp_events"
}
func (*WantIgmpEvents) GetCrcString() string {
	return "cfaccc1f"
}
func (*WantIgmpEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantIgmpEventsReply represents VPP binary API message 'want_igmp_events_reply':
//
//	"want_igmp_events_reply",
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
type WantIgmpEventsReply struct {
	Retval int32
}

func (*WantIgmpEventsReply) GetMessageName() string {
	return "want_igmp_events_reply"
}
func (*WantIgmpEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantIgmpEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpEvent represents VPP binary API message 'igmp_event':
//
//	"igmp_event",
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
//	    "vl_api_filter_mode_t",
//	    "filter"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "saddr"
//	],
//	[
//	    "vl_api_ip4_address_t",
//	    "gaddr"
//	],
//	{
//	    "crc": "0xb1122e50"
//	}
//
type IgmpEvent struct {
	SwIfIndex uint32
	Filter    FilterMode
	Saddr     IP4Address
	Gaddr     IP4Address
}

func (*IgmpEvent) GetMessageName() string {
	return "igmp_event"
}
func (*IgmpEvent) GetCrcString() string {
	return "b1122e50"
}
func (*IgmpEvent) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpGroupPrefixSet represents VPP binary API message 'igmp_group_prefix_set':
//
//	"igmp_group_prefix_set",
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
//	    "vl_api_group_prefix_t",
//	    "gp"
//	],
//	{
//	    "crc": "0x7ae82c26"
//	}
//
type IgmpGroupPrefixSet struct {
	Gp GroupPrefix
}

func (*IgmpGroupPrefixSet) GetMessageName() string {
	return "igmp_group_prefix_set"
}
func (*IgmpGroupPrefixSet) GetCrcString() string {
	return "7ae82c26"
}
func (*IgmpGroupPrefixSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpGroupPrefixSetReply represents VPP binary API message 'igmp_group_prefix_set_reply':
//
//	"igmp_group_prefix_set_reply",
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
type IgmpGroupPrefixSetReply struct {
	Retval int32
}

func (*IgmpGroupPrefixSetReply) GetMessageName() string {
	return "igmp_group_prefix_set_reply"
}
func (*IgmpGroupPrefixSetReply) GetCrcString() string {
	return "e8d4e804"
}
func (*IgmpGroupPrefixSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// IgmpGroupPrefixDump represents VPP binary API message 'igmp_group_prefix_dump':
//
//	"igmp_group_prefix_dump",
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
type IgmpGroupPrefixDump struct{}

func (*IgmpGroupPrefixDump) GetMessageName() string {
	return "igmp_group_prefix_dump"
}
func (*IgmpGroupPrefixDump) GetCrcString() string {
	return "51077d14"
}
func (*IgmpGroupPrefixDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// IgmpGroupPrefixDetails represents VPP binary API message 'igmp_group_prefix_details':
//
//	"igmp_group_prefix_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_group_prefix_t",
//	    "gp"
//	],
//	{
//	    "crc": "0xc6b0f93d"
//	}
//
type IgmpGroupPrefixDetails struct {
	Gp GroupPrefix
}

func (*IgmpGroupPrefixDetails) GetMessageName() string {
	return "igmp_group_prefix_details"
}
func (*IgmpGroupPrefixDetails) GetCrcString() string {
	return "c6b0f93d"
}
func (*IgmpGroupPrefixDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*IgmpListen)(nil), "igmp.IgmpListen")
	api.RegisterMessage((*IgmpListenReply)(nil), "igmp.IgmpListenReply")
	api.RegisterMessage((*IgmpEnableDisable)(nil), "igmp.IgmpEnableDisable")
	api.RegisterMessage((*IgmpEnableDisableReply)(nil), "igmp.IgmpEnableDisableReply")
	api.RegisterMessage((*IgmpProxyDeviceAddDel)(nil), "igmp.IgmpProxyDeviceAddDel")
	api.RegisterMessage((*IgmpProxyDeviceAddDelReply)(nil), "igmp.IgmpProxyDeviceAddDelReply")
	api.RegisterMessage((*IgmpProxyDeviceAddDelInterface)(nil), "igmp.IgmpProxyDeviceAddDelInterface")
	api.RegisterMessage((*IgmpProxyDeviceAddDelInterfaceReply)(nil), "igmp.IgmpProxyDeviceAddDelInterfaceReply")
	api.RegisterMessage((*IgmpDump)(nil), "igmp.IgmpDump")
	api.RegisterMessage((*IgmpDetails)(nil), "igmp.IgmpDetails")
	api.RegisterMessage((*IgmpClearInterface)(nil), "igmp.IgmpClearInterface")
	api.RegisterMessage((*IgmpClearInterfaceReply)(nil), "igmp.IgmpClearInterfaceReply")
	api.RegisterMessage((*WantIgmpEvents)(nil), "igmp.WantIgmpEvents")
	api.RegisterMessage((*WantIgmpEventsReply)(nil), "igmp.WantIgmpEventsReply")
	api.RegisterMessage((*IgmpEvent)(nil), "igmp.IgmpEvent")
	api.RegisterMessage((*IgmpGroupPrefixSet)(nil), "igmp.IgmpGroupPrefixSet")
	api.RegisterMessage((*IgmpGroupPrefixSetReply)(nil), "igmp.IgmpGroupPrefixSetReply")
	api.RegisterMessage((*IgmpGroupPrefixDump)(nil), "igmp.IgmpGroupPrefixDump")
	api.RegisterMessage((*IgmpGroupPrefixDetails)(nil), "igmp.IgmpGroupPrefixDetails")
}
