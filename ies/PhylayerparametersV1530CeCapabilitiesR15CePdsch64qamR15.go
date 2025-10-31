package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-PDSCH-64QAM-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CePdsch64qamR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CePdsch64qamR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CePdsch64qamR15EnumeratedSupported
)
