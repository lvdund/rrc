package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-pdsch-RepSubframe-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubframeR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubframeR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubframeR15EnumeratedSupported
)
