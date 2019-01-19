# vpp-binapi
**vpp-binapi** repository contains generated Go code for [VPP binary API](https://wiki.fd.io/view/VPP/The_VPP_API) generated with [GoVPP](https://github.com/FDio/govpp)'s binapi-generator.

## Versioning

The version of VPP that was used for generating the code is defined in [`VPP_VERSION`](VPP_VERSION) file.

Compatibility with VPP instance is verified at runtime by comparing version hash of all generated binapi messages with version hashes of connected VPP. This means that any change to a specific message will result in runtime incompatibility.
