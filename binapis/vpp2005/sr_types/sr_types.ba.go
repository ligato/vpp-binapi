// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: binapis/vpp2005/.vppapi/core/sr_types.api.json

/*
Package sr_types is a generated VPP binary API for 'sr_types' module.

It consists of:
	  3 enums
*/
package sr_types

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
	ModuleName = "sr_types"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x2ac0a9d5
)

// SrBehavior represents VPP binary API enum 'sr_behavior'.
type SrBehavior uint8

const (
	SR_BEHAVIOR_API_END     SrBehavior = 1
	SR_BEHAVIOR_API_X       SrBehavior = 2
	SR_BEHAVIOR_API_T       SrBehavior = 3
	SR_BEHAVIOR_API_D_FIRST SrBehavior = 4
	SR_BEHAVIOR_API_DX2     SrBehavior = 5
	SR_BEHAVIOR_API_DX6     SrBehavior = 6
	SR_BEHAVIOR_API_DX4     SrBehavior = 7
	SR_BEHAVIOR_API_DT6     SrBehavior = 8
	SR_BEHAVIOR_API_DT4     SrBehavior = 9
	SR_BEHAVIOR_API_LAST    SrBehavior = 10
)

var SrBehavior_name = map[uint8]string{
	1:  "SR_BEHAVIOR_API_END",
	2:  "SR_BEHAVIOR_API_X",
	3:  "SR_BEHAVIOR_API_T",
	4:  "SR_BEHAVIOR_API_D_FIRST",
	5:  "SR_BEHAVIOR_API_DX2",
	6:  "SR_BEHAVIOR_API_DX6",
	7:  "SR_BEHAVIOR_API_DX4",
	8:  "SR_BEHAVIOR_API_DT6",
	9:  "SR_BEHAVIOR_API_DT4",
	10: "SR_BEHAVIOR_API_LAST",
}

var SrBehavior_value = map[string]uint8{
	"SR_BEHAVIOR_API_END":     1,
	"SR_BEHAVIOR_API_X":       2,
	"SR_BEHAVIOR_API_T":       3,
	"SR_BEHAVIOR_API_D_FIRST": 4,
	"SR_BEHAVIOR_API_DX2":     5,
	"SR_BEHAVIOR_API_DX6":     6,
	"SR_BEHAVIOR_API_DX4":     7,
	"SR_BEHAVIOR_API_DT6":     8,
	"SR_BEHAVIOR_API_DT4":     9,
	"SR_BEHAVIOR_API_LAST":    10,
}

func (x SrBehavior) String() string {
	s, ok := SrBehavior_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SrPolicyOp represents VPP binary API enum 'sr_policy_op'.
type SrPolicyOp uint8

const (
	SR_POLICY_OP_API_NONE SrPolicyOp = 0
	SR_POLICY_OP_API_ADD  SrPolicyOp = 1
	SR_POLICY_OP_API_DEL  SrPolicyOp = 2
	SR_POLICY_OP_API_MOD  SrPolicyOp = 3
)

var SrPolicyOp_name = map[uint8]string{
	0: "SR_POLICY_OP_API_NONE",
	1: "SR_POLICY_OP_API_ADD",
	2: "SR_POLICY_OP_API_DEL",
	3: "SR_POLICY_OP_API_MOD",
}

var SrPolicyOp_value = map[string]uint8{
	"SR_POLICY_OP_API_NONE": 0,
	"SR_POLICY_OP_API_ADD":  1,
	"SR_POLICY_OP_API_DEL":  2,
	"SR_POLICY_OP_API_MOD":  3,
}

func (x SrPolicyOp) String() string {
	s, ok := SrPolicyOp_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// SrSteer represents VPP binary API enum 'sr_steer'.
type SrSteer uint8

const (
	SR_STEER_API_L2   SrSteer = 2
	SR_STEER_API_IPV4 SrSteer = 4
	SR_STEER_API_IPV6 SrSteer = 6
)

var SrSteer_name = map[uint8]string{
	2: "SR_STEER_API_L2",
	4: "SR_STEER_API_IPV4",
	6: "SR_STEER_API_IPV6",
}

var SrSteer_value = map[string]uint8{
	"SR_STEER_API_L2":   2,
	"SR_STEER_API_IPV4": 4,
	"SR_STEER_API_IPV6": 6,
}

func (x SrSteer) String() string {
	s, ok := SrSteer_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
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
