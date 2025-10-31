package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-twoDifferentTPC-Loop-PUCCH ::= ENUMERATED
type PhyParametersfrxDiffTwodifferenttpcLoopPucch struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTwodifferenttpcLoopPucchEnumeratedNothing = iota
	PhyParametersfrxDiffTwodifferenttpcLoopPucchEnumeratedSupported
)
