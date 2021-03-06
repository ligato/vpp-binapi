// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/core/netmap.api.json

/*
Package netmap is a generated VPP binary API for 'netmap' module.

It consists of:
	  4 messages
	  2 services
*/
package netmap

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
	ModuleName = "netmap"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x6ae7e12e
)

// NetmapCreate represents VPP binary API message 'netmap_create'.
type NetmapCreate struct {
	NetmapIfName    []byte `struc:"[64]byte"`
	HwAddr          []byte `struc:"[6]byte"`
	UseRandomHwAddr uint8
	IsPipe          uint8
	IsMaster        uint8
}

func (m *NetmapCreate) Reset()                        { *m = NetmapCreate{} }
func (*NetmapCreate) GetMessageName() string          { return "netmap_create" }
func (*NetmapCreate) GetCrcString() string            { return "5299044e" }
func (*NetmapCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// NetmapCreateReply represents VPP binary API message 'netmap_create_reply'.
type NetmapCreateReply struct {
	Retval int32
}

func (m *NetmapCreateReply) Reset()                        { *m = NetmapCreateReply{} }
func (*NetmapCreateReply) GetMessageName() string          { return "netmap_create_reply" }
func (*NetmapCreateReply) GetCrcString() string            { return "e8d4e804" }
func (*NetmapCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// NetmapDelete represents VPP binary API message 'netmap_delete'.
type NetmapDelete struct {
	NetmapIfName []byte `struc:"[64]byte"`
}

func (m *NetmapDelete) Reset()                        { *m = NetmapDelete{} }
func (*NetmapDelete) GetMessageName() string          { return "netmap_delete" }
func (*NetmapDelete) GetCrcString() string            { return "a8b6c201" }
func (*NetmapDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// NetmapDeleteReply represents VPP binary API message 'netmap_delete_reply'.
type NetmapDeleteReply struct {
	Retval int32
}

func (m *NetmapDeleteReply) Reset()                        { *m = NetmapDeleteReply{} }
func (*NetmapDeleteReply) GetMessageName() string          { return "netmap_delete_reply" }
func (*NetmapDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*NetmapDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*NetmapCreate)(nil), "netmap.NetmapCreate")
	api.RegisterMessage((*NetmapCreateReply)(nil), "netmap.NetmapCreateReply")
	api.RegisterMessage((*NetmapDelete)(nil), "netmap.NetmapDelete")
	api.RegisterMessage((*NetmapDeleteReply)(nil), "netmap.NetmapDeleteReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*NetmapCreate)(nil),
		(*NetmapCreateReply)(nil),
		(*NetmapDelete)(nil),
		(*NetmapDeleteReply)(nil),
	}
}

// RPCService represents RPC service API for netmap module.
type RPCService interface {
	NetmapCreate(ctx context.Context, in *NetmapCreate) (*NetmapCreateReply, error)
	NetmapDelete(ctx context.Context, in *NetmapDelete) (*NetmapDeleteReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) NetmapCreate(ctx context.Context, in *NetmapCreate) (*NetmapCreateReply, error) {
	out := new(NetmapCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NetmapDelete(ctx context.Context, in *NetmapDelete) (*NetmapDeleteReply, error) {
	out := new(NetmapDeleteReply)
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
