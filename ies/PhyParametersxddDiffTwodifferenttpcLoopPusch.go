package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-twoDifferentTPC-Loop-PUSCH ::= ENUMERATED
type PhyParametersxddDiffTwodifferenttpcLoopPusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffTwodifferenttpcLoopPuschEnumeratedNothing = iota
	PhyParametersxddDiffTwodifferenttpcLoopPuschEnumeratedSupported
)
