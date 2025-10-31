package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-twoDifferentTPC-Loop-PUSCH ::= ENUMERATED
type PhyParametersfrxDiffTwodifferenttpcLoopPusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTwodifferenttpcLoopPuschEnumeratedNothing = iota
	PhyParametersfrxDiffTwodifferenttpcLoopPuschEnumeratedSupported
)
