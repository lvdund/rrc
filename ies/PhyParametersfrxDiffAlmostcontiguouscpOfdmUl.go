package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-almostContiguousCP-OFDM-UL ::= ENUMERATED
type PhyParametersfrxDiffAlmostcontiguouscpOfdmUl struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffAlmostcontiguouscpOfdmUlEnumeratedNothing = iota
	PhyParametersfrxDiffAlmostcontiguouscpOfdmUlEnumeratedSupported
)
