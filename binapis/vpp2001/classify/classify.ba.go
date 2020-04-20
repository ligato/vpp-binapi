// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2001/.vppapi/core/classify.api.json

/*
Package classify is a generated VPP binary API for 'classify' module.

It consists of:
	  9 enums
	  1 alias
	 28 messages
	 14 services
*/
package classify

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
	ModuleName = "classify"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x13587952
)

// ClassifyAction represents VPP binary API enum 'classify_action'.
type ClassifyAction uint8

const (
	CLASSIFY_API_ACTION_NONE              ClassifyAction = 0
	CLASSIFY_API_ACTION_SET_IP4_FIB_INDEX ClassifyAction = 1
	CLASSIFY_API_ACTION_SET_IP6_FIB_INDEX ClassifyAction = 2
	CLASSIFY_API_ACTION_SET_METADATA      ClassifyAction = 3
)

var ClassifyAction_name = map[uint8]string{
	0: "CLASSIFY_API_ACTION_NONE",
	1: "CLASSIFY_API_ACTION_SET_IP4_FIB_INDEX",
	2: "CLASSIFY_API_ACTION_SET_IP6_FIB_INDEX",
	3: "CLASSIFY_API_ACTION_SET_METADATA",
}

var ClassifyAction_value = map[string]uint8{
	"CLASSIFY_API_ACTION_NONE":              0,
	"CLASSIFY_API_ACTION_SET_IP4_FIB_INDEX": 1,
	"CLASSIFY_API_ACTION_SET_IP6_FIB_INDEX": 2,
	"CLASSIFY_API_ACTION_SET_METADATA":      3,
}

func (x ClassifyAction) String() string {
	s, ok := ClassifyAction_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// FlowClassifyTable represents VPP binary API enum 'flow_classify_table'.
type FlowClassifyTable uint8

const (
	FLOW_CLASSIFY_API_TABLE_IP4 FlowClassifyTable = 1
	FLOW_CLASSIFY_API_TABLE_IP6 FlowClassifyTable = 2
)

var FlowClassifyTable_name = map[uint8]string{
	1: "FLOW_CLASSIFY_API_TABLE_IP4",
	2: "FLOW_CLASSIFY_API_TABLE_IP6",
}

var FlowClassifyTable_value = map[string]uint8{
	"FLOW_CLASSIFY_API_TABLE_IP4": 1,
	"FLOW_CLASSIFY_API_TABLE_IP6": 2,
}

func (x FlowClassifyTable) String() string {
	s, ok := FlowClassifyTable_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

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

// PolicerClassifyTable represents VPP binary API enum 'policer_classify_table'.
type PolicerClassifyTable uint8

const (
	POLICER_CLASSIFY_API_TABLE_IP4 PolicerClassifyTable = 1
	POLICER_CLASSIFY_API_TABLE_IP6 PolicerClassifyTable = 2
	POLICER_CLASSIFY_API_TABLE_L2  PolicerClassifyTable = 3
)

var PolicerClassifyTable_name = map[uint8]string{
	1: "POLICER_CLASSIFY_API_TABLE_IP4",
	2: "POLICER_CLASSIFY_API_TABLE_IP6",
	3: "POLICER_CLASSIFY_API_TABLE_L2",
}

var PolicerClassifyTable_value = map[string]uint8{
	"POLICER_CLASSIFY_API_TABLE_IP4": 1,
	"POLICER_CLASSIFY_API_TABLE_IP6": 2,
	"POLICER_CLASSIFY_API_TABLE_L2":  3,
}

func (x PolicerClassifyTable) String() string {
	s, ok := PolicerClassifyTable_name[uint8(x)]
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

// ClassifyAddDelSession represents VPP binary API message 'classify_add_del_session'.
type ClassifyAddDelSession struct {
	IsAdd        bool
	TableIndex   uint32
	HitNextIndex uint32
	OpaqueIndex  uint32
	Advance      int32
	Action       ClassifyAction
	Metadata     uint32
	MatchLen     uint32 `struc:"sizeof=Match"`
	Match        []byte
}

func (m *ClassifyAddDelSession) Reset()                        { *m = ClassifyAddDelSession{} }
func (*ClassifyAddDelSession) GetMessageName() string          { return "classify_add_del_session" }
func (*ClassifyAddDelSession) GetCrcString() string            { return "f20879f0" }
func (*ClassifyAddDelSession) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifyAddDelSessionReply represents VPP binary API message 'classify_add_del_session_reply'.
type ClassifyAddDelSessionReply struct {
	Retval int32
}

func (m *ClassifyAddDelSessionReply) Reset()                        { *m = ClassifyAddDelSessionReply{} }
func (*ClassifyAddDelSessionReply) GetMessageName() string          { return "classify_add_del_session_reply" }
func (*ClassifyAddDelSessionReply) GetCrcString() string            { return "e8d4e804" }
func (*ClassifyAddDelSessionReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifyAddDelTable represents VPP binary API message 'classify_add_del_table'.
type ClassifyAddDelTable struct {
	IsAdd             bool
	DelChain          bool
	TableIndex        uint32
	Nbuckets          uint32
	MemorySize        uint32
	SkipNVectors      uint32
	MatchNVectors     uint32
	NextTableIndex    uint32
	MissNextIndex     uint32
	CurrentDataFlag   uint8
	CurrentDataOffset int16
	MaskLen           uint32 `struc:"sizeof=Mask"`
	Mask              []byte
}

func (m *ClassifyAddDelTable) Reset()                        { *m = ClassifyAddDelTable{} }
func (*ClassifyAddDelTable) GetMessageName() string          { return "classify_add_del_table" }
func (*ClassifyAddDelTable) GetCrcString() string            { return "6849e39e" }
func (*ClassifyAddDelTable) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifyAddDelTableReply represents VPP binary API message 'classify_add_del_table_reply'.
type ClassifyAddDelTableReply struct {
	Retval        int32
	NewTableIndex uint32
	SkipNVectors  uint32
	MatchNVectors uint32
}

func (m *ClassifyAddDelTableReply) Reset()                        { *m = ClassifyAddDelTableReply{} }
func (*ClassifyAddDelTableReply) GetMessageName() string          { return "classify_add_del_table_reply" }
func (*ClassifyAddDelTableReply) GetCrcString() string            { return "05486349" }
func (*ClassifyAddDelTableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifySessionDetails represents VPP binary API message 'classify_session_details'.
type ClassifySessionDetails struct {
	Retval       int32
	TableID      uint32
	HitNextIndex uint32
	Advance      int32
	OpaqueIndex  uint32
	MatchLength  uint32 `struc:"sizeof=Match"`
	Match        []byte
}

func (m *ClassifySessionDetails) Reset()                        { *m = ClassifySessionDetails{} }
func (*ClassifySessionDetails) GetMessageName() string          { return "classify_session_details" }
func (*ClassifySessionDetails) GetCrcString() string            { return "60e3ef94" }
func (*ClassifySessionDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifySessionDump represents VPP binary API message 'classify_session_dump'.
type ClassifySessionDump struct {
	TableID uint32
}

func (m *ClassifySessionDump) Reset()                        { *m = ClassifySessionDump{} }
func (*ClassifySessionDump) GetMessageName() string          { return "classify_session_dump" }
func (*ClassifySessionDump) GetCrcString() string            { return "0cca2cd9" }
func (*ClassifySessionDump) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifySetInterfaceIPTable represents VPP binary API message 'classify_set_interface_ip_table'.
type ClassifySetInterfaceIPTable struct {
	IsIPv6     bool
	SwIfIndex  InterfaceIndex
	TableIndex uint32
}

func (m *ClassifySetInterfaceIPTable) Reset()                        { *m = ClassifySetInterfaceIPTable{} }
func (*ClassifySetInterfaceIPTable) GetMessageName() string          { return "classify_set_interface_ip_table" }
func (*ClassifySetInterfaceIPTable) GetCrcString() string            { return "e0b097c7" }
func (*ClassifySetInterfaceIPTable) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifySetInterfaceIPTableReply represents VPP binary API message 'classify_set_interface_ip_table_reply'.
type ClassifySetInterfaceIPTableReply struct {
	Retval int32
}

func (m *ClassifySetInterfaceIPTableReply) Reset() { *m = ClassifySetInterfaceIPTableReply{} }
func (*ClassifySetInterfaceIPTableReply) GetMessageName() string {
	return "classify_set_interface_ip_table_reply"
}
func (*ClassifySetInterfaceIPTableReply) GetCrcString() string            { return "e8d4e804" }
func (*ClassifySetInterfaceIPTableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifySetInterfaceL2Tables represents VPP binary API message 'classify_set_interface_l2_tables'.
type ClassifySetInterfaceL2Tables struct {
	SwIfIndex       InterfaceIndex
	IP4TableIndex   uint32
	IP6TableIndex   uint32
	OtherTableIndex uint32
	IsInput         bool
}

func (m *ClassifySetInterfaceL2Tables) Reset() { *m = ClassifySetInterfaceL2Tables{} }
func (*ClassifySetInterfaceL2Tables) GetMessageName() string {
	return "classify_set_interface_l2_tables"
}
func (*ClassifySetInterfaceL2Tables) GetCrcString() string            { return "5a6ddf65" }
func (*ClassifySetInterfaceL2Tables) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifySetInterfaceL2TablesReply represents VPP binary API message 'classify_set_interface_l2_tables_reply'.
type ClassifySetInterfaceL2TablesReply struct {
	Retval int32
}

func (m *ClassifySetInterfaceL2TablesReply) Reset() { *m = ClassifySetInterfaceL2TablesReply{} }
func (*ClassifySetInterfaceL2TablesReply) GetMessageName() string {
	return "classify_set_interface_l2_tables_reply"
}
func (*ClassifySetInterfaceL2TablesReply) GetCrcString() string            { return "e8d4e804" }
func (*ClassifySetInterfaceL2TablesReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifyTableByInterface represents VPP binary API message 'classify_table_by_interface'.
type ClassifyTableByInterface struct {
	SwIfIndex InterfaceIndex
}

func (m *ClassifyTableByInterface) Reset()                        { *m = ClassifyTableByInterface{} }
func (*ClassifyTableByInterface) GetMessageName() string          { return "classify_table_by_interface" }
func (*ClassifyTableByInterface) GetCrcString() string            { return "f9e6675e" }
func (*ClassifyTableByInterface) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifyTableByInterfaceReply represents VPP binary API message 'classify_table_by_interface_reply'.
type ClassifyTableByInterfaceReply struct {
	Retval     int32
	SwIfIndex  InterfaceIndex
	L2TableID  uint32
	IP4TableID uint32
	IP6TableID uint32
}

func (m *ClassifyTableByInterfaceReply) Reset() { *m = ClassifyTableByInterfaceReply{} }
func (*ClassifyTableByInterfaceReply) GetMessageName() string {
	return "classify_table_by_interface_reply"
}
func (*ClassifyTableByInterfaceReply) GetCrcString() string            { return "ed4197db" }
func (*ClassifyTableByInterfaceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifyTableIds represents VPP binary API message 'classify_table_ids'.
type ClassifyTableIds struct{}

func (m *ClassifyTableIds) Reset()                        { *m = ClassifyTableIds{} }
func (*ClassifyTableIds) GetMessageName() string          { return "classify_table_ids" }
func (*ClassifyTableIds) GetCrcString() string            { return "51077d14" }
func (*ClassifyTableIds) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifyTableIdsReply represents VPP binary API message 'classify_table_ids_reply'.
type ClassifyTableIdsReply struct {
	Retval int32
	Count  uint32 `struc:"sizeof=Ids"`
	Ids    []uint32
}

func (m *ClassifyTableIdsReply) Reset()                        { *m = ClassifyTableIdsReply{} }
func (*ClassifyTableIdsReply) GetMessageName() string          { return "classify_table_ids_reply" }
func (*ClassifyTableIdsReply) GetCrcString() string            { return "d1d20e1d" }
func (*ClassifyTableIdsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ClassifyTableInfo represents VPP binary API message 'classify_table_info'.
type ClassifyTableInfo struct {
	TableID uint32
}

func (m *ClassifyTableInfo) Reset()                        { *m = ClassifyTableInfo{} }
func (*ClassifyTableInfo) GetMessageName() string          { return "classify_table_info" }
func (*ClassifyTableInfo) GetCrcString() string            { return "0cca2cd9" }
func (*ClassifyTableInfo) GetMessageType() api.MessageType { return api.RequestMessage }

// ClassifyTableInfoReply represents VPP binary API message 'classify_table_info_reply'.
type ClassifyTableInfoReply struct {
	Retval         int32
	TableID        uint32
	Nbuckets       uint32
	MatchNVectors  uint32
	SkipNVectors   uint32
	ActiveSessions uint32
	NextTableIndex uint32
	MissNextIndex  uint32
	MaskLength     uint32 `struc:"sizeof=Mask"`
	Mask           []byte
}

func (m *ClassifyTableInfoReply) Reset()                        { *m = ClassifyTableInfoReply{} }
func (*ClassifyTableInfoReply) GetMessageName() string          { return "classify_table_info_reply" }
func (*ClassifyTableInfoReply) GetCrcString() string            { return "4a573c0e" }
func (*ClassifyTableInfoReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// FlowClassifyDetails represents VPP binary API message 'flow_classify_details'.
type FlowClassifyDetails struct {
	SwIfIndex  InterfaceIndex
	TableIndex uint32
}

func (m *FlowClassifyDetails) Reset()                        { *m = FlowClassifyDetails{} }
func (*FlowClassifyDetails) GetMessageName() string          { return "flow_classify_details" }
func (*FlowClassifyDetails) GetCrcString() string            { return "dfd08765" }
func (*FlowClassifyDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// FlowClassifyDump represents VPP binary API message 'flow_classify_dump'.
type FlowClassifyDump struct {
	Type      FlowClassifyTable
	SwIfIndex InterfaceIndex
}

func (m *FlowClassifyDump) Reset()                        { *m = FlowClassifyDump{} }
func (*FlowClassifyDump) GetMessageName() string          { return "flow_classify_dump" }
func (*FlowClassifyDump) GetCrcString() string            { return "8a6ad43d" }
func (*FlowClassifyDump) GetMessageType() api.MessageType { return api.RequestMessage }

// FlowClassifySetInterface represents VPP binary API message 'flow_classify_set_interface'.
type FlowClassifySetInterface struct {
	SwIfIndex     InterfaceIndex
	IP4TableIndex uint32
	IP6TableIndex uint32
	IsAdd         bool
}

func (m *FlowClassifySetInterface) Reset()                        { *m = FlowClassifySetInterface{} }
func (*FlowClassifySetInterface) GetMessageName() string          { return "flow_classify_set_interface" }
func (*FlowClassifySetInterface) GetCrcString() string            { return "b6192f1c" }
func (*FlowClassifySetInterface) GetMessageType() api.MessageType { return api.RequestMessage }

// FlowClassifySetInterfaceReply represents VPP binary API message 'flow_classify_set_interface_reply'.
type FlowClassifySetInterfaceReply struct {
	Retval int32
}

func (m *FlowClassifySetInterfaceReply) Reset() { *m = FlowClassifySetInterfaceReply{} }
func (*FlowClassifySetInterfaceReply) GetMessageName() string {
	return "flow_classify_set_interface_reply"
}
func (*FlowClassifySetInterfaceReply) GetCrcString() string            { return "e8d4e804" }
func (*FlowClassifySetInterfaceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// InputACLSetInterface represents VPP binary API message 'input_acl_set_interface'.
type InputACLSetInterface struct {
	SwIfIndex     InterfaceIndex
	IP4TableIndex uint32
	IP6TableIndex uint32
	L2TableIndex  uint32
	IsAdd         bool
}

func (m *InputACLSetInterface) Reset()                        { *m = InputACLSetInterface{} }
func (*InputACLSetInterface) GetMessageName() string          { return "input_acl_set_interface" }
func (*InputACLSetInterface) GetCrcString() string            { return "de7ad708" }
func (*InputACLSetInterface) GetMessageType() api.MessageType { return api.RequestMessage }

// InputACLSetInterfaceReply represents VPP binary API message 'input_acl_set_interface_reply'.
type InputACLSetInterfaceReply struct {
	Retval int32
}

func (m *InputACLSetInterfaceReply) Reset()                        { *m = InputACLSetInterfaceReply{} }
func (*InputACLSetInterfaceReply) GetMessageName() string          { return "input_acl_set_interface_reply" }
func (*InputACLSetInterfaceReply) GetCrcString() string            { return "e8d4e804" }
func (*InputACLSetInterfaceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// OutputACLSetInterface represents VPP binary API message 'output_acl_set_interface'.
type OutputACLSetInterface struct {
	SwIfIndex     InterfaceIndex
	IP4TableIndex uint32
	IP6TableIndex uint32
	L2TableIndex  uint32
	IsAdd         bool
}

func (m *OutputACLSetInterface) Reset()                        { *m = OutputACLSetInterface{} }
func (*OutputACLSetInterface) GetMessageName() string          { return "output_acl_set_interface" }
func (*OutputACLSetInterface) GetCrcString() string            { return "de7ad708" }
func (*OutputACLSetInterface) GetMessageType() api.MessageType { return api.RequestMessage }

// OutputACLSetInterfaceReply represents VPP binary API message 'output_acl_set_interface_reply'.
type OutputACLSetInterfaceReply struct {
	Retval int32
}

func (m *OutputACLSetInterfaceReply) Reset()                        { *m = OutputACLSetInterfaceReply{} }
func (*OutputACLSetInterfaceReply) GetMessageName() string          { return "output_acl_set_interface_reply" }
func (*OutputACLSetInterfaceReply) GetCrcString() string            { return "e8d4e804" }
func (*OutputACLSetInterfaceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// PolicerClassifyDetails represents VPP binary API message 'policer_classify_details'.
type PolicerClassifyDetails struct {
	SwIfIndex  InterfaceIndex
	TableIndex uint32
}

func (m *PolicerClassifyDetails) Reset()                        { *m = PolicerClassifyDetails{} }
func (*PolicerClassifyDetails) GetMessageName() string          { return "policer_classify_details" }
func (*PolicerClassifyDetails) GetCrcString() string            { return "dfd08765" }
func (*PolicerClassifyDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// PolicerClassifyDump represents VPP binary API message 'policer_classify_dump'.
type PolicerClassifyDump struct {
	Type      PolicerClassifyTable
	SwIfIndex InterfaceIndex
}

func (m *PolicerClassifyDump) Reset()                        { *m = PolicerClassifyDump{} }
func (*PolicerClassifyDump) GetMessageName() string          { return "policer_classify_dump" }
func (*PolicerClassifyDump) GetCrcString() string            { return "6bfe6603" }
func (*PolicerClassifyDump) GetMessageType() api.MessageType { return api.RequestMessage }

// PolicerClassifySetInterface represents VPP binary API message 'policer_classify_set_interface'.
type PolicerClassifySetInterface struct {
	SwIfIndex     InterfaceIndex
	IP4TableIndex uint32
	IP6TableIndex uint32
	L2TableIndex  uint32
	IsAdd         bool
}

func (m *PolicerClassifySetInterface) Reset()                        { *m = PolicerClassifySetInterface{} }
func (*PolicerClassifySetInterface) GetMessageName() string          { return "policer_classify_set_interface" }
func (*PolicerClassifySetInterface) GetCrcString() string            { return "de7ad708" }
func (*PolicerClassifySetInterface) GetMessageType() api.MessageType { return api.RequestMessage }

// PolicerClassifySetInterfaceReply represents VPP binary API message 'policer_classify_set_interface_reply'.
type PolicerClassifySetInterfaceReply struct {
	Retval int32
}

func (m *PolicerClassifySetInterfaceReply) Reset() { *m = PolicerClassifySetInterfaceReply{} }
func (*PolicerClassifySetInterfaceReply) GetMessageName() string {
	return "policer_classify_set_interface_reply"
}
func (*PolicerClassifySetInterfaceReply) GetCrcString() string            { return "e8d4e804" }
func (*PolicerClassifySetInterfaceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*ClassifyAddDelSession)(nil), "classify.ClassifyAddDelSession")
	api.RegisterMessage((*ClassifyAddDelSessionReply)(nil), "classify.ClassifyAddDelSessionReply")
	api.RegisterMessage((*ClassifyAddDelTable)(nil), "classify.ClassifyAddDelTable")
	api.RegisterMessage((*ClassifyAddDelTableReply)(nil), "classify.ClassifyAddDelTableReply")
	api.RegisterMessage((*ClassifySessionDetails)(nil), "classify.ClassifySessionDetails")
	api.RegisterMessage((*ClassifySessionDump)(nil), "classify.ClassifySessionDump")
	api.RegisterMessage((*ClassifySetInterfaceIPTable)(nil), "classify.ClassifySetInterfaceIPTable")
	api.RegisterMessage((*ClassifySetInterfaceIPTableReply)(nil), "classify.ClassifySetInterfaceIPTableReply")
	api.RegisterMessage((*ClassifySetInterfaceL2Tables)(nil), "classify.ClassifySetInterfaceL2Tables")
	api.RegisterMessage((*ClassifySetInterfaceL2TablesReply)(nil), "classify.ClassifySetInterfaceL2TablesReply")
	api.RegisterMessage((*ClassifyTableByInterface)(nil), "classify.ClassifyTableByInterface")
	api.RegisterMessage((*ClassifyTableByInterfaceReply)(nil), "classify.ClassifyTableByInterfaceReply")
	api.RegisterMessage((*ClassifyTableIds)(nil), "classify.ClassifyTableIds")
	api.RegisterMessage((*ClassifyTableIdsReply)(nil), "classify.ClassifyTableIdsReply")
	api.RegisterMessage((*ClassifyTableInfo)(nil), "classify.ClassifyTableInfo")
	api.RegisterMessage((*ClassifyTableInfoReply)(nil), "classify.ClassifyTableInfoReply")
	api.RegisterMessage((*FlowClassifyDetails)(nil), "classify.FlowClassifyDetails")
	api.RegisterMessage((*FlowClassifyDump)(nil), "classify.FlowClassifyDump")
	api.RegisterMessage((*FlowClassifySetInterface)(nil), "classify.FlowClassifySetInterface")
	api.RegisterMessage((*FlowClassifySetInterfaceReply)(nil), "classify.FlowClassifySetInterfaceReply")
	api.RegisterMessage((*InputACLSetInterface)(nil), "classify.InputACLSetInterface")
	api.RegisterMessage((*InputACLSetInterfaceReply)(nil), "classify.InputACLSetInterfaceReply")
	api.RegisterMessage((*OutputACLSetInterface)(nil), "classify.OutputACLSetInterface")
	api.RegisterMessage((*OutputACLSetInterfaceReply)(nil), "classify.OutputACLSetInterfaceReply")
	api.RegisterMessage((*PolicerClassifyDetails)(nil), "classify.PolicerClassifyDetails")
	api.RegisterMessage((*PolicerClassifyDump)(nil), "classify.PolicerClassifyDump")
	api.RegisterMessage((*PolicerClassifySetInterface)(nil), "classify.PolicerClassifySetInterface")
	api.RegisterMessage((*PolicerClassifySetInterfaceReply)(nil), "classify.PolicerClassifySetInterfaceReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*ClassifyAddDelSession)(nil),
		(*ClassifyAddDelSessionReply)(nil),
		(*ClassifyAddDelTable)(nil),
		(*ClassifyAddDelTableReply)(nil),
		(*ClassifySessionDetails)(nil),
		(*ClassifySessionDump)(nil),
		(*ClassifySetInterfaceIPTable)(nil),
		(*ClassifySetInterfaceIPTableReply)(nil),
		(*ClassifySetInterfaceL2Tables)(nil),
		(*ClassifySetInterfaceL2TablesReply)(nil),
		(*ClassifyTableByInterface)(nil),
		(*ClassifyTableByInterfaceReply)(nil),
		(*ClassifyTableIds)(nil),
		(*ClassifyTableIdsReply)(nil),
		(*ClassifyTableInfo)(nil),
		(*ClassifyTableInfoReply)(nil),
		(*FlowClassifyDetails)(nil),
		(*FlowClassifyDump)(nil),
		(*FlowClassifySetInterface)(nil),
		(*FlowClassifySetInterfaceReply)(nil),
		(*InputACLSetInterface)(nil),
		(*InputACLSetInterfaceReply)(nil),
		(*OutputACLSetInterface)(nil),
		(*OutputACLSetInterfaceReply)(nil),
		(*PolicerClassifyDetails)(nil),
		(*PolicerClassifyDump)(nil),
		(*PolicerClassifySetInterface)(nil),
		(*PolicerClassifySetInterfaceReply)(nil),
	}
}

// RPCService represents RPC service API for classify module.
type RPCService interface {
	DumpClassifySession(ctx context.Context, in *ClassifySessionDump) (RPCService_DumpClassifySessionClient, error)
	DumpFlowClassify(ctx context.Context, in *FlowClassifyDump) (RPCService_DumpFlowClassifyClient, error)
	DumpPolicerClassify(ctx context.Context, in *PolicerClassifyDump) (RPCService_DumpPolicerClassifyClient, error)
	ClassifyAddDelSession(ctx context.Context, in *ClassifyAddDelSession) (*ClassifyAddDelSessionReply, error)
	ClassifyAddDelTable(ctx context.Context, in *ClassifyAddDelTable) (*ClassifyAddDelTableReply, error)
	ClassifySetInterfaceIPTable(ctx context.Context, in *ClassifySetInterfaceIPTable) (*ClassifySetInterfaceIPTableReply, error)
	ClassifySetInterfaceL2Tables(ctx context.Context, in *ClassifySetInterfaceL2Tables) (*ClassifySetInterfaceL2TablesReply, error)
	ClassifyTableByInterface(ctx context.Context, in *ClassifyTableByInterface) (*ClassifyTableByInterfaceReply, error)
	ClassifyTableIds(ctx context.Context, in *ClassifyTableIds) (*ClassifyTableIdsReply, error)
	ClassifyTableInfo(ctx context.Context, in *ClassifyTableInfo) (*ClassifyTableInfoReply, error)
	FlowClassifySetInterface(ctx context.Context, in *FlowClassifySetInterface) (*FlowClassifySetInterfaceReply, error)
	InputACLSetInterface(ctx context.Context, in *InputACLSetInterface) (*InputACLSetInterfaceReply, error)
	OutputACLSetInterface(ctx context.Context, in *OutputACLSetInterface) (*OutputACLSetInterfaceReply, error)
	PolicerClassifySetInterface(ctx context.Context, in *PolicerClassifySetInterface) (*PolicerClassifySetInterfaceReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpClassifySession(ctx context.Context, in *ClassifySessionDump) (RPCService_DumpClassifySessionClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpClassifySessionClient{stream}
	return x, nil
}

type RPCService_DumpClassifySessionClient interface {
	Recv() (*ClassifySessionDetails, error)
}

type serviceClient_DumpClassifySessionClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpClassifySessionClient) Recv() (*ClassifySessionDetails, error) {
	m := new(ClassifySessionDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpFlowClassify(ctx context.Context, in *FlowClassifyDump) (RPCService_DumpFlowClassifyClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpFlowClassifyClient{stream}
	return x, nil
}

type RPCService_DumpFlowClassifyClient interface {
	Recv() (*FlowClassifyDetails, error)
}

type serviceClient_DumpFlowClassifyClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpFlowClassifyClient) Recv() (*FlowClassifyDetails, error) {
	m := new(FlowClassifyDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpPolicerClassify(ctx context.Context, in *PolicerClassifyDump) (RPCService_DumpPolicerClassifyClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpPolicerClassifyClient{stream}
	return x, nil
}

type RPCService_DumpPolicerClassifyClient interface {
	Recv() (*PolicerClassifyDetails, error)
}

type serviceClient_DumpPolicerClassifyClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpPolicerClassifyClient) Recv() (*PolicerClassifyDetails, error) {
	m := new(PolicerClassifyDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) ClassifyAddDelSession(ctx context.Context, in *ClassifyAddDelSession) (*ClassifyAddDelSessionReply, error) {
	out := new(ClassifyAddDelSessionReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ClassifyAddDelTable(ctx context.Context, in *ClassifyAddDelTable) (*ClassifyAddDelTableReply, error) {
	out := new(ClassifyAddDelTableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ClassifySetInterfaceIPTable(ctx context.Context, in *ClassifySetInterfaceIPTable) (*ClassifySetInterfaceIPTableReply, error) {
	out := new(ClassifySetInterfaceIPTableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ClassifySetInterfaceL2Tables(ctx context.Context, in *ClassifySetInterfaceL2Tables) (*ClassifySetInterfaceL2TablesReply, error) {
	out := new(ClassifySetInterfaceL2TablesReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ClassifyTableByInterface(ctx context.Context, in *ClassifyTableByInterface) (*ClassifyTableByInterfaceReply, error) {
	out := new(ClassifyTableByInterfaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ClassifyTableIds(ctx context.Context, in *ClassifyTableIds) (*ClassifyTableIdsReply, error) {
	out := new(ClassifyTableIdsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ClassifyTableInfo(ctx context.Context, in *ClassifyTableInfo) (*ClassifyTableInfoReply, error) {
	out := new(ClassifyTableInfoReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FlowClassifySetInterface(ctx context.Context, in *FlowClassifySetInterface) (*FlowClassifySetInterfaceReply, error) {
	out := new(FlowClassifySetInterfaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) InputACLSetInterface(ctx context.Context, in *InputACLSetInterface) (*InputACLSetInterfaceReply, error) {
	out := new(InputACLSetInterfaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) OutputACLSetInterface(ctx context.Context, in *OutputACLSetInterface) (*OutputACLSetInterfaceReply, error) {
	out := new(OutputACLSetInterfaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PolicerClassifySetInterface(ctx context.Context, in *PolicerClassifySetInterface) (*PolicerClassifySetInterfaceReply, error) {
	out := new(PolicerClassifySetInterfaceReply)
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
