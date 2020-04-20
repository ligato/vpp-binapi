// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/master/vppapi/plugins/rdma.api.json

/*
Package rdma is a generated VPP binary API for 'rdma' module.

It consists of:
	  7 enums
	  1 alias
	  4 messages
	  2 services
*/
package rdma

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "rdma"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x5ce233e0
)

// IfStatusFlags represents VPP binary API enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var IfStatusFlags_name = map[uint32]string{
	1: "IF_STATUS_API_FLAG_ADMIN_UP",
	2: "IF_STATUS_API_FLAG_LINK_UP",
}

var IfStatusFlags_value = map[string]uint32{
	"IF_STATUS_API_FLAG_ADMIN_UP": 1,
	"IF_STATUS_API_FLAG_LINK_UP":  2,
}

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IfType represents VPP binary API enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var IfType_name = map[uint32]string{
	1: "IF_API_TYPE_HARDWARE",
	2: "IF_API_TYPE_SUB",
	3: "IF_API_TYPE_P2P",
	4: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 1,
	"IF_API_TYPE_SUB":      2,
	"IF_API_TYPE_P2P":      3,
	"IF_API_TYPE_PIPE":     4,
}

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LinkDuplex represents VPP binary API enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var LinkDuplex_name = map[uint32]string{
	0: "LINK_DUPLEX_API_UNKNOWN",
	1: "LINK_DUPLEX_API_HALF",
	2: "LINK_DUPLEX_API_FULL",
}

var LinkDuplex_value = map[string]uint32{
	"LINK_DUPLEX_API_UNKNOWN": 0,
	"LINK_DUPLEX_API_HALF":    1,
	"LINK_DUPLEX_API_FULL":    2,
}

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var MtuProto_name = map[uint32]string{
	1: "MTU_PROTO_API_L3",
	2: "MTU_PROTO_API_IP4",
	3: "MTU_PROTO_API_IP6",
	4: "MTU_PROTO_API_MPLS",
	5: "MTU_PROTO_API_N",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   1,
	"MTU_PROTO_API_IP4":  2,
	"MTU_PROTO_API_IP6":  3,
	"MTU_PROTO_API_MPLS": 4,
	"MTU_PROTO_API_N":    5,
}

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RdmaMode represents VPP binary API enum 'rdma_mode'.
type RdmaMode uint32

const (
	RDMA_API_MODE_AUTO RdmaMode = 0
	RDMA_API_MODE_IBV  RdmaMode = 1
	RDMA_API_MODE_DV   RdmaMode = 2
)

var RdmaMode_name = map[uint32]string{
	0: "RDMA_API_MODE_AUTO",
	1: "RDMA_API_MODE_IBV",
	2: "RDMA_API_MODE_DV",
}

var RdmaMode_value = map[string]uint32{
	"RDMA_API_MODE_AUTO": 0,
	"RDMA_API_MODE_IBV":  1,
	"RDMA_API_MODE_DV":   2,
}

func (x RdmaMode) String() string {
	s, ok := RdmaMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// RxMode represents VPP binary API enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var RxMode_name = map[uint32]string{
	0: "RX_MODE_API_UNKNOWN",
	1: "RX_MODE_API_POLLING",
	2: "RX_MODE_API_INTERRUPT",
	3: "RX_MODE_API_ADAPTIVE",
	4: "RX_MODE_API_DEFAULT",
}

var RxMode_value = map[string]uint32{
	"RX_MODE_API_UNKNOWN":   0,
	"RX_MODE_API_POLLING":   1,
	"RX_MODE_API_INTERRUPT": 2,
	"RX_MODE_API_ADAPTIVE":  3,
	"RX_MODE_API_DEFAULT":   4,
}

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SubIfFlags represents VPP binary API enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var SubIfFlags_name = map[uint32]string{
	1:   "SUB_IF_API_FLAG_NO_TAGS",
	2:   "SUB_IF_API_FLAG_ONE_TAG",
	4:   "SUB_IF_API_FLAG_TWO_TAGS",
	8:   "SUB_IF_API_FLAG_DOT1AD",
	16:  "SUB_IF_API_FLAG_EXACT_MATCH",
	32:  "SUB_IF_API_FLAG_DEFAULT",
	64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
	128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
	254: "SUB_IF_API_FLAG_MASK_VNET",
	256: "SUB_IF_API_FLAG_DOT1AH",
}

var SubIfFlags_value = map[string]uint32{
	"SUB_IF_API_FLAG_NO_TAGS":           1,
	"SUB_IF_API_FLAG_ONE_TAG":           2,
	"SUB_IF_API_FLAG_TWO_TAGS":          4,
	"SUB_IF_API_FLAG_DOT1AD":            8,
	"SUB_IF_API_FLAG_EXACT_MATCH":       16,
	"SUB_IF_API_FLAG_DEFAULT":           32,
	"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
	"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
	"SUB_IF_API_FLAG_MASK_VNET":         254,
	"SUB_IF_API_FLAG_DOT1AH":            256,
}

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

// RdmaCreate represents VPP binary API message 'rdma_create'.
type RdmaCreate struct {
	HostIf  string `struc:"[64]byte"`
	Name    string `struc:"[64]byte"`
	RxqNum  uint16
	RxqSize uint16
	TxqSize uint16
	Mode    RdmaMode
}

func (m *RdmaCreate) Reset()                        { *m = RdmaCreate{} }
func (*RdmaCreate) GetMessageName() string          { return "rdma_create" }
func (*RdmaCreate) GetCrcString() string            { return "076fe418" }
func (*RdmaCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// RdmaCreateReply represents VPP binary API message 'rdma_create_reply'.
type RdmaCreateReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *RdmaCreateReply) Reset()                        { *m = RdmaCreateReply{} }
func (*RdmaCreateReply) GetMessageName() string          { return "rdma_create_reply" }
func (*RdmaCreateReply) GetCrcString() string            { return "5383d31f" }
func (*RdmaCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// RdmaDelete represents VPP binary API message 'rdma_delete'.
type RdmaDelete struct {
	SwIfIndex InterfaceIndex
}

func (m *RdmaDelete) Reset()                        { *m = RdmaDelete{} }
func (*RdmaDelete) GetMessageName() string          { return "rdma_delete" }
func (*RdmaDelete) GetCrcString() string            { return "f9e6675e" }
func (*RdmaDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// RdmaDeleteReply represents VPP binary API message 'rdma_delete_reply'.
type RdmaDeleteReply struct {
	Retval int32
}

func (m *RdmaDeleteReply) Reset()                        { *m = RdmaDeleteReply{} }
func (*RdmaDeleteReply) GetMessageName() string          { return "rdma_delete_reply" }
func (*RdmaDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*RdmaDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*RdmaCreate)(nil), "rdma.RdmaCreate")
	api.RegisterMessage((*RdmaCreateReply)(nil), "rdma.RdmaCreateReply")
	api.RegisterMessage((*RdmaDelete)(nil), "rdma.RdmaDelete")
	api.RegisterMessage((*RdmaDeleteReply)(nil), "rdma.RdmaDeleteReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*RdmaCreate)(nil),
		(*RdmaCreateReply)(nil),
		(*RdmaDelete)(nil),
		(*RdmaDeleteReply)(nil),
	}
}

// RPCService represents RPC service API for rdma module.
type RPCService interface {
	RdmaCreate(ctx context.Context, in *RdmaCreate) (*RdmaCreateReply, error)
	RdmaDelete(ctx context.Context, in *RdmaDelete) (*RdmaDeleteReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) RdmaCreate(ctx context.Context, in *RdmaCreate) (*RdmaCreateReply, error) {
	out := new(RdmaCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RdmaDelete(ctx context.Context, in *RdmaDelete) (*RdmaDeleteReply, error) {
	out := new(RdmaDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
