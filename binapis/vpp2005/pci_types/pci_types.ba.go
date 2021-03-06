// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/core/pci_types.api.json

/*
Package pci_types is a generated VPP binary API for 'pci_types' module.

It consists of:
	  1 type
*/
package pci_types

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
	ModuleName = "pci_types"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x5d418665
)

// PciAddress represents VPP binary API type 'pci_address'.
type PciAddress struct {
	Domain   uint16
	Bus      uint8
	Slot     uint8
	Function uint8
}

func (*PciAddress) GetTypeName() string { return "pci_address" }

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
