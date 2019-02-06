// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/dhcp.api.json

/*
 Package dhcp is a generated from VPP binary API module 'dhcp'.

 It contains following objects:
	 25 messages
	  5 types
	 11 services

*/
package dhcp

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
//	    "dhcp_client_config": {
//	        "reply": "dhcp_client_config_reply",
//	        "events": [
//	            "dhcp_compl_event"
//	        ]
//	    },
//	    "want_dhcp6_reply_events": {
//	        "reply": "want_dhcp6_reply_events_reply",
//	        "events": [
//	            "dhcp6_reply_event"
//	        ]
//	    },
//	    "want_dhcp6_pd_reply_events": {
//	        "reply": "want_dhcp6_pd_reply_events_reply",
//	        "events": [
//	            "dhcp6_pd_reply_event"
//	        ]
//	    },
//	    "dhcp_proxy_config": {
//	        "reply": "dhcp_proxy_config_reply"
//	    },
//	    "dhcp_proxy_set_vss": {
//	        "reply": "dhcp_proxy_set_vss_reply"
//	    },
//	    "dhcp_client_dump": {
//	        "reply": "dhcp_client_details",
//	        "stream": true
//	    },
//	    "dhcp_proxy_dump": {
//	        "reply": "dhcp_proxy_details",
//	        "stream": true
//	    },
//	    "dhcp6_duid_ll_set": {
//	        "reply": "dhcp6_duid_ll_set_reply"
//	    },
//	    "dhcp6_clients_enable_disable": {
//	        "reply": "dhcp6_clients_enable_disable_reply"
//	    },
//	    "dhcp6_send_client_message": {
//	        "reply": "dhcp6_send_client_message_reply"
//	    },
//	    "dhcp6_pd_send_client_message": {
//	        "reply": "dhcp6_pd_send_client_message_reply"
//	    }
//	},
//
type Services interface {
	DumpDHCPClient(*DHCPClientDump) ([]*DHCPClientDetails, error)
	DumpDHCPProxy(*DHCPProxyDump) ([]*DHCPProxyDetails, error)
	DHCP6ClientsEnableDisable(*DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error)
	DHCP6DuidLlSet(*DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error)
	DHCP6PdSendClientMessage(*DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error)
	DHCP6SendClientMessage(*DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error)
	DHCPClientConfig(*DHCPClientConfig) (*DHCPClientConfigReply, error)
	DHCPProxyConfig(*DHCPProxyConfig) (*DHCPProxyConfigReply, error)
	DHCPProxySetVss(*DHCPProxySetVss) (*DHCPProxySetVssReply, error)
	WantDHCP6PdReplyEvents(*WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error)
	WantDHCP6ReplyEvents(*WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error)
}

/* Types */

// DHCPClient represents VPP binary API type 'dhcp_client':
//
//	"dhcp_client",
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u8",
//	    "hostname",
//	    64
//	],
//	[
//	    "u8",
//	    "id",
//	    64
//	],
//	[
//	    "u8",
//	    "want_dhcp_event"
//	],
//	[
//	    "u8",
//	    "set_broadcast_flag"
//	],
//	[
//	    "u32",
//	    "pid"
//	],
//	{
//	    "crc": "0x527f7935"
//	}
//
type DHCPClient struct {
	SwIfIndex        uint32
	Hostname         []byte `struc:"[64]byte"`
	ID               []byte `struc:"[64]byte"`
	WantDHCPEvent    uint8
	SetBroadcastFlag uint8
	PID              uint32
}

func (*DHCPClient) GetTypeName() string {
	return "dhcp_client"
}
func (*DHCPClient) GetCrcString() string {
	return "527f7935"
}

// DHCPLease represents VPP binary API type 'dhcp_lease':
//
//	"dhcp_lease",
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u8",
//	    "state"
//	],
//	[
//	    "u8",
//	    "hostname",
//	    64
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "mask_width"
//	],
//	[
//	    "u8",
//	    "host_address",
//	    16
//	],
//	[
//	    "u8",
//	    "router_address",
//	    16
//	],
//	[
//	    "u8",
//	    "host_mac",
//	    6
//	],
//	{
//	    "crc": "0x61090276"
//	}
//
type DHCPLease struct {
	SwIfIndex     uint32
	State         uint8
	Hostname      []byte `struc:"[64]byte"`
	IsIPv6        uint8
	MaskWidth     uint8
	HostAddress   []byte `struc:"[16]byte"`
	RouterAddress []byte `struc:"[16]byte"`
	HostMac       []byte `struc:"[6]byte"`
}

func (*DHCPLease) GetTypeName() string {
	return "dhcp_lease"
}
func (*DHCPLease) GetCrcString() string {
	return "61090276"
}

// DHCPServer represents VPP binary API type 'dhcp_server':
//
//	"dhcp_server",
//	[
//	    "u32",
//	    "server_vrf_id"
//	],
//	[
//	    "u8",
//	    "dhcp_server",
//	    16
//	],
//	{
//	    "crc": "0xf16506c4"
//	}
//
type DHCPServer struct {
	ServerVrfID uint32
	DHCPServer  []byte `struc:"[16]byte"`
}

func (*DHCPServer) GetTypeName() string {
	return "dhcp_server"
}
func (*DHCPServer) GetCrcString() string {
	return "f16506c4"
}

// DHCP6AddressInfo represents VPP binary API type 'dhcp6_address_info':
//
//	"dhcp6_address_info",
//	[
//	    "u8",
//	    "address",
//	    16
//	],
//	[
//	    "u32",
//	    "valid_time"
//	],
//	[
//	    "u32",
//	    "preferred_time"
//	],
//	{
//	    "crc": "0xf3d501e2"
//	}
//
type DHCP6AddressInfo struct {
	Address       []byte `struc:"[16]byte"`
	ValidTime     uint32
	PreferredTime uint32
}

func (*DHCP6AddressInfo) GetTypeName() string {
	return "dhcp6_address_info"
}
func (*DHCP6AddressInfo) GetCrcString() string {
	return "f3d501e2"
}

// DHCP6PdPrefixInfo represents VPP binary API type 'dhcp6_pd_prefix_info':
//
//	"dhcp6_pd_prefix_info",
//	[
//	    "u8",
//	    "prefix",
//	    16
//	],
//	[
//	    "u8",
//	    "prefix_length"
//	],
//	[
//	    "u32",
//	    "valid_time"
//	],
//	[
//	    "u32",
//	    "preferred_time"
//	],
//	{
//	    "crc": "0xc459690e"
//	}
//
type DHCP6PdPrefixInfo struct {
	Prefix        []byte `struc:"[16]byte"`
	PrefixLength  uint8
	ValidTime     uint32
	PreferredTime uint32
}

func (*DHCP6PdPrefixInfo) GetTypeName() string {
	return "dhcp6_pd_prefix_info"
}
func (*DHCP6PdPrefixInfo) GetCrcString() string {
	return "c459690e"
}

/* Messages */

// DHCPProxyConfig represents VPP binary API message 'dhcp_proxy_config':
//
//	"dhcp_proxy_config",
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
//	    "rx_vrf_id"
//	],
//	[
//	    "u32",
//	    "server_vrf_id"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	[
//	    "u8",
//	    "dhcp_server",
//	    16
//	],
//	[
//	    "u8",
//	    "dhcp_src_address",
//	    16
//	],
//	{
//	    "crc": "0x6af4b645"
//	}
//
type DHCPProxyConfig struct {
	RxVrfID        uint32
	ServerVrfID    uint32
	IsIPv6         uint8
	IsAdd          uint8
	DHCPServer     []byte `struc:"[16]byte"`
	DHCPSrcAddress []byte `struc:"[16]byte"`
}

func (*DHCPProxyConfig) GetMessageName() string {
	return "dhcp_proxy_config"
}
func (*DHCPProxyConfig) GetCrcString() string {
	return "6af4b645"
}
func (*DHCPProxyConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxyConfigReply represents VPP binary API message 'dhcp_proxy_config_reply':
//
//	"dhcp_proxy_config_reply",
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
type DHCPProxyConfigReply struct {
	Retval int32
}

func (*DHCPProxyConfigReply) GetMessageName() string {
	return "dhcp_proxy_config_reply"
}
func (*DHCPProxyConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPProxyConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPProxySetVss represents VPP binary API message 'dhcp_proxy_set_vss':
//
//	"dhcp_proxy_set_vss",
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
//	    "tbl_id"
//	],
//	[
//	    "u8",
//	    "vss_type"
//	],
//	[
//	    "u8",
//	    "vpn_ascii_id",
//	    129
//	],
//	[
//	    "u32",
//	    "oui"
//	],
//	[
//	    "u32",
//	    "vpn_index"
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	{
//	    "crc": "0x606535aa"
//	}
//
type DHCPProxySetVss struct {
	TblID      uint32
	VssType    uint8
	VPNAsciiID []byte `struc:"[129]byte"`
	Oui        uint32
	VPNIndex   uint32
	IsIPv6     uint8
	IsAdd      uint8
}

func (*DHCPProxySetVss) GetMessageName() string {
	return "dhcp_proxy_set_vss"
}
func (*DHCPProxySetVss) GetCrcString() string {
	return "606535aa"
}
func (*DHCPProxySetVss) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxySetVssReply represents VPP binary API message 'dhcp_proxy_set_vss_reply':
//
//	"dhcp_proxy_set_vss_reply",
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
type DHCPProxySetVssReply struct {
	Retval int32
}

func (*DHCPProxySetVssReply) GetMessageName() string {
	return "dhcp_proxy_set_vss_reply"
}
func (*DHCPProxySetVssReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPProxySetVssReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientConfig represents VPP binary API message 'dhcp_client_config':
//
//	"dhcp_client_config",
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
//	    "vl_api_dhcp_client_t",
//	    "client"
//	],
//	{
//	    "crc": "0xc32ccdfe"
//	}
//
type DHCPClientConfig struct {
	IsAdd  uint8
	Client DHCPClient
}

func (*DHCPClientConfig) GetMessageName() string {
	return "dhcp_client_config"
}
func (*DHCPClientConfig) GetCrcString() string {
	return "c32ccdfe"
}
func (*DHCPClientConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPClientConfigReply represents VPP binary API message 'dhcp_client_config_reply':
//
//	"dhcp_client_config_reply",
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
type DHCPClientConfigReply struct {
	Retval int32
}

func (*DHCPClientConfigReply) GetMessageName() string {
	return "dhcp_client_config_reply"
}
func (*DHCPClientConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPClientConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPComplEvent represents VPP binary API message 'dhcp_compl_event':
//
//	"dhcp_compl_event",
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
//	    "pid"
//	],
//	[
//	    "vl_api_dhcp_lease_t",
//	    "lease"
//	],
//	{
//	    "crc": "0x2105c31b"
//	}
//
type DHCPComplEvent struct {
	PID   uint32
	Lease DHCPLease
}

func (*DHCPComplEvent) GetMessageName() string {
	return "dhcp_compl_event"
}
func (*DHCPComplEvent) GetCrcString() string {
	return "2105c31b"
}
func (*DHCPComplEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCPClientDump represents VPP binary API message 'dhcp_client_dump':
//
//	"dhcp_client_dump",
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
type DHCPClientDump struct{}

func (*DHCPClientDump) GetMessageName() string {
	return "dhcp_client_dump"
}
func (*DHCPClientDump) GetCrcString() string {
	return "51077d14"
}
func (*DHCPClientDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPClientDetails represents VPP binary API message 'dhcp_client_details':
//
//	"dhcp_client_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_dhcp_client_t",
//	    "client"
//	],
//	[
//	    "vl_api_dhcp_lease_t",
//	    "lease"
//	],
//	{
//	    "crc": "0x7ea3a745"
//	}
//
type DHCPClientDetails struct {
	Client DHCPClient
	Lease  DHCPLease
}

func (*DHCPClientDetails) GetMessageName() string {
	return "dhcp_client_details"
}
func (*DHCPClientDetails) GetCrcString() string {
	return "7ea3a745"
}
func (*DHCPClientDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPProxyDump represents VPP binary API message 'dhcp_proxy_dump':
//
//	"dhcp_proxy_dump",
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
//	    "is_ip6"
//	],
//	{
//	    "crc": "0x6fe91190"
//	}
//
type DHCPProxyDump struct {
	IsIP6 uint8
}

func (*DHCPProxyDump) GetMessageName() string {
	return "dhcp_proxy_dump"
}
func (*DHCPProxyDump) GetCrcString() string {
	return "6fe91190"
}
func (*DHCPProxyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxyDetails represents VPP binary API message 'dhcp_proxy_details':
//
//	"dhcp_proxy_details",
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
//	    "rx_vrf_id"
//	],
//	[
//	    "u32",
//	    "vss_oui"
//	],
//	[
//	    "u32",
//	    "vss_fib_id"
//	],
//	[
//	    "u8",
//	    "vss_type"
//	],
//	[
//	    "u8",
//	    "vss_vpn_ascii_id",
//	    129
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "dhcp_src_address",
//	    16
//	],
//	[
//	    "u8",
//	    "count"
//	],
//	[
//	    "vl_api_dhcp_server_t",
//	    "servers",
//	    0,
//	    "count"
//	],
//	{
//	    "crc": "0xa5f2ad84"
//	}
//
type DHCPProxyDetails struct {
	RxVrfID        uint32
	VssOui         uint32
	VssFibID       uint32
	VssType        uint8
	VssVPNAsciiID  []byte `struc:"[129]byte"`
	IsIPv6         uint8
	DHCPSrcAddress []byte `struc:"[16]byte"`
	Count          uint8  `struc:"sizeof=Servers"`
	Servers        []DHCPServer
}

func (*DHCPProxyDetails) GetMessageName() string {
	return "dhcp_proxy_details"
}
func (*DHCPProxyDetails) GetCrcString() string {
	return "a5f2ad84"
}
func (*DHCPProxyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6DuidLlSet represents VPP binary API message 'dhcp6_duid_ll_set':
//
//	"dhcp6_duid_ll_set",
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
//	    "duid_ll",
//	    10
//	],
//	{
//	    "crc": "0x0f6ca323"
//	}
//
type DHCP6DuidLlSet struct {
	DuidLl []byte `struc:"[10]byte"`
}

func (*DHCP6DuidLlSet) GetMessageName() string {
	return "dhcp6_duid_ll_set"
}
func (*DHCP6DuidLlSet) GetCrcString() string {
	return "0f6ca323"
}
func (*DHCP6DuidLlSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6DuidLlSetReply represents VPP binary API message 'dhcp6_duid_ll_set_reply':
//
//	"dhcp6_duid_ll_set_reply",
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
type DHCP6DuidLlSetReply struct {
	Retval int32
}

func (*DHCP6DuidLlSetReply) GetMessageName() string {
	return "dhcp6_duid_ll_set_reply"
}
func (*DHCP6DuidLlSetReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6DuidLlSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6ClientsEnableDisable represents VPP binary API message 'dhcp6_clients_enable_disable':
//
//	"dhcp6_clients_enable_disable",
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
//	{
//	    "crc": "0x8050327d"
//	}
//
type DHCP6ClientsEnableDisable struct {
	Enable uint8
}

func (*DHCP6ClientsEnableDisable) GetMessageName() string {
	return "dhcp6_clients_enable_disable"
}
func (*DHCP6ClientsEnableDisable) GetCrcString() string {
	return "8050327d"
}
func (*DHCP6ClientsEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6ClientsEnableDisableReply represents VPP binary API message 'dhcp6_clients_enable_disable_reply':
//
//	"dhcp6_clients_enable_disable_reply",
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
type DHCP6ClientsEnableDisableReply struct {
	Retval int32
}

func (*DHCP6ClientsEnableDisableReply) GetMessageName() string {
	return "dhcp6_clients_enable_disable_reply"
}
func (*DHCP6ClientsEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6ClientsEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6SendClientMessage represents VPP binary API message 'dhcp6_send_client_message':
//
//	"dhcp6_send_client_message",
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
//	    "u32",
//	    "server_index"
//	],
//	[
//	    "u32",
//	    "irt"
//	],
//	[
//	    "u32",
//	    "mrt"
//	],
//	[
//	    "u32",
//	    "mrc"
//	],
//	[
//	    "u32",
//	    "mrd"
//	],
//	[
//	    "u8",
//	    "stop"
//	],
//	[
//	    "u8",
//	    "msg_type"
//	],
//	[
//	    "u32",
//	    "T1"
//	],
//	[
//	    "u32",
//	    "T2"
//	],
//	[
//	    "u32",
//	    "n_addresses"
//	],
//	[
//	    "vl_api_dhcp6_address_info_t",
//	    "addresses",
//	    0,
//	    "n_addresses"
//	],
//	{
//	    "crc": "0xa13ae8c4"
//	}
//
type DHCP6SendClientMessage struct {
	SwIfIndex   uint32
	ServerIndex uint32
	Irt         uint32
	Mrt         uint32
	Mrc         uint32
	Mrd         uint32
	Stop        uint8
	MsgType     uint8
	T1          uint32
	T2          uint32
	NAddresses  uint32 `struc:"sizeof=Addresses"`
	Addresses   []DHCP6AddressInfo
}

func (*DHCP6SendClientMessage) GetMessageName() string {
	return "dhcp6_send_client_message"
}
func (*DHCP6SendClientMessage) GetCrcString() string {
	return "a13ae8c4"
}
func (*DHCP6SendClientMessage) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6SendClientMessageReply represents VPP binary API message 'dhcp6_send_client_message_reply':
//
//	"dhcp6_send_client_message_reply",
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
type DHCP6SendClientMessageReply struct {
	Retval int32
}

func (*DHCP6SendClientMessageReply) GetMessageName() string {
	return "dhcp6_send_client_message_reply"
}
func (*DHCP6SendClientMessageReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6SendClientMessageReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6PdSendClientMessage represents VPP binary API message 'dhcp6_pd_send_client_message':
//
//	"dhcp6_pd_send_client_message",
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
//	    "u32",
//	    "server_index"
//	],
//	[
//	    "u32",
//	    "irt"
//	],
//	[
//	    "u32",
//	    "mrt"
//	],
//	[
//	    "u32",
//	    "mrc"
//	],
//	[
//	    "u32",
//	    "mrd"
//	],
//	[
//	    "u8",
//	    "stop"
//	],
//	[
//	    "u8",
//	    "msg_type"
//	],
//	[
//	    "u32",
//	    "T1"
//	],
//	[
//	    "u32",
//	    "T2"
//	],
//	[
//	    "u32",
//	    "n_prefixes"
//	],
//	[
//	    "vl_api_dhcp6_pd_prefix_info_t",
//	    "prefixes",
//	    0,
//	    "n_prefixes"
//	],
//	{
//	    "crc": "0x5a063fbc"
//	}
//
type DHCP6PdSendClientMessage struct {
	SwIfIndex   uint32
	ServerIndex uint32
	Irt         uint32
	Mrt         uint32
	Mrc         uint32
	Mrd         uint32
	Stop        uint8
	MsgType     uint8
	T1          uint32
	T2          uint32
	NPrefixes   uint32 `struc:"sizeof=Prefixes"`
	Prefixes    []DHCP6PdPrefixInfo
}

func (*DHCP6PdSendClientMessage) GetMessageName() string {
	return "dhcp6_pd_send_client_message"
}
func (*DHCP6PdSendClientMessage) GetCrcString() string {
	return "5a063fbc"
}
func (*DHCP6PdSendClientMessage) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6PdSendClientMessageReply represents VPP binary API message 'dhcp6_pd_send_client_message_reply':
//
//	"dhcp6_pd_send_client_message_reply",
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
type DHCP6PdSendClientMessageReply struct {
	Retval int32
}

func (*DHCP6PdSendClientMessageReply) GetMessageName() string {
	return "dhcp6_pd_send_client_message_reply"
}
func (*DHCP6PdSendClientMessageReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6PdSendClientMessageReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantDHCP6ReplyEvents represents VPP binary API message 'want_dhcp6_reply_events':
//
//	"want_dhcp6_reply_events",
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
//	    "enable_disable"
//	],
//	[
//	    "u32",
//	    "pid"
//	],
//	{
//	    "crc": "0x05b454b5"
//	}
//
type WantDHCP6ReplyEvents struct {
	EnableDisable uint8
	PID           uint32
}

func (*WantDHCP6ReplyEvents) GetMessageName() string {
	return "want_dhcp6_reply_events"
}
func (*WantDHCP6ReplyEvents) GetCrcString() string {
	return "05b454b5"
}
func (*WantDHCP6ReplyEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantDHCP6ReplyEventsReply represents VPP binary API message 'want_dhcp6_reply_events_reply':
//
//	"want_dhcp6_reply_events_reply",
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
type WantDHCP6ReplyEventsReply struct {
	Retval int32
}

func (*WantDHCP6ReplyEventsReply) GetMessageName() string {
	return "want_dhcp6_reply_events_reply"
}
func (*WantDHCP6ReplyEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantDHCP6ReplyEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantDHCP6PdReplyEvents represents VPP binary API message 'want_dhcp6_pd_reply_events':
//
//	"want_dhcp6_pd_reply_events",
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
//	    "enable_disable"
//	],
//	[
//	    "u32",
//	    "pid"
//	],
//	{
//	    "crc": "0x05b454b5"
//	}
//
type WantDHCP6PdReplyEvents struct {
	EnableDisable uint8
	PID           uint32
}

func (*WantDHCP6PdReplyEvents) GetMessageName() string {
	return "want_dhcp6_pd_reply_events"
}
func (*WantDHCP6PdReplyEvents) GetCrcString() string {
	return "05b454b5"
}
func (*WantDHCP6PdReplyEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantDHCP6PdReplyEventsReply represents VPP binary API message 'want_dhcp6_pd_reply_events_reply':
//
//	"want_dhcp6_pd_reply_events_reply",
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
type WantDHCP6PdReplyEventsReply struct {
	Retval int32
}

func (*WantDHCP6PdReplyEventsReply) GetMessageName() string {
	return "want_dhcp6_pd_reply_events_reply"
}
func (*WantDHCP6PdReplyEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantDHCP6PdReplyEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6ReplyEvent represents VPP binary API message 'dhcp6_reply_event':
//
//	"dhcp6_reply_event",
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
//	    "pid"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u32",
//	    "server_index"
//	],
//	[
//	    "u8",
//	    "msg_type"
//	],
//	[
//	    "u32",
//	    "T1"
//	],
//	[
//	    "u32",
//	    "T2"
//	],
//	[
//	    "u16",
//	    "inner_status_code"
//	],
//	[
//	    "u16",
//	    "status_code"
//	],
//	[
//	    "u8",
//	    "preference"
//	],
//	[
//	    "u32",
//	    "n_addresses"
//	],
//	[
//	    "vl_api_dhcp6_address_info_t",
//	    "addresses",
//	    0,
//	    "n_addresses"
//	],
//	{
//	    "crc": "0xac4563f9"
//	}
//
type DHCP6ReplyEvent struct {
	PID             uint32
	SwIfIndex       uint32
	ServerIndex     uint32
	MsgType         uint8
	T1              uint32
	T2              uint32
	InnerStatusCode uint16
	StatusCode      uint16
	Preference      uint8
	NAddresses      uint32 `struc:"sizeof=Addresses"`
	Addresses       []DHCP6AddressInfo
}

func (*DHCP6ReplyEvent) GetMessageName() string {
	return "dhcp6_reply_event"
}
func (*DHCP6ReplyEvent) GetCrcString() string {
	return "ac4563f9"
}
func (*DHCP6ReplyEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCP6PdReplyEvent represents VPP binary API message 'dhcp6_pd_reply_event':
//
//	"dhcp6_pd_reply_event",
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
//	    "pid"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u32",
//	    "server_index"
//	],
//	[
//	    "u8",
//	    "msg_type"
//	],
//	[
//	    "u32",
//	    "T1"
//	],
//	[
//	    "u32",
//	    "T2"
//	],
//	[
//	    "u16",
//	    "inner_status_code"
//	],
//	[
//	    "u16",
//	    "status_code"
//	],
//	[
//	    "u8",
//	    "preference"
//	],
//	[
//	    "u32",
//	    "n_prefixes"
//	],
//	[
//	    "vl_api_dhcp6_pd_prefix_info_t",
//	    "prefixes",
//	    0,
//	    "n_prefixes"
//	],
//	{
//	    "crc": "0x48e73c36"
//	}
//
type DHCP6PdReplyEvent struct {
	PID             uint32
	SwIfIndex       uint32
	ServerIndex     uint32
	MsgType         uint8
	T1              uint32
	T2              uint32
	InnerStatusCode uint16
	StatusCode      uint16
	Preference      uint8
	NPrefixes       uint32 `struc:"sizeof=Prefixes"`
	Prefixes        []DHCP6PdPrefixInfo
}

func (*DHCP6PdReplyEvent) GetMessageName() string {
	return "dhcp6_pd_reply_event"
}
func (*DHCP6PdReplyEvent) GetCrcString() string {
	return "48e73c36"
}
func (*DHCP6PdReplyEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

func init() {
	api.RegisterMessage((*DHCPProxyConfig)(nil), "dhcp.DHCPProxyConfig")
	api.RegisterMessage((*DHCPProxyConfigReply)(nil), "dhcp.DHCPProxyConfigReply")
	api.RegisterMessage((*DHCPProxySetVss)(nil), "dhcp.DHCPProxySetVss")
	api.RegisterMessage((*DHCPProxySetVssReply)(nil), "dhcp.DHCPProxySetVssReply")
	api.RegisterMessage((*DHCPClientConfig)(nil), "dhcp.DHCPClientConfig")
	api.RegisterMessage((*DHCPClientConfigReply)(nil), "dhcp.DHCPClientConfigReply")
	api.RegisterMessage((*DHCPComplEvent)(nil), "dhcp.DHCPComplEvent")
	api.RegisterMessage((*DHCPClientDump)(nil), "dhcp.DHCPClientDump")
	api.RegisterMessage((*DHCPClientDetails)(nil), "dhcp.DHCPClientDetails")
	api.RegisterMessage((*DHCPProxyDump)(nil), "dhcp.DHCPProxyDump")
	api.RegisterMessage((*DHCPProxyDetails)(nil), "dhcp.DHCPProxyDetails")
	api.RegisterMessage((*DHCP6DuidLlSet)(nil), "dhcp.DHCP6DuidLlSet")
	api.RegisterMessage((*DHCP6DuidLlSetReply)(nil), "dhcp.DHCP6DuidLlSetReply")
	api.RegisterMessage((*DHCP6ClientsEnableDisable)(nil), "dhcp.DHCP6ClientsEnableDisable")
	api.RegisterMessage((*DHCP6ClientsEnableDisableReply)(nil), "dhcp.DHCP6ClientsEnableDisableReply")
	api.RegisterMessage((*DHCP6SendClientMessage)(nil), "dhcp.DHCP6SendClientMessage")
	api.RegisterMessage((*DHCP6SendClientMessageReply)(nil), "dhcp.DHCP6SendClientMessageReply")
	api.RegisterMessage((*DHCP6PdSendClientMessage)(nil), "dhcp.DHCP6PdSendClientMessage")
	api.RegisterMessage((*DHCP6PdSendClientMessageReply)(nil), "dhcp.DHCP6PdSendClientMessageReply")
	api.RegisterMessage((*WantDHCP6ReplyEvents)(nil), "dhcp.WantDHCP6ReplyEvents")
	api.RegisterMessage((*WantDHCP6ReplyEventsReply)(nil), "dhcp.WantDHCP6ReplyEventsReply")
	api.RegisterMessage((*WantDHCP6PdReplyEvents)(nil), "dhcp.WantDHCP6PdReplyEvents")
	api.RegisterMessage((*WantDHCP6PdReplyEventsReply)(nil), "dhcp.WantDHCP6PdReplyEventsReply")
	api.RegisterMessage((*DHCP6ReplyEvent)(nil), "dhcp.DHCP6ReplyEvent")
	api.RegisterMessage((*DHCP6PdReplyEvent)(nil), "dhcp.DHCP6PdReplyEvent")
}
