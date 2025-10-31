package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-semiStaticCFI-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiR15EnumeratedSupported
)
