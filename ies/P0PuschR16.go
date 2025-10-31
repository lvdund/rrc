package ies

import "rrc/utils"

// P0-PUSCH-r16 ::= utils.INTEGER (-16..15)
type P0PuschR16 struct {
	Value utils.INTEGER `lb:0,ub:15`
}
