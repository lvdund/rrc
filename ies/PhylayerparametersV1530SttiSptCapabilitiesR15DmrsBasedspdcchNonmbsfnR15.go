package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-dmrs-BasedSPDCCH-nonMBSFN-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchNonmbsfnR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchNonmbsfnR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchNonmbsfnR15EnumeratedSupported
)
