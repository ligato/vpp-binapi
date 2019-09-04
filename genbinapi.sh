#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"

vpp_img=${VPP_IMG:-'ligato/vpp-base:master'}

function log() {
	echo >&2 -ne "\e[1;30m"
	echo >&2 "$*"
	echo >&2 -ne "\e[0;0m"
}

function export_vppapi() {
	if [ ! -n "${NOPULL-}" ]; then
		log "# Pulling VPP image ${vpp_img}"
		docker pull "${vpp_img}"
	fi

	vpp_ver=$(docker run --rm "$vpp_img" dpkg-query -f '${Version}\n' -W vpp)

	log "# Exporting VPP API"
	echo "vpp $vpp_ver"

	docker run --rm -u $(id -u):$(id -g) \
		-v $(pwd)/vppapi:/vpp/api \
		"$vpp_img" \
		sh -c 'cp -r /usr/share/vpp/api/* /vpp/api'
}

function generate_binapi() {
	if [ ! -n "${NOINSTALL-}" ]; then
		log "# Installing binapi generator"
		go install git.fd.io/govpp.git/cmd/binapi-generator
	fi
	binapi-generator version

	log "# Generating binapi code"
	out=$(binapi-generator --input-dir=vppapi --output-dir=binapi 2>&1)

	warns=$(echo "$out" | grep -ci WARN)
	if [ "${warns}" -gt "0" ]; then
		echo -e "\e[0;33mDetected $warns warnings\e[0;0m"
	fi
}

function check_previous() {
	if [ -r VPP_VERSION ]; then
		prev_ver=$(cat VPP_VERSION)
		log "# Found previously generated for VPP $prev_ver"
	fi
}

cd ${SCRIPT_DIR}

#check_previous
rm -f VPP_VERSION
rm -rf vppapi/* binapi/*

export_vppapi
generate_binapi

echo "$vpp_ver" > VPP_VERSION
