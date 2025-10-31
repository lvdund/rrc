package ies

import "rrc/utils"

// SL-CBR-r14 ::= utils.INTEGER (0..100)
type SlCbrR14 struct {
	Value utils.INTEGER `lb:0,ub:100`
}
