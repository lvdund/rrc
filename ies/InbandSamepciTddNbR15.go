package ies

import "rrc/utils"

// Inband-SamePCI-TDD-NB-r15 ::= SEQUENCE
type InbandSamepciTddNbR15 struct {
	EutraCrsSequenceinfoR15 utils.INTEGER
	SibInbandlocationR15    utils.ENUMERATED
}
