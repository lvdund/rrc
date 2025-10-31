package ies

import "rrc/utils"

// SL-PQFI-r16 ::= utils.INTEGER (1..64)
type SlPqfiR16 struct {
	Value utils.INTEGER `lb:0,ub:64`
}
