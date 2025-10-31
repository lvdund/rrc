package ies

import "rrc/utils"

// ShortI-RNTI-r15 ::= BIT STRING (SIZE (24))
type ShortiRntiR15 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
