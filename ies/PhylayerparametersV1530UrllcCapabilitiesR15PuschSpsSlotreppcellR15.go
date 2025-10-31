package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-pusch-SPS-SlotRepPCell-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppcellR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppcellR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppcellR15EnumeratedSupported
)
