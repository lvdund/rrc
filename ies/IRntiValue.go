package ies

import "rrc/utils"

// I-RNTI-Value ::= BIT STRING (SIZE (40))
type IRntiValue struct {
	Value utils.BITSTRING `lb:40,ub:40`
}
