package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-pdsch-InLteControlRegionCE-ModeB-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610PdschInltecontrolregionceModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610PdschInltecontrolregionceModebR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610PdschInltecontrolregionceModebR16EnumeratedSupported
)
