// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/master/vppapi/core/policer.api.json

/*
Package policer is a generated VPP binary API for 'policer' module.

It consists of:
	  4 enums
	  1 type
	  4 messages
	  2 services
*/
package policer

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
	ModuleName = "policer"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd9188811
)

// Sse2QosActionType represents VPP binary API enum 'sse2_qos_action_type'.
type Sse2QosActionType uint8

const (
	SSE2_QOS_ACTION_API_DROP              Sse2QosActionType = 0
	SSE2_QOS_ACTION_API_TRANSMIT          Sse2QosActionType = 1
	SSE2_QOS_ACTION_API_MARK_AND_TRANSMIT Sse2QosActionType = 2
)

var Sse2QosActionType_name = map[uint8]string{
	0: "SSE2_QOS_ACTION_API_DROP",
	1: "SSE2_QOS_ACTION_API_TRANSMIT",
	2: "SSE2_QOS_ACTION_API_MARK_AND_TRANSMIT",
}

var Sse2QosActionType_value = map[string]uint8{
	"SSE2_QOS_ACTION_API_DROP":              0,
	"SSE2_QOS_ACTION_API_TRANSMIT":          1,
	"SSE2_QOS_ACTION_API_MARK_AND_TRANSMIT": 2,
}

func (x Sse2QosActionType) String() string {
	s, ok := Sse2QosActionType_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Sse2QosPolicerType represents VPP binary API enum 'sse2_qos_policer_type'.
type Sse2QosPolicerType uint8

const (
	SSE2_QOS_POLICER_TYPE_API_1R2C             Sse2QosPolicerType = 0
	SSE2_QOS_POLICER_TYPE_API_1R3C_RFC_2697    Sse2QosPolicerType = 1
	SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_2698    Sse2QosPolicerType = 2
	SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_4115    Sse2QosPolicerType = 3
	SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_MEF5CF1 Sse2QosPolicerType = 4
	SSE2_QOS_POLICER_TYPE_API_MAX              Sse2QosPolicerType = 5
)

var Sse2QosPolicerType_name = map[uint8]string{
	0: "SSE2_QOS_POLICER_TYPE_API_1R2C",
	1: "SSE2_QOS_POLICER_TYPE_API_1R3C_RFC_2697",
	2: "SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_2698",
	3: "SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_4115",
	4: "SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_MEF5CF1",
	5: "SSE2_QOS_POLICER_TYPE_API_MAX",
}

var Sse2QosPolicerType_value = map[string]uint8{
	"SSE2_QOS_POLICER_TYPE_API_1R2C":             0,
	"SSE2_QOS_POLICER_TYPE_API_1R3C_RFC_2697":    1,
	"SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_2698":    2,
	"SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_4115":    3,
	"SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_MEF5CF1": 4,
	"SSE2_QOS_POLICER_TYPE_API_MAX":              5,
}

func (x Sse2QosPolicerType) String() string {
	s, ok := Sse2QosPolicerType_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Sse2QosRateType represents VPP binary API enum 'sse2_qos_rate_type'.
type Sse2QosRateType uint8

const (
	SSE2_QOS_RATE_API_KBPS    Sse2QosRateType = 0
	SSE2_QOS_RATE_API_PPS     Sse2QosRateType = 1
	SSE2_QOS_RATE_API_INVALID Sse2QosRateType = 2
)

var Sse2QosRateType_name = map[uint8]string{
	0: "SSE2_QOS_RATE_API_KBPS",
	1: "SSE2_QOS_RATE_API_PPS",
	2: "SSE2_QOS_RATE_API_INVALID",
}

var Sse2QosRateType_value = map[string]uint8{
	"SSE2_QOS_RATE_API_KBPS":    0,
	"SSE2_QOS_RATE_API_PPS":     1,
	"SSE2_QOS_RATE_API_INVALID": 2,
}

func (x Sse2QosRateType) String() string {
	s, ok := Sse2QosRateType_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Sse2QosRoundType represents VPP binary API enum 'sse2_qos_round_type'.
type Sse2QosRoundType uint8

const (
	SSE2_QOS_ROUND_API_TO_CLOSEST Sse2QosRoundType = 0
	SSE2_QOS_ROUND_API_TO_UP      Sse2QosRoundType = 1
	SSE2_QOS_ROUND_API_TO_DOWN    Sse2QosRoundType = 2
	SSE2_QOS_ROUND_API_INVALID    Sse2QosRoundType = 3
)

var Sse2QosRoundType_name = map[uint8]string{
	0: "SSE2_QOS_ROUND_API_TO_CLOSEST",
	1: "SSE2_QOS_ROUND_API_TO_UP",
	2: "SSE2_QOS_ROUND_API_TO_DOWN",
	3: "SSE2_QOS_ROUND_API_INVALID",
}

var Sse2QosRoundType_value = map[string]uint8{
	"SSE2_QOS_ROUND_API_TO_CLOSEST": 0,
	"SSE2_QOS_ROUND_API_TO_UP":      1,
	"SSE2_QOS_ROUND_API_TO_DOWN":    2,
	"SSE2_QOS_ROUND_API_INVALID":    3,
}

func (x Sse2QosRoundType) String() string {
	s, ok := Sse2QosRoundType_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Sse2QosAction represents VPP binary API type 'sse2_qos_action'.
type Sse2QosAction struct {
	Type Sse2QosActionType
	Dscp uint8
}

func (*Sse2QosAction) GetTypeName() string { return "sse2_qos_action" }

// PolicerAddDel represents VPP binary API message 'policer_add_del'.
type PolicerAddDel struct {
	IsAdd         bool
	Name          string `struc:"[64]byte"`
	Cir           uint32
	Eir           uint32
	Cb            uint64
	Eb            uint64
	RateType      Sse2QosRateType
	RoundType     Sse2QosRoundType
	Type          Sse2QosPolicerType
	ColorAware    bool
	ConformAction Sse2QosAction
	ExceedAction  Sse2QosAction
	ViolateAction Sse2QosAction
}

func (m *PolicerAddDel) Reset()                        { *m = PolicerAddDel{} }
func (*PolicerAddDel) GetMessageName() string          { return "policer_add_del" }
func (*PolicerAddDel) GetCrcString() string            { return "cb948f6e" }
func (*PolicerAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// PolicerAddDelReply represents VPP binary API message 'policer_add_del_reply'.
type PolicerAddDelReply struct {
	Retval       int32
	PolicerIndex uint32
}

func (m *PolicerAddDelReply) Reset()                        { *m = PolicerAddDelReply{} }
func (*PolicerAddDelReply) GetMessageName() string          { return "policer_add_del_reply" }
func (*PolicerAddDelReply) GetCrcString() string            { return "a177cef2" }
func (*PolicerAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// PolicerDetails represents VPP binary API message 'policer_details'.
type PolicerDetails struct {
	Name               string `struc:"[64]byte"`
	Cir                uint32
	Eir                uint32
	Cb                 uint64
	Eb                 uint64
	RateType           Sse2QosRateType
	RoundType          Sse2QosRoundType
	Type               Sse2QosPolicerType
	ConformAction      Sse2QosAction
	ExceedAction       Sse2QosAction
	ViolateAction      Sse2QosAction
	SingleRate         bool
	ColorAware         bool
	Scale              uint32
	CirTokensPerPeriod uint32
	PirTokensPerPeriod uint32
	CurrentLimit       uint32
	CurrentBucket      uint32
	ExtendedLimit      uint32
	ExtendedBucket     uint32
	LastUpdateTime     uint64
}

func (m *PolicerDetails) Reset()                        { *m = PolicerDetails{} }
func (*PolicerDetails) GetMessageName() string          { return "policer_details" }
func (*PolicerDetails) GetCrcString() string            { return "a43f781a" }
func (*PolicerDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// PolicerDump represents VPP binary API message 'policer_dump'.
type PolicerDump struct {
	MatchNameValid bool
	MatchName      string `struc:"[64]byte"`
}

func (m *PolicerDump) Reset()                        { *m = PolicerDump{} }
func (*PolicerDump) GetMessageName() string          { return "policer_dump" }
func (*PolicerDump) GetCrcString() string            { return "35f1ae0f" }
func (*PolicerDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*PolicerAddDel)(nil), "policer.PolicerAddDel")
	api.RegisterMessage((*PolicerAddDelReply)(nil), "policer.PolicerAddDelReply")
	api.RegisterMessage((*PolicerDetails)(nil), "policer.PolicerDetails")
	api.RegisterMessage((*PolicerDump)(nil), "policer.PolicerDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PolicerAddDel)(nil),
		(*PolicerAddDelReply)(nil),
		(*PolicerDetails)(nil),
		(*PolicerDump)(nil),
	}
}

// RPCService represents RPC service API for policer module.
type RPCService interface {
	DumpPolicer(ctx context.Context, in *PolicerDump) (RPCService_DumpPolicerClient, error)
	PolicerAddDel(ctx context.Context, in *PolicerAddDel) (*PolicerAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpPolicer(ctx context.Context, in *PolicerDump) (RPCService_DumpPolicerClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpPolicerClient{stream}
	return x, nil
}

type RPCService_DumpPolicerClient interface {
	Recv() (*PolicerDetails, error)
}

type serviceClient_DumpPolicerClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpPolicerClient) Recv() (*PolicerDetails, error) {
	m := new(PolicerDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) PolicerAddDel(ctx context.Context, in *PolicerAddDel) (*PolicerAddDelReply, error) {
	out := new(PolicerAddDelReply)
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
