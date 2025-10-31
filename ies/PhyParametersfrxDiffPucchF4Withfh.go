package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pucch-F4-WithFH ::= ENUMERATED
type PhyParametersfrxDiffPucchF4Withfh struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPucchF4WithfhEnumeratedNothing = iota
	PhyParametersfrxDiffPucchF4WithfhEnumeratedSupported
)
