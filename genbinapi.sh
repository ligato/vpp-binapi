#!/bin/bash

set -euo pipefail

cd $(dirname "${BASH_SOURCE[0]}")

VPP_IMG=${VPP_IMG:-ligato/vpp-base:master}

echo "=> Pulling image: $VPP_IMG"
docker pull "$VPP_IMG"

echo "=> Installing GoVPP binapi generator"
go install -v git.fd.io/govpp.git/cmd/binapi-generator

docker run --rm "$VPP_IMG" cat /vpp/version > /tmp/VPP_VERSION

echo "Generating binapi code for VPP $(cat /tmp/VPP_VERSION)"

rm -rf vppapi/*
rm -rf binapi/*

docker run --rm -v $(pwd)/vppapi:/vppapi -w /vppapi -u $(id -u):$(id -g) "$VPP_IMG" \
	sh -c "cp -r /usr/share/vpp/api/* /vppapi"

binapi-generator --input-dir=vppapi --output-dir=binapi --debug

mv /tmp/VPP_VERSION .
