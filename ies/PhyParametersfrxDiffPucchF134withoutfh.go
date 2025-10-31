package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pucch-F1-3-4WithoutFH ::= ENUMERATED
type PhyParametersfrxDiffPucchF134withoutfh struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPucchF134withoutfhEnumeratedNothing = iota
	PhyParametersfrxDiffPucchF134withoutfhEnumeratedNotsupported
)
