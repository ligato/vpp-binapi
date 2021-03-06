// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/plugins/ioam_vxlan_gpe.api.json

/*
Package ioam_vxlan_gpe is a generated VPP binary API for 'ioam_vxlan_gpe' module.

It consists of:
	 12 messages
	  6 services
*/
package ioam_vxlan_gpe

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
	ModuleName = "ioam_vxlan_gpe"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3c50f9af
)

// VxlanGpeIoamDisable represents VPP binary API message 'vxlan_gpe_ioam_disable'.
type VxlanGpeIoamDisable struct {
	ID uint16
}

func (m *VxlanGpeIoamDisable) Reset()                        { *m = VxlanGpeIoamDisable{} }
func (*VxlanGpeIoamDisable) GetMessageName() string          { return "vxlan_gpe_ioam_disable" }
func (*VxlanGpeIoamDisable) GetCrcString() string            { return "6b16a45e" }
func (*VxlanGpeIoamDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamDisableReply represents VPP binary API message 'vxlan_gpe_ioam_disable_reply'.
type VxlanGpeIoamDisableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamDisableReply) Reset()                        { *m = VxlanGpeIoamDisableReply{} }
func (*VxlanGpeIoamDisableReply) GetMessageName() string          { return "vxlan_gpe_ioam_disable_reply" }
func (*VxlanGpeIoamDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamEnable represents VPP binary API message 'vxlan_gpe_ioam_enable'.
type VxlanGpeIoamEnable struct {
	ID          uint16
	TracePpc    uint8
	PowEnable   uint8
	TraceEnable uint8
}

func (m *VxlanGpeIoamEnable) Reset()                        { *m = VxlanGpeIoamEnable{} }
func (*VxlanGpeIoamEnable) GetMessageName() string          { return "vxlan_gpe_ioam_enable" }
func (*VxlanGpeIoamEnable) GetCrcString() string            { return "7727876d" }
func (*VxlanGpeIoamEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamEnableReply represents VPP binary API message 'vxlan_gpe_ioam_enable_reply'.
type VxlanGpeIoamEnableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamEnableReply) Reset()                        { *m = VxlanGpeIoamEnableReply{} }
func (*VxlanGpeIoamEnableReply) GetMessageName() string          { return "vxlan_gpe_ioam_enable_reply" }
func (*VxlanGpeIoamEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamTransitDisable represents VPP binary API message 'vxlan_gpe_ioam_transit_disable'.
type VxlanGpeIoamTransitDisable struct {
	OuterFibIndex uint32
	DstAddr       []byte `struc:"[16]byte"`
	IsIPv6        uint8
}

func (m *VxlanGpeIoamTransitDisable) Reset()                        { *m = VxlanGpeIoamTransitDisable{} }
func (*VxlanGpeIoamTransitDisable) GetMessageName() string          { return "vxlan_gpe_ioam_transit_disable" }
func (*VxlanGpeIoamTransitDisable) GetCrcString() string            { return "7c6ff202" }
func (*VxlanGpeIoamTransitDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamTransitDisableReply represents VPP binary API message 'vxlan_gpe_ioam_transit_disable_reply'.
type VxlanGpeIoamTransitDisableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamTransitDisableReply) Reset() { *m = VxlanGpeIoamTransitDisableReply{} }
func (*VxlanGpeIoamTransitDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_disable_reply"
}
func (*VxlanGpeIoamTransitDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamTransitDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamTransitEnable represents VPP binary API message 'vxlan_gpe_ioam_transit_enable'.
type VxlanGpeIoamTransitEnable struct {
	OuterFibIndex uint32
	DstAddr       []byte `struc:"[16]byte"`
	IsIPv6        uint8
}

func (m *VxlanGpeIoamTransitEnable) Reset()                        { *m = VxlanGpeIoamTransitEnable{} }
func (*VxlanGpeIoamTransitEnable) GetMessageName() string          { return "vxlan_gpe_ioam_transit_enable" }
func (*VxlanGpeIoamTransitEnable) GetCrcString() string            { return "7c6ff202" }
func (*VxlanGpeIoamTransitEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamTransitEnableReply represents VPP binary API message 'vxlan_gpe_ioam_transit_enable_reply'.
type VxlanGpeIoamTransitEnableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamTransitEnableReply) Reset() { *m = VxlanGpeIoamTransitEnableReply{} }
func (*VxlanGpeIoamTransitEnableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_transit_enable_reply"
}
func (*VxlanGpeIoamTransitEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamTransitEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamVniDisable represents VPP binary API message 'vxlan_gpe_ioam_vni_disable'.
type VxlanGpeIoamVniDisable struct {
	Vni    uint32
	Local  []byte `struc:"[16]byte"`
	Remote []byte `struc:"[16]byte"`
	IsIPv6 uint8
}

func (m *VxlanGpeIoamVniDisable) Reset()                        { *m = VxlanGpeIoamVniDisable{} }
func (*VxlanGpeIoamVniDisable) GetMessageName() string          { return "vxlan_gpe_ioam_vni_disable" }
func (*VxlanGpeIoamVniDisable) GetCrcString() string            { return "6d93fc5d" }
func (*VxlanGpeIoamVniDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamVniDisableReply represents VPP binary API message 'vxlan_gpe_ioam_vni_disable_reply'.
type VxlanGpeIoamVniDisableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamVniDisableReply) Reset() { *m = VxlanGpeIoamVniDisableReply{} }
func (*VxlanGpeIoamVniDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_vni_disable_reply"
}
func (*VxlanGpeIoamVniDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamVniDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VxlanGpeIoamVniEnable represents VPP binary API message 'vxlan_gpe_ioam_vni_enable'.
type VxlanGpeIoamVniEnable struct {
	Vni    uint32
	Local  []byte `struc:"[16]byte"`
	Remote []byte `struc:"[16]byte"`
	IsIPv6 uint8
}

func (m *VxlanGpeIoamVniEnable) Reset()                        { *m = VxlanGpeIoamVniEnable{} }
func (*VxlanGpeIoamVniEnable) GetMessageName() string          { return "vxlan_gpe_ioam_vni_enable" }
func (*VxlanGpeIoamVniEnable) GetCrcString() string            { return "6d93fc5d" }
func (*VxlanGpeIoamVniEnable) GetMessageType() api.MessageType { return api.RequestMessage }

// VxlanGpeIoamVniEnableReply represents VPP binary API message 'vxlan_gpe_ioam_vni_enable_reply'.
type VxlanGpeIoamVniEnableReply struct {
	Retval int32
}

func (m *VxlanGpeIoamVniEnableReply) Reset()                        { *m = VxlanGpeIoamVniEnableReply{} }
func (*VxlanGpeIoamVniEnableReply) GetMessageName() string          { return "vxlan_gpe_ioam_vni_enable_reply" }
func (*VxlanGpeIoamVniEnableReply) GetCrcString() string            { return "e8d4e804" }
func (*VxlanGpeIoamVniEnableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*VxlanGpeIoamDisable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamDisable")
	api.RegisterMessage((*VxlanGpeIoamDisableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamDisableReply")
	api.RegisterMessage((*VxlanGpeIoamEnable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamEnable")
	api.RegisterMessage((*VxlanGpeIoamEnableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamEnableReply")
	api.RegisterMessage((*VxlanGpeIoamTransitDisable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitDisable")
	api.RegisterMessage((*VxlanGpeIoamTransitDisableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitDisableReply")
	api.RegisterMessage((*VxlanGpeIoamTransitEnable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitEnable")
	api.RegisterMessage((*VxlanGpeIoamTransitEnableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamTransitEnableReply")
	api.RegisterMessage((*VxlanGpeIoamVniDisable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniDisable")
	api.RegisterMessage((*VxlanGpeIoamVniDisableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniDisableReply")
	api.RegisterMessage((*VxlanGpeIoamVniEnable)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniEnable")
	api.RegisterMessage((*VxlanGpeIoamVniEnableReply)(nil), "ioam_vxlan_gpe.VxlanGpeIoamVniEnableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*VxlanGpeIoamDisable)(nil),
		(*VxlanGpeIoamDisableReply)(nil),
		(*VxlanGpeIoamEnable)(nil),
		(*VxlanGpeIoamEnableReply)(nil),
		(*VxlanGpeIoamTransitDisable)(nil),
		(*VxlanGpeIoamTransitDisableReply)(nil),
		(*VxlanGpeIoamTransitEnable)(nil),
		(*VxlanGpeIoamTransitEnableReply)(nil),
		(*VxlanGpeIoamVniDisable)(nil),
		(*VxlanGpeIoamVniDisableReply)(nil),
		(*VxlanGpeIoamVniEnable)(nil),
		(*VxlanGpeIoamVniEnableReply)(nil),
	}
}

// RPCService represents RPC service API for ioam_vxlan_gpe module.
type RPCService interface {
	VxlanGpeIoamDisable(ctx context.Context, in *VxlanGpeIoamDisable) (*VxlanGpeIoamDisableReply, error)
	VxlanGpeIoamEnable(ctx context.Context, in *VxlanGpeIoamEnable) (*VxlanGpeIoamEnableReply, error)
	VxlanGpeIoamTransitDisable(ctx context.Context, in *VxlanGpeIoamTransitDisable) (*VxlanGpeIoamTransitDisableReply, error)
	VxlanGpeIoamTransitEnable(ctx context.Context, in *VxlanGpeIoamTransitEnable) (*VxlanGpeIoamTransitEnableReply, error)
	VxlanGpeIoamVniDisable(ctx context.Context, in *VxlanGpeIoamVniDisable) (*VxlanGpeIoamVniDisableReply, error)
	VxlanGpeIoamVniEnable(ctx context.Context, in *VxlanGpeIoamVniEnable) (*VxlanGpeIoamVniEnableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) VxlanGpeIoamDisable(ctx context.Context, in *VxlanGpeIoamDisable) (*VxlanGpeIoamDisableReply, error) {
	out := new(VxlanGpeIoamDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamEnable(ctx context.Context, in *VxlanGpeIoamEnable) (*VxlanGpeIoamEnableReply, error) {
	out := new(VxlanGpeIoamEnableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamTransitDisable(ctx context.Context, in *VxlanGpeIoamTransitDisable) (*VxlanGpeIoamTransitDisableReply, error) {
	out := new(VxlanGpeIoamTransitDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamTransitEnable(ctx context.Context, in *VxlanGpeIoamTransitEnable) (*VxlanGpeIoamTransitEnableReply, error) {
	out := new(VxlanGpeIoamTransitEnableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamVniDisable(ctx context.Context, in *VxlanGpeIoamVniDisable) (*VxlanGpeIoamVniDisableReply, error) {
	out := new(VxlanGpeIoamVniDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VxlanGpeIoamVniEnable(ctx context.Context, in *VxlanGpeIoamVniEnable) (*VxlanGpeIoamVniEnableReply, error) {
	out := new(VxlanGpeIoamVniEnableReply)
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
