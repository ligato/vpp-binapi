// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/plugins/l2e.api.json

/*
Package l2e is a generated VPP binary API for 'l2e' module.

It consists of:
	  2 messages
	  1 service
*/
package l2e

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
	ModuleName = "l2e"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3fc84d7c
)

// L2Emulation represents VPP binary API message 'l2_emulation'.
type L2Emulation struct {
	SwIfIndex uint32
	Enable    uint8
}

func (m *L2Emulation) Reset()                        { *m = L2Emulation{} }
func (*L2Emulation) GetMessageName() string          { return "l2_emulation" }
func (*L2Emulation) GetCrcString() string            { return "a36fadc0" }
func (*L2Emulation) GetMessageType() api.MessageType { return api.RequestMessage }

// L2EmulationReply represents VPP binary API message 'l2_emulation_reply'.
type L2EmulationReply struct {
	Retval int32
}

func (m *L2EmulationReply) Reset()                        { *m = L2EmulationReply{} }
func (*L2EmulationReply) GetMessageName() string          { return "l2_emulation_reply" }
func (*L2EmulationReply) GetCrcString() string            { return "e8d4e804" }
func (*L2EmulationReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*L2Emulation)(nil), "l2e.L2Emulation")
	api.RegisterMessage((*L2EmulationReply)(nil), "l2e.L2EmulationReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*L2Emulation)(nil),
		(*L2EmulationReply)(nil),
	}
}

// RPCService represents RPC service API for l2e module.
type RPCService interface {
	L2Emulation(ctx context.Context, in *L2Emulation) (*L2EmulationReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) L2Emulation(ctx context.Context, in *L2Emulation) (*L2EmulationReply, error) {
	out := new(L2EmulationReply)
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
