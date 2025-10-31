package ies

import "rrc/utils"

// AMF-Identifier ::= BIT STRING (SIZE (24))
type AmfIdentifier struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
