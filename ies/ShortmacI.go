package ies

import "rrc/utils"

// ShortMAC-I ::= BIT STRING (SIZE (16))
type ShortmacI struct {
	Value utils.BITSTRING `lb:16,ub:16`
}
