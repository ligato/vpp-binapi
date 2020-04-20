// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/plugins/flowprobe.api.json

/*
Package flowprobe is a generated VPP binary API for 'flowprobe' module.

It consists of:
	  4 messages
	  2 services
*/
package flowprobe

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
	ModuleName = "flowprobe"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x8c36cb50
)

// FlowprobeParams represents VPP binary API message 'flowprobe_params'.
type FlowprobeParams struct {
	RecordL2     uint8
	RecordL3     uint8
	RecordL4     uint8
	ActiveTimer  uint32
	PassiveTimer uint32
}

func (m *FlowprobeParams) Reset()                        { *m = FlowprobeParams{} }
func (*FlowprobeParams) GetMessageName() string          { return "flowprobe_params" }
func (*FlowprobeParams) GetCrcString() string            { return "574adc1c" }
func (*FlowprobeParams) GetMessageType() api.MessageType { return api.RequestMessage }

// FlowprobeParamsReply represents VPP binary API message 'flowprobe_params_reply'.
type FlowprobeParamsReply struct {
	Retval int32
}

func (m *FlowprobeParamsReply) Reset()                        { *m = FlowprobeParamsReply{} }
func (*FlowprobeParamsReply) GetMessageName() string          { return "flowprobe_params_reply" }
func (*FlowprobeParamsReply) GetCrcString() string            { return "e8d4e804" }
func (*FlowprobeParamsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// FlowprobeTxInterfaceAddDel represents VPP binary API message 'flowprobe_tx_interface_add_del'.
type FlowprobeTxInterfaceAddDel struct {
	IsAdd     uint8
	Which     uint8
	SwIfIndex uint32
}

func (m *FlowprobeTxInterfaceAddDel) Reset()                        { *m = FlowprobeTxInterfaceAddDel{} }
func (*FlowprobeTxInterfaceAddDel) GetMessageName() string          { return "flowprobe_tx_interface_add_del" }
func (*FlowprobeTxInterfaceAddDel) GetCrcString() string            { return "10c2fbab" }
func (*FlowprobeTxInterfaceAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// FlowprobeTxInterfaceAddDelReply represents VPP binary API message 'flowprobe_tx_interface_add_del_reply'.
type FlowprobeTxInterfaceAddDelReply struct {
	Retval int32
}

func (m *FlowprobeTxInterfaceAddDelReply) Reset() { *m = FlowprobeTxInterfaceAddDelReply{} }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageName() string {
	return "flowprobe_tx_interface_add_del_reply"
}
func (*FlowprobeTxInterfaceAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*FlowprobeParams)(nil), "flowprobe.FlowprobeParams")
	api.RegisterMessage((*FlowprobeParamsReply)(nil), "flowprobe.FlowprobeParamsReply")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDel)(nil), "flowprobe.FlowprobeTxInterfaceAddDel")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDelReply)(nil), "flowprobe.FlowprobeTxInterfaceAddDelReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FlowprobeParams)(nil),
		(*FlowprobeParamsReply)(nil),
		(*FlowprobeTxInterfaceAddDel)(nil),
		(*FlowprobeTxInterfaceAddDelReply)(nil),
	}
}

// RPCService represents RPC service API for flowprobe module.
type RPCService interface {
	FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error)
	FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error) {
	out := new(FlowprobeParamsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error) {
	out := new(FlowprobeTxInterfaceAddDelReply)
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
