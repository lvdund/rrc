package ies

import "rrc/utils"

// ShortI-RNTI-Value ::= BIT STRING (SIZE (24))
type ShortiRntiValue struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
