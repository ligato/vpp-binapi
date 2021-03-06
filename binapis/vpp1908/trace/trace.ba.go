// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/plugins/trace.api.json

/*
Package trace is a generated VPP binary API for 'trace' module.

It consists of:
	  6 messages
	  3 services
*/
package trace

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
	ModuleName = "trace"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3467fbcd
)

// TraceProfileAdd represents VPP binary API message 'trace_profile_add'.
type TraceProfileAdd struct {
	TraceType uint8
	NumElts   uint8
	TraceTsp  uint8
	NodeID    uint32
	AppData   uint32
}

func (m *TraceProfileAdd) Reset()                        { *m = TraceProfileAdd{} }
func (*TraceProfileAdd) GetMessageName() string          { return "trace_profile_add" }
func (*TraceProfileAdd) GetCrcString() string            { return "de08aa6d" }
func (*TraceProfileAdd) GetMessageType() api.MessageType { return api.RequestMessage }

// TraceProfileAddReply represents VPP binary API message 'trace_profile_add_reply'.
type TraceProfileAddReply struct {
	Retval int32
}

func (m *TraceProfileAddReply) Reset()                        { *m = TraceProfileAddReply{} }
func (*TraceProfileAddReply) GetMessageName() string          { return "trace_profile_add_reply" }
func (*TraceProfileAddReply) GetCrcString() string            { return "e8d4e804" }
func (*TraceProfileAddReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// TraceProfileDel represents VPP binary API message 'trace_profile_del'.
type TraceProfileDel struct{}

func (m *TraceProfileDel) Reset()                        { *m = TraceProfileDel{} }
func (*TraceProfileDel) GetMessageName() string          { return "trace_profile_del" }
func (*TraceProfileDel) GetCrcString() string            { return "51077d14" }
func (*TraceProfileDel) GetMessageType() api.MessageType { return api.RequestMessage }

// TraceProfileDelReply represents VPP binary API message 'trace_profile_del_reply'.
type TraceProfileDelReply struct {
	Retval int32
}

func (m *TraceProfileDelReply) Reset()                        { *m = TraceProfileDelReply{} }
func (*TraceProfileDelReply) GetMessageName() string          { return "trace_profile_del_reply" }
func (*TraceProfileDelReply) GetCrcString() string            { return "e8d4e804" }
func (*TraceProfileDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// TraceProfileShowConfig represents VPP binary API message 'trace_profile_show_config'.
type TraceProfileShowConfig struct{}

func (m *TraceProfileShowConfig) Reset()                        { *m = TraceProfileShowConfig{} }
func (*TraceProfileShowConfig) GetMessageName() string          { return "trace_profile_show_config" }
func (*TraceProfileShowConfig) GetCrcString() string            { return "51077d14" }
func (*TraceProfileShowConfig) GetMessageType() api.MessageType { return api.RequestMessage }

// TraceProfileShowConfigReply represents VPP binary API message 'trace_profile_show_config_reply'.
type TraceProfileShowConfigReply struct {
	Retval    int32
	TraceType uint8
	NumElts   uint8
	TraceTsp  uint8
	NodeID    uint32
	AppData   uint32
}

func (m *TraceProfileShowConfigReply) Reset()                        { *m = TraceProfileShowConfigReply{} }
func (*TraceProfileShowConfigReply) GetMessageName() string          { return "trace_profile_show_config_reply" }
func (*TraceProfileShowConfigReply) GetCrcString() string            { return "0f1d374c" }
func (*TraceProfileShowConfigReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*TraceProfileAdd)(nil), "trace.TraceProfileAdd")
	api.RegisterMessage((*TraceProfileAddReply)(nil), "trace.TraceProfileAddReply")
	api.RegisterMessage((*TraceProfileDel)(nil), "trace.TraceProfileDel")
	api.RegisterMessage((*TraceProfileDelReply)(nil), "trace.TraceProfileDelReply")
	api.RegisterMessage((*TraceProfileShowConfig)(nil), "trace.TraceProfileShowConfig")
	api.RegisterMessage((*TraceProfileShowConfigReply)(nil), "trace.TraceProfileShowConfigReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TraceProfileAdd)(nil),
		(*TraceProfileAddReply)(nil),
		(*TraceProfileDel)(nil),
		(*TraceProfileDelReply)(nil),
		(*TraceProfileShowConfig)(nil),
		(*TraceProfileShowConfigReply)(nil),
	}
}

// RPCService represents RPC service API for trace module.
type RPCService interface {
	TraceProfileAdd(ctx context.Context, in *TraceProfileAdd) (*TraceProfileAddReply, error)
	TraceProfileDel(ctx context.Context, in *TraceProfileDel) (*TraceProfileDelReply, error)
	TraceProfileShowConfig(ctx context.Context, in *TraceProfileShowConfig) (*TraceProfileShowConfigReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) TraceProfileAdd(ctx context.Context, in *TraceProfileAdd) (*TraceProfileAddReply, error) {
	out := new(TraceProfileAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TraceProfileDel(ctx context.Context, in *TraceProfileDel) (*TraceProfileDelReply, error) {
	out := new(TraceProfileDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TraceProfileShowConfig(ctx context.Context, in *TraceProfileShowConfig) (*TraceProfileShowConfigReply, error) {
	out := new(TraceProfileShowConfigReply)
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
