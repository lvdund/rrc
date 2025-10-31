package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-dl-SchedulingOffset-PDSCH-TypeA ::= ENUMERATED
type PhyParametersfrxDiffDlSchedulingoffsetPdschTypea struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffDlSchedulingoffsetPdschTypeaEnumeratedNothing = iota
	PhyParametersfrxDiffDlSchedulingoffsetPdschTypeaEnumeratedSupported
)
