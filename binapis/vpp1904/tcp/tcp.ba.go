// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/tcp.api.json

/*
Package tcp is a generated VPP binary API for 'tcp' module.

It consists of:
	  2 messages
	  1 service
*/
package tcp

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
	ModuleName = "tcp"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3e216e88
)

// TCPConfigureSrcAddresses represents VPP binary API message 'tcp_configure_src_addresses'.
type TCPConfigureSrcAddresses struct {
	IsIPv6       uint8
	VrfID        uint32
	FirstAddress []byte `struc:"[16]byte"`
	LastAddress  []byte `struc:"[16]byte"`
}

func (m *TCPConfigureSrcAddresses) Reset()                        { *m = TCPConfigureSrcAddresses{} }
func (*TCPConfigureSrcAddresses) GetMessageName() string          { return "tcp_configure_src_addresses" }
func (*TCPConfigureSrcAddresses) GetCrcString() string            { return "8c1f804f" }
func (*TCPConfigureSrcAddresses) GetMessageType() api.MessageType { return api.RequestMessage }

// TCPConfigureSrcAddressesReply represents VPP binary API message 'tcp_configure_src_addresses_reply'.
type TCPConfigureSrcAddressesReply struct {
	Retval int32
}

func (m *TCPConfigureSrcAddressesReply) Reset() { *m = TCPConfigureSrcAddressesReply{} }
func (*TCPConfigureSrcAddressesReply) GetMessageName() string {
	return "tcp_configure_src_addresses_reply"
}
func (*TCPConfigureSrcAddressesReply) GetCrcString() string            { return "e8d4e804" }
func (*TCPConfigureSrcAddressesReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*TCPConfigureSrcAddresses)(nil), "tcp.TCPConfigureSrcAddresses")
	api.RegisterMessage((*TCPConfigureSrcAddressesReply)(nil), "tcp.TCPConfigureSrcAddressesReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TCPConfigureSrcAddresses)(nil),
		(*TCPConfigureSrcAddressesReply)(nil),
	}
}

// RPCService represents RPC service API for tcp module.
type RPCService interface {
	TCPConfigureSrcAddresses(ctx context.Context, in *TCPConfigureSrcAddresses) (*TCPConfigureSrcAddressesReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) TCPConfigureSrcAddresses(ctx context.Context, in *TCPConfigureSrcAddresses) (*TCPConfigureSrcAddressesReply, error) {
	out := new(TCPConfigureSrcAddressesReply)
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
