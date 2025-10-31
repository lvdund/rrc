package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-crs-ChEstMPDCCH-ReciprocityTDD-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchReciprocitytddR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchReciprocitytddR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610CrsChestmpdcchReciprocitytddR16EnumeratedSupported
)
