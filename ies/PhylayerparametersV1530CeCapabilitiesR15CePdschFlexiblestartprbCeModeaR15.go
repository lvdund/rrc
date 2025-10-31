package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-PDSCH-FlexibleStartPRB-CE-ModeA-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CePdschFlexiblestartprbCeModeaR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CePdschFlexiblestartprbCeModeaR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CePdschFlexiblestartprbCeModeaR15EnumeratedSupported
)
