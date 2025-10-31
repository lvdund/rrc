package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-pusch-SPS-SubframeRepPSCell-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppscellR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppscellR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppscellR15EnumeratedSupported
)
