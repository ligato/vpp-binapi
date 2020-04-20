#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"

vpp_img=${VPP_IMG:-'ligato/vpp-base:master'}
binapi_dir=binapis/${1:-'master'}

function log() {
	echo >&2 -ne "\e[35;1m$*\e[0;0m\n"
}

function err() {
	echo >&2 -ne "\e[31;1mERROR: \e[0;31m$*\e[0;0m\n"
}

function group() {
	[ -n "${GITHUB_ACTIONS-}" ] && echo >&2 -e "##[group]$1"
	echo >&2 -e "$2"
	[ -n "${GITHUB_ACTIONS-}" ] && echo >&2 -e "##[endgroup]"
}

function export_vppapi() {
	if [ ! -n "${NOPULL-}" ]; then
		log "# Pulling VPP image ${vpp_img}"
		docker pull "${vpp_img}"
	fi
	
	vpp_ver=$(docker run --rm "$vpp_img" dpkg-query -f '${Version}\n' -W vpp)
	log "# Exporting VPP API ($vpp_ver)"

	mkdir -p "${binapi_dir}"/.vppapi
	
	docker run --rm \
		-u $(id -u):$(id -g) \
		-v $(pwd)/${binapi_dir}/.vppapi:/vpp/api \
		"$vpp_img" \
		sh -c 'cp -r /usr/share/vpp/api/* /vpp/api'
	
	vppapi_files=$(find "${binapi_dir}"/.vppapi -type f -name '*.api.json')
	log "$(echo "$vppapi_files" | wc -l) VPP API files"
	#group "$(echo "$vppapi_files" | wc -l) exported VPP API files" "$vppapi_files"
}

function generate_binapi() {
	if [ ! -n "${NOINSTALL-}" ]; then
		log "# Installing binapi-generator"
		set -x
		go env
		env | sort
		mkdir -p bin
		export BINDIR="$(pwd)/bin"
		GOBIN="$BINDIR" go install -v git.fd.io/govpp.git/cmd/binapi-generator
		export PATH="$BINDIR":$PATH
		echo "binapi-generator installed:"
		binapi-generator version
		set +x
	fi
	
	bingen_ver=$(binapi-generator -version)
	log "# Generating binapi ($bingen_ver)"
	
	out=$(binapi-generator --debug --input-dir=${binapi_dir}/.vppapi --output-dir=${binapi_dir} 2>&1)
	if [ "$?" -ne 0 ]; then
		err "Generating binapi failed!"
		echo "$out"
		exit 3
	fi
	
	binapi_files=$(find "${binapi_dir}" -type f -name '*.ba.go')
	#group "$(echo "$binapi_files" | wc -l) generated binapi Go files" "$binapi_files"
	log "$(echo "$binapi_files" | wc -l) binapi Go files generated"
	
	#group "Generator output" "$out"
	
	warns=$(echo "$out" | grep -ci WARN || true)
	if [ "${warns}" -gt "0" ]; then
		echo >&2 -e "\e[33;1m$warns warnings detected in output!\e[0;0m"
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
#rm -f ${binapi_dir}/VPP_VERSION
#rm -rf ${binapi_dir}/vppapi/* ${binapi_dir}/binapi/*
rm -rf ${binapi_dir}

export_vppapi
generate_binapi

echo -n "$vpp_ver" > ${binapi_dir}/.vppapi/VPP_VERSION
