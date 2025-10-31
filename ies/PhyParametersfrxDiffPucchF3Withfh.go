package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pucch-F3-WithFH ::= ENUMERATED
type PhyParametersfrxDiffPucchF3Withfh struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPucchF3WithfhEnumeratedNothing = iota
	PhyParametersfrxDiffPucchF3WithfhEnumeratedSupported
)
