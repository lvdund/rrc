package ies

import "rrc/utils"

// C-RNTI ::= BIT STRING (SIZE (16))
type CRnti struct {
	Value utils.BITSTRING
}
