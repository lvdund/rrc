package ies

import "rrc/utils"

// SL-Priority-r13 ::= utils.INTEGER (1..8)
type SlPriorityR13 struct {
	Value utils.INTEGER `lb:0,ub:8`
}
