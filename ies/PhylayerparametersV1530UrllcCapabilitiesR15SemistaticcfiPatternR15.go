package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-semiStaticCFI-Pattern-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiPatternR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiPatternR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiPatternR15EnumeratedSupported
)
