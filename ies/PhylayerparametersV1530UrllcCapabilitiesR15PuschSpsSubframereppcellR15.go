package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-pusch-SPS-SubframeRepPCell-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppcellR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppcellR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppcellR15EnumeratedSupported
)
