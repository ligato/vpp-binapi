// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/sr_mpls.api.json

/*
Package sr_mpls is a generated VPP binary API for 'sr_mpls' module.

It consists of:
	 10 messages
	  5 services
*/
package sr_mpls

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
	ModuleName = "sr_mpls"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x93240385
)

// SrMplsPolicyAdd represents VPP binary API message 'sr_mpls_policy_add'.
type SrMplsPolicyAdd struct {
	Bsid      uint32
	Weight    uint32
	Type      uint8
	NSegments uint8 `struc:"sizeof=Segments"`
	Segments  []uint32
}

func (*SrMplsPolicyAdd) GetMessageName() string {
	return "sr_mpls_policy_add"
}
func (*SrMplsPolicyAdd) GetCrcString() string {
	return "6f5b21cc"
}
func (*SrMplsPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyAddReply represents VPP binary API message 'sr_mpls_policy_add_reply'.
type SrMplsPolicyAddReply struct {
	Retval int32
}

func (*SrMplsPolicyAddReply) GetMessageName() string {
	return "sr_mpls_policy_add_reply"
}
func (*SrMplsPolicyAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsPolicyAssignEndpointColor represents VPP binary API message 'sr_mpls_policy_assign_endpoint_color'.
type SrMplsPolicyAssignEndpointColor struct {
	Bsid         uint32
	Endpoint     []byte `struc:"[16]byte"`
	EndpointType uint8
	Color        uint32
}

func (*SrMplsPolicyAssignEndpointColor) GetMessageName() string {
	return "sr_mpls_policy_assign_endpoint_color"
}
func (*SrMplsPolicyAssignEndpointColor) GetCrcString() string {
	return "6c82a6da"
}
func (*SrMplsPolicyAssignEndpointColor) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyAssignEndpointColorReply represents VPP binary API message 'sr_mpls_policy_assign_endpoint_color_reply'.
type SrMplsPolicyAssignEndpointColorReply struct {
	Retval int32
}

func (*SrMplsPolicyAssignEndpointColorReply) GetMessageName() string {
	return "sr_mpls_policy_assign_endpoint_color_reply"
}
func (*SrMplsPolicyAssignEndpointColorReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyAssignEndpointColorReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsPolicyDel represents VPP binary API message 'sr_mpls_policy_del'.
type SrMplsPolicyDel struct {
	Bsid uint32
}

func (*SrMplsPolicyDel) GetMessageName() string {
	return "sr_mpls_policy_del"
}
func (*SrMplsPolicyDel) GetCrcString() string {
	return "e29d34fa"
}
func (*SrMplsPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyDelReply represents VPP binary API message 'sr_mpls_policy_del_reply'.
type SrMplsPolicyDelReply struct {
	Retval int32
}

func (*SrMplsPolicyDelReply) GetMessageName() string {
	return "sr_mpls_policy_del_reply"
}
func (*SrMplsPolicyDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsPolicyMod represents VPP binary API message 'sr_mpls_policy_mod'.
type SrMplsPolicyMod struct {
	Bsid      uint32
	Operation uint8
	SlIndex   uint32
	Weight    uint32
	NSegments uint8 `struc:"sizeof=Segments"`
	Segments  []uint32
}

func (*SrMplsPolicyMod) GetMessageName() string {
	return "sr_mpls_policy_mod"
}
func (*SrMplsPolicyMod) GetCrcString() string {
	return "09d338ac"
}
func (*SrMplsPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsPolicyModReply represents VPP binary API message 'sr_mpls_policy_mod_reply'.
type SrMplsPolicyModReply struct {
	Retval int32
}

func (*SrMplsPolicyModReply) GetMessageName() string {
	return "sr_mpls_policy_mod_reply"
}
func (*SrMplsPolicyModReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SrMplsSteeringAddDel represents VPP binary API message 'sr_mpls_steering_add_del'.
type SrMplsSteeringAddDel struct {
	IsDel       uint8
	Bsid        uint32
	TableID     uint32
	PrefixAddr  []byte `struc:"[16]byte"`
	MaskWidth   uint32
	TrafficType uint8
	NextHop     []byte `struc:"[16]byte"`
	NhType      uint8
	Color       uint32
	CoBits      uint8
	VPNLabel    uint32
}

func (*SrMplsSteeringAddDel) GetMessageName() string {
	return "sr_mpls_steering_add_del"
}
func (*SrMplsSteeringAddDel) GetCrcString() string {
	return "1591f94a"
}
func (*SrMplsSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SrMplsSteeringAddDelReply represents VPP binary API message 'sr_mpls_steering_add_del_reply'.
type SrMplsSteeringAddDelReply struct {
	Retval int32
}

func (*SrMplsSteeringAddDelReply) GetMessageName() string {
	return "sr_mpls_steering_add_del_reply"
}
func (*SrMplsSteeringAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SrMplsSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SrMplsPolicyAdd)(nil), "sr_mpls.SrMplsPolicyAdd")
	api.RegisterMessage((*SrMplsPolicyAddReply)(nil), "sr_mpls.SrMplsPolicyAddReply")
	api.RegisterMessage((*SrMplsPolicyAssignEndpointColor)(nil), "sr_mpls.SrMplsPolicyAssignEndpointColor")
	api.RegisterMessage((*SrMplsPolicyAssignEndpointColorReply)(nil), "sr_mpls.SrMplsPolicyAssignEndpointColorReply")
	api.RegisterMessage((*SrMplsPolicyDel)(nil), "sr_mpls.SrMplsPolicyDel")
	api.RegisterMessage((*SrMplsPolicyDelReply)(nil), "sr_mpls.SrMplsPolicyDelReply")
	api.RegisterMessage((*SrMplsPolicyMod)(nil), "sr_mpls.SrMplsPolicyMod")
	api.RegisterMessage((*SrMplsPolicyModReply)(nil), "sr_mpls.SrMplsPolicyModReply")
	api.RegisterMessage((*SrMplsSteeringAddDel)(nil), "sr_mpls.SrMplsSteeringAddDel")
	api.RegisterMessage((*SrMplsSteeringAddDelReply)(nil), "sr_mpls.SrMplsSteeringAddDelReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SrMplsPolicyAdd)(nil),
		(*SrMplsPolicyAddReply)(nil),
		(*SrMplsPolicyAssignEndpointColor)(nil),
		(*SrMplsPolicyAssignEndpointColorReply)(nil),
		(*SrMplsPolicyDel)(nil),
		(*SrMplsPolicyDelReply)(nil),
		(*SrMplsPolicyMod)(nil),
		(*SrMplsPolicyModReply)(nil),
		(*SrMplsSteeringAddDel)(nil),
		(*SrMplsSteeringAddDelReply)(nil),
	}
}

// RPCService represents RPC service API for sr_mpls module.
type RPCService interface {
	SrMplsPolicyAdd(ctx context.Context, in *SrMplsPolicyAdd) (*SrMplsPolicyAddReply, error)
	SrMplsPolicyAssignEndpointColor(ctx context.Context, in *SrMplsPolicyAssignEndpointColor) (*SrMplsPolicyAssignEndpointColorReply, error)
	SrMplsPolicyDel(ctx context.Context, in *SrMplsPolicyDel) (*SrMplsPolicyDelReply, error)
	SrMplsPolicyMod(ctx context.Context, in *SrMplsPolicyMod) (*SrMplsPolicyModReply, error)
	SrMplsSteeringAddDel(ctx context.Context, in *SrMplsSteeringAddDel) (*SrMplsSteeringAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) SrMplsPolicyAdd(ctx context.Context, in *SrMplsPolicyAdd) (*SrMplsPolicyAddReply, error) {
	out := new(SrMplsPolicyAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrMplsPolicyAssignEndpointColor(ctx context.Context, in *SrMplsPolicyAssignEndpointColor) (*SrMplsPolicyAssignEndpointColorReply, error) {
	out := new(SrMplsPolicyAssignEndpointColorReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrMplsPolicyDel(ctx context.Context, in *SrMplsPolicyDel) (*SrMplsPolicyDelReply, error) {
	out := new(SrMplsPolicyDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrMplsPolicyMod(ctx context.Context, in *SrMplsPolicyMod) (*SrMplsPolicyModReply, error) {
	out := new(SrMplsPolicyModReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrMplsSteeringAddDel(ctx context.Context, in *SrMplsSteeringAddDel) (*SrMplsSteeringAddDelReply, error) {
	out := new(SrMplsSteeringAddDelReply)
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
