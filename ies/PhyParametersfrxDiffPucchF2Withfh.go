package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pucch-F2-WithFH ::= ENUMERATED
type PhyParametersfrxDiffPucchF2Withfh struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPucchF2WithfhEnumeratedNothing = iota
	PhyParametersfrxDiffPucchF2WithfhEnumeratedSupported
)
