package ies

import "rrc/utils"

// Phy-ParametersCommon-type2-HARQ-ACK-Codebook-r16 ::= ENUMERATED
type PhyParameterscommonType2HarqAckCodebookR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonType2HarqAckCodebookR16EnumeratedNothing = iota
	PhyParameterscommonType2HarqAckCodebookR16EnumeratedSupported
)
