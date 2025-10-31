package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-ul-SchedulingOffset ::= ENUMERATED
type PhyParametersxddDiffUlSchedulingoffset struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffUlSchedulingoffsetEnumeratedNothing = iota
	PhyParametersxddDiffUlSchedulingoffsetEnumeratedSupported
)
