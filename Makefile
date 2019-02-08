VPP_REPO_URL ?= https://github.com/FDio/vpp
VPP_VERSION ?= $(shell cd vpp && git describe --always --tags)

VPPAPIGEN=vpp/src/tools/vppapigen/vppapigen
BINAPIGEN=binapi-generator

VPPVER_FILE=VPP_VERSION
BINAPI_DIR=binapi
VPPAPI_DIR=vppapi

default: cleanup prepare generate

cleanup:
	@echo "=> cleaning up.."
	@rm -rf $(VPPVER_FILE) $(VPPAPI_DIR)/* $(BINAPI_DIR)/*

prepare: prepare-api generate-json

prepare-api:
	@echo "=> collecting VPP API files from VPP ${VPP_VERSION}"
	@rm -rf $(VPPAPI_DIR)/*.api
	@find vpp/src -type f -name '*.api' -printf "/bin/cp -v %p $(VPPAPI_DIR)/%f\n" | xargs -0 sh -c
	@echo "=> `find $(VPPAPI_DIR) -type f -name '*.api' | wc -l` files collected"

generate-json:
	@echo "=> generating JSON files.."
	@rm -rf $(VPPAPI_DIR)/*.api.json
	@find $(VPPAPI_DIR) -type f -name '*.api' -printf "echo %p \
	  && $(VPPAPIGEN) --includedir vpp/src --input %p --output $(VPPAPI_DIR)/%f.json JSON\n" | xargs -0 sh -c
	@echo "=> JSON files generated"

generate: generate-binapi patch-binapi

generate-binapi:
	@echo "=> generating binapi for VPP ${VPP_VERSION}"
	@rm -rf $(BINAPI_DIR)/* $(VPPVER_FILE)
	@binapi-generator --input-dir=$(VPPAPI_DIR) --output-dir=$(BINAPI_DIR) --include-apiver
	@echo ${VPP_VERSION} > $(VPPVER_FILE)
	@echo "=> binapi generated"

patch-binapi:
	@echo "=> patching binapi.."
	@find patches -maxdepth 1 -type f -name '*.patch' -exec patch --no-backup-if-mismatch -p1 -i {} \;
	@echo "=> binapi patched"

.PHONY: default cleanup \
 	prepare prepare-api generate-json \
 	generate generate-binapi patch-binapi \
