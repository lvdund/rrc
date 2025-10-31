package ies

import "rrc/utils"

// RSRQ-Range ::= utils.INTEGER (0..34)
type RsrqRange struct {
	Value utils.INTEGER `lb:0,ub:34`
}
