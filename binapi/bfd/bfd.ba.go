// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/bfd.api.json

/*
 Package bfd is a generated from VPP binary API module 'bfd'.

 It contains following objects:
	 28 messages
	 14 services

*/
package bfd

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0xa37bd695

// Services represents VPP binary API services:
//
//	"services": {
//	    "bfd_udp_mod": {
//	        "reply": "bfd_udp_mod_reply"
//	    },
//	    "bfd_auth_del_key": {
//	        "reply": "bfd_auth_del_key_reply"
//	    },
//	    "bfd_udp_del_echo_source": {
//	        "reply": "bfd_udp_del_echo_source_reply"
//	    },
//	    "bfd_auth_keys_dump": {
//	        "reply": "bfd_auth_keys_details",
//	        "stream": true
//	    },
//	    "bfd_udp_get_echo_source": {
//	        "reply": "bfd_udp_get_echo_source_reply"
//	    },
//	    "want_bfd_events": {
//	        "reply": "want_bfd_events_reply"
//	    },
//	    "bfd_udp_auth_deactivate": {
//	        "reply": "bfd_udp_auth_deactivate_reply"
//	    },
//	    "bfd_auth_set_key": {
//	        "reply": "bfd_auth_set_key_reply"
//	    },
//	    "bfd_udp_auth_activate": {
//	        "reply": "bfd_udp_auth_activate_reply"
//	    },
//	    "bfd_udp_session_dump": {
//	        "reply": "bfd_udp_session_details",
//	        "stream": true
//	    },
//	    "bfd_udp_add": {
//	        "reply": "bfd_udp_add_reply"
//	    },
//	    "bfd_udp_del": {
//	        "reply": "bfd_udp_del_reply"
//	    },
//	    "bfd_udp_set_echo_source": {
//	        "reply": "bfd_udp_set_echo_source_reply"
//	    },
//	    "bfd_udp_session_set_flags": {
//	        "reply": "bfd_udp_session_set_flags_reply"
//	    }
//	},
//
type Services interface {
	DumpBfdAuthKeys(*BfdAuthKeysDump) ([]*BfdAuthKeysDetails, error)
	DumpBfdUDPSession(*BfdUDPSessionDump) ([]*BfdUDPSessionDetails, error)
	BfdAuthDelKey(*BfdAuthDelKey) (*BfdAuthDelKeyReply, error)
	BfdAuthSetKey(*BfdAuthSetKey) (*BfdAuthSetKeyReply, error)
	BfdUDPAdd(*BfdUDPAdd) (*BfdUDPAddReply, error)
	BfdUDPAuthActivate(*BfdUDPAuthActivate) (*BfdUDPAuthActivateReply, error)
	BfdUDPAuthDeactivate(*BfdUDPAuthDeactivate) (*BfdUDPAuthDeactivateReply, error)
	BfdUDPDel(*BfdUDPDel) (*BfdUDPDelReply, error)
	BfdUDPDelEchoSource(*BfdUDPDelEchoSource) (*BfdUDPDelEchoSourceReply, error)
	BfdUDPGetEchoSource(*BfdUDPGetEchoSource) (*BfdUDPGetEchoSourceReply, error)
	BfdUDPMod(*BfdUDPMod) (*BfdUDPModReply, error)
	BfdUDPSessionSetFlags(*BfdUDPSessionSetFlags) (*BfdUDPSessionSetFlagsReply, error)
	BfdUDPSetEchoSource(*BfdUDPSetEchoSource) (*BfdUDPSetEchoSourceReply, error)
	WantBfdEvents(*WantBfdEvents) (*WantBfdEventsReply, error)
}

/* Messages */

// BfdUDPSetEchoSource represents VPP binary API message 'bfd_udp_set_echo_source':
//
//	"bfd_udp_set_echo_source",
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
type BfdUDPSetEchoSource struct {
	SwIfIndex uint32
}

func (*BfdUDPSetEchoSource) GetMessageName() string {
	return "bfd_udp_set_echo_source"
}
func (*BfdUDPSetEchoSource) GetCrcString() string {
	return "529cb13f"
}
func (*BfdUDPSetEchoSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPSetEchoSourceReply represents VPP binary API message 'bfd_udp_set_echo_source_reply':
//
//	"bfd_udp_set_echo_source_reply",
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
type BfdUDPSetEchoSourceReply struct {
	Retval int32
}

func (*BfdUDPSetEchoSourceReply) GetMessageName() string {
	return "bfd_udp_set_echo_source_reply"
}
func (*BfdUDPSetEchoSourceReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPSetEchoSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPDelEchoSource represents VPP binary API message 'bfd_udp_del_echo_source':
//
//	"bfd_udp_del_echo_source",
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
type BfdUDPDelEchoSource struct{}

func (*BfdUDPDelEchoSource) GetMessageName() string {
	return "bfd_udp_del_echo_source"
}
func (*BfdUDPDelEchoSource) GetCrcString() string {
	return "51077d14"
}
func (*BfdUDPDelEchoSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPDelEchoSourceReply represents VPP binary API message 'bfd_udp_del_echo_source_reply':
//
//	"bfd_udp_del_echo_source_reply",
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
type BfdUDPDelEchoSourceReply struct {
	Retval int32
}

func (*BfdUDPDelEchoSourceReply) GetMessageName() string {
	return "bfd_udp_del_echo_source_reply"
}
func (*BfdUDPDelEchoSourceReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPDelEchoSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPGetEchoSource represents VPP binary API message 'bfd_udp_get_echo_source':
//
//	"bfd_udp_get_echo_source",
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
type BfdUDPGetEchoSource struct{}

func (*BfdUDPGetEchoSource) GetMessageName() string {
	return "bfd_udp_get_echo_source"
}
func (*BfdUDPGetEchoSource) GetCrcString() string {
	return "51077d14"
}
func (*BfdUDPGetEchoSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPGetEchoSourceReply represents VPP binary API message 'bfd_udp_get_echo_source_reply':
//
//	"bfd_udp_get_echo_source_reply",
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
//	    "u8",
//	    "is_set"
//	],
//	[
//	    "u8",
//	    "have_usable_ip4"
//	],
//	[
//	    "u8",
//	    "ip4_addr",
//	    4
//	],
//	[
//	    "u8",
//	    "have_usable_ip6"
//	],
//	[
//	    "u8",
//	    "ip6_addr",
//	    16
//	],
//	{
//	    "crc": "0x6924ca6b"
//	}
//
type BfdUDPGetEchoSourceReply struct {
	Retval        int32
	SwIfIndex     uint32
	IsSet         uint8
	HaveUsableIP4 uint8
	IP4Addr       []byte `struc:"[4]byte"`
	HaveUsableIP6 uint8
	IP6Addr       []byte `struc:"[16]byte"`
}

func (*BfdUDPGetEchoSourceReply) GetMessageName() string {
	return "bfd_udp_get_echo_source_reply"
}
func (*BfdUDPGetEchoSourceReply) GetCrcString() string {
	return "6924ca6b"
}
func (*BfdUDPGetEchoSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPAdd represents VPP binary API message 'bfd_udp_add':
//
//	"bfd_udp_add",
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
//	    "desired_min_tx"
//	],
//	[
//	    "u32",
//	    "required_min_rx"
//	],
//	[
//	    "u8",
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "detect_mult"
//	],
//	[
//	    "u8",
//	    "is_authenticated"
//	],
//	[
//	    "u8",
//	    "bfd_key_id"
//	],
//	[
//	    "u32",
//	    "conf_key_id"
//	],
//	{
//	    "crc": "0x61cf1850"
//	}
//
type BfdUDPAdd struct {
	SwIfIndex       uint32
	DesiredMinTx    uint32
	RequiredMinRx   uint32
	LocalAddr       []byte `struc:"[16]byte"`
	PeerAddr        []byte `struc:"[16]byte"`
	IsIPv6          uint8
	DetectMult      uint8
	IsAuthenticated uint8
	BfdKeyID        uint8
	ConfKeyID       uint32
}

func (*BfdUDPAdd) GetMessageName() string {
	return "bfd_udp_add"
}
func (*BfdUDPAdd) GetCrcString() string {
	return "61cf1850"
}
func (*BfdUDPAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPAddReply represents VPP binary API message 'bfd_udp_add_reply':
//
//	"bfd_udp_add_reply",
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
type BfdUDPAddReply struct {
	Retval int32
}

func (*BfdUDPAddReply) GetMessageName() string {
	return "bfd_udp_add_reply"
}
func (*BfdUDPAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPMod represents VPP binary API message 'bfd_udp_mod':
//
//	"bfd_udp_mod",
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
//	    "desired_min_tx"
//	],
//	[
//	    "u32",
//	    "required_min_rx"
//	],
//	[
//	    "u8",
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "detect_mult"
//	],
//	{
//	    "crc": "0x6049bf47"
//	}
//
type BfdUDPMod struct {
	SwIfIndex     uint32
	DesiredMinTx  uint32
	RequiredMinRx uint32
	LocalAddr     []byte `struc:"[16]byte"`
	PeerAddr      []byte `struc:"[16]byte"`
	IsIPv6        uint8
	DetectMult    uint8
}

func (*BfdUDPMod) GetMessageName() string {
	return "bfd_udp_mod"
}
func (*BfdUDPMod) GetCrcString() string {
	return "6049bf47"
}
func (*BfdUDPMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPModReply represents VPP binary API message 'bfd_udp_mod_reply':
//
//	"bfd_udp_mod_reply",
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
type BfdUDPModReply struct {
	Retval int32
}

func (*BfdUDPModReply) GetMessageName() string {
	return "bfd_udp_mod_reply"
}
func (*BfdUDPModReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPDel represents VPP binary API message 'bfd_udp_del':
//
//	"bfd_udp_del",
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
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	{
//	    "crc": "0xc9e9cc5a"
//	}
//
type BfdUDPDel struct {
	SwIfIndex uint32
	LocalAddr []byte `struc:"[16]byte"`
	PeerAddr  []byte `struc:"[16]byte"`
	IsIPv6    uint8
}

func (*BfdUDPDel) GetMessageName() string {
	return "bfd_udp_del"
}
func (*BfdUDPDel) GetCrcString() string {
	return "c9e9cc5a"
}
func (*BfdUDPDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPDelReply represents VPP binary API message 'bfd_udp_del_reply':
//
//	"bfd_udp_del_reply",
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
type BfdUDPDelReply struct {
	Retval int32
}

func (*BfdUDPDelReply) GetMessageName() string {
	return "bfd_udp_del_reply"
}
func (*BfdUDPDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPSessionDump represents VPP binary API message 'bfd_udp_session_dump':
//
//	"bfd_udp_session_dump",
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
type BfdUDPSessionDump struct{}

func (*BfdUDPSessionDump) GetMessageName() string {
	return "bfd_udp_session_dump"
}
func (*BfdUDPSessionDump) GetCrcString() string {
	return "51077d14"
}
func (*BfdUDPSessionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPSessionDetails represents VPP binary API message 'bfd_udp_session_details':
//
//	"bfd_udp_session_details",
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
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "state"
//	],
//	[
//	    "u8",
//	    "is_authenticated"
//	],
//	[
//	    "u8",
//	    "bfd_key_id"
//	],
//	[
//	    "u32",
//	    "conf_key_id"
//	],
//	[
//	    "u32",
//	    "required_min_rx"
//	],
//	[
//	    "u32",
//	    "desired_min_tx"
//	],
//	[
//	    "u8",
//	    "detect_mult"
//	],
//	{
//	    "crc": "0x837bb0ed"
//	}
//
type BfdUDPSessionDetails struct {
	SwIfIndex       uint32
	LocalAddr       []byte `struc:"[16]byte"`
	PeerAddr        []byte `struc:"[16]byte"`
	IsIPv6          uint8
	State           uint8
	IsAuthenticated uint8
	BfdKeyID        uint8
	ConfKeyID       uint32
	RequiredMinRx   uint32
	DesiredMinTx    uint32
	DetectMult      uint8
}

func (*BfdUDPSessionDetails) GetMessageName() string {
	return "bfd_udp_session_details"
}
func (*BfdUDPSessionDetails) GetCrcString() string {
	return "837bb0ed"
}
func (*BfdUDPSessionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPSessionSetFlags represents VPP binary API message 'bfd_udp_session_set_flags':
//
//	"bfd_udp_session_set_flags",
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
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "admin_up_down"
//	],
//	{
//	    "crc": "0x667d6e7c"
//	}
//
type BfdUDPSessionSetFlags struct {
	SwIfIndex   uint32
	LocalAddr   []byte `struc:"[16]byte"`
	PeerAddr    []byte `struc:"[16]byte"`
	IsIPv6      uint8
	AdminUpDown uint8
}

func (*BfdUDPSessionSetFlags) GetMessageName() string {
	return "bfd_udp_session_set_flags"
}
func (*BfdUDPSessionSetFlags) GetCrcString() string {
	return "667d6e7c"
}
func (*BfdUDPSessionSetFlags) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPSessionSetFlagsReply represents VPP binary API message 'bfd_udp_session_set_flags_reply':
//
//	"bfd_udp_session_set_flags_reply",
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
type BfdUDPSessionSetFlagsReply struct {
	Retval int32
}

func (*BfdUDPSessionSetFlagsReply) GetMessageName() string {
	return "bfd_udp_session_set_flags_reply"
}
func (*BfdUDPSessionSetFlagsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPSessionSetFlagsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantBfdEvents represents VPP binary API message 'want_bfd_events':
//
//	"want_bfd_events",
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
//	    "enable_disable"
//	],
//	[
//	    "u32",
//	    "pid"
//	],
//	{
//	    "crc": "0x476f5a08"
//	}
//
type WantBfdEvents struct {
	EnableDisable uint32
	PID           uint32
}

func (*WantBfdEvents) GetMessageName() string {
	return "want_bfd_events"
}
func (*WantBfdEvents) GetCrcString() string {
	return "476f5a08"
}
func (*WantBfdEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantBfdEventsReply represents VPP binary API message 'want_bfd_events_reply':
//
//	"want_bfd_events_reply",
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
type WantBfdEventsReply struct {
	Retval int32
}

func (*WantBfdEventsReply) GetMessageName() string {
	return "want_bfd_events_reply"
}
func (*WantBfdEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantBfdEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdAuthSetKey represents VPP binary API message 'bfd_auth_set_key':
//
//	"bfd_auth_set_key",
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
//	    "conf_key_id"
//	],
//	[
//	    "u8",
//	    "key_len"
//	],
//	[
//	    "u8",
//	    "auth_type"
//	],
//	[
//	    "u8",
//	    "key",
//	    20
//	],
//	{
//	    "crc": "0x690b8877"
//	}
//
type BfdAuthSetKey struct {
	ConfKeyID uint32
	KeyLen    uint8
	AuthType  uint8
	Key       []byte `struc:"[20]byte"`
}

func (*BfdAuthSetKey) GetMessageName() string {
	return "bfd_auth_set_key"
}
func (*BfdAuthSetKey) GetCrcString() string {
	return "690b8877"
}
func (*BfdAuthSetKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdAuthSetKeyReply represents VPP binary API message 'bfd_auth_set_key_reply':
//
//	"bfd_auth_set_key_reply",
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
type BfdAuthSetKeyReply struct {
	Retval int32
}

func (*BfdAuthSetKeyReply) GetMessageName() string {
	return "bfd_auth_set_key_reply"
}
func (*BfdAuthSetKeyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdAuthSetKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdAuthDelKey represents VPP binary API message 'bfd_auth_del_key':
//
//	"bfd_auth_del_key",
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
//	    "conf_key_id"
//	],
//	{
//	    "crc": "0x65310b22"
//	}
//
type BfdAuthDelKey struct {
	ConfKeyID uint32
}

func (*BfdAuthDelKey) GetMessageName() string {
	return "bfd_auth_del_key"
}
func (*BfdAuthDelKey) GetCrcString() string {
	return "65310b22"
}
func (*BfdAuthDelKey) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdAuthDelKeyReply represents VPP binary API message 'bfd_auth_del_key_reply':
//
//	"bfd_auth_del_key_reply",
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
type BfdAuthDelKeyReply struct {
	Retval int32
}

func (*BfdAuthDelKeyReply) GetMessageName() string {
	return "bfd_auth_del_key_reply"
}
func (*BfdAuthDelKeyReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdAuthDelKeyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdAuthKeysDump represents VPP binary API message 'bfd_auth_keys_dump':
//
//	"bfd_auth_keys_dump",
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
type BfdAuthKeysDump struct{}

func (*BfdAuthKeysDump) GetMessageName() string {
	return "bfd_auth_keys_dump"
}
func (*BfdAuthKeysDump) GetCrcString() string {
	return "51077d14"
}
func (*BfdAuthKeysDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdAuthKeysDetails represents VPP binary API message 'bfd_auth_keys_details':
//
//	"bfd_auth_keys_details",
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
//	    "conf_key_id"
//	],
//	[
//	    "u32",
//	    "use_count"
//	],
//	[
//	    "u8",
//	    "auth_type"
//	],
//	{
//	    "crc": "0x84130e9f"
//	}
//
type BfdAuthKeysDetails struct {
	ConfKeyID uint32
	UseCount  uint32
	AuthType  uint8
}

func (*BfdAuthKeysDetails) GetMessageName() string {
	return "bfd_auth_keys_details"
}
func (*BfdAuthKeysDetails) GetCrcString() string {
	return "84130e9f"
}
func (*BfdAuthKeysDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPAuthActivate represents VPP binary API message 'bfd_udp_auth_activate':
//
//	"bfd_udp_auth_activate",
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
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "is_delayed"
//	],
//	[
//	    "u8",
//	    "bfd_key_id"
//	],
//	[
//	    "u32",
//	    "conf_key_id"
//	],
//	{
//	    "crc": "0x1bae0947"
//	}
//
type BfdUDPAuthActivate struct {
	SwIfIndex uint32
	LocalAddr []byte `struc:"[16]byte"`
	PeerAddr  []byte `struc:"[16]byte"`
	IsIPv6    uint8
	IsDelayed uint8
	BfdKeyID  uint8
	ConfKeyID uint32
}

func (*BfdUDPAuthActivate) GetMessageName() string {
	return "bfd_udp_auth_activate"
}
func (*BfdUDPAuthActivate) GetCrcString() string {
	return "1bae0947"
}
func (*BfdUDPAuthActivate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPAuthActivateReply represents VPP binary API message 'bfd_udp_auth_activate_reply':
//
//	"bfd_udp_auth_activate_reply",
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
type BfdUDPAuthActivateReply struct {
	Retval int32
}

func (*BfdUDPAuthActivateReply) GetMessageName() string {
	return "bfd_udp_auth_activate_reply"
}
func (*BfdUDPAuthActivateReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPAuthActivateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BfdUDPAuthDeactivate represents VPP binary API message 'bfd_udp_auth_deactivate':
//
//	"bfd_udp_auth_deactivate",
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
//	    "local_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "peer_addr",
//	    16
//	],
//	[
//	    "u8",
//	    "is_ipv6"
//	],
//	[
//	    "u8",
//	    "is_delayed"
//	],
//	{
//	    "crc": "0x10661991"
//	}
//
type BfdUDPAuthDeactivate struct {
	SwIfIndex uint32
	LocalAddr []byte `struc:"[16]byte"`
	PeerAddr  []byte `struc:"[16]byte"`
	IsIPv6    uint8
	IsDelayed uint8
}

func (*BfdUDPAuthDeactivate) GetMessageName() string {
	return "bfd_udp_auth_deactivate"
}
func (*BfdUDPAuthDeactivate) GetCrcString() string {
	return "10661991"
}
func (*BfdUDPAuthDeactivate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BfdUDPAuthDeactivateReply represents VPP binary API message 'bfd_udp_auth_deactivate_reply':
//
//	"bfd_udp_auth_deactivate_reply",
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
type BfdUDPAuthDeactivateReply struct {
	Retval int32
}

func (*BfdUDPAuthDeactivateReply) GetMessageName() string {
	return "bfd_udp_auth_deactivate_reply"
}
func (*BfdUDPAuthDeactivateReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BfdUDPAuthDeactivateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*BfdUDPSetEchoSource)(nil), "bfd.BfdUDPSetEchoSource")
	api.RegisterMessage((*BfdUDPSetEchoSourceReply)(nil), "bfd.BfdUDPSetEchoSourceReply")
	api.RegisterMessage((*BfdUDPDelEchoSource)(nil), "bfd.BfdUDPDelEchoSource")
	api.RegisterMessage((*BfdUDPDelEchoSourceReply)(nil), "bfd.BfdUDPDelEchoSourceReply")
	api.RegisterMessage((*BfdUDPGetEchoSource)(nil), "bfd.BfdUDPGetEchoSource")
	api.RegisterMessage((*BfdUDPGetEchoSourceReply)(nil), "bfd.BfdUDPGetEchoSourceReply")
	api.RegisterMessage((*BfdUDPAdd)(nil), "bfd.BfdUDPAdd")
	api.RegisterMessage((*BfdUDPAddReply)(nil), "bfd.BfdUDPAddReply")
	api.RegisterMessage((*BfdUDPMod)(nil), "bfd.BfdUDPMod")
	api.RegisterMessage((*BfdUDPModReply)(nil), "bfd.BfdUDPModReply")
	api.RegisterMessage((*BfdUDPDel)(nil), "bfd.BfdUDPDel")
	api.RegisterMessage((*BfdUDPDelReply)(nil), "bfd.BfdUDPDelReply")
	api.RegisterMessage((*BfdUDPSessionDump)(nil), "bfd.BfdUDPSessionDump")
	api.RegisterMessage((*BfdUDPSessionDetails)(nil), "bfd.BfdUDPSessionDetails")
	api.RegisterMessage((*BfdUDPSessionSetFlags)(nil), "bfd.BfdUDPSessionSetFlags")
	api.RegisterMessage((*BfdUDPSessionSetFlagsReply)(nil), "bfd.BfdUDPSessionSetFlagsReply")
	api.RegisterMessage((*WantBfdEvents)(nil), "bfd.WantBfdEvents")
	api.RegisterMessage((*WantBfdEventsReply)(nil), "bfd.WantBfdEventsReply")
	api.RegisterMessage((*BfdAuthSetKey)(nil), "bfd.BfdAuthSetKey")
	api.RegisterMessage((*BfdAuthSetKeyReply)(nil), "bfd.BfdAuthSetKeyReply")
	api.RegisterMessage((*BfdAuthDelKey)(nil), "bfd.BfdAuthDelKey")
	api.RegisterMessage((*BfdAuthDelKeyReply)(nil), "bfd.BfdAuthDelKeyReply")
	api.RegisterMessage((*BfdAuthKeysDump)(nil), "bfd.BfdAuthKeysDump")
	api.RegisterMessage((*BfdAuthKeysDetails)(nil), "bfd.BfdAuthKeysDetails")
	api.RegisterMessage((*BfdUDPAuthActivate)(nil), "bfd.BfdUDPAuthActivate")
	api.RegisterMessage((*BfdUDPAuthActivateReply)(nil), "bfd.BfdUDPAuthActivateReply")
	api.RegisterMessage((*BfdUDPAuthDeactivate)(nil), "bfd.BfdUDPAuthDeactivate")
	api.RegisterMessage((*BfdUDPAuthDeactivateReply)(nil), "bfd.BfdUDPAuthDeactivateReply")
}
