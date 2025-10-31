package ies

import "rrc/utils"

// SL-CBR-r16 ::= utils.INTEGER (0..100)
type SlCbrR16 struct {
	Value utils.INTEGER `lb:0,ub:100`
}
