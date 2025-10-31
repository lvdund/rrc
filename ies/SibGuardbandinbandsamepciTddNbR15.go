package ies

import "rrc/utils"

// SIB-GuardbandInbandSamePCI-TDD-NB-r15 ::= SEQUENCE
type SibGuardbandinbandsamepciTddNbR15 struct {
	Spare utils.BITSTRING `lb:1,ub:1`
}
