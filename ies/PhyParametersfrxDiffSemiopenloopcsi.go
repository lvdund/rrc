package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-semiOpenLoopCSI ::= ENUMERATED
type PhyParametersfrxDiffSemiopenloopcsi struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSemiopenloopcsiEnumeratedNothing = iota
	PhyParametersfrxDiffSemiopenloopcsiEnumeratedSupported
)
