package ies

import "rrc/utils"

// NSAG-ID-r17 ::= BIT STRING (SIZE (8))
type NsagIdR17 struct {
	Value utils.BITSTRING `lb:8,ub:8`
}
