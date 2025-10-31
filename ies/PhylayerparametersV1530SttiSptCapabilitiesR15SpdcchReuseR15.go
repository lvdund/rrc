package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-spdcch-Reuse-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchReuseR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchReuseR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchReuseR15EnumeratedSupported
)
