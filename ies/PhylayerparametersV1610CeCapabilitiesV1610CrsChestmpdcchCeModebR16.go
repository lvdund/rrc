package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-crs-ChEstMPDCCH-CE-ModeB-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModebR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModebR16EnumeratedSupported
)
