package ies

import "rrc/utils"

// Inband-SamePCI-TDD-NB-r15 ::= SEQUENCE
type InbandSamepciTddNbR15 struct {
	EutraCrsSequenceinfoR15 utils.INTEGER `lb:0,ub:31`
	SibInbandlocationR15    InbandSamepciTddNbR15SibInbandlocationR15
}
