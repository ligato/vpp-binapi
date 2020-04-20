// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/core/bfd.api.json

/*
Package bfd is a generated VPP binary API for 'bfd' module.

It consists of:
	 28 messages
	 14 services
*/
package bfd

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
	ModuleName = "bfd"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x5c3ba394
)

// BfdAuthDelKey represents VPP binary API message 'bfd_auth_del_key'.
type BfdAuthDelKey struct {
	ConfKeyID uint32
}

func (m *BfdAuthDelKey) Reset()                        { *m = BfdAuthDelKey{} }
func (*BfdAuthDelKey) GetMessageName() string          { return "bfd_auth_del_key" }
func (*BfdAuthDelKey) GetCrcString() string            { return "65310b22" }
func (*BfdAuthDelKey) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdAuthDelKeyReply represents VPP binary API message 'bfd_auth_del_key_reply'.
type BfdAuthDelKeyReply struct {
	Retval int32
}

func (m *BfdAuthDelKeyReply) Reset()                        { *m = BfdAuthDelKeyReply{} }
func (*BfdAuthDelKeyReply) GetMessageName() string          { return "bfd_auth_del_key_reply" }
func (*BfdAuthDelKeyReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdAuthDelKeyReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdAuthKeysDetails represents VPP binary API message 'bfd_auth_keys_details'.
type BfdAuthKeysDetails struct {
	ConfKeyID uint32
	UseCount  uint32
	AuthType  uint8
}

func (m *BfdAuthKeysDetails) Reset()                        { *m = BfdAuthKeysDetails{} }
func (*BfdAuthKeysDetails) GetMessageName() string          { return "bfd_auth_keys_details" }
func (*BfdAuthKeysDetails) GetCrcString() string            { return "84130e9f" }
func (*BfdAuthKeysDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdAuthKeysDump represents VPP binary API message 'bfd_auth_keys_dump'.
type BfdAuthKeysDump struct{}

func (m *BfdAuthKeysDump) Reset()                        { *m = BfdAuthKeysDump{} }
func (*BfdAuthKeysDump) GetMessageName() string          { return "bfd_auth_keys_dump" }
func (*BfdAuthKeysDump) GetCrcString() string            { return "51077d14" }
func (*BfdAuthKeysDump) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdAuthSetKey represents VPP binary API message 'bfd_auth_set_key'.
type BfdAuthSetKey struct {
	ConfKeyID uint32
	KeyLen    uint8
	AuthType  uint8
	Key       []byte `struc:"[20]byte"`
}

func (m *BfdAuthSetKey) Reset()                        { *m = BfdAuthSetKey{} }
func (*BfdAuthSetKey) GetMessageName() string          { return "bfd_auth_set_key" }
func (*BfdAuthSetKey) GetCrcString() string            { return "690b8877" }
func (*BfdAuthSetKey) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdAuthSetKeyReply represents VPP binary API message 'bfd_auth_set_key_reply'.
type BfdAuthSetKeyReply struct {
	Retval int32
}

func (m *BfdAuthSetKeyReply) Reset()                        { *m = BfdAuthSetKeyReply{} }
func (*BfdAuthSetKeyReply) GetMessageName() string          { return "bfd_auth_set_key_reply" }
func (*BfdAuthSetKeyReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdAuthSetKeyReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPAdd represents VPP binary API message 'bfd_udp_add'.
type BfdUDPAdd struct {
	SwIfIndex       uint32
	DesiredMinTx    uint32
	RequiredMinRx   uint32
	LocalAddr       []byte `struc:"[16]byte"`
	PeerAddr        []byte `struc:"[16]byte"`
	IsIPv6          uint8
	DetectMult      uint8
	IsAuthenticated uint8
	BfdKeyID        uint8
	ConfKeyID       uint32
}

func (m *BfdUDPAdd) Reset()                        { *m = BfdUDPAdd{} }
func (*BfdUDPAdd) GetMessageName() string          { return "bfd_udp_add" }
func (*BfdUDPAdd) GetCrcString() string            { return "61cf1850" }
func (*BfdUDPAdd) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPAddReply represents VPP binary API message 'bfd_udp_add_reply'.
type BfdUDPAddReply struct {
	Retval int32
}

func (m *BfdUDPAddReply) Reset()                        { *m = BfdUDPAddReply{} }
func (*BfdUDPAddReply) GetMessageName() string          { return "bfd_udp_add_reply" }
func (*BfdUDPAddReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPAddReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPAuthActivate represents VPP binary API message 'bfd_udp_auth_activate'.
type BfdUDPAuthActivate struct {
	SwIfIndex uint32
	LocalAddr []byte `struc:"[16]byte"`
	PeerAddr  []byte `struc:"[16]byte"`
	IsIPv6    uint8
	IsDelayed uint8
	BfdKeyID  uint8
	ConfKeyID uint32
}

func (m *BfdUDPAuthActivate) Reset()                        { *m = BfdUDPAuthActivate{} }
func (*BfdUDPAuthActivate) GetMessageName() string          { return "bfd_udp_auth_activate" }
func (*BfdUDPAuthActivate) GetCrcString() string            { return "1bae0947" }
func (*BfdUDPAuthActivate) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPAuthActivateReply represents VPP binary API message 'bfd_udp_auth_activate_reply'.
type BfdUDPAuthActivateReply struct {
	Retval int32
}

func (m *BfdUDPAuthActivateReply) Reset()                        { *m = BfdUDPAuthActivateReply{} }
func (*BfdUDPAuthActivateReply) GetMessageName() string          { return "bfd_udp_auth_activate_reply" }
func (*BfdUDPAuthActivateReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPAuthActivateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPAuthDeactivate represents VPP binary API message 'bfd_udp_auth_deactivate'.
type BfdUDPAuthDeactivate struct {
	SwIfIndex uint32
	LocalAddr []byte `struc:"[16]byte"`
	PeerAddr  []byte `struc:"[16]byte"`
	IsIPv6    uint8
	IsDelayed uint8
}

func (m *BfdUDPAuthDeactivate) Reset()                        { *m = BfdUDPAuthDeactivate{} }
func (*BfdUDPAuthDeactivate) GetMessageName() string          { return "bfd_udp_auth_deactivate" }
func (*BfdUDPAuthDeactivate) GetCrcString() string            { return "10661991" }
func (*BfdUDPAuthDeactivate) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPAuthDeactivateReply represents VPP binary API message 'bfd_udp_auth_deactivate_reply'.
type BfdUDPAuthDeactivateReply struct {
	Retval int32
}

func (m *BfdUDPAuthDeactivateReply) Reset()                        { *m = BfdUDPAuthDeactivateReply{} }
func (*BfdUDPAuthDeactivateReply) GetMessageName() string          { return "bfd_udp_auth_deactivate_reply" }
func (*BfdUDPAuthDeactivateReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPAuthDeactivateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPDel represents VPP binary API message 'bfd_udp_del'.
type BfdUDPDel struct {
	SwIfIndex uint32
	LocalAddr []byte `struc:"[16]byte"`
	PeerAddr  []byte `struc:"[16]byte"`
	IsIPv6    uint8
}

func (m *BfdUDPDel) Reset()                        { *m = BfdUDPDel{} }
func (*BfdUDPDel) GetMessageName() string          { return "bfd_udp_del" }
func (*BfdUDPDel) GetCrcString() string            { return "c9e9cc5a" }
func (*BfdUDPDel) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPDelEchoSource represents VPP binary API message 'bfd_udp_del_echo_source'.
type BfdUDPDelEchoSource struct{}

func (m *BfdUDPDelEchoSource) Reset()                        { *m = BfdUDPDelEchoSource{} }
func (*BfdUDPDelEchoSource) GetMessageName() string          { return "bfd_udp_del_echo_source" }
func (*BfdUDPDelEchoSource) GetCrcString() string            { return "51077d14" }
func (*BfdUDPDelEchoSource) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPDelEchoSourceReply represents VPP binary API message 'bfd_udp_del_echo_source_reply'.
type BfdUDPDelEchoSourceReply struct {
	Retval int32
}

func (m *BfdUDPDelEchoSourceReply) Reset()                        { *m = BfdUDPDelEchoSourceReply{} }
func (*BfdUDPDelEchoSourceReply) GetMessageName() string          { return "bfd_udp_del_echo_source_reply" }
func (*BfdUDPDelEchoSourceReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPDelEchoSourceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPDelReply represents VPP binary API message 'bfd_udp_del_reply'.
type BfdUDPDelReply struct {
	Retval int32
}

func (m *BfdUDPDelReply) Reset()                        { *m = BfdUDPDelReply{} }
func (*BfdUDPDelReply) GetMessageName() string          { return "bfd_udp_del_reply" }
func (*BfdUDPDelReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPGetEchoSource represents VPP binary API message 'bfd_udp_get_echo_source'.
type BfdUDPGetEchoSource struct{}

func (m *BfdUDPGetEchoSource) Reset()                        { *m = BfdUDPGetEchoSource{} }
func (*BfdUDPGetEchoSource) GetMessageName() string          { return "bfd_udp_get_echo_source" }
func (*BfdUDPGetEchoSource) GetCrcString() string            { return "51077d14" }
func (*BfdUDPGetEchoSource) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPGetEchoSourceReply represents VPP binary API message 'bfd_udp_get_echo_source_reply'.
type BfdUDPGetEchoSourceReply struct {
	Retval        int32
	SwIfIndex     uint32
	IsSet         uint8
	HaveUsableIP4 uint8
	IP4Addr       []byte `struc:"[4]byte"`
	HaveUsableIP6 uint8
	IP6Addr       []byte `struc:"[16]byte"`
}

func (m *BfdUDPGetEchoSourceReply) Reset()                        { *m = BfdUDPGetEchoSourceReply{} }
func (*BfdUDPGetEchoSourceReply) GetMessageName() string          { return "bfd_udp_get_echo_source_reply" }
func (*BfdUDPGetEchoSourceReply) GetCrcString() string            { return "6924ca6b" }
func (*BfdUDPGetEchoSourceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPMod represents VPP binary API message 'bfd_udp_mod'.
type BfdUDPMod struct {
	SwIfIndex     uint32
	DesiredMinTx  uint32
	RequiredMinRx uint32
	LocalAddr     []byte `struc:"[16]byte"`
	PeerAddr      []byte `struc:"[16]byte"`
	IsIPv6        uint8
	DetectMult    uint8
}

func (m *BfdUDPMod) Reset()                        { *m = BfdUDPMod{} }
func (*BfdUDPMod) GetMessageName() string          { return "bfd_udp_mod" }
func (*BfdUDPMod) GetCrcString() string            { return "6049bf47" }
func (*BfdUDPMod) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPModReply represents VPP binary API message 'bfd_udp_mod_reply'.
type BfdUDPModReply struct {
	Retval int32
}

func (m *BfdUDPModReply) Reset()                        { *m = BfdUDPModReply{} }
func (*BfdUDPModReply) GetMessageName() string          { return "bfd_udp_mod_reply" }
func (*BfdUDPModReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPModReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPSessionDetails represents VPP binary API message 'bfd_udp_session_details'.
type BfdUDPSessionDetails struct {
	SwIfIndex       uint32
	LocalAddr       []byte `struc:"[16]byte"`
	PeerAddr        []byte `struc:"[16]byte"`
	IsIPv6          uint8
	State           uint8
	IsAuthenticated uint8
	BfdKeyID        uint8
	ConfKeyID       uint32
	RequiredMinRx   uint32
	DesiredMinTx    uint32
	DetectMult      uint8
}

func (m *BfdUDPSessionDetails) Reset()                        { *m = BfdUDPSessionDetails{} }
func (*BfdUDPSessionDetails) GetMessageName() string          { return "bfd_udp_session_details" }
func (*BfdUDPSessionDetails) GetCrcString() string            { return "837bb0ed" }
func (*BfdUDPSessionDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPSessionDump represents VPP binary API message 'bfd_udp_session_dump'.
type BfdUDPSessionDump struct{}

func (m *BfdUDPSessionDump) Reset()                        { *m = BfdUDPSessionDump{} }
func (*BfdUDPSessionDump) GetMessageName() string          { return "bfd_udp_session_dump" }
func (*BfdUDPSessionDump) GetCrcString() string            { return "51077d14" }
func (*BfdUDPSessionDump) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPSessionSetFlags represents VPP binary API message 'bfd_udp_session_set_flags'.
type BfdUDPSessionSetFlags struct {
	SwIfIndex   uint32
	LocalAddr   []byte `struc:"[16]byte"`
	PeerAddr    []byte `struc:"[16]byte"`
	IsIPv6      uint8
	AdminUpDown uint8
}

func (m *BfdUDPSessionSetFlags) Reset()                        { *m = BfdUDPSessionSetFlags{} }
func (*BfdUDPSessionSetFlags) GetMessageName() string          { return "bfd_udp_session_set_flags" }
func (*BfdUDPSessionSetFlags) GetCrcString() string            { return "667d6e7c" }
func (*BfdUDPSessionSetFlags) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPSessionSetFlagsReply represents VPP binary API message 'bfd_udp_session_set_flags_reply'.
type BfdUDPSessionSetFlagsReply struct {
	Retval int32
}

func (m *BfdUDPSessionSetFlagsReply) Reset()                        { *m = BfdUDPSessionSetFlagsReply{} }
func (*BfdUDPSessionSetFlagsReply) GetMessageName() string          { return "bfd_udp_session_set_flags_reply" }
func (*BfdUDPSessionSetFlagsReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPSessionSetFlagsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BfdUDPSetEchoSource represents VPP binary API message 'bfd_udp_set_echo_source'.
type BfdUDPSetEchoSource struct {
	SwIfIndex uint32
}

func (m *BfdUDPSetEchoSource) Reset()                        { *m = BfdUDPSetEchoSource{} }
func (*BfdUDPSetEchoSource) GetMessageName() string          { return "bfd_udp_set_echo_source" }
func (*BfdUDPSetEchoSource) GetCrcString() string            { return "529cb13f" }
func (*BfdUDPSetEchoSource) GetMessageType() api.MessageType { return api.RequestMessage }

// BfdUDPSetEchoSourceReply represents VPP binary API message 'bfd_udp_set_echo_source_reply'.
type BfdUDPSetEchoSourceReply struct {
	Retval int32
}

func (m *BfdUDPSetEchoSourceReply) Reset()                        { *m = BfdUDPSetEchoSourceReply{} }
func (*BfdUDPSetEchoSourceReply) GetMessageName() string          { return "bfd_udp_set_echo_source_reply" }
func (*BfdUDPSetEchoSourceReply) GetCrcString() string            { return "e8d4e804" }
func (*BfdUDPSetEchoSourceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// WantBfdEvents represents VPP binary API message 'want_bfd_events'.
type WantBfdEvents struct {
	EnableDisable uint32
	PID           uint32
}

func (m *WantBfdEvents) Reset()                        { *m = WantBfdEvents{} }
func (*WantBfdEvents) GetMessageName() string          { return "want_bfd_events" }
func (*WantBfdEvents) GetCrcString() string            { return "476f5a08" }
func (*WantBfdEvents) GetMessageType() api.MessageType { return api.RequestMessage }

// WantBfdEventsReply represents VPP binary API message 'want_bfd_events_reply'.
type WantBfdEventsReply struct {
	Retval int32
}

func (m *WantBfdEventsReply) Reset()                        { *m = WantBfdEventsReply{} }
func (*WantBfdEventsReply) GetMessageName() string          { return "want_bfd_events_reply" }
func (*WantBfdEventsReply) GetCrcString() string            { return "e8d4e804" }
func (*WantBfdEventsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*BfdAuthDelKey)(nil), "bfd.BfdAuthDelKey")
	api.RegisterMessage((*BfdAuthDelKeyReply)(nil), "bfd.BfdAuthDelKeyReply")
	api.RegisterMessage((*BfdAuthKeysDetails)(nil), "bfd.BfdAuthKeysDetails")
	api.RegisterMessage((*BfdAuthKeysDump)(nil), "bfd.BfdAuthKeysDump")
	api.RegisterMessage((*BfdAuthSetKey)(nil), "bfd.BfdAuthSetKey")
	api.RegisterMessage((*BfdAuthSetKeyReply)(nil), "bfd.BfdAuthSetKeyReply")
	api.RegisterMessage((*BfdUDPAdd)(nil), "bfd.BfdUDPAdd")
	api.RegisterMessage((*BfdUDPAddReply)(nil), "bfd.BfdUDPAddReply")
	api.RegisterMessage((*BfdUDPAuthActivate)(nil), "bfd.BfdUDPAuthActivate")
	api.RegisterMessage((*BfdUDPAuthActivateReply)(nil), "bfd.BfdUDPAuthActivateReply")
	api.RegisterMessage((*BfdUDPAuthDeactivate)(nil), "bfd.BfdUDPAuthDeactivate")
	api.RegisterMessage((*BfdUDPAuthDeactivateReply)(nil), "bfd.BfdUDPAuthDeactivateReply")
	api.RegisterMessage((*BfdUDPDel)(nil), "bfd.BfdUDPDel")
	api.RegisterMessage((*BfdUDPDelEchoSource)(nil), "bfd.BfdUDPDelEchoSource")
	api.RegisterMessage((*BfdUDPDelEchoSourceReply)(nil), "bfd.BfdUDPDelEchoSourceReply")
	api.RegisterMessage((*BfdUDPDelReply)(nil), "bfd.BfdUDPDelReply")
	api.RegisterMessage((*BfdUDPGetEchoSource)(nil), "bfd.BfdUDPGetEchoSource")
	api.RegisterMessage((*BfdUDPGetEchoSourceReply)(nil), "bfd.BfdUDPGetEchoSourceReply")
	api.RegisterMessage((*BfdUDPMod)(nil), "bfd.BfdUDPMod")
	api.RegisterMessage((*BfdUDPModReply)(nil), "bfd.BfdUDPModReply")
	api.RegisterMessage((*BfdUDPSessionDetails)(nil), "bfd.BfdUDPSessionDetails")
	api.RegisterMessage((*BfdUDPSessionDump)(nil), "bfd.BfdUDPSessionDump")
	api.RegisterMessage((*BfdUDPSessionSetFlags)(nil), "bfd.BfdUDPSessionSetFlags")
	api.RegisterMessage((*BfdUDPSessionSetFlagsReply)(nil), "bfd.BfdUDPSessionSetFlagsReply")
	api.RegisterMessage((*BfdUDPSetEchoSource)(nil), "bfd.BfdUDPSetEchoSource")
	api.RegisterMessage((*BfdUDPSetEchoSourceReply)(nil), "bfd.BfdUDPSetEchoSourceReply")
	api.RegisterMessage((*WantBfdEvents)(nil), "bfd.WantBfdEvents")
	api.RegisterMessage((*WantBfdEventsReply)(nil), "bfd.WantBfdEventsReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BfdAuthDelKey)(nil),
		(*BfdAuthDelKeyReply)(nil),
		(*BfdAuthKeysDetails)(nil),
		(*BfdAuthKeysDump)(nil),
		(*BfdAuthSetKey)(nil),
		(*BfdAuthSetKeyReply)(nil),
		(*BfdUDPAdd)(nil),
		(*BfdUDPAddReply)(nil),
		(*BfdUDPAuthActivate)(nil),
		(*BfdUDPAuthActivateReply)(nil),
		(*BfdUDPAuthDeactivate)(nil),
		(*BfdUDPAuthDeactivateReply)(nil),
		(*BfdUDPDel)(nil),
		(*BfdUDPDelEchoSource)(nil),
		(*BfdUDPDelEchoSourceReply)(nil),
		(*BfdUDPDelReply)(nil),
		(*BfdUDPGetEchoSource)(nil),
		(*BfdUDPGetEchoSourceReply)(nil),
		(*BfdUDPMod)(nil),
		(*BfdUDPModReply)(nil),
		(*BfdUDPSessionDetails)(nil),
		(*BfdUDPSessionDump)(nil),
		(*BfdUDPSessionSetFlags)(nil),
		(*BfdUDPSessionSetFlagsReply)(nil),
		(*BfdUDPSetEchoSource)(nil),
		(*BfdUDPSetEchoSourceReply)(nil),
		(*WantBfdEvents)(nil),
		(*WantBfdEventsReply)(nil),
	}
}

// RPCService represents RPC service API for bfd module.
type RPCService interface {
	DumpBfdAuthKeys(ctx context.Context, in *BfdAuthKeysDump) (RPCService_DumpBfdAuthKeysClient, error)
	DumpBfdUDPSession(ctx context.Context, in *BfdUDPSessionDump) (RPCService_DumpBfdUDPSessionClient, error)
	BfdAuthDelKey(ctx context.Context, in *BfdAuthDelKey) (*BfdAuthDelKeyReply, error)
	BfdAuthSetKey(ctx context.Context, in *BfdAuthSetKey) (*BfdAuthSetKeyReply, error)
	BfdUDPAdd(ctx context.Context, in *BfdUDPAdd) (*BfdUDPAddReply, error)
	BfdUDPAuthActivate(ctx context.Context, in *BfdUDPAuthActivate) (*BfdUDPAuthActivateReply, error)
	BfdUDPAuthDeactivate(ctx context.Context, in *BfdUDPAuthDeactivate) (*BfdUDPAuthDeactivateReply, error)
	BfdUDPDel(ctx context.Context, in *BfdUDPDel) (*BfdUDPDelReply, error)
	BfdUDPDelEchoSource(ctx context.Context, in *BfdUDPDelEchoSource) (*BfdUDPDelEchoSourceReply, error)
	BfdUDPGetEchoSource(ctx context.Context, in *BfdUDPGetEchoSource) (*BfdUDPGetEchoSourceReply, error)
	BfdUDPMod(ctx context.Context, in *BfdUDPMod) (*BfdUDPModReply, error)
	BfdUDPSessionSetFlags(ctx context.Context, in *BfdUDPSessionSetFlags) (*BfdUDPSessionSetFlagsReply, error)
	BfdUDPSetEchoSource(ctx context.Context, in *BfdUDPSetEchoSource) (*BfdUDPSetEchoSourceReply, error)
	WantBfdEvents(ctx context.Context, in *WantBfdEvents) (*WantBfdEventsReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpBfdAuthKeys(ctx context.Context, in *BfdAuthKeysDump) (RPCService_DumpBfdAuthKeysClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBfdAuthKeysClient{stream}
	return x, nil
}

type RPCService_DumpBfdAuthKeysClient interface {
	Recv() (*BfdAuthKeysDetails, error)
}

type serviceClient_DumpBfdAuthKeysClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBfdAuthKeysClient) Recv() (*BfdAuthKeysDetails, error) {
	m := new(BfdAuthKeysDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpBfdUDPSession(ctx context.Context, in *BfdUDPSessionDump) (RPCService_DumpBfdUDPSessionClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpBfdUDPSessionClient{stream}
	return x, nil
}

type RPCService_DumpBfdUDPSessionClient interface {
	Recv() (*BfdUDPSessionDetails, error)
}

type serviceClient_DumpBfdUDPSessionClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpBfdUDPSessionClient) Recv() (*BfdUDPSessionDetails, error) {
	m := new(BfdUDPSessionDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) BfdAuthDelKey(ctx context.Context, in *BfdAuthDelKey) (*BfdAuthDelKeyReply, error) {
	out := new(BfdAuthDelKeyReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdAuthSetKey(ctx context.Context, in *BfdAuthSetKey) (*BfdAuthSetKeyReply, error) {
	out := new(BfdAuthSetKeyReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPAdd(ctx context.Context, in *BfdUDPAdd) (*BfdUDPAddReply, error) {
	out := new(BfdUDPAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPAuthActivate(ctx context.Context, in *BfdUDPAuthActivate) (*BfdUDPAuthActivateReply, error) {
	out := new(BfdUDPAuthActivateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPAuthDeactivate(ctx context.Context, in *BfdUDPAuthDeactivate) (*BfdUDPAuthDeactivateReply, error) {
	out := new(BfdUDPAuthDeactivateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPDel(ctx context.Context, in *BfdUDPDel) (*BfdUDPDelReply, error) {
	out := new(BfdUDPDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPDelEchoSource(ctx context.Context, in *BfdUDPDelEchoSource) (*BfdUDPDelEchoSourceReply, error) {
	out := new(BfdUDPDelEchoSourceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPGetEchoSource(ctx context.Context, in *BfdUDPGetEchoSource) (*BfdUDPGetEchoSourceReply, error) {
	out := new(BfdUDPGetEchoSourceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPMod(ctx context.Context, in *BfdUDPMod) (*BfdUDPModReply, error) {
	out := new(BfdUDPModReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPSessionSetFlags(ctx context.Context, in *BfdUDPSessionSetFlags) (*BfdUDPSessionSetFlagsReply, error) {
	out := new(BfdUDPSessionSetFlagsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BfdUDPSetEchoSource(ctx context.Context, in *BfdUDPSetEchoSource) (*BfdUDPSetEchoSourceReply, error) {
	out := new(BfdUDPSetEchoSourceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) WantBfdEvents(ctx context.Context, in *WantBfdEvents) (*WantBfdEventsReply, error) {
	out := new(WantBfdEventsReply)
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
