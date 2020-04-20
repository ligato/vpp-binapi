// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/vhost_user.api.json

/*
Package vhost_user is a generated VPP binary API for 'vhost_user' module.

It consists of:
	  8 messages
	  4 services
*/
package vhost_user

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
	ModuleName = "vhost_user"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x584ddcc0
)

// CreateVhostUserIf represents VPP binary API message 'create_vhost_user_if'.
type CreateVhostUserIf struct {
	IsServer            uint8
	SockFilename        []byte `struc:"[256]byte"`
	Renumber            uint8
	DisableMrgRxbuf     uint8
	DisableIndirectDesc uint8
	CustomDevInstance   uint32
	UseCustomMac        uint8
	MacAddress          []byte `struc:"[6]byte"`
	Tag                 []byte `struc:"[64]byte"`
}

func (m *CreateVhostUserIf) Reset()                        { *m = CreateVhostUserIf{} }
func (*CreateVhostUserIf) GetMessageName() string          { return "create_vhost_user_if" }
func (*CreateVhostUserIf) GetCrcString() string            { return "bd230b87" }
func (*CreateVhostUserIf) GetMessageType() api.MessageType { return api.RequestMessage }

// CreateVhostUserIfReply represents VPP binary API message 'create_vhost_user_if_reply'.
type CreateVhostUserIfReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (m *CreateVhostUserIfReply) Reset()                        { *m = CreateVhostUserIfReply{} }
func (*CreateVhostUserIfReply) GetMessageName() string          { return "create_vhost_user_if_reply" }
func (*CreateVhostUserIfReply) GetCrcString() string            { return "fda5941f" }
func (*CreateVhostUserIfReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// DeleteVhostUserIf represents VPP binary API message 'delete_vhost_user_if'.
type DeleteVhostUserIf struct {
	SwIfIndex uint32
}

func (m *DeleteVhostUserIf) Reset()                        { *m = DeleteVhostUserIf{} }
func (*DeleteVhostUserIf) GetMessageName() string          { return "delete_vhost_user_if" }
func (*DeleteVhostUserIf) GetCrcString() string            { return "529cb13f" }
func (*DeleteVhostUserIf) GetMessageType() api.MessageType { return api.RequestMessage }

// DeleteVhostUserIfReply represents VPP binary API message 'delete_vhost_user_if_reply'.
type DeleteVhostUserIfReply struct {
	Retval int32
}

func (m *DeleteVhostUserIfReply) Reset()                        { *m = DeleteVhostUserIfReply{} }
func (*DeleteVhostUserIfReply) GetMessageName() string          { return "delete_vhost_user_if_reply" }
func (*DeleteVhostUserIfReply) GetCrcString() string            { return "e8d4e804" }
func (*DeleteVhostUserIfReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// ModifyVhostUserIf represents VPP binary API message 'modify_vhost_user_if'.
type ModifyVhostUserIf struct {
	SwIfIndex         uint32
	IsServer          uint8
	SockFilename      []byte `struc:"[256]byte"`
	Renumber          uint8
	CustomDevInstance uint32
}

func (m *ModifyVhostUserIf) Reset()                        { *m = ModifyVhostUserIf{} }
func (*ModifyVhostUserIf) GetMessageName() string          { return "modify_vhost_user_if" }
func (*ModifyVhostUserIf) GetCrcString() string            { return "80a9eaaa" }
func (*ModifyVhostUserIf) GetMessageType() api.MessageType { return api.RequestMessage }

// ModifyVhostUserIfReply represents VPP binary API message 'modify_vhost_user_if_reply'.
type ModifyVhostUserIfReply struct {
	Retval int32
}

func (m *ModifyVhostUserIfReply) Reset()                        { *m = ModifyVhostUserIfReply{} }
func (*ModifyVhostUserIfReply) GetMessageName() string          { return "modify_vhost_user_if_reply" }
func (*ModifyVhostUserIfReply) GetCrcString() string            { return "e8d4e804" }
func (*ModifyVhostUserIfReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceVhostUserDetails represents VPP binary API message 'sw_interface_vhost_user_details'.
type SwInterfaceVhostUserDetails struct {
	SwIfIndex      uint32
	InterfaceName  []byte `struc:"[64]byte"`
	VirtioNetHdrSz uint32
	Features       uint64
	IsServer       uint8
	SockFilename   []byte `struc:"[256]byte"`
	NumRegions     uint32
	SockErrno      int32
}

func (m *SwInterfaceVhostUserDetails) Reset()                        { *m = SwInterfaceVhostUserDetails{} }
func (*SwInterfaceVhostUserDetails) GetMessageName() string          { return "sw_interface_vhost_user_details" }
func (*SwInterfaceVhostUserDetails) GetCrcString() string            { return "91ff3307" }
func (*SwInterfaceVhostUserDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceVhostUserDump represents VPP binary API message 'sw_interface_vhost_user_dump'.
type SwInterfaceVhostUserDump struct{}

func (m *SwInterfaceVhostUserDump) Reset()                        { *m = SwInterfaceVhostUserDump{} }
func (*SwInterfaceVhostUserDump) GetMessageName() string          { return "sw_interface_vhost_user_dump" }
func (*SwInterfaceVhostUserDump) GetCrcString() string            { return "51077d14" }
func (*SwInterfaceVhostUserDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*CreateVhostUserIf)(nil), "vhost_user.CreateVhostUserIf")
	api.RegisterMessage((*CreateVhostUserIfReply)(nil), "vhost_user.CreateVhostUserIfReply")
	api.RegisterMessage((*DeleteVhostUserIf)(nil), "vhost_user.DeleteVhostUserIf")
	api.RegisterMessage((*DeleteVhostUserIfReply)(nil), "vhost_user.DeleteVhostUserIfReply")
	api.RegisterMessage((*ModifyVhostUserIf)(nil), "vhost_user.ModifyVhostUserIf")
	api.RegisterMessage((*ModifyVhostUserIfReply)(nil), "vhost_user.ModifyVhostUserIfReply")
	api.RegisterMessage((*SwInterfaceVhostUserDetails)(nil), "vhost_user.SwInterfaceVhostUserDetails")
	api.RegisterMessage((*SwInterfaceVhostUserDump)(nil), "vhost_user.SwInterfaceVhostUserDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CreateVhostUserIf)(nil),
		(*CreateVhostUserIfReply)(nil),
		(*DeleteVhostUserIf)(nil),
		(*DeleteVhostUserIfReply)(nil),
		(*ModifyVhostUserIf)(nil),
		(*ModifyVhostUserIfReply)(nil),
		(*SwInterfaceVhostUserDetails)(nil),
		(*SwInterfaceVhostUserDump)(nil),
	}
}

// RPCService represents RPC service API for vhost_user module.
type RPCService interface {
	DumpSwInterfaceVhostUser(ctx context.Context, in *SwInterfaceVhostUserDump) (RPCService_DumpSwInterfaceVhostUserClient, error)
	CreateVhostUserIf(ctx context.Context, in *CreateVhostUserIf) (*CreateVhostUserIfReply, error)
	DeleteVhostUserIf(ctx context.Context, in *DeleteVhostUserIf) (*DeleteVhostUserIfReply, error)
	ModifyVhostUserIf(ctx context.Context, in *ModifyVhostUserIf) (*ModifyVhostUserIfReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwInterfaceVhostUser(ctx context.Context, in *SwInterfaceVhostUserDump) (RPCService_DumpSwInterfaceVhostUserClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceVhostUserClient{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceVhostUserClient interface {
	Recv() (*SwInterfaceVhostUserDetails, error)
}

type serviceClient_DumpSwInterfaceVhostUserClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceVhostUserClient) Recv() (*SwInterfaceVhostUserDetails, error) {
	m := new(SwInterfaceVhostUserDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) CreateVhostUserIf(ctx context.Context, in *CreateVhostUserIf) (*CreateVhostUserIfReply, error) {
	out := new(CreateVhostUserIfReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteVhostUserIf(ctx context.Context, in *DeleteVhostUserIf) (*DeleteVhostUserIfReply, error) {
	out := new(DeleteVhostUserIfReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ModifyVhostUserIf(ctx context.Context, in *ModifyVhostUserIf) (*ModifyVhostUserIfReply, error) {
	out := new(ModifyVhostUserIfReply)
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
