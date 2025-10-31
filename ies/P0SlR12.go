package ies

import "rrc/utils"

// P0-SL-r12 ::= utils.INTEGER (-126..31)
type P0SlR12 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
