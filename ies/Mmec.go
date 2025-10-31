package ies

import "rrc/utils"

// MMEC ::= BIT STRING (SIZE (8))
type Mmec struct {
	Value utils.BITSTRING `lb:8,ub:8`
}
