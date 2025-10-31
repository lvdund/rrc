package ies

import "rrc/utils"

// UplinkHARQ-mode-r17 ::= BIT STRING (SIZE (32))
type UplinkharqModeR17 struct {
	Value utils.BITSTRING `lb:32,ub:32`
}
