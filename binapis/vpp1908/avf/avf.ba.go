// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/plugins/avf.api.json

/*
Package avf is a generated VPP binary API for 'avf' module.

It consists of:
	  4 messages
	  2 services
*/
package avf

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
	ModuleName = "avf"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3dffe16a
)

// AvfCreate represents VPP binary API message 'avf_create'.
type AvfCreate struct {
	PciAddr    uint32
	EnableElog int32
	RxqNum     uint16
	RxqSize    uint16
	TxqSize    uint16
}

func (m *AvfCreate) Reset()                        { *m = AvfCreate{} }
func (*AvfCreate) GetMessageName() string          { return "avf_create" }
func (*AvfCreate) GetCrcString() string            { return "daab8ae2" }
func (*AvfCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// AvfCreateReply represents VPP binary API message 'avf_create_reply'.
type AvfCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (m *AvfCreateReply) Reset()                        { *m = AvfCreateReply{} }
func (*AvfCreateReply) GetMessageName() string          { return "avf_create_reply" }
func (*AvfCreateReply) GetCrcString() string            { return "fda5941f" }
func (*AvfCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// AvfDelete represents VPP binary API message 'avf_delete'.
type AvfDelete struct {
	SwIfIndex uint32
}

func (m *AvfDelete) Reset()                        { *m = AvfDelete{} }
func (*AvfDelete) GetMessageName() string          { return "avf_delete" }
func (*AvfDelete) GetCrcString() string            { return "529cb13f" }
func (*AvfDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// AvfDeleteReply represents VPP binary API message 'avf_delete_reply'.
type AvfDeleteReply struct {
	Retval int32
}

func (m *AvfDeleteReply) Reset()                        { *m = AvfDeleteReply{} }
func (*AvfDeleteReply) GetMessageName() string          { return "avf_delete_reply" }
func (*AvfDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*AvfDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*AvfCreate)(nil), "avf.AvfCreate")
	api.RegisterMessage((*AvfCreateReply)(nil), "avf.AvfCreateReply")
	api.RegisterMessage((*AvfDelete)(nil), "avf.AvfDelete")
	api.RegisterMessage((*AvfDeleteReply)(nil), "avf.AvfDeleteReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AvfCreate)(nil),
		(*AvfCreateReply)(nil),
		(*AvfDelete)(nil),
		(*AvfDeleteReply)(nil),
	}
}

// RPCService represents RPC service API for avf module.
type RPCService interface {
	AvfCreate(ctx context.Context, in *AvfCreate) (*AvfCreateReply, error)
	AvfDelete(ctx context.Context, in *AvfDelete) (*AvfDeleteReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) AvfCreate(ctx context.Context, in *AvfCreate) (*AvfCreateReply, error) {
	out := new(AvfCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AvfDelete(ctx context.Context, in *AvfDelete) (*AvfDeleteReply, error) {
	out := new(AvfDeleteReply)
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
