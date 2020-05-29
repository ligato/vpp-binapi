// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/plugins/memif.api.json

/*
Package memif is a generated VPP binary API for 'memif' module.

It consists of:
	  8 enums
	  2 aliases
	 10 messages
	  5 services
*/
package memif

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "memif"
	// APIVersion is the API version of this module.
	APIVersion = "3.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x1a1c95b8
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
	IF_API_TYPE_HARDWARE IfType = 0
	IF_API_TYPE_SUB      IfType = 1
	IF_API_TYPE_P2P      IfType = 2
	IF_API_TYPE_PIPE     IfType = 3
)

var IfType_name = map[uint32]string{
	0: "IF_API_TYPE_HARDWARE",
	1: "IF_API_TYPE_SUB",
	2: "IF_API_TYPE_P2P",
	3: "IF_API_TYPE_PIPE",
}

var IfType_value = map[string]uint32{
	"IF_API_TYPE_HARDWARE": 0,
	"IF_API_TYPE_SUB":      1,
	"IF_API_TYPE_P2P":      2,
	"IF_API_TYPE_PIPE":     3,
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

// MemifMode represents VPP binary API enum 'memif_mode'.
type MemifMode uint32

const (
	MEMIF_MODE_API_ETHERNET    MemifMode = 0
	MEMIF_MODE_API_IP          MemifMode = 1
	MEMIF_MODE_API_PUNT_INJECT MemifMode = 2
)

var MemifMode_name = map[uint32]string{
	0: "MEMIF_MODE_API_ETHERNET",
	1: "MEMIF_MODE_API_IP",
	2: "MEMIF_MODE_API_PUNT_INJECT",
}

var MemifMode_value = map[string]uint32{
	"MEMIF_MODE_API_ETHERNET":    0,
	"MEMIF_MODE_API_IP":          1,
	"MEMIF_MODE_API_PUNT_INJECT": 2,
}

func (x MemifMode) String() string {
	s, ok := MemifMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MemifRole represents VPP binary API enum 'memif_role'.
type MemifRole uint32

const (
	MEMIF_ROLE_API_MASTER MemifRole = 0
	MEMIF_ROLE_API_SLAVE  MemifRole = 1
)

var MemifRole_name = map[uint32]string{
	0: "MEMIF_ROLE_API_MASTER",
	1: "MEMIF_ROLE_API_SLAVE",
}

var MemifRole_value = map[string]uint32{
	"MEMIF_ROLE_API_MASTER": 0,
	"MEMIF_ROLE_API_SLAVE":  1,
}

func (x MemifRole) String() string {
	s, ok := MemifRole_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// MtuProto represents VPP binary API enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 0
	MTU_PROTO_API_IP4  MtuProto = 1
	MTU_PROTO_API_IP6  MtuProto = 2
	MTU_PROTO_API_MPLS MtuProto = 3
)

var MtuProto_name = map[uint32]string{
	0: "MTU_PROTO_API_L3",
	1: "MTU_PROTO_API_IP4",
	2: "MTU_PROTO_API_IP6",
	3: "MTU_PROTO_API_MPLS",
}

var MtuProto_value = map[string]uint32{
	"MTU_PROTO_API_L3":   0,
	"MTU_PROTO_API_IP4":  1,
	"MTU_PROTO_API_IP6":  2,
	"MTU_PROTO_API_MPLS": 3,
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

// MemifCreate represents VPP binary API message 'memif_create'.
type MemifCreate struct {
	Role       MemifRole
	Mode       MemifMode
	RxQueues   uint8
	TxQueues   uint8
	ID         uint32
	SocketID   uint32
	RingSize   uint32
	BufferSize uint16
	NoZeroCopy bool
	HwAddr     MacAddress
	Secret     string `struc:"[24]byte"`
}

func (m *MemifCreate) Reset()                        { *m = MemifCreate{} }
func (*MemifCreate) GetMessageName() string          { return "memif_create" }
func (*MemifCreate) GetCrcString() string            { return "b1b25061" }
func (*MemifCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// MemifCreateReply represents VPP binary API message 'memif_create_reply'.
type MemifCreateReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *MemifCreateReply) Reset()                        { *m = MemifCreateReply{} }
func (*MemifCreateReply) GetMessageName() string          { return "memif_create_reply" }
func (*MemifCreateReply) GetCrcString() string            { return "5383d31f" }
func (*MemifCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MemifDelete represents VPP binary API message 'memif_delete'.
type MemifDelete struct {
	SwIfIndex InterfaceIndex
}

func (m *MemifDelete) Reset()                        { *m = MemifDelete{} }
func (*MemifDelete) GetMessageName() string          { return "memif_delete" }
func (*MemifDelete) GetCrcString() string            { return "f9e6675e" }
func (*MemifDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// MemifDeleteReply represents VPP binary API message 'memif_delete_reply'.
type MemifDeleteReply struct {
	Retval int32
}

func (m *MemifDeleteReply) Reset()                        { *m = MemifDeleteReply{} }
func (*MemifDeleteReply) GetMessageName() string          { return "memif_delete_reply" }
func (*MemifDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*MemifDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MemifDetails represents VPP binary API message 'memif_details'.
type MemifDetails struct {
	SwIfIndex  InterfaceIndex
	HwAddr     MacAddress
	ID         uint32
	Role       MemifRole
	Mode       MemifMode
	ZeroCopy   bool
	SocketID   uint32
	RingSize   uint32
	BufferSize uint16
	Flags      IfStatusFlags
	IfName     string `struc:"[64]byte"`
}

func (m *MemifDetails) Reset()                        { *m = MemifDetails{} }
func (*MemifDetails) GetMessageName() string          { return "memif_details" }
func (*MemifDetails) GetCrcString() string            { return "d0382c4c" }
func (*MemifDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MemifDump represents VPP binary API message 'memif_dump'.
type MemifDump struct{}

func (m *MemifDump) Reset()                        { *m = MemifDump{} }
func (*MemifDump) GetMessageName() string          { return "memif_dump" }
func (*MemifDump) GetCrcString() string            { return "51077d14" }
func (*MemifDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MemifSocketFilenameAddDel represents VPP binary API message 'memif_socket_filename_add_del'.
type MemifSocketFilenameAddDel struct {
	IsAdd          bool
	SocketID       uint32
	SocketFilename string `struc:"[108]byte"`
}

func (m *MemifSocketFilenameAddDel) Reset()                        { *m = MemifSocketFilenameAddDel{} }
func (*MemifSocketFilenameAddDel) GetMessageName() string          { return "memif_socket_filename_add_del" }
func (*MemifSocketFilenameAddDel) GetCrcString() string            { return "a2ce1a10" }
func (*MemifSocketFilenameAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MemifSocketFilenameAddDelReply represents VPP binary API message 'memif_socket_filename_add_del_reply'.
type MemifSocketFilenameAddDelReply struct {
	Retval int32
}

func (m *MemifSocketFilenameAddDelReply) Reset() { *m = MemifSocketFilenameAddDelReply{} }
func (*MemifSocketFilenameAddDelReply) GetMessageName() string {
	return "memif_socket_filename_add_del_reply"
}
func (*MemifSocketFilenameAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*MemifSocketFilenameAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MemifSocketFilenameDetails represents VPP binary API message 'memif_socket_filename_details'.
type MemifSocketFilenameDetails struct {
	SocketID       uint32
	SocketFilename string `struc:"[108]byte"`
}

func (m *MemifSocketFilenameDetails) Reset()                        { *m = MemifSocketFilenameDetails{} }
func (*MemifSocketFilenameDetails) GetMessageName() string          { return "memif_socket_filename_details" }
func (*MemifSocketFilenameDetails) GetCrcString() string            { return "7ff326f7" }
func (*MemifSocketFilenameDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MemifSocketFilenameDump represents VPP binary API message 'memif_socket_filename_dump'.
type MemifSocketFilenameDump struct{}

func (m *MemifSocketFilenameDump) Reset()                        { *m = MemifSocketFilenameDump{} }
func (*MemifSocketFilenameDump) GetMessageName() string          { return "memif_socket_filename_dump" }
func (*MemifSocketFilenameDump) GetCrcString() string            { return "51077d14" }
func (*MemifSocketFilenameDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*MemifCreate)(nil), "memif.MemifCreate")
	api.RegisterMessage((*MemifCreateReply)(nil), "memif.MemifCreateReply")
	api.RegisterMessage((*MemifDelete)(nil), "memif.MemifDelete")
	api.RegisterMessage((*MemifDeleteReply)(nil), "memif.MemifDeleteReply")
	api.RegisterMessage((*MemifDetails)(nil), "memif.MemifDetails")
	api.RegisterMessage((*MemifDump)(nil), "memif.MemifDump")
	api.RegisterMessage((*MemifSocketFilenameAddDel)(nil), "memif.MemifSocketFilenameAddDel")
	api.RegisterMessage((*MemifSocketFilenameAddDelReply)(nil), "memif.MemifSocketFilenameAddDelReply")
	api.RegisterMessage((*MemifSocketFilenameDetails)(nil), "memif.MemifSocketFilenameDetails")
	api.RegisterMessage((*MemifSocketFilenameDump)(nil), "memif.MemifSocketFilenameDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MemifCreate)(nil),
		(*MemifCreateReply)(nil),
		(*MemifDelete)(nil),
		(*MemifDeleteReply)(nil),
		(*MemifDetails)(nil),
		(*MemifDump)(nil),
		(*MemifSocketFilenameAddDel)(nil),
		(*MemifSocketFilenameAddDelReply)(nil),
		(*MemifSocketFilenameDetails)(nil),
		(*MemifSocketFilenameDump)(nil),
	}
}

// RPCService represents RPC service API for memif module.
type RPCService interface {
	DumpMemif(ctx context.Context, in *MemifDump) (RPCService_DumpMemifClient, error)
	DumpMemifSocketFilename(ctx context.Context, in *MemifSocketFilenameDump) (RPCService_DumpMemifSocketFilenameClient, error)
	MemifCreate(ctx context.Context, in *MemifCreate) (*MemifCreateReply, error)
	MemifDelete(ctx context.Context, in *MemifDelete) (*MemifDeleteReply, error)
	MemifSocketFilenameAddDel(ctx context.Context, in *MemifSocketFilenameAddDel) (*MemifSocketFilenameAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMemif(ctx context.Context, in *MemifDump) (RPCService_DumpMemifClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMemifClient{stream}
	return x, nil
}

type RPCService_DumpMemifClient interface {
	Recv() (*MemifDetails, error)
}

type serviceClient_DumpMemifClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMemifClient) Recv() (*MemifDetails, error) {
	m := new(MemifDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMemifSocketFilename(ctx context.Context, in *MemifSocketFilenameDump) (RPCService_DumpMemifSocketFilenameClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMemifSocketFilenameClient{stream}
	return x, nil
}

type RPCService_DumpMemifSocketFilenameClient interface {
	Recv() (*MemifSocketFilenameDetails, error)
}

type serviceClient_DumpMemifSocketFilenameClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMemifSocketFilenameClient) Recv() (*MemifSocketFilenameDetails, error) {
	m := new(MemifSocketFilenameDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MemifCreate(ctx context.Context, in *MemifCreate) (*MemifCreateReply, error) {
	out := new(MemifCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MemifDelete(ctx context.Context, in *MemifDelete) (*MemifDeleteReply, error) {
	out := new(MemifDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MemifSocketFilenameAddDel(ctx context.Context, in *MemifSocketFilenameAddDel) (*MemifSocketFilenameAddDelReply, error) {
	out := new(MemifSocketFilenameAddDelReply)
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
