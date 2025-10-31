package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-spdcch-differentRS-types-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchDifferentrsTypesR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchDifferentrsTypesR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15SpdcchDifferentrsTypesR15EnumeratedSupported
)
