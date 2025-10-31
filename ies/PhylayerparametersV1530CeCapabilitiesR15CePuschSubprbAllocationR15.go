package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-PUSCH-SubPRB-Allocation-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CePuschSubprbAllocationR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CePuschSubprbAllocationR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CePuschSubprbAllocationR15EnumeratedSupported
)
