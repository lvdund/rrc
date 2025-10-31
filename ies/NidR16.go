package ies

import "rrc/utils"

// NID-r16 ::= BIT STRING (SIZE (44))
type NidR16 struct {
	Value utils.BITSTRING `lb:44,ub:44`
}
