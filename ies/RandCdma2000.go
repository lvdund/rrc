package ies

import "rrc/utils"

// RAND-CDMA2000 ::= BIT STRING (SIZE (32))
type RandCdma2000 struct {
	Value utils.BITSTRING `lb:32,ub:32`
}
