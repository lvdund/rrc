package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-dl-SchedulingOffset-PDSCH-TypeB ::= ENUMERATED
type PhyParametersfrxDiffDlSchedulingoffsetPdschTypeb struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffDlSchedulingoffsetPdschTypebEnumeratedNothing = iota
	PhyParametersfrxDiffDlSchedulingoffsetPdschTypebEnumeratedSupported
)
