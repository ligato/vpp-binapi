// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/plugins/ct6.api.json

/*
Package ct6 is a generated VPP binary API for 'ct6' module.

It consists of:
	  2 messages
	  1 service
*/
package ct6

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
	ModuleName = "ct6"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xf5b1e79f
)

// Ct6EnableDisable represents VPP binary API message 'ct6_enable_disable'.
type Ct6EnableDisable struct {
	EnableDisable uint8
	IsInside      uint8
	SwIfIndex     uint32
}

func (m *Ct6EnableDisable) Reset()                        { *m = Ct6EnableDisable{} }
func (*Ct6EnableDisable) GetMessageName() string          { return "ct6_enable_disable" }
func (*Ct6EnableDisable) GetCrcString() string            { return "8c8159a0" }
func (*Ct6EnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// Ct6EnableDisableReply represents VPP binary API message 'ct6_enable_disable_reply'.
type Ct6EnableDisableReply struct {
	Retval int32
}

func (m *Ct6EnableDisableReply) Reset()                        { *m = Ct6EnableDisableReply{} }
func (*Ct6EnableDisableReply) GetMessageName() string          { return "ct6_enable_disable_reply" }
func (*Ct6EnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*Ct6EnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*Ct6EnableDisable)(nil), "ct6.Ct6EnableDisable")
	api.RegisterMessage((*Ct6EnableDisableReply)(nil), "ct6.Ct6EnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Ct6EnableDisable)(nil),
		(*Ct6EnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for ct6 module.
type RPCService interface {
	Ct6EnableDisable(ctx context.Context, in *Ct6EnableDisable) (*Ct6EnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) Ct6EnableDisable(ctx context.Context, in *Ct6EnableDisable) (*Ct6EnableDisableReply, error) {
	out := new(Ct6EnableDisableReply)
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
