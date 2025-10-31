package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pucch-F0-2WithoutFH ::= ENUMERATED
type PhyParametersfrxDiffPucchF02withoutfh struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPucchF02withoutfhEnumeratedNothing = iota
	PhyParametersfrxDiffPucchF02withoutfhEnumeratedNotsupported
)
