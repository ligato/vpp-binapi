// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: binapi/bier.api.json

/*
 Package bier is a generated from VPP binary API module 'bier'.

 It contains following objects:
	 22 messages
	  3 types
	 11 services

*/
package bier

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
//	    "bier_imp_add": {
//	        "reply": "bier_imp_add_reply"
//	    },
//	    "bier_table_dump": {
//	        "reply": "bier_table_details",
//	        "stream": true
//	    },
//	    "bier_imp_dump": {
//	        "reply": "bier_imp_details",
//	        "stream": true
//	    },
//	    "bier_disp_entry_add_del": {
//	        "reply": "bier_disp_entry_add_del_reply"
//	    },
//	    "bier_imp_del": {
//	        "reply": "bier_imp_del_reply"
//	    },
//	    "bier_disp_table_dump": {
//	        "reply": "bier_disp_table_details",
//	        "stream": true
//	    },
//	    "bier_table_add_del": {
//	        "reply": "bier_table_add_del_reply"
//	    },
//	    "bier_route_add_del": {
//	        "reply": "bier_route_add_del_reply"
//	    },
//	    "bier_route_dump": {
//	        "reply": "bier_route_details",
//	        "stream": true
//	    },
//	    "bier_disp_table_add_del": {
//	        "reply": "bier_disp_table_add_del_reply"
//	    },
//	    "bier_disp_entry_dump": {
//	        "reply": "bier_disp_entry_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpBierDispEntry(*BierDispEntryDump) ([]*BierDispEntryDetails, error)
	DumpBierDispTable(*BierDispTableDump) ([]*BierDispTableDetails, error)
	DumpBierImp(*BierImpDump) ([]*BierImpDetails, error)
	DumpBierRoute(*BierRouteDump) ([]*BierRouteDetails, error)
	DumpBierTable(*BierTableDump) ([]*BierTableDetails, error)
	BierDispEntryAddDel(*BierDispEntryAddDel) (*BierDispEntryAddDelReply, error)
	BierDispTableAddDel(*BierDispTableAddDel) (*BierDispTableAddDelReply, error)
	BierImpAdd(*BierImpAdd) (*BierImpAddReply, error)
	BierImpDel(*BierImpDel) (*BierImpDelReply, error)
	BierRouteAddDel(*BierRouteAddDel) (*BierRouteAddDelReply, error)
	BierTableAddDel(*BierTableAddDel) (*BierTableAddDelReply, error)
}

/* Types */

// FibMplsLabel represents VPP binary API type 'fib_mpls_label':
//
//	"fib_mpls_label",
//	[
//	    "u8",
//	    "is_uniform"
//	],
//	[
//	    "u32",
//	    "label"
//	],
//	[
//	    "u8",
//	    "ttl"
//	],
//	[
//	    "u8",
//	    "exp"
//	],
//	{
//	    "crc": "0xc93bf35c"
//	}
//
type FibMplsLabel struct {
	IsUniform uint8
	Label     uint32
	TTL       uint8
	Exp       uint8
}

func (*FibMplsLabel) GetTypeName() string {
	return "fib_mpls_label"
}
func (*FibMplsLabel) GetCrcString() string {
	return "c93bf35c"
}

// FibPath represents VPP binary API type 'fib_path':
//
//	"fib_path",
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u32",
//	    "table_id"
//	],
//	[
//	    "u8",
//	    "weight"
//	],
//	[
//	    "u8",
//	    "preference"
//	],
//	[
//	    "u8",
//	    "is_local"
//	],
//	[
//	    "u8",
//	    "is_drop"
//	],
//	[
//	    "u8",
//	    "is_udp_encap"
//	],
//	[
//	    "u8",
//	    "is_unreach"
//	],
//	[
//	    "u8",
//	    "is_prohibit"
//	],
//	[
//	    "u8",
//	    "is_resolve_host"
//	],
//	[
//	    "u8",
//	    "is_resolve_attached"
//	],
//	[
//	    "u8",
//	    "is_dvr"
//	],
//	[
//	    "u8",
//	    "is_source_lookup"
//	],
//	[
//	    "u8",
//	    "is_interface_rx"
//	],
//	[
//	    "u8",
//	    "afi"
//	],
//	[
//	    "u8",
//	    "next_hop",
//	    16
//	],
//	[
//	    "u32",
//	    "next_hop_id"
//	],
//	[
//	    "u32",
//	    "rpf_id"
//	],
//	[
//	    "u32",
//	    "via_label"
//	],
//	[
//	    "u8",
//	    "n_labels"
//	],
//	[
//	    "vl_api_fib_mpls_label_t",
//	    "label_stack",
//	    16
//	],
//	{
//	    "crc": "0xba7a81f0"
//	}
//
type FibPath struct {
	SwIfIndex         uint32
	TableID           uint32
	Weight            uint8
	Preference        uint8
	IsLocal           uint8
	IsDrop            uint8
	IsUDPEncap        uint8
	IsUnreach         uint8
	IsProhibit        uint8
	IsResolveHost     uint8
	IsResolveAttached uint8
	IsDvr             uint8
	IsSourceLookup    uint8
	IsInterfaceRx     uint8
	Afi               uint8
	NextHop           []byte `struc:"[16]byte"`
	NextHopID         uint32
	RpfID             uint32
	ViaLabel          uint32
	NLabels           uint8
	LabelStack        []FibMplsLabel `struc:"[16]FibMplsLabel"`
}

func (*FibPath) GetTypeName() string {
	return "fib_path"
}
func (*FibPath) GetCrcString() string {
	return "ba7a81f0"
}

// BierTableID represents VPP binary API type 'bier_table_id':
//
//	"bier_table_id",
//	[
//	    "u8",
//	    "bt_set"
//	],
//	[
//	    "u8",
//	    "bt_sub_domain"
//	],
//	[
//	    "u8",
//	    "bt_hdr_len_id"
//	],
//	{
//	    "crc": "0x435c691d"
//	}
//
type BierTableID struct {
	BtSet       uint8
	BtSubDomain uint8
	BtHdrLenID  uint8
}

func (*BierTableID) GetTypeName() string {
	return "bier_table_id"
}
func (*BierTableID) GetCrcString() string {
	return "435c691d"
}

/* Messages */

// BierTableAddDel represents VPP binary API message 'bier_table_add_del':
//
//	"bier_table_add_del",
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
//	    "vl_api_bier_table_id_t",
//	    "bt_tbl_id"
//	],
//	[
//	    "u32",
//	    "bt_label"
//	],
//	[
//	    "u8",
//	    "bt_is_add"
//	],
//	{
//	    "crc": "0x4f7963b6"
//	}
//
type BierTableAddDel struct {
	BtTblID BierTableID
	BtLabel uint32
	BtIsAdd uint8
}

func (*BierTableAddDel) GetMessageName() string {
	return "bier_table_add_del"
}
func (*BierTableAddDel) GetCrcString() string {
	return "4f7963b6"
}
func (*BierTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierTableAddDelReply represents VPP binary API message 'bier_table_add_del_reply':
//
//	"bier_table_add_del_reply",
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
type BierTableAddDelReply struct {
	Retval int32
}

func (*BierTableAddDelReply) GetMessageName() string {
	return "bier_table_add_del_reply"
}
func (*BierTableAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierTableDump represents VPP binary API message 'bier_table_dump':
//
//	"bier_table_dump",
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
type BierTableDump struct{}

func (*BierTableDump) GetMessageName() string {
	return "bier_table_dump"
}
func (*BierTableDump) GetCrcString() string {
	return "51077d14"
}
func (*BierTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierTableDetails represents VPP binary API message 'bier_table_details':
//
//	"bier_table_details",
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
//	    "bt_label"
//	],
//	[
//	    "vl_api_bier_table_id_t",
//	    "bt_tbl_id"
//	],
//	{
//	    "crc": "0xb7a41e40"
//	}
//
type BierTableDetails struct {
	BtLabel uint32
	BtTblID BierTableID
}

func (*BierTableDetails) GetMessageName() string {
	return "bier_table_details"
}
func (*BierTableDetails) GetCrcString() string {
	return "b7a41e40"
}
func (*BierTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierRouteAddDel represents VPP binary API message 'bier_route_add_del':
//
//	"bier_route_add_del",
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
//	    "br_bp"
//	],
//	[
//	    "u8",
//	    "br_is_add"
//	],
//	[
//	    "u8",
//	    "br_is_replace"
//	],
//	[
//	    "vl_api_bier_table_id_t",
//	    "br_tbl_id"
//	],
//	[
//	    "u8",
//	    "br_n_paths"
//	],
//	[
//	    "vl_api_fib_path_t",
//	    "br_paths",
//	    0,
//	    "br_n_paths"
//	],
//	{
//	    "crc": "0x0b83f5d2"
//	}
//
type BierRouteAddDel struct {
	BrBp        uint32
	BrIsAdd     uint8
	BrIsReplace uint8
	BrTblID     BierTableID
	BrNPaths    uint8 `struc:"sizeof=BrPaths"`
	BrPaths     []FibPath
}

func (*BierRouteAddDel) GetMessageName() string {
	return "bier_route_add_del"
}
func (*BierRouteAddDel) GetCrcString() string {
	return "0b83f5d2"
}
func (*BierRouteAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierRouteAddDelReply represents VPP binary API message 'bier_route_add_del_reply':
//
//	"bier_route_add_del_reply",
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
type BierRouteAddDelReply struct {
	Retval int32
}

func (*BierRouteAddDelReply) GetMessageName() string {
	return "bier_route_add_del_reply"
}
func (*BierRouteAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierRouteAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierRouteDump represents VPP binary API message 'bier_route_dump':
//
//	"bier_route_dump",
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
//	    "vl_api_bier_table_id_t",
//	    "br_tbl_id"
//	],
//	{
//	    "crc": "0x43af6370"
//	}
//
type BierRouteDump struct {
	BrTblID BierTableID
}

func (*BierRouteDump) GetMessageName() string {
	return "bier_route_dump"
}
func (*BierRouteDump) GetCrcString() string {
	return "43af6370"
}
func (*BierRouteDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierRouteDetails represents VPP binary API message 'bier_route_details':
//
//	"bier_route_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u16",
//	    "br_bp"
//	],
//	[
//	    "vl_api_bier_table_id_t",
//	    "br_tbl_id"
//	],
//	[
//	    "u32",
//	    "br_n_paths"
//	],
//	[
//	    "vl_api_fib_path_t",
//	    "br_paths",
//	    0,
//	    "br_n_paths"
//	],
//	{
//	    "crc": "0x82c4087b"
//	}
//
type BierRouteDetails struct {
	BrBp     uint16
	BrTblID  BierTableID
	BrNPaths uint32 `struc:"sizeof=BrPaths"`
	BrPaths  []FibPath
}

func (*BierRouteDetails) GetMessageName() string {
	return "bier_route_details"
}
func (*BierRouteDetails) GetCrcString() string {
	return "82c4087b"
}
func (*BierRouteDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierImpAdd represents VPP binary API message 'bier_imp_add':
//
//	"bier_imp_add",
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
//	    "vl_api_bier_table_id_t",
//	    "bi_tbl_id"
//	],
//	[
//	    "u16",
//	    "bi_src"
//	],
//	[
//	    "u8",
//	    "bi_n_bytes"
//	],
//	[
//	    "u8",
//	    "bi_bytes",
//	    0,
//	    "bi_n_bytes"
//	],
//	{
//	    "crc": "0x5d62e9ce"
//	}
//
type BierImpAdd struct {
	BiTblID  BierTableID
	BiSrc    uint16
	BiNBytes uint8 `struc:"sizeof=BiBytes"`
	BiBytes  []byte
}

func (*BierImpAdd) GetMessageName() string {
	return "bier_imp_add"
}
func (*BierImpAdd) GetCrcString() string {
	return "5d62e9ce"
}
func (*BierImpAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierImpAddReply represents VPP binary API message 'bier_imp_add_reply':
//
//	"bier_imp_add_reply",
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
//	    "bi_index"
//	],
//	{
//	    "crc": "0xd49c5793"
//	}
//
type BierImpAddReply struct {
	Retval  int32
	BiIndex uint32
}

func (*BierImpAddReply) GetMessageName() string {
	return "bier_imp_add_reply"
}
func (*BierImpAddReply) GetCrcString() string {
	return "d49c5793"
}
func (*BierImpAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierImpDel represents VPP binary API message 'bier_imp_del':
//
//	"bier_imp_del",
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
//	    "bi_index"
//	],
//	{
//	    "crc": "0x7d45edf6"
//	}
//
type BierImpDel struct {
	BiIndex uint32
}

func (*BierImpDel) GetMessageName() string {
	return "bier_imp_del"
}
func (*BierImpDel) GetCrcString() string {
	return "7d45edf6"
}
func (*BierImpDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierImpDelReply represents VPP binary API message 'bier_imp_del_reply':
//
//	"bier_imp_del_reply",
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
type BierImpDelReply struct {
	Retval int32
}

func (*BierImpDelReply) GetMessageName() string {
	return "bier_imp_del_reply"
}
func (*BierImpDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierImpDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierImpDump represents VPP binary API message 'bier_imp_dump':
//
//	"bier_imp_dump",
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
type BierImpDump struct{}

func (*BierImpDump) GetMessageName() string {
	return "bier_imp_dump"
}
func (*BierImpDump) GetCrcString() string {
	return "51077d14"
}
func (*BierImpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierImpDetails represents VPP binary API message 'bier_imp_details':
//
//	"bier_imp_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "vl_api_bier_table_id_t",
//	    "bi_tbl_id"
//	],
//	[
//	    "u16",
//	    "bi_src"
//	],
//	[
//	    "u8",
//	    "bi_n_bytes"
//	],
//	[
//	    "u8",
//	    "bi_bytes",
//	    0,
//	    "bi_n_bytes"
//	],
//	{
//	    "crc": "0x104b1f23"
//	}
//
type BierImpDetails struct {
	BiTblID  BierTableID
	BiSrc    uint16
	BiNBytes uint8 `struc:"sizeof=BiBytes"`
	BiBytes  []byte
}

func (*BierImpDetails) GetMessageName() string {
	return "bier_imp_details"
}
func (*BierImpDetails) GetCrcString() string {
	return "104b1f23"
}
func (*BierImpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispTableAddDel represents VPP binary API message 'bier_disp_table_add_del':
//
//	"bier_disp_table_add_del",
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
//	    "bdt_tbl_id"
//	],
//	[
//	    "u8",
//	    "bdt_is_add"
//	],
//	{
//	    "crc": "0x7671b2cb"
//	}
//
type BierDispTableAddDel struct {
	BdtTblID uint32
	BdtIsAdd uint8
}

func (*BierDispTableAddDel) GetMessageName() string {
	return "bier_disp_table_add_del"
}
func (*BierDispTableAddDel) GetCrcString() string {
	return "7671b2cb"
}
func (*BierDispTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispTableAddDelReply represents VPP binary API message 'bier_disp_table_add_del_reply':
//
//	"bier_disp_table_add_del_reply",
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
type BierDispTableAddDelReply struct {
	Retval int32
}

func (*BierDispTableAddDelReply) GetMessageName() string {
	return "bier_disp_table_add_del_reply"
}
func (*BierDispTableAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierDispTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispTableDump represents VPP binary API message 'bier_disp_table_dump':
//
//	"bier_disp_table_dump",
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
type BierDispTableDump struct{}

func (*BierDispTableDump) GetMessageName() string {
	return "bier_disp_table_dump"
}
func (*BierDispTableDump) GetCrcString() string {
	return "51077d14"
}
func (*BierDispTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispTableDetails represents VPP binary API message 'bier_disp_table_details':
//
//	"bier_disp_table_details",
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
//	    "bdt_tbl_id"
//	],
//	{
//	    "crc": "0xd27942c0"
//	}
//
type BierDispTableDetails struct {
	BdtTblID uint32
}

func (*BierDispTableDetails) GetMessageName() string {
	return "bier_disp_table_details"
}
func (*BierDispTableDetails) GetCrcString() string {
	return "d27942c0"
}
func (*BierDispTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispEntryAddDel represents VPP binary API message 'bier_disp_entry_add_del':
//
//	"bier_disp_entry_add_del",
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
//	    "u16",
//	    "bde_bp"
//	],
//	[
//	    "u32",
//	    "bde_tbl_id"
//	],
//	[
//	    "u8",
//	    "bde_is_add"
//	],
//	[
//	    "u8",
//	    "bde_payload_proto"
//	],
//	[
//	    "u8",
//	    "bde_n_paths"
//	],
//	[
//	    "vl_api_fib_path_t",
//	    "bde_paths",
//	    0,
//	    "bde_n_paths"
//	],
//	{
//	    "crc": "0x4260a077"
//	}
//
type BierDispEntryAddDel struct {
	BdeBp           uint16
	BdeTblID        uint32
	BdeIsAdd        uint8
	BdePayloadProto uint8
	BdeNPaths       uint8 `struc:"sizeof=BdePaths"`
	BdePaths        []FibPath
}

func (*BierDispEntryAddDel) GetMessageName() string {
	return "bier_disp_entry_add_del"
}
func (*BierDispEntryAddDel) GetCrcString() string {
	return "4260a077"
}
func (*BierDispEntryAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispEntryAddDelReply represents VPP binary API message 'bier_disp_entry_add_del_reply':
//
//	"bier_disp_entry_add_del_reply",
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
type BierDispEntryAddDelReply struct {
	Retval int32
}

func (*BierDispEntryAddDelReply) GetMessageName() string {
	return "bier_disp_entry_add_del_reply"
}
func (*BierDispEntryAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BierDispEntryAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BierDispEntryDump represents VPP binary API message 'bier_disp_entry_dump':
//
//	"bier_disp_entry_dump",
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
//	    "bde_tbl_id"
//	],
//	{
//	    "crc": "0xb5fa54ad"
//	}
//
type BierDispEntryDump struct {
	BdeTblID uint32
}

func (*BierDispEntryDump) GetMessageName() string {
	return "bier_disp_entry_dump"
}
func (*BierDispEntryDump) GetCrcString() string {
	return "b5fa54ad"
}
func (*BierDispEntryDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BierDispEntryDetails represents VPP binary API message 'bier_disp_entry_details':
//
//	"bier_disp_entry_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u16",
//	    "bde_bp"
//	],
//	[
//	    "u32",
//	    "bde_tbl_id"
//	],
//	[
//	    "u8",
//	    "bde_is_add"
//	],
//	[
//	    "u8",
//	    "bde_payload_proto"
//	],
//	[
//	    "u8",
//	    "bde_n_paths"
//	],
//	[
//	    "vl_api_fib_path_t",
//	    "bde_paths",
//	    0,
//	    "bde_n_paths"
//	],
//	{
//	    "crc": "0xb9b10a59"
//	}
//
type BierDispEntryDetails struct {
	BdeBp           uint16
	BdeTblID        uint32
	BdeIsAdd        uint8
	BdePayloadProto uint8
	BdeNPaths       uint8 `struc:"sizeof=BdePaths"`
	BdePaths        []FibPath
}

func (*BierDispEntryDetails) GetMessageName() string {
	return "bier_disp_entry_details"
}
func (*BierDispEntryDetails) GetCrcString() string {
	return "b9b10a59"
}
func (*BierDispEntryDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*BierTableAddDel)(nil), "bier.BierTableAddDel")
	api.RegisterMessage((*BierTableAddDelReply)(nil), "bier.BierTableAddDelReply")
	api.RegisterMessage((*BierTableDump)(nil), "bier.BierTableDump")
	api.RegisterMessage((*BierTableDetails)(nil), "bier.BierTableDetails")
	api.RegisterMessage((*BierRouteAddDel)(nil), "bier.BierRouteAddDel")
	api.RegisterMessage((*BierRouteAddDelReply)(nil), "bier.BierRouteAddDelReply")
	api.RegisterMessage((*BierRouteDump)(nil), "bier.BierRouteDump")
	api.RegisterMessage((*BierRouteDetails)(nil), "bier.BierRouteDetails")
	api.RegisterMessage((*BierImpAdd)(nil), "bier.BierImpAdd")
	api.RegisterMessage((*BierImpAddReply)(nil), "bier.BierImpAddReply")
	api.RegisterMessage((*BierImpDel)(nil), "bier.BierImpDel")
	api.RegisterMessage((*BierImpDelReply)(nil), "bier.BierImpDelReply")
	api.RegisterMessage((*BierImpDump)(nil), "bier.BierImpDump")
	api.RegisterMessage((*BierImpDetails)(nil), "bier.BierImpDetails")
	api.RegisterMessage((*BierDispTableAddDel)(nil), "bier.BierDispTableAddDel")
	api.RegisterMessage((*BierDispTableAddDelReply)(nil), "bier.BierDispTableAddDelReply")
	api.RegisterMessage((*BierDispTableDump)(nil), "bier.BierDispTableDump")
	api.RegisterMessage((*BierDispTableDetails)(nil), "bier.BierDispTableDetails")
	api.RegisterMessage((*BierDispEntryAddDel)(nil), "bier.BierDispEntryAddDel")
	api.RegisterMessage((*BierDispEntryAddDelReply)(nil), "bier.BierDispEntryAddDelReply")
	api.RegisterMessage((*BierDispEntryDump)(nil), "bier.BierDispEntryDump")
	api.RegisterMessage((*BierDispEntryDetails)(nil), "bier.BierDispEntryDetails")
}
