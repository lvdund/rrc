package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-tpc-PUCCH-RNTI ::= ENUMERATED
type PhyParametersfrxDiffTpcPucchRnti struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTpcPucchRntiEnumeratedNothing = iota
	PhyParametersfrxDiffTpcPucchRntiEnumeratedSupported
)
