package ies

import "rrc/utils"

// Inband-DifferentPCI-TDD-NB-r15-sib-InbandLocation-r15 ::= ENUMERATED
type InbandDifferentpciTddNbR15SibInbandlocationR15 struct {
	Value utils.ENUMERATED
}

const (
	InbandDifferentpciTddNbR15SibInbandlocationR15EnumeratedNothing = iota
	InbandDifferentpciTddNbR15SibInbandlocationR15EnumeratedLower
	InbandDifferentpciTddNbR15SibInbandlocationR15EnumeratedHigher
)
