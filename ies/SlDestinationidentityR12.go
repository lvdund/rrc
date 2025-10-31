package ies

import "rrc/utils"

// SL-DestinationIdentity-r12 ::= BIT STRING (SIZE (24))
type SlDestinationidentityR12 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
