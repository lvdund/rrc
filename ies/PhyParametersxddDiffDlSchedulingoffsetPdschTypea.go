package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-dl-SchedulingOffset-PDSCH-TypeA ::= ENUMERATED
type PhyParametersxddDiffDlSchedulingoffsetPdschTypea struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffDlSchedulingoffsetPdschTypeaEnumeratedNothing = iota
	PhyParametersxddDiffDlSchedulingoffsetPdschTypeaEnumeratedSupported
)
