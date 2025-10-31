package ies

import "rrc/utils"

// Standalone-NB-r13 ::= SEQUENCE
type StandaloneNbR13 struct {
	Spare utils.BITSTRING `lb:5,ub:5`
}
