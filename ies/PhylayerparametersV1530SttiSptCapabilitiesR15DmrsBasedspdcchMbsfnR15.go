package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-dmrs-BasedSPDCCH-MBSFN-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchMbsfnR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchMbsfnR15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15DmrsBasedspdcchMbsfnR15EnumeratedSupported
)
