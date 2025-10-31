package ies

import "rrc/utils"

// AMF-Identifier-r15 ::= BIT STRING (SIZE (24))
type AmfIdentifierR15 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
