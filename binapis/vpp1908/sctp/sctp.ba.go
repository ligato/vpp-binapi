// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/plugins/sctp.api.json

/*
Package sctp is a generated VPP binary API for 'sctp' module.

It consists of:
	  6 messages
	  3 services
*/
package sctp

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
	ModuleName = "sctp"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x301e027b
)

// SctpAddSrcDstConnection represents VPP binary API message 'sctp_add_src_dst_connection'.
type SctpAddSrcDstConnection struct {
	IsIPv6     uint8
	VrfID      uint32
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
}

func (m *SctpAddSrcDstConnection) Reset()                        { *m = SctpAddSrcDstConnection{} }
func (*SctpAddSrcDstConnection) GetMessageName() string          { return "sctp_add_src_dst_connection" }
func (*SctpAddSrcDstConnection) GetCrcString() string            { return "45c59b2f" }
func (*SctpAddSrcDstConnection) GetMessageType() api.MessageType { return api.RequestMessage }

// SctpAddSrcDstConnectionReply represents VPP binary API message 'sctp_add_src_dst_connection_reply'.
type SctpAddSrcDstConnectionReply struct {
	Retval int32
}

func (m *SctpAddSrcDstConnectionReply) Reset() { *m = SctpAddSrcDstConnectionReply{} }
func (*SctpAddSrcDstConnectionReply) GetMessageName() string {
	return "sctp_add_src_dst_connection_reply"
}
func (*SctpAddSrcDstConnectionReply) GetCrcString() string            { return "e8d4e804" }
func (*SctpAddSrcDstConnectionReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SctpConfig represents VPP binary API message 'sctp_config'.
type SctpConfig struct {
	NeverDelaySack uint8
	NeverBundle    uint8
}

func (m *SctpConfig) Reset()                        { *m = SctpConfig{} }
func (*SctpConfig) GetMessageName() string          { return "sctp_config" }
func (*SctpConfig) GetCrcString() string            { return "7617eced" }
func (*SctpConfig) GetMessageType() api.MessageType { return api.RequestMessage }

// SctpConfigReply represents VPP binary API message 'sctp_config_reply'.
type SctpConfigReply struct {
	Retval int32
}

func (m *SctpConfigReply) Reset()                        { *m = SctpConfigReply{} }
func (*SctpConfigReply) GetMessageName() string          { return "sctp_config_reply" }
func (*SctpConfigReply) GetCrcString() string            { return "e8d4e804" }
func (*SctpConfigReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SctpDelSrcDstConnection represents VPP binary API message 'sctp_del_src_dst_connection'.
type SctpDelSrcDstConnection struct {
	IsIPv6     uint8
	VrfID      uint32
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
}

func (m *SctpDelSrcDstConnection) Reset()                        { *m = SctpDelSrcDstConnection{} }
func (*SctpDelSrcDstConnection) GetMessageName() string          { return "sctp_del_src_dst_connection" }
func (*SctpDelSrcDstConnection) GetCrcString() string            { return "45c59b2f" }
func (*SctpDelSrcDstConnection) GetMessageType() api.MessageType { return api.RequestMessage }

// SctpDelSrcDstConnectionReply represents VPP binary API message 'sctp_del_src_dst_connection_reply'.
type SctpDelSrcDstConnectionReply struct {
	Retval int32
}

func (m *SctpDelSrcDstConnectionReply) Reset() { *m = SctpDelSrcDstConnectionReply{} }
func (*SctpDelSrcDstConnectionReply) GetMessageName() string {
	return "sctp_del_src_dst_connection_reply"
}
func (*SctpDelSrcDstConnectionReply) GetCrcString() string            { return "e8d4e804" }
func (*SctpDelSrcDstConnectionReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*SctpAddSrcDstConnection)(nil), "sctp.SctpAddSrcDstConnection")
	api.RegisterMessage((*SctpAddSrcDstConnectionReply)(nil), "sctp.SctpAddSrcDstConnectionReply")
	api.RegisterMessage((*SctpConfig)(nil), "sctp.SctpConfig")
	api.RegisterMessage((*SctpConfigReply)(nil), "sctp.SctpConfigReply")
	api.RegisterMessage((*SctpDelSrcDstConnection)(nil), "sctp.SctpDelSrcDstConnection")
	api.RegisterMessage((*SctpDelSrcDstConnectionReply)(nil), "sctp.SctpDelSrcDstConnectionReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SctpAddSrcDstConnection)(nil),
		(*SctpAddSrcDstConnectionReply)(nil),
		(*SctpConfig)(nil),
		(*SctpConfigReply)(nil),
		(*SctpDelSrcDstConnection)(nil),
		(*SctpDelSrcDstConnectionReply)(nil),
	}
}

// RPCService represents RPC service API for sctp module.
type RPCService interface {
	SctpAddSrcDstConnection(ctx context.Context, in *SctpAddSrcDstConnection) (*SctpAddSrcDstConnectionReply, error)
	SctpConfig(ctx context.Context, in *SctpConfig) (*SctpConfigReply, error)
	SctpDelSrcDstConnection(ctx context.Context, in *SctpDelSrcDstConnection) (*SctpDelSrcDstConnectionReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) SctpAddSrcDstConnection(ctx context.Context, in *SctpAddSrcDstConnection) (*SctpAddSrcDstConnectionReply, error) {
	out := new(SctpAddSrcDstConnectionReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SctpConfig(ctx context.Context, in *SctpConfig) (*SctpConfigReply, error) {
	out := new(SctpConfigReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SctpDelSrcDstConnection(ctx context.Context, in *SctpDelSrcDstConnection) (*SctpDelSrcDstConnectionReply, error) {
	out := new(SctpDelSrcDstConnectionReply)
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
