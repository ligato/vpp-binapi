// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1904/.vppapi/core/virtio.api.json

/*
Package virtio is a generated VPP binary API for 'virtio' module.

It consists of:
	  6 messages
	  3 services
*/
package virtio

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
	ModuleName = "virtio"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x394fa590
)

// SwInterfaceVirtioPciDetails represents VPP binary API message 'sw_interface_virtio_pci_details'.
type SwInterfaceVirtioPciDetails struct {
	SwIfIndex uint32
	PciAddr   uint32
	MacAddr   []byte `struc:"[6]byte"`
	TxRingSz  uint16
	RxRingSz  uint16
	Features  uint64
}

func (m *SwInterfaceVirtioPciDetails) Reset()                        { *m = SwInterfaceVirtioPciDetails{} }
func (*SwInterfaceVirtioPciDetails) GetMessageName() string          { return "sw_interface_virtio_pci_details" }
func (*SwInterfaceVirtioPciDetails) GetCrcString() string            { return "d3a3ade3" }
func (*SwInterfaceVirtioPciDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceVirtioPciDump represents VPP binary API message 'sw_interface_virtio_pci_dump'.
type SwInterfaceVirtioPciDump struct{}

func (m *SwInterfaceVirtioPciDump) Reset()                        { *m = SwInterfaceVirtioPciDump{} }
func (*SwInterfaceVirtioPciDump) GetMessageName() string          { return "sw_interface_virtio_pci_dump" }
func (*SwInterfaceVirtioPciDump) GetCrcString() string            { return "51077d14" }
func (*SwInterfaceVirtioPciDump) GetMessageType() api.MessageType { return api.RequestMessage }

// VirtioPciCreate represents VPP binary API message 'virtio_pci_create'.
type VirtioPciCreate struct {
	PciAddr      uint32
	UseRandomMac uint8
	MacAddress   []byte `struc:"[6]byte"`
	TxRingSz     uint16
	RxRingSz     uint16
	Features     uint64
}

func (m *VirtioPciCreate) Reset()                        { *m = VirtioPciCreate{} }
func (*VirtioPciCreate) GetMessageName() string          { return "virtio_pci_create" }
func (*VirtioPciCreate) GetCrcString() string            { return "78e17a09" }
func (*VirtioPciCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// VirtioPciCreateReply represents VPP binary API message 'virtio_pci_create_reply'.
type VirtioPciCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (m *VirtioPciCreateReply) Reset()                        { *m = VirtioPciCreateReply{} }
func (*VirtioPciCreateReply) GetMessageName() string          { return "virtio_pci_create_reply" }
func (*VirtioPciCreateReply) GetCrcString() string            { return "fda5941f" }
func (*VirtioPciCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// VirtioPciDelete represents VPP binary API message 'virtio_pci_delete'.
type VirtioPciDelete struct {
	SwIfIndex uint32
}

func (m *VirtioPciDelete) Reset()                        { *m = VirtioPciDelete{} }
func (*VirtioPciDelete) GetMessageName() string          { return "virtio_pci_delete" }
func (*VirtioPciDelete) GetCrcString() string            { return "529cb13f" }
func (*VirtioPciDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// VirtioPciDeleteReply represents VPP binary API message 'virtio_pci_delete_reply'.
type VirtioPciDeleteReply struct {
	Retval int32
}

func (m *VirtioPciDeleteReply) Reset()                        { *m = VirtioPciDeleteReply{} }
func (*VirtioPciDeleteReply) GetMessageName() string          { return "virtio_pci_delete_reply" }
func (*VirtioPciDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*VirtioPciDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*SwInterfaceVirtioPciDetails)(nil), "virtio.SwInterfaceVirtioPciDetails")
	api.RegisterMessage((*SwInterfaceVirtioPciDump)(nil), "virtio.SwInterfaceVirtioPciDump")
	api.RegisterMessage((*VirtioPciCreate)(nil), "virtio.VirtioPciCreate")
	api.RegisterMessage((*VirtioPciCreateReply)(nil), "virtio.VirtioPciCreateReply")
	api.RegisterMessage((*VirtioPciDelete)(nil), "virtio.VirtioPciDelete")
	api.RegisterMessage((*VirtioPciDeleteReply)(nil), "virtio.VirtioPciDeleteReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceVirtioPciDetails)(nil),
		(*SwInterfaceVirtioPciDump)(nil),
		(*VirtioPciCreate)(nil),
		(*VirtioPciCreateReply)(nil),
		(*VirtioPciDelete)(nil),
		(*VirtioPciDeleteReply)(nil),
	}
}

// RPCService represents RPC service API for virtio module.
type RPCService interface {
	DumpSwInterfaceVirtioPci(ctx context.Context, in *SwInterfaceVirtioPciDump) (RPCService_DumpSwInterfaceVirtioPciClient, error)
	VirtioPciCreate(ctx context.Context, in *VirtioPciCreate) (*VirtioPciCreateReply, error)
	VirtioPciDelete(ctx context.Context, in *VirtioPciDelete) (*VirtioPciDeleteReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwInterfaceVirtioPci(ctx context.Context, in *SwInterfaceVirtioPciDump) (RPCService_DumpSwInterfaceVirtioPciClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceVirtioPciClient{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceVirtioPciClient interface {
	Recv() (*SwInterfaceVirtioPciDetails, error)
}

type serviceClient_DumpSwInterfaceVirtioPciClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceVirtioPciClient) Recv() (*SwInterfaceVirtioPciDetails, error) {
	m := new(SwInterfaceVirtioPciDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) VirtioPciCreate(ctx context.Context, in *VirtioPciCreate) (*VirtioPciCreateReply, error) {
	out := new(VirtioPciCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VirtioPciDelete(ctx context.Context, in *VirtioPciDelete) (*VirtioPciDeleteReply, error) {
	out := new(VirtioPciDeleteReply)
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
