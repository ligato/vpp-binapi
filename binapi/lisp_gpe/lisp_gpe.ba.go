// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/lisp_gpe.api.json

/*
Package lisp_gpe is a generated VPP binary API for 'lisp_gpe' module.

It consists of:
	  3 types
	 20 messages
	 10 services
*/
package lisp_gpe

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
	ModuleName = "lisp_gpe"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xc6e77c9
)

// GpeFwdEntry represents VPP binary API type 'gpe_fwd_entry'.
type GpeFwdEntry struct {
	FwdEntryIndex uint32
	DpTable       uint32
	EidType       uint8
	LeidPrefixLen uint8
	ReidPrefixLen uint8
	Leid          []byte `struc:"[16]byte"`
	Reid          []byte `struc:"[16]byte"`
	Vni           uint32
	Action        uint8
}

func (*GpeFwdEntry) GetTypeName() string {
	return "gpe_fwd_entry"
}

// GpeLocator represents VPP binary API type 'gpe_locator'.
type GpeLocator struct {
	IsIP4  uint8
	Weight uint8
	Addr   []byte `struc:"[16]byte"`
}

func (*GpeLocator) GetTypeName() string {
	return "gpe_locator"
}

// GpeNativeFwdRpath represents VPP binary API type 'gpe_native_fwd_rpath'.
type GpeNativeFwdRpath struct {
	FibIndex    uint32
	NhSwIfIndex uint32
	IsIP4       uint8
	NhAddr      []byte `struc:"[16]byte"`
}

func (*GpeNativeFwdRpath) GetTypeName() string {
	return "gpe_native_fwd_rpath"
}

// GpeAddDelFwdEntry represents VPP binary API message 'gpe_add_del_fwd_entry'.
type GpeAddDelFwdEntry struct {
	IsAdd   uint8
	EidType uint8
	RmtEid  []byte `struc:"[16]byte"`
	LclEid  []byte `struc:"[16]byte"`
	RmtLen  uint8
	LclLen  uint8
	Vni     uint32
	DpTable uint32
	Action  uint8
	LocNum  uint32 `struc:"sizeof=Locs"`
	Locs    []GpeLocator
}

func (*GpeAddDelFwdEntry) GetMessageName() string {
	return "gpe_add_del_fwd_entry"
}
func (*GpeAddDelFwdEntry) GetCrcString() string {
	return "afbf857a"
}
func (*GpeAddDelFwdEntry) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeAddDelFwdEntryReply represents VPP binary API message 'gpe_add_del_fwd_entry_reply'.
type GpeAddDelFwdEntryReply struct {
	Retval        int32
	FwdEntryIndex uint32
}

func (*GpeAddDelFwdEntryReply) GetMessageName() string {
	return "gpe_add_del_fwd_entry_reply"
}
func (*GpeAddDelFwdEntryReply) GetCrcString() string {
	return "efe5f176"
}
func (*GpeAddDelFwdEntryReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeAddDelIface represents VPP binary API message 'gpe_add_del_iface'.
type GpeAddDelIface struct {
	IsAdd   uint8
	IsL2    uint8
	DpTable uint32
	Vni     uint32
}

func (*GpeAddDelIface) GetMessageName() string {
	return "gpe_add_del_iface"
}
func (*GpeAddDelIface) GetCrcString() string {
	return "42d6b533"
}
func (*GpeAddDelIface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeAddDelIfaceReply represents VPP binary API message 'gpe_add_del_iface_reply'.
type GpeAddDelIfaceReply struct {
	Retval int32
}

func (*GpeAddDelIfaceReply) GetMessageName() string {
	return "gpe_add_del_iface_reply"
}
func (*GpeAddDelIfaceReply) GetCrcString() string {
	return "e8d4e804"
}
func (*GpeAddDelIfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeAddDelNativeFwdRpath represents VPP binary API message 'gpe_add_del_native_fwd_rpath'.
type GpeAddDelNativeFwdRpath struct {
	IsAdd       uint8
	TableID     uint32
	NhSwIfIndex uint32
	IsIP4       uint8
	NhAddr      []byte `struc:"[16]byte"`
}

func (*GpeAddDelNativeFwdRpath) GetMessageName() string {
	return "gpe_add_del_native_fwd_rpath"
}
func (*GpeAddDelNativeFwdRpath) GetCrcString() string {
	return "bfc42b8f"
}
func (*GpeAddDelNativeFwdRpath) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeAddDelNativeFwdRpathReply represents VPP binary API message 'gpe_add_del_native_fwd_rpath_reply'.
type GpeAddDelNativeFwdRpathReply struct {
	Retval int32
}

func (*GpeAddDelNativeFwdRpathReply) GetMessageName() string {
	return "gpe_add_del_native_fwd_rpath_reply"
}
func (*GpeAddDelNativeFwdRpathReply) GetCrcString() string {
	return "e8d4e804"
}
func (*GpeAddDelNativeFwdRpathReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeEnableDisable represents VPP binary API message 'gpe_enable_disable'.
type GpeEnableDisable struct {
	IsEn uint8
}

func (*GpeEnableDisable) GetMessageName() string {
	return "gpe_enable_disable"
}
func (*GpeEnableDisable) GetCrcString() string {
	return "eb0e943b"
}
func (*GpeEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeEnableDisableReply represents VPP binary API message 'gpe_enable_disable_reply'.
type GpeEnableDisableReply struct {
	Retval int32
}

func (*GpeEnableDisableReply) GetMessageName() string {
	return "gpe_enable_disable_reply"
}
func (*GpeEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*GpeEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeFwdEntriesGet represents VPP binary API message 'gpe_fwd_entries_get'.
type GpeFwdEntriesGet struct {
	Vni uint32
}

func (*GpeFwdEntriesGet) GetMessageName() string {
	return "gpe_fwd_entries_get"
}
func (*GpeFwdEntriesGet) GetCrcString() string {
	return "8d1f2fe9"
}
func (*GpeFwdEntriesGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeFwdEntriesGetReply represents VPP binary API message 'gpe_fwd_entries_get_reply'.
type GpeFwdEntriesGetReply struct {
	Retval  int32
	Count   uint32 `struc:"sizeof=Entries"`
	Entries []GpeFwdEntry
}

func (*GpeFwdEntriesGetReply) GetMessageName() string {
	return "gpe_fwd_entries_get_reply"
}
func (*GpeFwdEntriesGetReply) GetCrcString() string {
	return "07b02c34"
}
func (*GpeFwdEntriesGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeFwdEntryPathDetails represents VPP binary API message 'gpe_fwd_entry_path_details'.
type GpeFwdEntryPathDetails struct {
	LclLoc GpeLocator
	RmtLoc GpeLocator
}

func (*GpeFwdEntryPathDetails) GetMessageName() string {
	return "gpe_fwd_entry_path_details"
}
func (*GpeFwdEntryPathDetails) GetCrcString() string {
	return "a9bdc1f1"
}
func (*GpeFwdEntryPathDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeFwdEntryPathDump represents VPP binary API message 'gpe_fwd_entry_path_dump'.
type GpeFwdEntryPathDump struct {
	FwdEntryIndex uint32
}

func (*GpeFwdEntryPathDump) GetMessageName() string {
	return "gpe_fwd_entry_path_dump"
}
func (*GpeFwdEntryPathDump) GetCrcString() string {
	return "39bce980"
}
func (*GpeFwdEntryPathDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeFwdEntryVnisGet represents VPP binary API message 'gpe_fwd_entry_vnis_get'.
type GpeFwdEntryVnisGet struct{}

func (*GpeFwdEntryVnisGet) GetMessageName() string {
	return "gpe_fwd_entry_vnis_get"
}
func (*GpeFwdEntryVnisGet) GetCrcString() string {
	return "51077d14"
}
func (*GpeFwdEntryVnisGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeFwdEntryVnisGetReply represents VPP binary API message 'gpe_fwd_entry_vnis_get_reply'.
type GpeFwdEntryVnisGetReply struct {
	Retval int32
	Count  uint32 `struc:"sizeof=Vnis"`
	Vnis   []uint32
}

func (*GpeFwdEntryVnisGetReply) GetMessageName() string {
	return "gpe_fwd_entry_vnis_get_reply"
}
func (*GpeFwdEntryVnisGetReply) GetCrcString() string {
	return "aa70da20"
}
func (*GpeFwdEntryVnisGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeGetEncapMode represents VPP binary API message 'gpe_get_encap_mode'.
type GpeGetEncapMode struct{}

func (*GpeGetEncapMode) GetMessageName() string {
	return "gpe_get_encap_mode"
}
func (*GpeGetEncapMode) GetCrcString() string {
	return "51077d14"
}
func (*GpeGetEncapMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeGetEncapModeReply represents VPP binary API message 'gpe_get_encap_mode_reply'.
type GpeGetEncapModeReply struct {
	Retval    int32
	EncapMode uint8
}

func (*GpeGetEncapModeReply) GetMessageName() string {
	return "gpe_get_encap_mode_reply"
}
func (*GpeGetEncapModeReply) GetCrcString() string {
	return "36e3f7ca"
}
func (*GpeGetEncapModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeNativeFwdRpathsGet represents VPP binary API message 'gpe_native_fwd_rpaths_get'.
type GpeNativeFwdRpathsGet struct {
	IsIP4 uint8
}

func (*GpeNativeFwdRpathsGet) GetMessageName() string {
	return "gpe_native_fwd_rpaths_get"
}
func (*GpeNativeFwdRpathsGet) GetCrcString() string {
	return "c5e0d91b"
}
func (*GpeNativeFwdRpathsGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeNativeFwdRpathsGetReply represents VPP binary API message 'gpe_native_fwd_rpaths_get_reply'.
type GpeNativeFwdRpathsGetReply struct {
	Retval  int32
	Count   uint32 `struc:"sizeof=Entries"`
	Entries []GpeNativeFwdRpath
}

func (*GpeNativeFwdRpathsGetReply) GetMessageName() string {
	return "gpe_native_fwd_rpaths_get_reply"
}
func (*GpeNativeFwdRpathsGetReply) GetCrcString() string {
	return "1e4536e3"
}
func (*GpeNativeFwdRpathsGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GpeSetEncapMode represents VPP binary API message 'gpe_set_encap_mode'.
type GpeSetEncapMode struct {
	Mode uint8
}

func (*GpeSetEncapMode) GetMessageName() string {
	return "gpe_set_encap_mode"
}
func (*GpeSetEncapMode) GetCrcString() string {
	return "f3f93ce9"
}
func (*GpeSetEncapMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GpeSetEncapModeReply represents VPP binary API message 'gpe_set_encap_mode_reply'.
type GpeSetEncapModeReply struct {
	Retval int32
}

func (*GpeSetEncapModeReply) GetMessageName() string {
	return "gpe_set_encap_mode_reply"
}
func (*GpeSetEncapModeReply) GetCrcString() string {
	return "e8d4e804"
}
func (*GpeSetEncapModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*GpeAddDelFwdEntry)(nil), "lisp_gpe.GpeAddDelFwdEntry")
	api.RegisterMessage((*GpeAddDelFwdEntryReply)(nil), "lisp_gpe.GpeAddDelFwdEntryReply")
	api.RegisterMessage((*GpeAddDelIface)(nil), "lisp_gpe.GpeAddDelIface")
	api.RegisterMessage((*GpeAddDelIfaceReply)(nil), "lisp_gpe.GpeAddDelIfaceReply")
	api.RegisterMessage((*GpeAddDelNativeFwdRpath)(nil), "lisp_gpe.GpeAddDelNativeFwdRpath")
	api.RegisterMessage((*GpeAddDelNativeFwdRpathReply)(nil), "lisp_gpe.GpeAddDelNativeFwdRpathReply")
	api.RegisterMessage((*GpeEnableDisable)(nil), "lisp_gpe.GpeEnableDisable")
	api.RegisterMessage((*GpeEnableDisableReply)(nil), "lisp_gpe.GpeEnableDisableReply")
	api.RegisterMessage((*GpeFwdEntriesGet)(nil), "lisp_gpe.GpeFwdEntriesGet")
	api.RegisterMessage((*GpeFwdEntriesGetReply)(nil), "lisp_gpe.GpeFwdEntriesGetReply")
	api.RegisterMessage((*GpeFwdEntryPathDetails)(nil), "lisp_gpe.GpeFwdEntryPathDetails")
	api.RegisterMessage((*GpeFwdEntryPathDump)(nil), "lisp_gpe.GpeFwdEntryPathDump")
	api.RegisterMessage((*GpeFwdEntryVnisGet)(nil), "lisp_gpe.GpeFwdEntryVnisGet")
	api.RegisterMessage((*GpeFwdEntryVnisGetReply)(nil), "lisp_gpe.GpeFwdEntryVnisGetReply")
	api.RegisterMessage((*GpeGetEncapMode)(nil), "lisp_gpe.GpeGetEncapMode")
	api.RegisterMessage((*GpeGetEncapModeReply)(nil), "lisp_gpe.GpeGetEncapModeReply")
	api.RegisterMessage((*GpeNativeFwdRpathsGet)(nil), "lisp_gpe.GpeNativeFwdRpathsGet")
	api.RegisterMessage((*GpeNativeFwdRpathsGetReply)(nil), "lisp_gpe.GpeNativeFwdRpathsGetReply")
	api.RegisterMessage((*GpeSetEncapMode)(nil), "lisp_gpe.GpeSetEncapMode")
	api.RegisterMessage((*GpeSetEncapModeReply)(nil), "lisp_gpe.GpeSetEncapModeReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GpeAddDelFwdEntry)(nil),
		(*GpeAddDelFwdEntryReply)(nil),
		(*GpeAddDelIface)(nil),
		(*GpeAddDelIfaceReply)(nil),
		(*GpeAddDelNativeFwdRpath)(nil),
		(*GpeAddDelNativeFwdRpathReply)(nil),
		(*GpeEnableDisable)(nil),
		(*GpeEnableDisableReply)(nil),
		(*GpeFwdEntriesGet)(nil),
		(*GpeFwdEntriesGetReply)(nil),
		(*GpeFwdEntryPathDetails)(nil),
		(*GpeFwdEntryPathDump)(nil),
		(*GpeFwdEntryVnisGet)(nil),
		(*GpeFwdEntryVnisGetReply)(nil),
		(*GpeGetEncapMode)(nil),
		(*GpeGetEncapModeReply)(nil),
		(*GpeNativeFwdRpathsGet)(nil),
		(*GpeNativeFwdRpathsGetReply)(nil),
		(*GpeSetEncapMode)(nil),
		(*GpeSetEncapModeReply)(nil),
	}
}

// RPCService represents RPC service API for lisp_gpe module.
type RPCService interface {
	DumpGpeFwdEntryPath(ctx context.Context, in *GpeFwdEntryPathDump) (RPCService_DumpGpeFwdEntryPathClient, error)
	GpeAddDelFwdEntry(ctx context.Context, in *GpeAddDelFwdEntry) (*GpeAddDelFwdEntryReply, error)
	GpeAddDelIface(ctx context.Context, in *GpeAddDelIface) (*GpeAddDelIfaceReply, error)
	GpeAddDelNativeFwdRpath(ctx context.Context, in *GpeAddDelNativeFwdRpath) (*GpeAddDelNativeFwdRpathReply, error)
	GpeEnableDisable(ctx context.Context, in *GpeEnableDisable) (*GpeEnableDisableReply, error)
	GpeFwdEntriesGet(ctx context.Context, in *GpeFwdEntriesGet) (*GpeFwdEntriesGetReply, error)
	GpeFwdEntryVnisGet(ctx context.Context, in *GpeFwdEntryVnisGet) (*GpeFwdEntryVnisGetReply, error)
	GpeGetEncapMode(ctx context.Context, in *GpeGetEncapMode) (*GpeGetEncapModeReply, error)
	GpeNativeFwdRpathsGet(ctx context.Context, in *GpeNativeFwdRpathsGet) (*GpeNativeFwdRpathsGetReply, error)
	GpeSetEncapMode(ctx context.Context, in *GpeSetEncapMode) (*GpeSetEncapModeReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpGpeFwdEntryPath(ctx context.Context, in *GpeFwdEntryPathDump) (RPCService_DumpGpeFwdEntryPathClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpGpeFwdEntryPathClient{stream}
	return x, nil
}

type RPCService_DumpGpeFwdEntryPathClient interface {
	Recv() (*GpeFwdEntryPathDetails, error)
}

type serviceClient_DumpGpeFwdEntryPathClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpGpeFwdEntryPathClient) Recv() (*GpeFwdEntryPathDetails, error) {
	m := new(GpeFwdEntryPathDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) GpeAddDelFwdEntry(ctx context.Context, in *GpeAddDelFwdEntry) (*GpeAddDelFwdEntryReply, error) {
	out := new(GpeAddDelFwdEntryReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeAddDelIface(ctx context.Context, in *GpeAddDelIface) (*GpeAddDelIfaceReply, error) {
	out := new(GpeAddDelIfaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeAddDelNativeFwdRpath(ctx context.Context, in *GpeAddDelNativeFwdRpath) (*GpeAddDelNativeFwdRpathReply, error) {
	out := new(GpeAddDelNativeFwdRpathReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeEnableDisable(ctx context.Context, in *GpeEnableDisable) (*GpeEnableDisableReply, error) {
	out := new(GpeEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeFwdEntriesGet(ctx context.Context, in *GpeFwdEntriesGet) (*GpeFwdEntriesGetReply, error) {
	out := new(GpeFwdEntriesGetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeFwdEntryVnisGet(ctx context.Context, in *GpeFwdEntryVnisGet) (*GpeFwdEntryVnisGetReply, error) {
	out := new(GpeFwdEntryVnisGetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeGetEncapMode(ctx context.Context, in *GpeGetEncapMode) (*GpeGetEncapModeReply, error) {
	out := new(GpeGetEncapModeReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeNativeFwdRpathsGet(ctx context.Context, in *GpeNativeFwdRpathsGet) (*GpeNativeFwdRpathsGetReply, error) {
	out := new(GpeNativeFwdRpathsGetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GpeSetEncapMode(ctx context.Context, in *GpeSetEncapMode) (*GpeSetEncapModeReply, error) {
	out := new(GpeSetEncapModeReply)
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
