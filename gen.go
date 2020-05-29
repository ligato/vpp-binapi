// +build generate

//go:generate env VPP_IMG=ligato/vpp-base:20.05 ./genbinapi.sh vpp2005
//go:generate env VPP_IMG=ligato/vpp-base:20.01 ./genbinapi.sh vpp2001
//go:generate env VPP_IMG=ligato/vpp-base:19.08 ./genbinapi.sh vpp1908
//go:generate env VPP_IMG=ligato/vpp-base:19.04 ./genbinapi.sh vpp1904

package binapi
