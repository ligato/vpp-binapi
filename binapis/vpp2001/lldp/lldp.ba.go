// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2001/.vppapi/core/lldp.api.json

/*
Package lldp is a generated VPP binary API for 'lldp' module.

It consists of:
	  4 messages
	  2 services
*/
package lldp

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
	ModuleName = "lldp"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xabb21dd0
)

// LldpConfig represents VPP binary API message 'lldp_config'.
type LldpConfig struct {
	SystemName []byte `struc:"[256]byte"`
	TxHold     uint32
	TxInterval uint32
}

func (m *LldpConfig) Reset()                        { *m = LldpConfig{} }
func (*LldpConfig) GetMessageName() string          { return "lldp_config" }
func (*LldpConfig) GetCrcString() string            { return "2410286f" }
func (*LldpConfig) GetMessageType() api.MessageType { return api.RequestMessage }

// LldpConfigReply represents VPP binary API message 'lldp_config_reply'.
type LldpConfigReply struct {
	Retval int32
}

func (m *LldpConfigReply) Reset()                        { *m = LldpConfigReply{} }
func (*LldpConfigReply) GetMessageName() string          { return "lldp_config_reply" }
func (*LldpConfigReply) GetCrcString() string            { return "e8d4e804" }
func (*LldpConfigReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceSetLldp represents VPP binary API message 'sw_interface_set_lldp'.
type SwInterfaceSetLldp struct {
	SwIfIndex uint32
	PortDesc  []byte `struc:"[256]byte"`
	MgmtIP4   []byte `struc:"[4]byte"`
	MgmtIP6   []byte `struc:"[16]byte"`
	MgmtOid   []byte `struc:"[128]byte"`
	Enable    uint8
}

func (m *SwInterfaceSetLldp) Reset()                        { *m = SwInterfaceSetLldp{} }
func (*SwInterfaceSetLldp) GetMessageName() string          { return "sw_interface_set_lldp" }
func (*SwInterfaceSetLldp) GetCrcString() string            { return "2d85d156" }
func (*SwInterfaceSetLldp) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetLldpReply represents VPP binary API message 'sw_interface_set_lldp_reply'.
type SwInterfaceSetLldpReply struct {
	Retval int32
}

func (m *SwInterfaceSetLldpReply) Reset()                        { *m = SwInterfaceSetLldpReply{} }
func (*SwInterfaceSetLldpReply) GetMessageName() string          { return "sw_interface_set_lldp_reply" }
func (*SwInterfaceSetLldpReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSetLldpReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*LldpConfig)(nil), "lldp.LldpConfig")
	api.RegisterMessage((*LldpConfigReply)(nil), "lldp.LldpConfigReply")
	api.RegisterMessage((*SwInterfaceSetLldp)(nil), "lldp.SwInterfaceSetLldp")
	api.RegisterMessage((*SwInterfaceSetLldpReply)(nil), "lldp.SwInterfaceSetLldpReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LldpConfig)(nil),
		(*LldpConfigReply)(nil),
		(*SwInterfaceSetLldp)(nil),
		(*SwInterfaceSetLldpReply)(nil),
	}
}

// RPCService represents RPC service API for lldp module.
type RPCService interface {
	LldpConfig(ctx context.Context, in *LldpConfig) (*LldpConfigReply, error)
	SwInterfaceSetLldp(ctx context.Context, in *SwInterfaceSetLldp) (*SwInterfaceSetLldpReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) LldpConfig(ctx context.Context, in *LldpConfig) (*LldpConfigReply, error) {
	out := new(LldpConfigReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceSetLldp(ctx context.Context, in *SwInterfaceSetLldp) (*SwInterfaceSetLldpReply, error) {
	out := new(SwInterfaceSetLldpReply)
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
