package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-twoDifferentTPC-Loop-PUCCH ::= ENUMERATED
type PhyParametersxddDiffTwodifferenttpcLoopPucch struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffTwodifferenttpcLoopPucchEnumeratedNothing = iota
	PhyParametersxddDiffTwodifferenttpcLoopPucchEnumeratedSupported
)
