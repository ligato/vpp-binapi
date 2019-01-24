VPP_VERSION ?= $(shell cd vpp && git describe --always --tags --dirty)

generate-binapi:
	@echo "=> generating binapi ${VPP_VERSION}"
	#bash ./generate.sh
	@echo ${VPP_VERSION} > VPP_VERSION

apply-patches:
	@echo "=> applying patches"
	find patches -maxdepth 1 -type f -name '*.patch' -exec patch --no-backup-if-mismatch -p1 -i {} \;


.PHONY: generate-binapi apply-patches

