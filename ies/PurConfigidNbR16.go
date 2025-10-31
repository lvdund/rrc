package ies

import "rrc/utils"

// PUR-ConfigID-NB-r16 ::= BIT STRING (SIZE (20))
type PurConfigidNbR16 struct {
	Value utils.BITSTRING `lb:20,ub:20`
}
