package ies

import "rrc/utils"

// SL-DestinationIdentity-r16 ::= BIT STRING (SIZE (24))
type SlDestinationidentityR16 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
