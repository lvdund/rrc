package ies

import "rrc/utils"

// PUR-ConfigID-r16 ::= BIT STRING (SIZE (20))
type PurConfigidR16 struct {
	Value utils.BITSTRING
}
