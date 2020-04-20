// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/core/feature.api.json

/*
Package feature is a generated VPP binary API for 'feature' module.

It consists of:
	  2 messages
	  1 service
*/
package feature

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
	ModuleName = "feature"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x5aec37eb
)

// FeatureEnableDisable represents VPP binary API message 'feature_enable_disable'.
type FeatureEnableDisable struct {
	SwIfIndex   uint32
	Enable      uint8
	ArcName     []byte `struc:"[64]byte"`
	FeatureName []byte `struc:"[64]byte"`
}

func (m *FeatureEnableDisable) Reset()                        { *m = FeatureEnableDisable{} }
func (*FeatureEnableDisable) GetMessageName() string          { return "feature_enable_disable" }
func (*FeatureEnableDisable) GetCrcString() string            { return "f525e210" }
func (*FeatureEnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// FeatureEnableDisableReply represents VPP binary API message 'feature_enable_disable_reply'.
type FeatureEnableDisableReply struct {
	Retval int32
}

func (m *FeatureEnableDisableReply) Reset()                        { *m = FeatureEnableDisableReply{} }
func (*FeatureEnableDisableReply) GetMessageName() string          { return "feature_enable_disable_reply" }
func (*FeatureEnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*FeatureEnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*FeatureEnableDisable)(nil), "feature.FeatureEnableDisable")
	api.RegisterMessage((*FeatureEnableDisableReply)(nil), "feature.FeatureEnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FeatureEnableDisable)(nil),
		(*FeatureEnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for feature module.
type RPCService interface {
	FeatureEnableDisable(ctx context.Context, in *FeatureEnableDisable) (*FeatureEnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) FeatureEnableDisable(ctx context.Context, in *FeatureEnableDisable) (*FeatureEnableDisableReply, error) {
	out := new(FeatureEnableDisableReply)
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
