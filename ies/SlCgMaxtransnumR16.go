package ies

import "rrc/utils"

// SL-CG-MaxTransNum-r16 ::= SEQUENCE
type SlCgMaxtransnumR16 struct {
	SlPriorityR16    utils.INTEGER `lb:0,ub:8`
	SlMaxtransnumR16 utils.INTEGER `lb:0,ub:32`
}
