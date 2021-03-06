// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2001/.vppapi/plugins/http_static.api.json

/*
Package http_static is a generated VPP binary API for 'http_static' module.

It consists of:
	  2 messages
	  1 service
*/
package http_static

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
	ModuleName = "http_static"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd29e72e9
)

// HTTPStaticEnable represents VPP binary API message 'http_static_enable'.
type HTTPStaticEnable struct {
	FifoSize           uint32
	CacheSizeLimit     uint32
	PreallocFifos      uint32
	PrivateSegmentSize uint32
	WwwRoot            string `struc:"[256]byte"`
	URI                string `struc:"[256]byte"`
}

func (m *HTTPStaticEnable) Reset()                        { *m = HTTPStaticEnable{} }
func (*HTTPStaticEnable) GetMessageName() string          { return "http_static_enable" }
func (*HTTPStaticEnable) GetCrcString() string            { return "075f8292" }
func (*HTTPStaticEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// HTTPStaticEnableReply represents VPP binary API message 'http_static_enable_reply'.
type HTTPStaticEnableReply struct {
	Retval int32
}

func (m *HTTPStaticEnableReply) Reset()                        { *m = HTTPStaticEnableReply{} }
func (*HTTPStaticEnableReply) GetMessageName() string          { return "http_static_enable_reply" }
func (*HTTPStaticEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*HTTPStaticEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*HTTPStaticEnable)(nil), "http_static.HTTPStaticEnable")
	api.RegisterMessage((*HTTPStaticEnableReply)(nil), "http_static.HTTPStaticEnableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*HTTPStaticEnable)(nil),
		(*HTTPStaticEnableReply)(nil),
	}
}

// RPCService represents RPC service API for http_static module.
type RPCService interface {
	HTTPStaticEnable(ctx context.Context, in *HTTPStaticEnable) (*HTTPStaticEnableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) HTTPStaticEnable(ctx context.Context, in *HTTPStaticEnable) (*HTTPStaticEnableReply, error) {
	out := new(HTTPStaticEnableReply)
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
