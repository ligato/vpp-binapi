// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/p2p_ethernet.api.json

/*
Package p2p_ethernet is a generated VPP binary API for 'p2p_ethernet' module.

It consists of:
	  6 enums
	  2 aliases
	  4 messages
	  2 services
*/
package p2p_ethernet

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
	ModuleName = "p2p_ethernet"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xe33357cb
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

// MacAddress represents VPP binary API alias 'mac_address'.
type MacAddress [6]uint8

// P2pEthernetAdd represents VPP binary API message 'p2p_ethernet_add'.
type P2pEthernetAdd struct {
	ParentIfIndex InterfaceIndex
	SubifID       uint32
	RemoteMac     MacAddress
}

func (*P2pEthernetAdd) GetMessageName() string {
	return "p2p_ethernet_add"
}
func (*P2pEthernetAdd) GetCrcString() string {
	return "85f5000b"
}
func (*P2pEthernetAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// P2pEthernetAddReply represents VPP binary API message 'p2p_ethernet_add_reply'.
type P2pEthernetAddReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (*P2pEthernetAddReply) GetMessageName() string {
	return "p2p_ethernet_add_reply"
}
func (*P2pEthernetAddReply) GetCrcString() string {
	return "903324db"
}
func (*P2pEthernetAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// P2pEthernetDel represents VPP binary API message 'p2p_ethernet_del'.
type P2pEthernetDel struct {
	ParentIfIndex InterfaceIndex
	RemoteMac     MacAddress
}

func (*P2pEthernetDel) GetMessageName() string {
	return "p2p_ethernet_del"
}
func (*P2pEthernetDel) GetCrcString() string {
	return "1f90f942"
}
func (*P2pEthernetDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// P2pEthernetDelReply represents VPP binary API message 'p2p_ethernet_del_reply'.
type P2pEthernetDelReply struct {
	Retval int32
}

func (*P2pEthernetDelReply) GetMessageName() string {
	return "p2p_ethernet_del_reply"
}
func (*P2pEthernetDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*P2pEthernetDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*P2pEthernetAdd)(nil), "p2p_ethernet.P2pEthernetAdd")
	api.RegisterMessage((*P2pEthernetAddReply)(nil), "p2p_ethernet.P2pEthernetAddReply")
	api.RegisterMessage((*P2pEthernetDel)(nil), "p2p_ethernet.P2pEthernetDel")
	api.RegisterMessage((*P2pEthernetDelReply)(nil), "p2p_ethernet.P2pEthernetDelReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*P2pEthernetAdd)(nil),
		(*P2pEthernetAddReply)(nil),
		(*P2pEthernetDel)(nil),
		(*P2pEthernetDelReply)(nil),
	}
}

// RPCService represents RPC service API for p2p_ethernet module.
type RPCService interface {
	P2pEthernetAdd(ctx context.Context, in *P2pEthernetAdd) (*P2pEthernetAddReply, error)
	P2pEthernetDel(ctx context.Context, in *P2pEthernetDel) (*P2pEthernetDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) P2pEthernetAdd(ctx context.Context, in *P2pEthernetAdd) (*P2pEthernetAddReply, error) {
	out := new(P2pEthernetAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) P2pEthernetDel(ctx context.Context, in *P2pEthernetDel) (*P2pEthernetDelReply, error) {
	out := new(P2pEthernetDelReply)
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
