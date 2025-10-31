package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-pdsch-RepSubslot-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubslotR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubslotR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubslotR15EnumeratedSupported
)
