// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/mpls.api.json

/*
Package mpls is a generated VPP binary API for 'mpls' module.

It consists of:
	  2 types
	 14 messages
	  7 services
*/
package mpls

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
	ModuleName = "mpls"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xe988bf4
)

// FibMplsLabel represents VPP binary API type 'fib_mpls_label'.
type FibMplsLabel struct {
	IsUniform uint8
	Label     uint32
	TTL       uint8
	Exp       uint8
}

func (*FibMplsLabel) GetTypeName() string  { return "fib_mpls_label" }
func (*FibMplsLabel) GetCrcString() string { return "c93bf35c" }

// FibPath represents VPP binary API type 'fib_path'.
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

func (*FibPath) GetTypeName() string  { return "fib_path" }
func (*FibPath) GetCrcString() string { return "ba7a81f0" }

// MplsFibDetails represents VPP binary API message 'mpls_fib_details'.
type MplsFibDetails struct {
	TableID   uint32
	TableName []byte `struc:"[64]byte"`
	EosBit    uint8
	Label     uint32
	Count     uint32 `struc:"sizeof=Path"`
	Path      []FibPath
}

func (m *MplsFibDetails) Reset()                        { *m = MplsFibDetails{} }
func (*MplsFibDetails) GetMessageName() string          { return "mpls_fib_details" }
func (*MplsFibDetails) GetCrcString() string            { return "4404bf64" }
func (*MplsFibDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsFibDump represents VPP binary API message 'mpls_fib_dump'.
type MplsFibDump struct{}

func (m *MplsFibDump) Reset()                        { *m = MplsFibDump{} }
func (*MplsFibDump) GetMessageName() string          { return "mpls_fib_dump" }
func (*MplsFibDump) GetCrcString() string            { return "51077d14" }
func (*MplsFibDump) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsIPBindUnbind represents VPP binary API message 'mpls_ip_bind_unbind'.
type MplsIPBindUnbind struct {
	MbMplsTableID   uint32
	MbLabel         uint32
	MbIPTableID     uint32
	MbIsBind        uint8
	MbIsIP4         uint8
	MbAddressLength uint8
	MbAddress       []byte `struc:"[16]byte"`
}

func (m *MplsIPBindUnbind) Reset()                        { *m = MplsIPBindUnbind{} }
func (*MplsIPBindUnbind) GetMessageName() string          { return "mpls_ip_bind_unbind" }
func (*MplsIPBindUnbind) GetCrcString() string            { return "f0405565" }
func (*MplsIPBindUnbind) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsIPBindUnbindReply represents VPP binary API message 'mpls_ip_bind_unbind_reply'.
type MplsIPBindUnbindReply struct {
	Retval int32
}

func (m *MplsIPBindUnbindReply) Reset()                        { *m = MplsIPBindUnbindReply{} }
func (*MplsIPBindUnbindReply) GetMessageName() string          { return "mpls_ip_bind_unbind_reply" }
func (*MplsIPBindUnbindReply) GetCrcString() string            { return "e8d4e804" }
func (*MplsIPBindUnbindReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsRouteAddDel represents VPP binary API message 'mpls_route_add_del'.
type MplsRouteAddDel struct {
	MrLabel                uint32
	MrEos                  uint8
	MrTableID              uint32
	MrClassifyTableIndex   uint32
	MrIsAdd                uint8
	MrIsClassify           uint8
	MrIsMulticast          uint8
	MrIsMultipath          uint8
	MrIsResolveHost        uint8
	MrIsResolveAttached    uint8
	MrIsInterfaceRx        uint8
	MrIsRpfID              uint8
	MrNextHopProto         uint8
	MrNextHopWeight        uint8
	MrNextHopPreference    uint8
	MrNextHop              []byte `struc:"[16]byte"`
	MrNextHopNOutLabels    uint8  `struc:"sizeof=MrNextHopOutLabelStack"`
	MrNextHopSwIfIndex     uint32
	MrNextHopTableID       uint32
	MrNextHopViaLabel      uint32
	MrNextHopOutLabelStack []FibMplsLabel
}

func (m *MplsRouteAddDel) Reset()                        { *m = MplsRouteAddDel{} }
func (*MplsRouteAddDel) GetMessageName() string          { return "mpls_route_add_del" }
func (*MplsRouteAddDel) GetCrcString() string            { return "d4cc6a8d" }
func (*MplsRouteAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsRouteAddDelReply represents VPP binary API message 'mpls_route_add_del_reply'.
type MplsRouteAddDelReply struct {
	Retval     int32
	StatsIndex uint32
}

func (m *MplsRouteAddDelReply) Reset()                        { *m = MplsRouteAddDelReply{} }
func (*MplsRouteAddDelReply) GetMessageName() string          { return "mpls_route_add_del_reply" }
func (*MplsRouteAddDelReply) GetCrcString() string            { return "1992deab" }
func (*MplsRouteAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTableAddDel represents VPP binary API message 'mpls_table_add_del'.
type MplsTableAddDel struct {
	MtTableID uint32
	MtIsAdd   uint8
	MtName    []byte `struc:"[64]byte"`
}

func (m *MplsTableAddDel) Reset()                        { *m = MplsTableAddDel{} }
func (*MplsTableAddDel) GetMessageName() string          { return "mpls_table_add_del" }
func (*MplsTableAddDel) GetCrcString() string            { return "83cf0340" }
func (*MplsTableAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsTableAddDelReply represents VPP binary API message 'mpls_table_add_del_reply'.
type MplsTableAddDelReply struct {
	Retval int32
}

func (m *MplsTableAddDelReply) Reset()                        { *m = MplsTableAddDelReply{} }
func (*MplsTableAddDelReply) GetMessageName() string          { return "mpls_table_add_del_reply" }
func (*MplsTableAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*MplsTableAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTunnelAddDel represents VPP binary API message 'mpls_tunnel_add_del'.
type MplsTunnelAddDel struct {
	MtSwIfIndex            uint32
	MtIsAdd                uint8
	MtL2Only               uint8
	MtIsMulticast          uint8
	MtNextHopProtoIsIP4    uint8
	MtNextHopWeight        uint8
	MtNextHopPreference    uint8
	MtNextHop              []byte `struc:"[16]byte"`
	MtNextHopNOutLabels    uint8  `struc:"sizeof=MtNextHopOutLabelStack"`
	MtNextHopViaLabel      uint32
	MtNextHopSwIfIndex     uint32
	MtNextHopTableID       uint32
	MtNextHopOutLabelStack []FibMplsLabel
}

func (m *MplsTunnelAddDel) Reset()                        { *m = MplsTunnelAddDel{} }
func (*MplsTunnelAddDel) GetMessageName() string          { return "mpls_tunnel_add_del" }
func (*MplsTunnelAddDel) GetCrcString() string            { return "d02d9e06" }
func (*MplsTunnelAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// MplsTunnelAddDelReply represents VPP binary API message 'mpls_tunnel_add_del_reply'.
type MplsTunnelAddDelReply struct {
	Retval      int32
	SwIfIndex   uint32
	TunnelIndex uint32
}

func (m *MplsTunnelAddDelReply) Reset()                        { *m = MplsTunnelAddDelReply{} }
func (*MplsTunnelAddDelReply) GetMessageName() string          { return "mpls_tunnel_add_del_reply" }
func (*MplsTunnelAddDelReply) GetCrcString() string            { return "cc62a1ce" }
func (*MplsTunnelAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTunnelDetails represents VPP binary API message 'mpls_tunnel_details'.
type MplsTunnelDetails struct {
	MtSwIfIndex   uint32
	MtTunnelIndex uint32
	MtL2Only      uint8
	MtIsMulticast uint8
	MtCount       uint32 `struc:"sizeof=MtPaths"`
	MtPaths       []FibPath
}

func (m *MplsTunnelDetails) Reset()                        { *m = MplsTunnelDetails{} }
func (*MplsTunnelDetails) GetMessageName() string          { return "mpls_tunnel_details" }
func (*MplsTunnelDetails) GetCrcString() string            { return "7c2070cf" }
func (*MplsTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// MplsTunnelDump represents VPP binary API message 'mpls_tunnel_dump'.
type MplsTunnelDump struct {
	SwIfIndex uint32
}

func (m *MplsTunnelDump) Reset()                        { *m = MplsTunnelDump{} }
func (*MplsTunnelDump) GetMessageName() string          { return "mpls_tunnel_dump" }
func (*MplsTunnelDump) GetCrcString() string            { return "529cb13f" }
func (*MplsTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetMplsEnable represents VPP binary API message 'sw_interface_set_mpls_enable'.
type SwInterfaceSetMplsEnable struct {
	SwIfIndex uint32
	Enable    uint8
}

func (m *SwInterfaceSetMplsEnable) Reset()                        { *m = SwInterfaceSetMplsEnable{} }
func (*SwInterfaceSetMplsEnable) GetMessageName() string          { return "sw_interface_set_mpls_enable" }
func (*SwInterfaceSetMplsEnable) GetCrcString() string            { return "a36fadc0" }
func (*SwInterfaceSetMplsEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetMplsEnableReply represents VPP binary API message 'sw_interface_set_mpls_enable_reply'.
type SwInterfaceSetMplsEnableReply struct {
	Retval int32
}

func (m *SwInterfaceSetMplsEnableReply) Reset() { *m = SwInterfaceSetMplsEnableReply{} }
func (*SwInterfaceSetMplsEnableReply) GetMessageName() string {
	return "sw_interface_set_mpls_enable_reply"
}
func (*SwInterfaceSetMplsEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSetMplsEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*MplsFibDetails)(nil), "mpls.MplsFibDetails")
	api.RegisterMessage((*MplsFibDump)(nil), "mpls.MplsFibDump")
	api.RegisterMessage((*MplsIPBindUnbind)(nil), "mpls.MplsIPBindUnbind")
	api.RegisterMessage((*MplsIPBindUnbindReply)(nil), "mpls.MplsIPBindUnbindReply")
	api.RegisterMessage((*MplsRouteAddDel)(nil), "mpls.MplsRouteAddDel")
	api.RegisterMessage((*MplsRouteAddDelReply)(nil), "mpls.MplsRouteAddDelReply")
	api.RegisterMessage((*MplsTableAddDel)(nil), "mpls.MplsTableAddDel")
	api.RegisterMessage((*MplsTableAddDelReply)(nil), "mpls.MplsTableAddDelReply")
	api.RegisterMessage((*MplsTunnelAddDel)(nil), "mpls.MplsTunnelAddDel")
	api.RegisterMessage((*MplsTunnelAddDelReply)(nil), "mpls.MplsTunnelAddDelReply")
	api.RegisterMessage((*MplsTunnelDetails)(nil), "mpls.MplsTunnelDetails")
	api.RegisterMessage((*MplsTunnelDump)(nil), "mpls.MplsTunnelDump")
	api.RegisterMessage((*SwInterfaceSetMplsEnable)(nil), "mpls.SwInterfaceSetMplsEnable")
	api.RegisterMessage((*SwInterfaceSetMplsEnableReply)(nil), "mpls.SwInterfaceSetMplsEnableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MplsFibDetails)(nil),
		(*MplsFibDump)(nil),
		(*MplsIPBindUnbind)(nil),
		(*MplsIPBindUnbindReply)(nil),
		(*MplsRouteAddDel)(nil),
		(*MplsRouteAddDelReply)(nil),
		(*MplsTableAddDel)(nil),
		(*MplsTableAddDelReply)(nil),
		(*MplsTunnelAddDel)(nil),
		(*MplsTunnelAddDelReply)(nil),
		(*MplsTunnelDetails)(nil),
		(*MplsTunnelDump)(nil),
		(*SwInterfaceSetMplsEnable)(nil),
		(*SwInterfaceSetMplsEnableReply)(nil),
	}
}

// RPCService represents RPC service API for mpls module.
type RPCService interface {
	DumpMplsFib(ctx context.Context, in *MplsFibDump) (RPCService_DumpMplsFibClient, error)
	DumpMplsTunnel(ctx context.Context, in *MplsTunnelDump) (RPCService_DumpMplsTunnelClient, error)
	MplsIPBindUnbind(ctx context.Context, in *MplsIPBindUnbind) (*MplsIPBindUnbindReply, error)
	MplsRouteAddDel(ctx context.Context, in *MplsRouteAddDel) (*MplsRouteAddDelReply, error)
	MplsTableAddDel(ctx context.Context, in *MplsTableAddDel) (*MplsTableAddDelReply, error)
	MplsTunnelAddDel(ctx context.Context, in *MplsTunnelAddDel) (*MplsTunnelAddDelReply, error)
	SwInterfaceSetMplsEnable(ctx context.Context, in *SwInterfaceSetMplsEnable) (*SwInterfaceSetMplsEnableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMplsFib(ctx context.Context, in *MplsFibDump) (RPCService_DumpMplsFibClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMplsFibClient{stream}
	return x, nil
}

type RPCService_DumpMplsFibClient interface {
	Recv() (*MplsFibDetails, error)
}

type serviceClient_DumpMplsFibClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMplsFibClient) Recv() (*MplsFibDetails, error) {
	m := new(MplsFibDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMplsTunnel(ctx context.Context, in *MplsTunnelDump) (RPCService_DumpMplsTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMplsTunnelClient{stream}
	return x, nil
}

type RPCService_DumpMplsTunnelClient interface {
	Recv() (*MplsTunnelDetails, error)
}

type serviceClient_DumpMplsTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMplsTunnelClient) Recv() (*MplsTunnelDetails, error) {
	m := new(MplsTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MplsIPBindUnbind(ctx context.Context, in *MplsIPBindUnbind) (*MplsIPBindUnbindReply, error) {
	out := new(MplsIPBindUnbindReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MplsRouteAddDel(ctx context.Context, in *MplsRouteAddDel) (*MplsRouteAddDelReply, error) {
	out := new(MplsRouteAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MplsTableAddDel(ctx context.Context, in *MplsTableAddDel) (*MplsTableAddDelReply, error) {
	out := new(MplsTableAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MplsTunnelAddDel(ctx context.Context, in *MplsTunnelAddDel) (*MplsTunnelAddDelReply, error) {
	out := new(MplsTunnelAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceSetMplsEnable(ctx context.Context, in *SwInterfaceSetMplsEnable) (*SwInterfaceSetMplsEnableReply, error) {
	out := new(SwInterfaceSetMplsEnableReply)
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
