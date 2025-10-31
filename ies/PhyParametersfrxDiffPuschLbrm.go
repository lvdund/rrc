package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pusch-LBRM ::= ENUMERATED
type PhyParametersfrxDiffPuschLbrm struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPuschLbrmEnumeratedNothing = iota
	PhyParametersfrxDiffPuschLbrmEnumeratedSupported
)
