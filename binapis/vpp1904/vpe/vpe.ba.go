// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/vpe.api.json

/*
Package vpe is a generated VPP binary API for 'vpe' module.

It consists of:
	  1 type
	 18 messages
	  9 services
*/
package vpe

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
	ModuleName = "vpe"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x2cc8d629
)

// ThreadData represents VPP binary API type 'thread_data'.
type ThreadData struct {
	ID        uint32
	Name      []byte `struc:"[64]byte"`
	Type      []byte `struc:"[64]byte"`
	PID       uint32
	CPUID     uint32
	Core      uint32
	CPUSocket uint32
}

func (*ThreadData) GetTypeName() string  { return "thread_data" }
func (*ThreadData) GetCrcString() string { return "0f57094e" }

// AddNodeNext represents VPP binary API message 'add_node_next'.
type AddNodeNext struct {
	NodeName []byte `struc:"[64]byte"`
	NextName []byte `struc:"[64]byte"`
}

func (m *AddNodeNext) Reset()                        { *m = AddNodeNext{} }
func (*AddNodeNext) GetMessageName() string          { return "add_node_next" }
func (*AddNodeNext) GetCrcString() string            { return "9ab92f7a" }
func (*AddNodeNext) GetMessageType() api.MessageType { return api.RequestMessage }

// AddNodeNextReply represents VPP binary API message 'add_node_next_reply'.
type AddNodeNextReply struct {
	Retval    int32
	NextIndex uint32
}

func (m *AddNodeNextReply) Reset()                        { *m = AddNodeNextReply{} }
func (*AddNodeNextReply) GetMessageName() string          { return "add_node_next_reply" }
func (*AddNodeNextReply) GetCrcString() string            { return "2ed75f32" }
func (*AddNodeNextReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// Cli represents VPP binary API message 'cli'.
type Cli struct {
	CmdInShmem uint64
}

func (m *Cli) Reset()                        { *m = Cli{} }
func (*Cli) GetMessageName() string          { return "cli" }
func (*Cli) GetCrcString() string            { return "23bfbfff" }
func (*Cli) GetMessageType() api.MessageType { return api.RequestMessage }

// CliInband represents VPP binary API message 'cli_inband'.
type CliInband struct {
	XXX_CmdLen uint32 `struc:"sizeof=Cmd"`
	Cmd        string
}

func (m *CliInband) Reset()                        { *m = CliInband{} }
func (*CliInband) GetMessageName() string          { return "cli_inband" }
func (*CliInband) GetCrcString() string            { return "b1ad59b3" }
func (*CliInband) GetMessageType() api.MessageType { return api.RequestMessage }

// CliInbandReply represents VPP binary API message 'cli_inband_reply'.
type CliInbandReply struct {
	Retval       int32
	XXX_ReplyLen uint32 `struc:"sizeof=Reply"`
	Reply        string
}

func (m *CliInbandReply) Reset()                        { *m = CliInbandReply{} }
func (*CliInbandReply) GetMessageName() string          { return "cli_inband_reply" }
func (*CliInbandReply) GetCrcString() string            { return "6d3c80a4" }
func (*CliInbandReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// CliReply represents VPP binary API message 'cli_reply'.
type CliReply struct {
	Retval       int32
	ReplyInShmem uint64
}

func (m *CliReply) Reset()                        { *m = CliReply{} }
func (*CliReply) GetMessageName() string          { return "cli_reply" }
func (*CliReply) GetCrcString() string            { return "06d68297" }
func (*CliReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ControlPing represents VPP binary API message 'control_ping'.
type ControlPing struct{}

func (m *ControlPing) Reset()                        { *m = ControlPing{} }
func (*ControlPing) GetMessageName() string          { return "control_ping" }
func (*ControlPing) GetCrcString() string            { return "51077d14" }
func (*ControlPing) GetMessageType() api.MessageType { return api.RequestMessage }

// ControlPingReply represents VPP binary API message 'control_ping_reply'.
type ControlPingReply struct {
	Retval      int32
	ClientIndex uint32
	VpePID      uint32
}

func (m *ControlPingReply) Reset()                        { *m = ControlPingReply{} }
func (*ControlPingReply) GetMessageName() string          { return "control_ping_reply" }
func (*ControlPingReply) GetCrcString() string            { return "f6b0b8ca" }
func (*ControlPingReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GetNextIndex represents VPP binary API message 'get_next_index'.
type GetNextIndex struct {
	NodeName []byte `struc:"[64]byte"`
	NextName []byte `struc:"[64]byte"`
}

func (m *GetNextIndex) Reset()                        { *m = GetNextIndex{} }
func (*GetNextIndex) GetMessageName() string          { return "get_next_index" }
func (*GetNextIndex) GetCrcString() string            { return "9ab92f7a" }
func (*GetNextIndex) GetMessageType() api.MessageType { return api.RequestMessage }

// GetNextIndexReply represents VPP binary API message 'get_next_index_reply'.
type GetNextIndexReply struct {
	Retval    int32
	NextIndex uint32
}

func (m *GetNextIndexReply) Reset()                        { *m = GetNextIndexReply{} }
func (*GetNextIndexReply) GetMessageName() string          { return "get_next_index_reply" }
func (*GetNextIndexReply) GetCrcString() string            { return "2ed75f32" }
func (*GetNextIndexReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GetNodeGraph represents VPP binary API message 'get_node_graph'.
type GetNodeGraph struct{}

func (m *GetNodeGraph) Reset()                        { *m = GetNodeGraph{} }
func (*GetNodeGraph) GetMessageName() string          { return "get_node_graph" }
func (*GetNodeGraph) GetCrcString() string            { return "51077d14" }
func (*GetNodeGraph) GetMessageType() api.MessageType { return api.RequestMessage }

// GetNodeGraphReply represents VPP binary API message 'get_node_graph_reply'.
type GetNodeGraphReply struct {
	Retval       int32
	ReplyInShmem uint64
}

func (m *GetNodeGraphReply) Reset()                        { *m = GetNodeGraphReply{} }
func (*GetNodeGraphReply) GetMessageName() string          { return "get_node_graph_reply" }
func (*GetNodeGraphReply) GetCrcString() string            { return "06d68297" }
func (*GetNodeGraphReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GetNodeIndex represents VPP binary API message 'get_node_index'.
type GetNodeIndex struct {
	NodeName []byte `struc:"[64]byte"`
}

func (m *GetNodeIndex) Reset()                        { *m = GetNodeIndex{} }
func (*GetNodeIndex) GetMessageName() string          { return "get_node_index" }
func (*GetNodeIndex) GetCrcString() string            { return "6c9a495d" }
func (*GetNodeIndex) GetMessageType() api.MessageType { return api.RequestMessage }

// GetNodeIndexReply represents VPP binary API message 'get_node_index_reply'.
type GetNodeIndexReply struct {
	Retval    int32
	NodeIndex uint32
}

func (m *GetNodeIndexReply) Reset()                        { *m = GetNodeIndexReply{} }
func (*GetNodeIndexReply) GetMessageName() string          { return "get_node_index_reply" }
func (*GetNodeIndexReply) GetCrcString() string            { return "a8600b89" }
func (*GetNodeIndexReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ShowThreads represents VPP binary API message 'show_threads'.
type ShowThreads struct{}

func (m *ShowThreads) Reset()                        { *m = ShowThreads{} }
func (*ShowThreads) GetMessageName() string          { return "show_threads" }
func (*ShowThreads) GetCrcString() string            { return "51077d14" }
func (*ShowThreads) GetMessageType() api.MessageType { return api.RequestMessage }

// ShowThreadsReply represents VPP binary API message 'show_threads_reply'.
type ShowThreadsReply struct {
	Retval     int32
	Count      uint32 `struc:"sizeof=ThreadData"`
	ThreadData []ThreadData
}

func (m *ShowThreadsReply) Reset()                        { *m = ShowThreadsReply{} }
func (*ShowThreadsReply) GetMessageName() string          { return "show_threads_reply" }
func (*ShowThreadsReply) GetCrcString() string            { return "6942fb35" }
func (*ShowThreadsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ShowVersion represents VPP binary API message 'show_version'.
type ShowVersion struct{}

func (m *ShowVersion) Reset()                        { *m = ShowVersion{} }
func (*ShowVersion) GetMessageName() string          { return "show_version" }
func (*ShowVersion) GetCrcString() string            { return "51077d14" }
func (*ShowVersion) GetMessageType() api.MessageType { return api.RequestMessage }

// ShowVersionReply represents VPP binary API message 'show_version_reply'.
type ShowVersionReply struct {
	Retval                int32
	XXX_ProgramLen        uint32 `struc:"sizeof=Program"`
	Program               string
	XXX_VersionLen        uint32 `struc:"sizeof=Version"`
	Version               string
	XXX_BuildDateLen      uint32 `struc:"sizeof=BuildDate"`
	BuildDate             string
	XXX_BuildDirectoryLen uint32 `struc:"sizeof=BuildDirectory"`
	BuildDirectory        string
}

func (m *ShowVersionReply) Reset()                        { *m = ShowVersionReply{} }
func (*ShowVersionReply) GetMessageName() string          { return "show_version_reply" }
func (*ShowVersionReply) GetCrcString() string            { return "b9bcf6df" }
func (*ShowVersionReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*AddNodeNext)(nil), "vpe.AddNodeNext")
	api.RegisterMessage((*AddNodeNextReply)(nil), "vpe.AddNodeNextReply")
	api.RegisterMessage((*Cli)(nil), "vpe.Cli")
	api.RegisterMessage((*CliInband)(nil), "vpe.CliInband")
	api.RegisterMessage((*CliInbandReply)(nil), "vpe.CliInbandReply")
	api.RegisterMessage((*CliReply)(nil), "vpe.CliReply")
	api.RegisterMessage((*ControlPing)(nil), "vpe.ControlPing")
	api.RegisterMessage((*ControlPingReply)(nil), "vpe.ControlPingReply")
	api.RegisterMessage((*GetNextIndex)(nil), "vpe.GetNextIndex")
	api.RegisterMessage((*GetNextIndexReply)(nil), "vpe.GetNextIndexReply")
	api.RegisterMessage((*GetNodeGraph)(nil), "vpe.GetNodeGraph")
	api.RegisterMessage((*GetNodeGraphReply)(nil), "vpe.GetNodeGraphReply")
	api.RegisterMessage((*GetNodeIndex)(nil), "vpe.GetNodeIndex")
	api.RegisterMessage((*GetNodeIndexReply)(nil), "vpe.GetNodeIndexReply")
	api.RegisterMessage((*ShowThreads)(nil), "vpe.ShowThreads")
	api.RegisterMessage((*ShowThreadsReply)(nil), "vpe.ShowThreadsReply")
	api.RegisterMessage((*ShowVersion)(nil), "vpe.ShowVersion")
	api.RegisterMessage((*ShowVersionReply)(nil), "vpe.ShowVersionReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AddNodeNext)(nil),
		(*AddNodeNextReply)(nil),
		(*Cli)(nil),
		(*CliInband)(nil),
		(*CliInbandReply)(nil),
		(*CliReply)(nil),
		(*ControlPing)(nil),
		(*ControlPingReply)(nil),
		(*GetNextIndex)(nil),
		(*GetNextIndexReply)(nil),
		(*GetNodeGraph)(nil),
		(*GetNodeGraphReply)(nil),
		(*GetNodeIndex)(nil),
		(*GetNodeIndexReply)(nil),
		(*ShowThreads)(nil),
		(*ShowThreadsReply)(nil),
		(*ShowVersion)(nil),
		(*ShowVersionReply)(nil),
	}
}

// RPCService represents RPC service API for vpe module.
type RPCService interface {
	AddNodeNext(ctx context.Context, in *AddNodeNext) (*AddNodeNextReply, error)
	Cli(ctx context.Context, in *Cli) (*CliReply, error)
	CliInband(ctx context.Context, in *CliInband) (*CliInbandReply, error)
	ControlPing(ctx context.Context, in *ControlPing) (*ControlPingReply, error)
	GetNextIndex(ctx context.Context, in *GetNextIndex) (*GetNextIndexReply, error)
	GetNodeGraph(ctx context.Context, in *GetNodeGraph) (*GetNodeGraphReply, error)
	GetNodeIndex(ctx context.Context, in *GetNodeIndex) (*GetNodeIndexReply, error)
	ShowThreads(ctx context.Context, in *ShowThreads) (*ShowThreadsReply, error)
	ShowVersion(ctx context.Context, in *ShowVersion) (*ShowVersionReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) AddNodeNext(ctx context.Context, in *AddNodeNext) (*AddNodeNextReply, error) {
	out := new(AddNodeNextReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Cli(ctx context.Context, in *Cli) (*CliReply, error) {
	out := new(CliReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CliInband(ctx context.Context, in *CliInband) (*CliInbandReply, error) {
	out := new(CliInbandReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ControlPing(ctx context.Context, in *ControlPing) (*ControlPingReply, error) {
	out := new(ControlPingReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNextIndex(ctx context.Context, in *GetNextIndex) (*GetNextIndexReply, error) {
	out := new(GetNextIndexReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNodeGraph(ctx context.Context, in *GetNodeGraph) (*GetNodeGraphReply, error) {
	out := new(GetNodeGraphReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNodeIndex(ctx context.Context, in *GetNodeIndex) (*GetNodeIndexReply, error) {
	out := new(GetNodeIndexReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ShowThreads(ctx context.Context, in *ShowThreads) (*ShowThreadsReply, error) {
	out := new(ShowThreadsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ShowVersion(ctx context.Context, in *ShowVersion) (*ShowVersionReply, error) {
	out := new(ShowVersionReply)
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
