package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-PUSCH-FlexibleStartPRB-CE-ModeB-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CePuschFlexiblestartprbCeModebR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CePuschFlexiblestartprbCeModebR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CePuschFlexiblestartprbCeModebR15EnumeratedSupported
)
