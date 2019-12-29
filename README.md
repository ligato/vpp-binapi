<h1 align="center">vpp-binapi</h1>

<p align="center">
  <a href="https://github.com/ligato/vpp-binapi/actions?query=workflow%3A%22Generate+binapi%22"><img src="https://github.com/ligato/vpp-binapi/workflows/Generate%20binapi/badge.svg" alt="Workflow"></a>
</p>

<p align="center">The <b>vpp-binapi</b> generates Go code for the <a href="https://wiki.fd.io/view/VPP/The_VPP_API">VPP binary API</a> using <a href="https://github.com/FDio/govpp">GoVPP</a>'s binapi-generator. </p>

---

## Versioning

The version of VPP that was used for generating the code is defined in [`VPP_VERSION`](VPP_VERSION) file.

Compatibility with VPP instance is verified at runtime by comparing version hash of all generated binapi messages with version hashes of connected VPP. This means that any change to a specific message will result in runtime incompatibility.
