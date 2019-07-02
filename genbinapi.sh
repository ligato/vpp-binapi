#!/bin/bash

set -exuo pipefail

cd $(dirname "${BASH_SOURCE[0]}")

VPP_IMG=${VPP_IMG:-ligato/vpp-base:master}

go install -v git.fd.io/govpp.git/cmd/binapi-generator

docker run --rm "$VPP_IMG" cat /vpp/version > /tmp/VPP_VERSION

echo "Generating binapi for VPP $(cat /tmp/VPP_VERSION)"

rm -rf vppapi/*
docker run --rm -v $(pwd)/vppapi:/vppapi -w /vppapi -u $(id -u):$(id -g) "$VPP_IMG" \
	sh -c "cp -r /usr/share/vpp/api/* /vppapi"

rm -rf binapi/*
binapi-generator --input-dir=vppapi --output-dir=binapi --include-services --debug

mv /tmp/VPP_VERSION .
