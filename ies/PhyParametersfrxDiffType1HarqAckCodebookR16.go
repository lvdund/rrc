package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-type1-HARQ-ACK-Codebook-r16 ::= ENUMERATED
type PhyParametersfrxDiffType1HarqAckCodebookR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffType1HarqAckCodebookR16EnumeratedNothing = iota
	PhyParametersfrxDiffType1HarqAckCodebookR16EnumeratedSupported
)
