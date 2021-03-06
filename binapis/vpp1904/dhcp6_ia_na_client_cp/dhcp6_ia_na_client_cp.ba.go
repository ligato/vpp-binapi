// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/dhcp6_ia_na_client_cp.api.json

/*
Package dhcp6_ia_na_client_cp is a generated VPP binary API for 'dhcp6_ia_na_client_cp' module.

It consists of:
	  2 messages
	  1 service
*/
package dhcp6_ia_na_client_cp

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
	ModuleName = "dhcp6_ia_na_client_cp"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3fc84d7c
)

// DHCP6ClientEnableDisable represents VPP binary API message 'dhcp6_client_enable_disable'.
type DHCP6ClientEnableDisable struct {
	SwIfIndex uint32
	Enable    uint8
}

func (m *DHCP6ClientEnableDisable) Reset()                        { *m = DHCP6ClientEnableDisable{} }
func (*DHCP6ClientEnableDisable) GetMessageName() string          { return "dhcp6_client_enable_disable" }
func (*DHCP6ClientEnableDisable) GetCrcString() string            { return "a36fadc0" }
func (*DHCP6ClientEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// DHCP6ClientEnableDisableReply represents VPP binary API message 'dhcp6_client_enable_disable_reply'.
type DHCP6ClientEnableDisableReply struct {
	Retval int32
}

func (m *DHCP6ClientEnableDisableReply) Reset() { *m = DHCP6ClientEnableDisableReply{} }
func (*DHCP6ClientEnableDisableReply) GetMessageName() string {
	return "dhcp6_client_enable_disable_reply"
}
func (*DHCP6ClientEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*DHCP6ClientEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*DHCP6ClientEnableDisable)(nil), "dhcp6_ia_na_client_cp.DHCP6ClientEnableDisable")
	api.RegisterMessage((*DHCP6ClientEnableDisableReply)(nil), "dhcp6_ia_na_client_cp.DHCP6ClientEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DHCP6ClientEnableDisable)(nil),
		(*DHCP6ClientEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for dhcp6_ia_na_client_cp module.
type RPCService interface {
	DHCP6ClientEnableDisable(ctx context.Context, in *DHCP6ClientEnableDisable) (*DHCP6ClientEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DHCP6ClientEnableDisable(ctx context.Context, in *DHCP6ClientEnableDisable) (*DHCP6ClientEnableDisableReply, error) {
	out := new(DHCP6ClientEnableDisableReply)
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
