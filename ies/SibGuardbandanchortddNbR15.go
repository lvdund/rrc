package ies

import "rrc/utils"

// SIB-GuardbandAnchorTDD-NB-r15 ::= SEQUENCE
type SibGuardbandanchortddNbR15 struct {
	Spare utils.BITSTRING `lb:1,ub:1`
}
