// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/core/span.api.json

/*
Package span is a generated VPP binary API for 'span' module.

It consists of:
	  4 messages
	  2 services
*/
package span

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
	ModuleName = "span"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x2a11ff26
)

// SwInterfaceSpanDetails represents VPP binary API message 'sw_interface_span_details'.
type SwInterfaceSpanDetails struct {
	SwIfIndexFrom uint32
	SwIfIndexTo   uint32
	State         uint8
	IsL2          uint8
}

func (m *SwInterfaceSpanDetails) Reset()                        { *m = SwInterfaceSpanDetails{} }
func (*SwInterfaceSpanDetails) GetMessageName() string          { return "sw_interface_span_details" }
func (*SwInterfaceSpanDetails) GetCrcString() string            { return "23966371" }
func (*SwInterfaceSpanDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceSpanDump represents VPP binary API message 'sw_interface_span_dump'.
type SwInterfaceSpanDump struct {
	IsL2 uint8
}

func (m *SwInterfaceSpanDump) Reset()                        { *m = SwInterfaceSpanDump{} }
func (*SwInterfaceSpanDump) GetMessageName() string          { return "sw_interface_span_dump" }
func (*SwInterfaceSpanDump) GetCrcString() string            { return "67c54650" }
func (*SwInterfaceSpanDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSpanEnableDisable represents VPP binary API message 'sw_interface_span_enable_disable'.
type SwInterfaceSpanEnableDisable struct {
	SwIfIndexFrom uint32
	SwIfIndexTo   uint32
	State         uint8
	IsL2          uint8
}

func (m *SwInterfaceSpanEnableDisable) Reset() { *m = SwInterfaceSpanEnableDisable{} }
func (*SwInterfaceSpanEnableDisable) GetMessageName() string {
	return "sw_interface_span_enable_disable"
}
func (*SwInterfaceSpanEnableDisable) GetCrcString() string            { return "7216258d" }
func (*SwInterfaceSpanEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSpanEnableDisableReply represents VPP binary API message 'sw_interface_span_enable_disable_reply'.
type SwInterfaceSpanEnableDisableReply struct {
	Retval int32
}

func (m *SwInterfaceSpanEnableDisableReply) Reset() { *m = SwInterfaceSpanEnableDisableReply{} }
func (*SwInterfaceSpanEnableDisableReply) GetMessageName() string {
	return "sw_interface_span_enable_disable_reply"
}
func (*SwInterfaceSpanEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSpanEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*SwInterfaceSpanDetails)(nil), "span.SwInterfaceSpanDetails")
	api.RegisterMessage((*SwInterfaceSpanDump)(nil), "span.SwInterfaceSpanDump")
	api.RegisterMessage((*SwInterfaceSpanEnableDisable)(nil), "span.SwInterfaceSpanEnableDisable")
	api.RegisterMessage((*SwInterfaceSpanEnableDisableReply)(nil), "span.SwInterfaceSpanEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceSpanDetails)(nil),
		(*SwInterfaceSpanDump)(nil),
		(*SwInterfaceSpanEnableDisable)(nil),
		(*SwInterfaceSpanEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for span module.
type RPCService interface {
	DumpSwInterfaceSpan(ctx context.Context, in *SwInterfaceSpanDump) (RPCService_DumpSwInterfaceSpanClient, error)
	SwInterfaceSpanEnableDisable(ctx context.Context, in *SwInterfaceSpanEnableDisable) (*SwInterfaceSpanEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwInterfaceSpan(ctx context.Context, in *SwInterfaceSpanDump) (RPCService_DumpSwInterfaceSpanClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceSpanClient{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceSpanClient interface {
	Recv() (*SwInterfaceSpanDetails, error)
}

type serviceClient_DumpSwInterfaceSpanClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceSpanClient) Recv() (*SwInterfaceSpanDetails, error) {
	m := new(SwInterfaceSpanDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) SwInterfaceSpanEnableDisable(ctx context.Context, in *SwInterfaceSpanEnableDisable) (*SwInterfaceSpanEnableDisableReply, error) {
	out := new(SwInterfaceSpanEnableDisableReply)
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
