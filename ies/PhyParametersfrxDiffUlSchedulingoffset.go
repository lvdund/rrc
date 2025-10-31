package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-ul-SchedulingOffset ::= ENUMERATED
type PhyParametersfrxDiffUlSchedulingoffset struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffUlSchedulingoffsetEnumeratedNothing = iota
	PhyParametersfrxDiffUlSchedulingoffsetEnumeratedSupported
)
