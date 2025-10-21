package ies

import "rrc/utils"

// CSG-Identity ::= BIT STRING (SIZE (27))
type CsgIdentity struct {
	Value utils.BITSTRING
}
