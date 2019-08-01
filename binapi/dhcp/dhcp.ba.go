// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: vppapi/core/dhcp.api.json

/*
Package dhcp is a generated VPP binary API for 'dhcp' module.

It consists of:
	  4 enums
	  2 aliases
	 11 types
	  1 union
	 25 messages
	 11 services
*/
package dhcp

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
	ModuleName = "dhcp"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x8c5157ba
)

// AddressFamily represents VPP binary API enum 'address_family'.
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

var AddressFamily_name = map[uint32]string{
	0: "ADDRESS_IP4",
	1: "ADDRESS_IP6",
}

var AddressFamily_value = map[string]uint32{
	"ADDRESS_IP4": 0,
	"ADDRESS_IP6": 1,
}

func (x AddressFamily) String() string {
	s, ok := AddressFamily_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPDscp represents VPP binary API enum 'ip_dscp'.
type IPDscp uint8

const (
	IP_API_DSCP_CS0  IPDscp = 0
	IP_API_DSCP_CS1  IPDscp = 8
	IP_API_DSCP_AF11 IPDscp = 10
	IP_API_DSCP_AF12 IPDscp = 12
	IP_API_DSCP_AF13 IPDscp = 14
	IP_API_DSCP_CS2  IPDscp = 16
	IP_API_DSCP_AF21 IPDscp = 18
	IP_API_DSCP_AF22 IPDscp = 20
	IP_API_DSCP_AF23 IPDscp = 22
	IP_API_DSCP_CS3  IPDscp = 24
	IP_API_DSCP_AF31 IPDscp = 26
	IP_API_DSCP_AF32 IPDscp = 28
	IP_API_DSCP_AF33 IPDscp = 30
	IP_API_DSCP_CS4  IPDscp = 32
	IP_API_DSCP_AF41 IPDscp = 34
	IP_API_DSCP_AF42 IPDscp = 36
	IP_API_DSCP_AF43 IPDscp = 38
	IP_API_DSCP_CS5  IPDscp = 40
	IP_API_DSCP_EF   IPDscp = 46
	IP_API_DSCP_CS6  IPDscp = 48
	IP_API_DSCP_CS7  IPDscp = 50
)

var IPDscp_name = map[uint8]string{
	0:  "IP_API_DSCP_CS0",
	8:  "IP_API_DSCP_CS1",
	10: "IP_API_DSCP_AF11",
	12: "IP_API_DSCP_AF12",
	14: "IP_API_DSCP_AF13",
	16: "IP_API_DSCP_CS2",
	18: "IP_API_DSCP_AF21",
	20: "IP_API_DSCP_AF22",
	22: "IP_API_DSCP_AF23",
	24: "IP_API_DSCP_CS3",
	26: "IP_API_DSCP_AF31",
	28: "IP_API_DSCP_AF32",
	30: "IP_API_DSCP_AF33",
	32: "IP_API_DSCP_CS4",
	34: "IP_API_DSCP_AF41",
	36: "IP_API_DSCP_AF42",
	38: "IP_API_DSCP_AF43",
	40: "IP_API_DSCP_CS5",
	46: "IP_API_DSCP_EF",
	48: "IP_API_DSCP_CS6",
	50: "IP_API_DSCP_CS7",
}

var IPDscp_value = map[string]uint8{
	"IP_API_DSCP_CS0":  0,
	"IP_API_DSCP_CS1":  8,
	"IP_API_DSCP_AF11": 10,
	"IP_API_DSCP_AF12": 12,
	"IP_API_DSCP_AF13": 14,
	"IP_API_DSCP_CS2":  16,
	"IP_API_DSCP_AF21": 18,
	"IP_API_DSCP_AF22": 20,
	"IP_API_DSCP_AF23": 22,
	"IP_API_DSCP_CS3":  24,
	"IP_API_DSCP_AF31": 26,
	"IP_API_DSCP_AF32": 28,
	"IP_API_DSCP_AF33": 30,
	"IP_API_DSCP_CS4":  32,
	"IP_API_DSCP_AF41": 34,
	"IP_API_DSCP_AF42": 36,
	"IP_API_DSCP_AF43": 38,
	"IP_API_DSCP_CS5":  40,
	"IP_API_DSCP_EF":   46,
	"IP_API_DSCP_CS6":  48,
	"IP_API_DSCP_CS7":  50,
}

func (x IPDscp) String() string {
	s, ok := IPDscp_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPEcn represents VPP binary API enum 'ip_ecn'.
type IPEcn uint8

const (
	IP_API_ECN_NONE IPEcn = 0
	IP_API_ECN_ECT0 IPEcn = 1
	IP_API_ECN_ECT1 IPEcn = 2
	IP_API_ECN_CE   IPEcn = 3
)

var IPEcn_name = map[uint8]string{
	0: "IP_API_ECN_NONE",
	1: "IP_API_ECN_ECT0",
	2: "IP_API_ECN_ECT1",
	3: "IP_API_ECN_CE",
}

var IPEcn_value = map[string]uint8{
	"IP_API_ECN_NONE": 0,
	"IP_API_ECN_ECT0": 1,
	"IP_API_ECN_ECT1": 2,
	"IP_API_ECN_CE":   3,
}

func (x IPEcn) String() string {
	s, ok := IPEcn_name[uint8(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IPProto represents VPP binary API enum 'ip_proto'.
type IPProto uint32

const (
	IP_API_PROTO_HOPOPT   IPProto = 0
	IP_API_PROTO_ICMP     IPProto = 1
	IP_API_PROTO_IGMP     IPProto = 2
	IP_API_PROTO_TCP      IPProto = 6
	IP_API_PROTO_UDP      IPProto = 17
	IP_API_PROTO_GRE      IPProto = 47
	IP_API_PROTO_AH       IPProto = 50
	IP_API_PROTO_ESP      IPProto = 51
	IP_API_PROTO_EIGRP    IPProto = 88
	IP_API_PROTO_OSPF     IPProto = 89
	IP_API_PROTO_SCTP     IPProto = 132
	IP_API_PROTO_RESERVED IPProto = 255
)

var IPProto_name = map[uint32]string{
	0:   "IP_API_PROTO_HOPOPT",
	1:   "IP_API_PROTO_ICMP",
	2:   "IP_API_PROTO_IGMP",
	6:   "IP_API_PROTO_TCP",
	17:  "IP_API_PROTO_UDP",
	47:  "IP_API_PROTO_GRE",
	50:  "IP_API_PROTO_AH",
	51:  "IP_API_PROTO_ESP",
	88:  "IP_API_PROTO_EIGRP",
	89:  "IP_API_PROTO_OSPF",
	132: "IP_API_PROTO_SCTP",
	255: "IP_API_PROTO_RESERVED",
}

var IPProto_value = map[string]uint32{
	"IP_API_PROTO_HOPOPT":   0,
	"IP_API_PROTO_ICMP":     1,
	"IP_API_PROTO_IGMP":     2,
	"IP_API_PROTO_TCP":      6,
	"IP_API_PROTO_UDP":      17,
	"IP_API_PROTO_GRE":      47,
	"IP_API_PROTO_AH":       50,
	"IP_API_PROTO_ESP":      51,
	"IP_API_PROTO_EIGRP":    88,
	"IP_API_PROTO_OSPF":     89,
	"IP_API_PROTO_SCTP":     132,
	"IP_API_PROTO_RESERVED": 255,
}

func (x IPProto) String() string {
	s, ok := IPProto_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// IP4Address represents VPP binary API alias 'ip4_address'.
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address'.
type IP6Address [16]uint8

// Address represents VPP binary API type 'address'.
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}

// DHCP6AddressInfo represents VPP binary API type 'dhcp6_address_info'.
type DHCP6AddressInfo struct {
	Address       []byte `struc:"[16]byte"`
	ValidTime     uint32
	PreferredTime uint32
}

func (*DHCP6AddressInfo) GetTypeName() string {
	return "dhcp6_address_info"
}

// DHCP6PdPrefixInfo represents VPP binary API type 'dhcp6_pd_prefix_info'.
type DHCP6PdPrefixInfo struct {
	Prefix        []byte `struc:"[16]byte"`
	PrefixLength  uint8
	ValidTime     uint32
	PreferredTime uint32
}

func (*DHCP6PdPrefixInfo) GetTypeName() string {
	return "dhcp6_pd_prefix_info"
}

// DHCPClient represents VPP binary API type 'dhcp_client'.
type DHCPClient struct {
	SwIfIndex        uint32
	Hostname         []byte `struc:"[64]byte"`
	ID               []byte `struc:"[64]byte"`
	WantDHCPEvent    uint8
	SetBroadcastFlag uint8
	Dscp             IPDscp
	PID              uint32
}

func (*DHCPClient) GetTypeName() string {
	return "dhcp_client"
}

// DHCPLease represents VPP binary API type 'dhcp_lease'.
type DHCPLease struct {
	SwIfIndex     uint32
	State         uint8
	Hostname      []byte `struc:"[64]byte"`
	IsIPv6        uint8
	MaskWidth     uint8
	HostAddress   []byte `struc:"[16]byte"`
	RouterAddress []byte `struc:"[16]byte"`
	HostMac       []byte `struc:"[6]byte"`
	Count         uint8  `struc:"sizeof=DomainServer"`
	DomainServer  []DomainServer
}

func (*DHCPLease) GetTypeName() string {
	return "dhcp_lease"
}

// DHCPServer represents VPP binary API type 'dhcp_server'.
type DHCPServer struct {
	ServerVrfID uint32
	DHCPServer  []byte `struc:"[16]byte"`
}

func (*DHCPServer) GetTypeName() string {
	return "dhcp_server"
}

// DomainServer represents VPP binary API type 'domain_server'.
type DomainServer struct {
	Address []byte `struc:"[16]byte"`
}

func (*DomainServer) GetTypeName() string {
	return "domain_server"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix'.
type IP4Prefix struct {
	Address IP4Address
	Len     uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix'.
type IP6Prefix struct {
	Address IP6Address
	Len     uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}

// Mprefix represents VPP binary API type 'mprefix'.
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}

// Prefix represents VPP binary API type 'prefix'.
type Prefix struct {
	Address Address
	Len     uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}

// AddressUnion represents VPP binary API union 'address_union'.
type AddressUnion struct {
	XXX_UnionData [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

// DHCP6ClientsEnableDisable represents VPP binary API message 'dhcp6_clients_enable_disable'.
type DHCP6ClientsEnableDisable struct {
	Enable uint8
}

func (*DHCP6ClientsEnableDisable) GetMessageName() string {
	return "dhcp6_clients_enable_disable"
}
func (*DHCP6ClientsEnableDisable) GetCrcString() string {
	return "8050327d"
}
func (*DHCP6ClientsEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6ClientsEnableDisableReply represents VPP binary API message 'dhcp6_clients_enable_disable_reply'.
type DHCP6ClientsEnableDisableReply struct {
	Retval int32
}

func (*DHCP6ClientsEnableDisableReply) GetMessageName() string {
	return "dhcp6_clients_enable_disable_reply"
}
func (*DHCP6ClientsEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6ClientsEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6DuidLlSet represents VPP binary API message 'dhcp6_duid_ll_set'.
type DHCP6DuidLlSet struct {
	DuidLl []byte `struc:"[10]byte"`
}

func (*DHCP6DuidLlSet) GetMessageName() string {
	return "dhcp6_duid_ll_set"
}
func (*DHCP6DuidLlSet) GetCrcString() string {
	return "0f6ca323"
}
func (*DHCP6DuidLlSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6DuidLlSetReply represents VPP binary API message 'dhcp6_duid_ll_set_reply'.
type DHCP6DuidLlSetReply struct {
	Retval int32
}

func (*DHCP6DuidLlSetReply) GetMessageName() string {
	return "dhcp6_duid_ll_set_reply"
}
func (*DHCP6DuidLlSetReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6DuidLlSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6PdReplyEvent represents VPP binary API message 'dhcp6_pd_reply_event'.
type DHCP6PdReplyEvent struct {
	PID             uint32
	SwIfIndex       uint32
	ServerIndex     uint32
	MsgType         uint8
	T1              uint32
	T2              uint32
	InnerStatusCode uint16
	StatusCode      uint16
	Preference      uint8
	NPrefixes       uint32 `struc:"sizeof=Prefixes"`
	Prefixes        []DHCP6PdPrefixInfo
}

func (*DHCP6PdReplyEvent) GetMessageName() string {
	return "dhcp6_pd_reply_event"
}
func (*DHCP6PdReplyEvent) GetCrcString() string {
	return "0e53217a"
}
func (*DHCP6PdReplyEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCP6PdSendClientMessage represents VPP binary API message 'dhcp6_pd_send_client_message'.
type DHCP6PdSendClientMessage struct {
	SwIfIndex   uint32
	ServerIndex uint32
	Irt         uint32
	Mrt         uint32
	Mrc         uint32
	Mrd         uint32
	Stop        uint8
	MsgType     uint8
	T1          uint32
	T2          uint32
	NPrefixes   uint32 `struc:"sizeof=Prefixes"`
	Prefixes    []DHCP6PdPrefixInfo
}

func (*DHCP6PdSendClientMessage) GetMessageName() string {
	return "dhcp6_pd_send_client_message"
}
func (*DHCP6PdSendClientMessage) GetCrcString() string {
	return "dadbfe97"
}
func (*DHCP6PdSendClientMessage) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6PdSendClientMessageReply represents VPP binary API message 'dhcp6_pd_send_client_message_reply'.
type DHCP6PdSendClientMessageReply struct {
	Retval int32
}

func (*DHCP6PdSendClientMessageReply) GetMessageName() string {
	return "dhcp6_pd_send_client_message_reply"
}
func (*DHCP6PdSendClientMessageReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6PdSendClientMessageReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6ReplyEvent represents VPP binary API message 'dhcp6_reply_event'.
type DHCP6ReplyEvent struct {
	PID             uint32
	SwIfIndex       uint32
	ServerIndex     uint32
	MsgType         uint8
	T1              uint32
	T2              uint32
	InnerStatusCode uint16
	StatusCode      uint16
	Preference      uint8
	NAddresses      uint32 `struc:"sizeof=Addresses"`
	Addresses       []DHCP6AddressInfo
}

func (*DHCP6ReplyEvent) GetMessageName() string {
	return "dhcp6_reply_event"
}
func (*DHCP6ReplyEvent) GetCrcString() string {
	return "8a34e0f5"
}
func (*DHCP6ReplyEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCP6SendClientMessage represents VPP binary API message 'dhcp6_send_client_message'.
type DHCP6SendClientMessage struct {
	SwIfIndex   uint32
	ServerIndex uint32
	Irt         uint32
	Mrt         uint32
	Mrc         uint32
	Mrd         uint32
	Stop        uint8
	MsgType     uint8
	T1          uint32
	T2          uint32
	NAddresses  uint32 `struc:"sizeof=Addresses"`
	Addresses   []DHCP6AddressInfo
}

func (*DHCP6SendClientMessage) GetMessageName() string {
	return "dhcp6_send_client_message"
}
func (*DHCP6SendClientMessage) GetCrcString() string {
	return "993f872f"
}
func (*DHCP6SendClientMessage) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6SendClientMessageReply represents VPP binary API message 'dhcp6_send_client_message_reply'.
type DHCP6SendClientMessageReply struct {
	Retval int32
}

func (*DHCP6SendClientMessageReply) GetMessageName() string {
	return "dhcp6_send_client_message_reply"
}
func (*DHCP6SendClientMessageReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6SendClientMessageReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientConfig represents VPP binary API message 'dhcp_client_config'.
type DHCPClientConfig struct {
	IsAdd  uint8
	Client DHCPClient
}

func (*DHCPClientConfig) GetMessageName() string {
	return "dhcp_client_config"
}
func (*DHCPClientConfig) GetCrcString() string {
	return "87a429e7"
}
func (*DHCPClientConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPClientConfigReply represents VPP binary API message 'dhcp_client_config_reply'.
type DHCPClientConfigReply struct {
	Retval int32
}

func (*DHCPClientConfigReply) GetMessageName() string {
	return "dhcp_client_config_reply"
}
func (*DHCPClientConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPClientConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientDetails represents VPP binary API message 'dhcp_client_details'.
type DHCPClientDetails struct {
	Client DHCPClient
	Lease  DHCPLease
}

func (*DHCPClientDetails) GetMessageName() string {
	return "dhcp_client_details"
}
func (*DHCPClientDetails) GetCrcString() string {
	return "4a95a2ad"
}
func (*DHCPClientDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientDump represents VPP binary API message 'dhcp_client_dump'.
type DHCPClientDump struct{}

func (*DHCPClientDump) GetMessageName() string {
	return "dhcp_client_dump"
}
func (*DHCPClientDump) GetCrcString() string {
	return "51077d14"
}
func (*DHCPClientDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPComplEvent represents VPP binary API message 'dhcp_compl_event'.
type DHCPComplEvent struct {
	PID   uint32
	Lease DHCPLease
}

func (*DHCPComplEvent) GetMessageName() string {
	return "dhcp_compl_event"
}
func (*DHCPComplEvent) GetCrcString() string {
	return "ed1e53d7"
}
func (*DHCPComplEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCPProxyConfig represents VPP binary API message 'dhcp_proxy_config'.
type DHCPProxyConfig struct {
	RxVrfID        uint32
	ServerVrfID    uint32
	IsIPv6         uint8
	IsAdd          uint8
	DHCPServer     []byte `struc:"[16]byte"`
	DHCPSrcAddress []byte `struc:"[16]byte"`
}

func (*DHCPProxyConfig) GetMessageName() string {
	return "dhcp_proxy_config"
}
func (*DHCPProxyConfig) GetCrcString() string {
	return "6af4b645"
}
func (*DHCPProxyConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxyConfigReply represents VPP binary API message 'dhcp_proxy_config_reply'.
type DHCPProxyConfigReply struct {
	Retval int32
}

func (*DHCPProxyConfigReply) GetMessageName() string {
	return "dhcp_proxy_config_reply"
}
func (*DHCPProxyConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPProxyConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPProxyDetails represents VPP binary API message 'dhcp_proxy_details'.
type DHCPProxyDetails struct {
	RxVrfID        uint32
	VssOui         uint32
	VssFibID       uint32
	VssType        uint8
	VssVPNAsciiID  []byte `struc:"[129]byte"`
	IsIPv6         uint8
	DHCPSrcAddress []byte `struc:"[16]byte"`
	Count          uint8  `struc:"sizeof=Servers"`
	Servers        []DHCPServer
}

func (*DHCPProxyDetails) GetMessageName() string {
	return "dhcp_proxy_details"
}
func (*DHCPProxyDetails) GetCrcString() string {
	return "e6c45917"
}
func (*DHCPProxyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPProxyDump represents VPP binary API message 'dhcp_proxy_dump'.
type DHCPProxyDump struct {
	IsIP6 uint8
}

func (*DHCPProxyDump) GetMessageName() string {
	return "dhcp_proxy_dump"
}
func (*DHCPProxyDump) GetCrcString() string {
	return "6fe91190"
}
func (*DHCPProxyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxySetVss represents VPP binary API message 'dhcp_proxy_set_vss'.
type DHCPProxySetVss struct {
	TblID      uint32
	VssType    uint8
	VPNAsciiID []byte `struc:"[129]byte"`
	Oui        uint32
	VPNIndex   uint32
	IsIPv6     uint8
	IsAdd      uint8
}

func (*DHCPProxySetVss) GetMessageName() string {
	return "dhcp_proxy_set_vss"
}
func (*DHCPProxySetVss) GetCrcString() string {
	return "606535aa"
}
func (*DHCPProxySetVss) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxySetVssReply represents VPP binary API message 'dhcp_proxy_set_vss_reply'.
type DHCPProxySetVssReply struct {
	Retval int32
}

func (*DHCPProxySetVssReply) GetMessageName() string {
	return "dhcp_proxy_set_vss_reply"
}
func (*DHCPProxySetVssReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPProxySetVssReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantDHCP6PdReplyEvents represents VPP binary API message 'want_dhcp6_pd_reply_events'.
type WantDHCP6PdReplyEvents struct {
	EnableDisable uint8
	PID           uint32
}

func (*WantDHCP6PdReplyEvents) GetMessageName() string {
	return "want_dhcp6_pd_reply_events"
}
func (*WantDHCP6PdReplyEvents) GetCrcString() string {
	return "05b454b5"
}
func (*WantDHCP6PdReplyEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantDHCP6PdReplyEventsReply represents VPP binary API message 'want_dhcp6_pd_reply_events_reply'.
type WantDHCP6PdReplyEventsReply struct {
	Retval int32
}

func (*WantDHCP6PdReplyEventsReply) GetMessageName() string {
	return "want_dhcp6_pd_reply_events_reply"
}
func (*WantDHCP6PdReplyEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantDHCP6PdReplyEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantDHCP6ReplyEvents represents VPP binary API message 'want_dhcp6_reply_events'.
type WantDHCP6ReplyEvents struct {
	EnableDisable uint8
	PID           uint32
}

func (*WantDHCP6ReplyEvents) GetMessageName() string {
	return "want_dhcp6_reply_events"
}
func (*WantDHCP6ReplyEvents) GetCrcString() string {
	return "05b454b5"
}
func (*WantDHCP6ReplyEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantDHCP6ReplyEventsReply represents VPP binary API message 'want_dhcp6_reply_events_reply'.
type WantDHCP6ReplyEventsReply struct {
	Retval int32
}

func (*WantDHCP6ReplyEventsReply) GetMessageName() string {
	return "want_dhcp6_reply_events_reply"
}
func (*WantDHCP6ReplyEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantDHCP6ReplyEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*DHCP6ClientsEnableDisable)(nil), "dhcp.DHCP6ClientsEnableDisable")
	api.RegisterMessage((*DHCP6ClientsEnableDisableReply)(nil), "dhcp.DHCP6ClientsEnableDisableReply")
	api.RegisterMessage((*DHCP6DuidLlSet)(nil), "dhcp.DHCP6DuidLlSet")
	api.RegisterMessage((*DHCP6DuidLlSetReply)(nil), "dhcp.DHCP6DuidLlSetReply")
	api.RegisterMessage((*DHCP6PdReplyEvent)(nil), "dhcp.DHCP6PdReplyEvent")
	api.RegisterMessage((*DHCP6PdSendClientMessage)(nil), "dhcp.DHCP6PdSendClientMessage")
	api.RegisterMessage((*DHCP6PdSendClientMessageReply)(nil), "dhcp.DHCP6PdSendClientMessageReply")
	api.RegisterMessage((*DHCP6ReplyEvent)(nil), "dhcp.DHCP6ReplyEvent")
	api.RegisterMessage((*DHCP6SendClientMessage)(nil), "dhcp.DHCP6SendClientMessage")
	api.RegisterMessage((*DHCP6SendClientMessageReply)(nil), "dhcp.DHCP6SendClientMessageReply")
	api.RegisterMessage((*DHCPClientConfig)(nil), "dhcp.DHCPClientConfig")
	api.RegisterMessage((*DHCPClientConfigReply)(nil), "dhcp.DHCPClientConfigReply")
	api.RegisterMessage((*DHCPClientDetails)(nil), "dhcp.DHCPClientDetails")
	api.RegisterMessage((*DHCPClientDump)(nil), "dhcp.DHCPClientDump")
	api.RegisterMessage((*DHCPComplEvent)(nil), "dhcp.DHCPComplEvent")
	api.RegisterMessage((*DHCPProxyConfig)(nil), "dhcp.DHCPProxyConfig")
	api.RegisterMessage((*DHCPProxyConfigReply)(nil), "dhcp.DHCPProxyConfigReply")
	api.RegisterMessage((*DHCPProxyDetails)(nil), "dhcp.DHCPProxyDetails")
	api.RegisterMessage((*DHCPProxyDump)(nil), "dhcp.DHCPProxyDump")
	api.RegisterMessage((*DHCPProxySetVss)(nil), "dhcp.DHCPProxySetVss")
	api.RegisterMessage((*DHCPProxySetVssReply)(nil), "dhcp.DHCPProxySetVssReply")
	api.RegisterMessage((*WantDHCP6PdReplyEvents)(nil), "dhcp.WantDHCP6PdReplyEvents")
	api.RegisterMessage((*WantDHCP6PdReplyEventsReply)(nil), "dhcp.WantDHCP6PdReplyEventsReply")
	api.RegisterMessage((*WantDHCP6ReplyEvents)(nil), "dhcp.WantDHCP6ReplyEvents")
	api.RegisterMessage((*WantDHCP6ReplyEventsReply)(nil), "dhcp.WantDHCP6ReplyEventsReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DHCP6ClientsEnableDisable)(nil),
		(*DHCP6ClientsEnableDisableReply)(nil),
		(*DHCP6DuidLlSet)(nil),
		(*DHCP6DuidLlSetReply)(nil),
		(*DHCP6PdReplyEvent)(nil),
		(*DHCP6PdSendClientMessage)(nil),
		(*DHCP6PdSendClientMessageReply)(nil),
		(*DHCP6ReplyEvent)(nil),
		(*DHCP6SendClientMessage)(nil),
		(*DHCP6SendClientMessageReply)(nil),
		(*DHCPClientConfig)(nil),
		(*DHCPClientConfigReply)(nil),
		(*DHCPClientDetails)(nil),
		(*DHCPClientDump)(nil),
		(*DHCPComplEvent)(nil),
		(*DHCPProxyConfig)(nil),
		(*DHCPProxyConfigReply)(nil),
		(*DHCPProxyDetails)(nil),
		(*DHCPProxyDump)(nil),
		(*DHCPProxySetVss)(nil),
		(*DHCPProxySetVssReply)(nil),
		(*WantDHCP6PdReplyEvents)(nil),
		(*WantDHCP6PdReplyEventsReply)(nil),
		(*WantDHCP6ReplyEvents)(nil),
		(*WantDHCP6ReplyEventsReply)(nil),
	}
}

// RPCService represents RPC service API for dhcp module.
type RPCService interface {
	DumpDHCPClient(ctx context.Context, in *DHCPClientDump) (RPCService_DumpDHCPClientClient, error)
	DumpDHCPProxy(ctx context.Context, in *DHCPProxyDump) (RPCService_DumpDHCPProxyClient, error)
	DHCP6ClientsEnableDisable(ctx context.Context, in *DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error)
	DHCP6DuidLlSet(ctx context.Context, in *DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error)
	DHCP6PdSendClientMessage(ctx context.Context, in *DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error)
	DHCP6SendClientMessage(ctx context.Context, in *DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error)
	DHCPClientConfig(ctx context.Context, in *DHCPClientConfig) (*DHCPClientConfigReply, error)
	DHCPProxyConfig(ctx context.Context, in *DHCPProxyConfig) (*DHCPProxyConfigReply, error)
	DHCPProxySetVss(ctx context.Context, in *DHCPProxySetVss) (*DHCPProxySetVssReply, error)
	WantDHCP6PdReplyEvents(ctx context.Context, in *WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error)
	WantDHCP6ReplyEvents(ctx context.Context, in *WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpDHCPClient(ctx context.Context, in *DHCPClientDump) (RPCService_DumpDHCPClientClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpDHCPClientClient{stream}
	return x, nil
}

type RPCService_DumpDHCPClientClient interface {
	Recv() (*DHCPClientDetails, error)
}

type serviceClient_DumpDHCPClientClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpDHCPClientClient) Recv() (*DHCPClientDetails, error) {
	m := new(DHCPClientDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpDHCPProxy(ctx context.Context, in *DHCPProxyDump) (RPCService_DumpDHCPProxyClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpDHCPProxyClient{stream}
	return x, nil
}

type RPCService_DumpDHCPProxyClient interface {
	Recv() (*DHCPProxyDetails, error)
}

type serviceClient_DumpDHCPProxyClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpDHCPProxyClient) Recv() (*DHCPProxyDetails, error) {
	m := new(DHCPProxyDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DHCP6ClientsEnableDisable(ctx context.Context, in *DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error) {
	out := new(DHCP6ClientsEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCP6DuidLlSet(ctx context.Context, in *DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error) {
	out := new(DHCP6DuidLlSetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCP6PdSendClientMessage(ctx context.Context, in *DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error) {
	out := new(DHCP6PdSendClientMessageReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCP6SendClientMessage(ctx context.Context, in *DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error) {
	out := new(DHCP6SendClientMessageReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPClientConfig(ctx context.Context, in *DHCPClientConfig) (*DHCPClientConfigReply, error) {
	out := new(DHCPClientConfigReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPProxyConfig(ctx context.Context, in *DHCPProxyConfig) (*DHCPProxyConfigReply, error) {
	out := new(DHCPProxyConfigReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPProxySetVss(ctx context.Context, in *DHCPProxySetVss) (*DHCPProxySetVssReply, error) {
	out := new(DHCPProxySetVssReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) WantDHCP6PdReplyEvents(ctx context.Context, in *WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error) {
	out := new(WantDHCP6PdReplyEventsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) WantDHCP6ReplyEvents(ctx context.Context, in *WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error) {
	out := new(WantDHCP6ReplyEventsReply)
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
