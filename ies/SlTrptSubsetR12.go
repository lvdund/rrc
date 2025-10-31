package ies

import "rrc/utils"

// SL-TRPT-Subset-r12 ::= BIT STRING (SIZE (3..5))
type SlTrptSubsetR12 struct {
	Value utils.BITSTRING `lb:3,ub:5`
}
