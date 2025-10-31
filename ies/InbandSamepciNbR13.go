package ies

import "rrc/utils"

// Inband-SamePCI-NB-r13 ::= SEQUENCE
type InbandSamepciNbR13 struct {
	EutraCrsSequenceinfoR13 utils.INTEGER `lb:0,ub:31`
}
