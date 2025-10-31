package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-dl-SchedulingOffset-PDSCH-TypeB ::= ENUMERATED
type PhyParametersxddDiffDlSchedulingoffsetPdschTypeb struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffDlSchedulingoffsetPdschTypebEnumeratedNothing = iota
	PhyParametersxddDiffDlSchedulingoffsetPdschTypebEnumeratedSupported
)
