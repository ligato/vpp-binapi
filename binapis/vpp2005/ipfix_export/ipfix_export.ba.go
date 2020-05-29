// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/core/ipfix_export.api.json

/*
Package ipfix_export is a generated VPP binary API for 'ipfix_export' module.

It consists of:
	  4 enums
	  5 aliases
	  6 types
	  1 union
	 14 messages
	  7 services
*/
package ipfix_export

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
	ModuleName = "ipfix_export"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xb7e9cad2
)

// AddressFamily represents VPP binary API enum 'address_family'.
type AddressFamily uint8

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var AddressFamily_name = map[uint8]string{
	0: "ADDRESS_IP4",
	1: "ADDRESS_IP6",
}

var AddressFamily_value = map[string]uint8{
	"ADDRESS_IP4": 0,
	"ADDRESS_IP6": 1,
}

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPDscp represents VPP binary API enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var IPDscp_name = map[uint8]string{
	0:  "IP_API_DSCP_CS0",
	8:  "IP_API_DSCP_CS1",
	10: "IP_API_DSCP_AF11",
	12: "IP_API_DSCP_AF12",
	14: "IP_API_DSCP_AF13",
	16: "IP_API_DSCP_CS2",
	18: "IP_API_DSCP_AF21",
	20: "IP_API_DSCP_AF22",
	22: "IP_API_DSCP_AF23",
	24: "IP_API_DSCP_CS3",
	26: "IP_API_DSCP_AF31",
	28: "IP_API_DSCP_AF32",
	30: "IP_API_DSCP_AF33",
	32: "IP_API_DSCP_CS4",
	34: "IP_API_DSCP_AF41",
	36: "IP_API_DSCP_AF42",
	38: "IP_API_DSCP_AF43",
	40: "IP_API_DSCP_CS5",
	46: "IP_API_DSCP_EF",
	48: "IP_API_DSCP_CS6",
	50: "IP_API_DSCP_CS7",
}

var IPDscp_value = map[string]uint8{
	"IP_API_DSCP_CS0":  0,
	"IP_API_DSCP_CS1":  8,
	"IP_API_DSCP_AF11": 10,
	"IP_API_DSCP_AF12": 12,
	"IP_API_DSCP_AF13": 14,
	"IP_API_DSCP_CS2":  16,
	"IP_API_DSCP_AF21": 18,
	"IP_API_DSCP_AF22": 20,
	"IP_API_DSCP_AF23": 22,
	"IP_API_DSCP_CS3":  24,
	"IP_API_DSCP_AF31": 26,
	"IP_API_DSCP_AF32": 28,
	"IP_API_DSCP_AF33": 30,
	"IP_API_DSCP_CS4":  32,
	"IP_API_DSCP_AF41": 34,
	"IP_API_DSCP_AF42": 36,
	"IP_API_DSCP_AF43": 38,
	"IP_API_DSCP_CS5":  40,
	"IP_API_DSCP_EF":   46,
	"IP_API_DSCP_CS6":  48,
	"IP_API_DSCP_CS7":  50,
}

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPEcn represents VPP binary API enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var IPEcn_name = map[uint8]string{
	0: "IP_API_ECN_NONE",
	1: "IP_API_ECN_ECT0",
	2: "IP_API_ECN_ECT1",
	3: "IP_API_ECN_CE",
}

var IPEcn_value = map[string]uint8{
	"IP_API_ECN_NONE": 0,
	"IP_API_ECN_ECT0": 1,
	"IP_API_ECN_ECT1": 2,
	"IP_API_ECN_CE":   3,
}

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPProto represents VPP binary API enum 'ip_proto'.
type IPProto uint8

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_ESP      IPProto = 50
	IP_API_PROTO_AH       IPProto = 51
	IP_API_PROTO_ICMP6    IPProto = 58
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var IPProto_name = map[uint8]string{
	0:   "IP_API_PROTO_HOPOPT",
	1:   "IP_API_PROTO_ICMP",
	2:   "IP_API_PROTO_IGMP",
	6:   "IP_API_PROTO_TCP",
	17:  "IP_API_PROTO_UDP",
	47:  "IP_API_PROTO_GRE",
	50:  "IP_API_PROTO_ESP",
	51:  "IP_API_PROTO_AH",
	58:  "IP_API_PROTO_ICMP6",
	88:  "IP_API_PROTO_EIGRP",
	89:  "IP_API_PROTO_OSPF",
	132: "IP_API_PROTO_SCTP",
	255: "IP_API_PROTO_RESERVED",
}

var IPProto_value = map[string]uint8{
	"IP_API_PROTO_HOPOPT":   0,
	"IP_API_PROTO_ICMP":     1,
	"IP_API_PROTO_IGMP":     2,
	"IP_API_PROTO_TCP":      6,
	"IP_API_PROTO_UDP":      17,
	"IP_API_PROTO_GRE":      47,
	"IP_API_PROTO_ESP":      50,
	"IP_API_PROTO_AH":       51,
	"IP_API_PROTO_ICMP6":    58,
	"IP_API_PROTO_EIGRP":    88,
	"IP_API_PROTO_OSPF":     89,
	"IP_API_PROTO_SCTP":     132,
	"IP_API_PROTO_RESERVED": 255,
}

func (x IPProto) String() string {
	s, ok := IPProto_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// AddressWithPrefix represents VPP binary API alias 'address_with_prefix'.
type AddressWithPrefix Prefix

// IP4Address represents VPP binary API alias 'ip4_address'.
type IP4Address [4]uint8

// IP4AddressWithPrefix represents VPP binary API alias 'ip4_address_with_prefix'.
type IP4AddressWithPrefix IP4Prefix

// IP6Address represents VPP binary API alias 'ip6_address'.
type IP6Address [16]uint8

// IP6AddressWithPrefix represents VPP binary API alias 'ip6_address_with_prefix'.
type IP6AddressWithPrefix IP6Prefix

// Address represents VPP binary API type 'address'.
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string { return "address" }

// IP4Prefix represents VPP binary API type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address
	Len     uint8
}

func (*IP4Prefix) GetTypeName() string { return "ip4_prefix" }

// IP6Prefix represents VPP binary API type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address
	Len     uint8
}

func (*IP6Prefix) GetTypeName() string { return "ip6_prefix" }

// Mprefix represents VPP binary API type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string { return "mprefix" }

// Prefix represents VPP binary API type 'prefix'.
type Prefix struct {
	Address Address
	Len     uint8
}

func (*Prefix) GetTypeName() string { return "prefix" }

// PrefixMatcher represents VPP binary API type 'prefix_matcher'.
type PrefixMatcher struct {
	Le uint8
	Ge uint8
}

func (*PrefixMatcher) GetTypeName() string { return "prefix_matcher" }

// AddressUnion represents VPP binary API union 'address_union'.
type AddressUnion struct {
	XXX_UnionData [16]byte
}

func (*AddressUnion) GetTypeName() string { return "address_union" }

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

// IpfixClassifyStreamDetails represents VPP binary API message 'ipfix_classify_stream_details'.
type IpfixClassifyStreamDetails struct {
	DomainID uint32
	SrcPort  uint16
}

func (m *IpfixClassifyStreamDetails) Reset()                        { *m = IpfixClassifyStreamDetails{} }
func (*IpfixClassifyStreamDetails) GetMessageName() string          { return "ipfix_classify_stream_details" }
func (*IpfixClassifyStreamDetails) GetCrcString() string            { return "2903539d" }
func (*IpfixClassifyStreamDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// IpfixClassifyStreamDump represents VPP binary API message 'ipfix_classify_stream_dump'.
type IpfixClassifyStreamDump struct{}

func (m *IpfixClassifyStreamDump) Reset()                        { *m = IpfixClassifyStreamDump{} }
func (*IpfixClassifyStreamDump) GetMessageName() string          { return "ipfix_classify_stream_dump" }
func (*IpfixClassifyStreamDump) GetCrcString() string            { return "51077d14" }
func (*IpfixClassifyStreamDump) GetMessageType() api.MessageType { return api.RequestMessage }

// IpfixClassifyTableAddDel represents VPP binary API message 'ipfix_classify_table_add_del'.
type IpfixClassifyTableAddDel struct {
	TableID           uint32
	IPVersion         AddressFamily
	TransportProtocol IPProto
	IsAdd             bool
}

func (m *IpfixClassifyTableAddDel) Reset()                        { *m = IpfixClassifyTableAddDel{} }
func (*IpfixClassifyTableAddDel) GetMessageName() string          { return "ipfix_classify_table_add_del" }
func (*IpfixClassifyTableAddDel) GetCrcString() string            { return "3e449bb9" }
func (*IpfixClassifyTableAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// IpfixClassifyTableAddDelReply represents VPP binary API message 'ipfix_classify_table_add_del_reply'.
type IpfixClassifyTableAddDelReply struct {
	Retval int32
}

func (m *IpfixClassifyTableAddDelReply) Reset() { *m = IpfixClassifyTableAddDelReply{} }
func (*IpfixClassifyTableAddDelReply) GetMessageName() string {
	return "ipfix_classify_table_add_del_reply"
}
func (*IpfixClassifyTableAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*IpfixClassifyTableAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// IpfixClassifyTableDetails represents VPP binary API message 'ipfix_classify_table_details'.
type IpfixClassifyTableDetails struct {
	TableID           uint32
	IPVersion         AddressFamily
	TransportProtocol IPProto
}

func (m *IpfixClassifyTableDetails) Reset()                        { *m = IpfixClassifyTableDetails{} }
func (*IpfixClassifyTableDetails) GetMessageName() string          { return "ipfix_classify_table_details" }
func (*IpfixClassifyTableDetails) GetCrcString() string            { return "1af8c28c" }
func (*IpfixClassifyTableDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// IpfixClassifyTableDump represents VPP binary API message 'ipfix_classify_table_dump'.
type IpfixClassifyTableDump struct{}

func (m *IpfixClassifyTableDump) Reset()                        { *m = IpfixClassifyTableDump{} }
func (*IpfixClassifyTableDump) GetMessageName() string          { return "ipfix_classify_table_dump" }
func (*IpfixClassifyTableDump) GetCrcString() string            { return "51077d14" }
func (*IpfixClassifyTableDump) GetMessageType() api.MessageType { return api.RequestMessage }

// IpfixExporterDetails represents VPP binary API message 'ipfix_exporter_details'.
type IpfixExporterDetails struct {
	CollectorAddress Address
	CollectorPort    uint16
	SrcAddress       Address
	VrfID            uint32
	PathMtu          uint32
	TemplateInterval uint32
	UDPChecksum      bool
}

func (m *IpfixExporterDetails) Reset()                        { *m = IpfixExporterDetails{} }
func (*IpfixExporterDetails) GetMessageName() string          { return "ipfix_exporter_details" }
func (*IpfixExporterDetails) GetCrcString() string            { return "11e07413" }
func (*IpfixExporterDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// IpfixExporterDump represents VPP binary API message 'ipfix_exporter_dump'.
type IpfixExporterDump struct{}

func (m *IpfixExporterDump) Reset()                        { *m = IpfixExporterDump{} }
func (*IpfixExporterDump) GetMessageName() string          { return "ipfix_exporter_dump" }
func (*IpfixExporterDump) GetCrcString() string            { return "51077d14" }
func (*IpfixExporterDump) GetMessageType() api.MessageType { return api.RequestMessage }

// IpfixFlush represents VPP binary API message 'ipfix_flush'.
type IpfixFlush struct{}

func (m *IpfixFlush) Reset()                        { *m = IpfixFlush{} }
func (*IpfixFlush) GetMessageName() string          { return "ipfix_flush" }
func (*IpfixFlush) GetCrcString() string            { return "51077d14" }
func (*IpfixFlush) GetMessageType() api.MessageType { return api.RequestMessage }

// IpfixFlushReply represents VPP binary API message 'ipfix_flush_reply'.
type IpfixFlushReply struct {
	Retval int32
}

func (m *IpfixFlushReply) Reset()                        { *m = IpfixFlushReply{} }
func (*IpfixFlushReply) GetMessageName() string          { return "ipfix_flush_reply" }
func (*IpfixFlushReply) GetCrcString() string            { return "e8d4e804" }
func (*IpfixFlushReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SetIpfixClassifyStream represents VPP binary API message 'set_ipfix_classify_stream'.
type SetIpfixClassifyStream struct {
	DomainID uint32
	SrcPort  uint16
}

func (m *SetIpfixClassifyStream) Reset()                        { *m = SetIpfixClassifyStream{} }
func (*SetIpfixClassifyStream) GetMessageName() string          { return "set_ipfix_classify_stream" }
func (*SetIpfixClassifyStream) GetCrcString() string            { return "c9cbe053" }
func (*SetIpfixClassifyStream) GetMessageType() api.MessageType { return api.RequestMessage }

// SetIpfixClassifyStreamReply represents VPP binary API message 'set_ipfix_classify_stream_reply'.
type SetIpfixClassifyStreamReply struct {
	Retval int32
}

func (m *SetIpfixClassifyStreamReply) Reset()                        { *m = SetIpfixClassifyStreamReply{} }
func (*SetIpfixClassifyStreamReply) GetMessageName() string          { return "set_ipfix_classify_stream_reply" }
func (*SetIpfixClassifyStreamReply) GetCrcString() string            { return "e8d4e804" }
func (*SetIpfixClassifyStreamReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SetIpfixExporter represents VPP binary API message 'set_ipfix_exporter'.
type SetIpfixExporter struct {
	CollectorAddress Address
	CollectorPort    uint16
	SrcAddress       Address
	VrfID            uint32
	PathMtu          uint32
	TemplateInterval uint32
	UDPChecksum      bool
}

func (m *SetIpfixExporter) Reset()                        { *m = SetIpfixExporter{} }
func (*SetIpfixExporter) GetMessageName() string          { return "set_ipfix_exporter" }
func (*SetIpfixExporter) GetCrcString() string            { return "69284e07" }
func (*SetIpfixExporter) GetMessageType() api.MessageType { return api.RequestMessage }

// SetIpfixExporterReply represents VPP binary API message 'set_ipfix_exporter_reply'.
type SetIpfixExporterReply struct {
	Retval int32
}

func (m *SetIpfixExporterReply) Reset()                        { *m = SetIpfixExporterReply{} }
func (*SetIpfixExporterReply) GetMessageName() string          { return "set_ipfix_exporter_reply" }
func (*SetIpfixExporterReply) GetCrcString() string            { return "e8d4e804" }
func (*SetIpfixExporterReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*IpfixClassifyStreamDetails)(nil), "ipfix_export.IpfixClassifyStreamDetails")
	api.RegisterMessage((*IpfixClassifyStreamDump)(nil), "ipfix_export.IpfixClassifyStreamDump")
	api.RegisterMessage((*IpfixClassifyTableAddDel)(nil), "ipfix_export.IpfixClassifyTableAddDel")
	api.RegisterMessage((*IpfixClassifyTableAddDelReply)(nil), "ipfix_export.IpfixClassifyTableAddDelReply")
	api.RegisterMessage((*IpfixClassifyTableDetails)(nil), "ipfix_export.IpfixClassifyTableDetails")
	api.RegisterMessage((*IpfixClassifyTableDump)(nil), "ipfix_export.IpfixClassifyTableDump")
	api.RegisterMessage((*IpfixExporterDetails)(nil), "ipfix_export.IpfixExporterDetails")
	api.RegisterMessage((*IpfixExporterDump)(nil), "ipfix_export.IpfixExporterDump")
	api.RegisterMessage((*IpfixFlush)(nil), "ipfix_export.IpfixFlush")
	api.RegisterMessage((*IpfixFlushReply)(nil), "ipfix_export.IpfixFlushReply")
	api.RegisterMessage((*SetIpfixClassifyStream)(nil), "ipfix_export.SetIpfixClassifyStream")
	api.RegisterMessage((*SetIpfixClassifyStreamReply)(nil), "ipfix_export.SetIpfixClassifyStreamReply")
	api.RegisterMessage((*SetIpfixExporter)(nil), "ipfix_export.SetIpfixExporter")
	api.RegisterMessage((*SetIpfixExporterReply)(nil), "ipfix_export.SetIpfixExporterReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IpfixClassifyStreamDetails)(nil),
		(*IpfixClassifyStreamDump)(nil),
		(*IpfixClassifyTableAddDel)(nil),
		(*IpfixClassifyTableAddDelReply)(nil),
		(*IpfixClassifyTableDetails)(nil),
		(*IpfixClassifyTableDump)(nil),
		(*IpfixExporterDetails)(nil),
		(*IpfixExporterDump)(nil),
		(*IpfixFlush)(nil),
		(*IpfixFlushReply)(nil),
		(*SetIpfixClassifyStream)(nil),
		(*SetIpfixClassifyStreamReply)(nil),
		(*SetIpfixExporter)(nil),
		(*SetIpfixExporterReply)(nil),
	}
}

// RPCService represents RPC service API for ipfix_export module.
type RPCService interface {
	DumpIpfixClassifyStream(ctx context.Context, in *IpfixClassifyStreamDump) (RPCService_DumpIpfixClassifyStreamClient, error)
	DumpIpfixClassifyTable(ctx context.Context, in *IpfixClassifyTableDump) (RPCService_DumpIpfixClassifyTableClient, error)
	DumpIpfixExporter(ctx context.Context, in *IpfixExporterDump) (RPCService_DumpIpfixExporterClient, error)
	IpfixClassifyTableAddDel(ctx context.Context, in *IpfixClassifyTableAddDel) (*IpfixClassifyTableAddDelReply, error)
	IpfixFlush(ctx context.Context, in *IpfixFlush) (*IpfixFlushReply, error)
	SetIpfixClassifyStream(ctx context.Context, in *SetIpfixClassifyStream) (*SetIpfixClassifyStreamReply, error)
	SetIpfixExporter(ctx context.Context, in *SetIpfixExporter) (*SetIpfixExporterReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpIpfixClassifyStream(ctx context.Context, in *IpfixClassifyStreamDump) (RPCService_DumpIpfixClassifyStreamClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIpfixClassifyStreamClient{stream}
	return x, nil
}

type RPCService_DumpIpfixClassifyStreamClient interface {
	Recv() (*IpfixClassifyStreamDetails, error)
}

type serviceClient_DumpIpfixClassifyStreamClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIpfixClassifyStreamClient) Recv() (*IpfixClassifyStreamDetails, error) {
	m := new(IpfixClassifyStreamDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIpfixClassifyTable(ctx context.Context, in *IpfixClassifyTableDump) (RPCService_DumpIpfixClassifyTableClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIpfixClassifyTableClient{stream}
	return x, nil
}

type RPCService_DumpIpfixClassifyTableClient interface {
	Recv() (*IpfixClassifyTableDetails, error)
}

type serviceClient_DumpIpfixClassifyTableClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIpfixClassifyTableClient) Recv() (*IpfixClassifyTableDetails, error) {
	m := new(IpfixClassifyTableDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIpfixExporter(ctx context.Context, in *IpfixExporterDump) (RPCService_DumpIpfixExporterClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIpfixExporterClient{stream}
	return x, nil
}

type RPCService_DumpIpfixExporterClient interface {
	Recv() (*IpfixExporterDetails, error)
}

type serviceClient_DumpIpfixExporterClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIpfixExporterClient) Recv() (*IpfixExporterDetails, error) {
	m := new(IpfixExporterDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) IpfixClassifyTableAddDel(ctx context.Context, in *IpfixClassifyTableAddDel) (*IpfixClassifyTableAddDelReply, error) {
	out := new(IpfixClassifyTableAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IpfixFlush(ctx context.Context, in *IpfixFlush) (*IpfixFlushReply, error) {
	out := new(IpfixFlushReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SetIpfixClassifyStream(ctx context.Context, in *SetIpfixClassifyStream) (*SetIpfixClassifyStreamReply, error) {
	out := new(SetIpfixClassifyStreamReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SetIpfixExporter(ctx context.Context, in *SetIpfixExporter) (*SetIpfixExporterReply, error) {
	out := new(SetIpfixExporterReply)
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
