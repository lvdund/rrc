package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-crs-ChEstMPDCCH-CE-ModeA-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModeaR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchCeModeaR16EnumeratedSupported
)
