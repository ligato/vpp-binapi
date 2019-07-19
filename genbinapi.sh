#!/bin/bash
set -eu

cd $(dirname "${BASH_SOURCE[0]}")

VPP_IMG=${VPP_IMG:-ligato/vpp-base:master}

function generate_binapi() {
	local vppimg=$1

	echo "=> Pulling VPP image: $vppimg"
	docker pull "$vppimg"

	vppver=$(docker run --rm "$vppimg" cat /vpp/version)

	echo "Generating binapi for VPP $vppver"

	rm -rf vppapi/* binapi/*

	docker run --rm -v $(pwd)/vppapi:/vppapi -w /vppapi -u $(id -u):$(id -g) "$vppimg" \
		sh -c "cp -r /usr/share/vpp/api/* /vppapi"

	binapi-generator --input-dir=vppapi --output-dir=binapi --debug

	echo "$vppver" > ./VPP_VERSION
}

echo "=> Installing binapi generator"
go install -v git.fd.io/govpp.git/cmd/binapi-generator

generate_binapi $VPP_IMG
