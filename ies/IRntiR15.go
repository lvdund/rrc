package ies

import "rrc/utils"

// I-RNTI-r15 ::= BIT STRING (SIZE (40))
type IRntiR15 struct {
	Value utils.BITSTRING `lb:40,ub:40`
}
