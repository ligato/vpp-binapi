// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/pg.api.json

/*
Package pg is a generated VPP binary API for 'pg' module.

It consists of:
	  6 messages
	  3 services
*/
package pg

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
	ModuleName = "pg"
	// APIVersion is the API version of this module.
	APIVersion = "1.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x953cbbb2
)

// PgCapture represents VPP binary API message 'pg_capture'.
type PgCapture struct {
	InterfaceID    uint32
	IsEnabled      uint8
	Count          uint32
	PcapNameLength uint32 `struc:"sizeof=PcapFileName"`
	PcapFileName   []byte
}

func (*PgCapture) GetMessageName() string {
	return "pg_capture"
}
func (*PgCapture) GetCrcString() string {
	return "453da78d"
}
func (*PgCapture) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PgCaptureReply represents VPP binary API message 'pg_capture_reply'.
type PgCaptureReply struct {
	Retval int32
}

func (*PgCaptureReply) GetMessageName() string {
	return "pg_capture_reply"
}
func (*PgCaptureReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PgCaptureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PgCreateInterface represents VPP binary API message 'pg_create_interface'.
type PgCreateInterface struct {
	InterfaceID uint32
	GsoEnabled  uint8
	GsoSize     uint32
}

func (*PgCreateInterface) GetMessageName() string {
	return "pg_create_interface"
}
func (*PgCreateInterface) GetCrcString() string {
	return "b1ecff05"
}
func (*PgCreateInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PgCreateInterfaceReply represents VPP binary API message 'pg_create_interface_reply'.
type PgCreateInterfaceReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*PgCreateInterfaceReply) GetMessageName() string {
	return "pg_create_interface_reply"
}
func (*PgCreateInterfaceReply) GetCrcString() string {
	return "fda5941f"
}
func (*PgCreateInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// PgEnableDisable represents VPP binary API message 'pg_enable_disable'.
type PgEnableDisable struct {
	IsEnabled        uint8
	StreamNameLength uint32 `struc:"sizeof=StreamName"`
	StreamName       []byte
}

func (*PgEnableDisable) GetMessageName() string {
	return "pg_enable_disable"
}
func (*PgEnableDisable) GetCrcString() string {
	return "0cb71d10"
}
func (*PgEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// PgEnableDisableReply represents VPP binary API message 'pg_enable_disable_reply'.
type PgEnableDisableReply struct {
	Retval int32
}

func (*PgEnableDisableReply) GetMessageName() string {
	return "pg_enable_disable_reply"
}
func (*PgEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PgEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*PgCapture)(nil), "pg.PgCapture")
	api.RegisterMessage((*PgCaptureReply)(nil), "pg.PgCaptureReply")
	api.RegisterMessage((*PgCreateInterface)(nil), "pg.PgCreateInterface")
	api.RegisterMessage((*PgCreateInterfaceReply)(nil), "pg.PgCreateInterfaceReply")
	api.RegisterMessage((*PgEnableDisable)(nil), "pg.PgEnableDisable")
	api.RegisterMessage((*PgEnableDisableReply)(nil), "pg.PgEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PgCapture)(nil),
		(*PgCaptureReply)(nil),
		(*PgCreateInterface)(nil),
		(*PgCreateInterfaceReply)(nil),
		(*PgEnableDisable)(nil),
		(*PgEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for pg module.
type RPCService interface {
	PgCapture(ctx context.Context, in *PgCapture) (*PgCaptureReply, error)
	PgCreateInterface(ctx context.Context, in *PgCreateInterface) (*PgCreateInterfaceReply, error)
	PgEnableDisable(ctx context.Context, in *PgEnableDisable) (*PgEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) PgCapture(ctx context.Context, in *PgCapture) (*PgCaptureReply, error) {
	out := new(PgCaptureReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PgCreateInterface(ctx context.Context, in *PgCreateInterface) (*PgCreateInterfaceReply, error) {
	out := new(PgCreateInterfaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PgEnableDisable(ctx context.Context, in *PgEnableDisable) (*PgEnableDisableReply, error) {
	out := new(PgEnableDisableReply)
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
