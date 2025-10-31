package ies

import "rrc/utils"

// Inband-SamePCI-TDD-NB-r15-sib-InbandLocation-r15 ::= ENUMERATED
type InbandSamepciTddNbR15SibInbandlocationR15 struct {
	Value utils.ENUMERATED
}

const (
	InbandSamepciTddNbR15SibInbandlocationR15EnumeratedNothing = iota
	InbandSamepciTddNbR15SibInbandlocationR15EnumeratedLower
	InbandSamepciTddNbR15SibInbandlocationR15EnumeratedHigher
)
