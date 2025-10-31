package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-multipleCORESET ::= ENUMERATED
type PhyParametersfrxDiffMultiplecoreset struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMultiplecoresetEnumeratedNothing = iota
	PhyParametersfrxDiffMultiplecoresetEnumeratedSupported
)
