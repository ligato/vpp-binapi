// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/oam.api.json

/*
Package oam is a generated VPP binary API for 'oam' module.

It consists of:
	  5 messages
	  2 services
*/
package oam

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
	ModuleName = "oam"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x6aac8f57
)

// OamAddDel represents VPP binary API message 'oam_add_del'.
type OamAddDel struct {
	VrfID      uint32
	SrcAddress []byte `struc:"[4]byte"`
	DstAddress []byte `struc:"[4]byte"`
	IsAdd      uint8
}

func (m *OamAddDel) Reset()                        { *m = OamAddDel{} }
func (*OamAddDel) GetMessageName() string          { return "oam_add_del" }
func (*OamAddDel) GetCrcString() string            { return "3d7fcf96" }
func (*OamAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// OamAddDelReply represents VPP binary API message 'oam_add_del_reply'.
type OamAddDelReply struct {
	Retval int32
}

func (m *OamAddDelReply) Reset()                        { *m = OamAddDelReply{} }
func (*OamAddDelReply) GetMessageName() string          { return "oam_add_del_reply" }
func (*OamAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*OamAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// OamEvent represents VPP binary API message 'oam_event'.
type OamEvent struct {
	DstAddress []byte `struc:"[4]byte"`
	State      uint8
}

func (m *OamEvent) Reset()                        { *m = OamEvent{} }
func (*OamEvent) GetMessageName() string          { return "oam_event" }
func (*OamEvent) GetCrcString() string            { return "a450141f" }
func (*OamEvent) GetMessageType() api.MessageType { return api.OtherMessage }

// WantOamEvents represents VPP binary API message 'want_oam_events'.
type WantOamEvents struct {
	EnableDisable uint32
	PID           uint32
}

func (m *WantOamEvents) Reset()                        { *m = WantOamEvents{} }
func (*WantOamEvents) GetMessageName() string          { return "want_oam_events" }
func (*WantOamEvents) GetCrcString() string            { return "476f5a08" }
func (*WantOamEvents) GetMessageType() api.MessageType { return api.RequestMessage }

// WantOamEventsReply represents VPP binary API message 'want_oam_events_reply'.
type WantOamEventsReply struct {
	Retval int32
}

func (m *WantOamEventsReply) Reset()                        { *m = WantOamEventsReply{} }
func (*WantOamEventsReply) GetMessageName() string          { return "want_oam_events_reply" }
func (*WantOamEventsReply) GetCrcString() string            { return "e8d4e804" }
func (*WantOamEventsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*OamAddDel)(nil), "oam.OamAddDel")
	api.RegisterMessage((*OamAddDelReply)(nil), "oam.OamAddDelReply")
	api.RegisterMessage((*OamEvent)(nil), "oam.OamEvent")
	api.RegisterMessage((*WantOamEvents)(nil), "oam.WantOamEvents")
	api.RegisterMessage((*WantOamEventsReply)(nil), "oam.WantOamEventsReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*OamAddDel)(nil),
		(*OamAddDelReply)(nil),
		(*OamEvent)(nil),
		(*WantOamEvents)(nil),
		(*WantOamEventsReply)(nil),
	}
}

// RPCService represents RPC service API for oam module.
type RPCService interface {
	OamAddDel(ctx context.Context, in *OamAddDel) (*OamAddDelReply, error)
	WantOamEvents(ctx context.Context, in *WantOamEvents) (*WantOamEventsReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) OamAddDel(ctx context.Context, in *OamAddDel) (*OamAddDelReply, error) {
	out := new(OamAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) WantOamEvents(ctx context.Context, in *WantOamEvents) (*WantOamEventsReply, error) {
	out := new(WantOamEventsReply)
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
