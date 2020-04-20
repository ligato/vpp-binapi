// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/plugins/ioam_cache.api.json

/*
Package ioam_cache is a generated VPP binary API for 'ioam_cache' module.

It consists of:
	  2 messages
	  1 service
*/
package ioam_cache

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
	ModuleName = "ioam_cache"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xb7452f41
)

// IoamCacheIP6EnableDisable represents VPP binary API message 'ioam_cache_ip6_enable_disable'.
type IoamCacheIP6EnableDisable struct {
	IsDisable uint8
}

func (m *IoamCacheIP6EnableDisable) Reset()                        { *m = IoamCacheIP6EnableDisable{} }
func (*IoamCacheIP6EnableDisable) GetMessageName() string          { return "ioam_cache_ip6_enable_disable" }
func (*IoamCacheIP6EnableDisable) GetCrcString() string            { return "22324d89" }
func (*IoamCacheIP6EnableDisable) GetMessageType() api.MessageType { return api.RequestMessage }

// IoamCacheIP6EnableDisableReply represents VPP binary API message 'ioam_cache_ip6_enable_disable_reply'.
type IoamCacheIP6EnableDisableReply struct {
	Retval int32
}

func (m *IoamCacheIP6EnableDisableReply) Reset() { *m = IoamCacheIP6EnableDisableReply{} }
func (*IoamCacheIP6EnableDisableReply) GetMessageName() string {
	return "ioam_cache_ip6_enable_disable_reply"
}
func (*IoamCacheIP6EnableDisableReply) GetCrcString() string            { return "e8d4e804" }
func (*IoamCacheIP6EnableDisableReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*IoamCacheIP6EnableDisable)(nil), "ioam_cache.IoamCacheIP6EnableDisable")
	api.RegisterMessage((*IoamCacheIP6EnableDisableReply)(nil), "ioam_cache.IoamCacheIP6EnableDisableReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IoamCacheIP6EnableDisable)(nil),
		(*IoamCacheIP6EnableDisableReply)(nil),
	}
}

// RPCService represents RPC service API for ioam_cache module.
type RPCService interface {
	IoamCacheIP6EnableDisable(ctx context.Context, in *IoamCacheIP6EnableDisable) (*IoamCacheIP6EnableDisableReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) IoamCacheIP6EnableDisable(ctx context.Context, in *IoamCacheIP6EnableDisable) (*IoamCacheIP6EnableDisableReply, error) {
	out := new(IoamCacheIP6EnableDisableReply)
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
