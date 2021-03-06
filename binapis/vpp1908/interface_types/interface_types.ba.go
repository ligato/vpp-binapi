// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp1908/.vppapi/core/interface_types.api.json

/*
Package interface_types is a generated VPP binary API for 'interface_types' module.

It consists of:
	  1 alias
*/
package interface_types

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
	ModuleName = "interface_types"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xfa73ea25
)

// InterfaceIndex represents VPP binary API alias 'interface_index'.
type InterfaceIndex uint32

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
