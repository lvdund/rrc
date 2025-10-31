package ies

import "rrc/utils"

// Hysteresis ::= utils.INTEGER (0..30)
type Hysteresis struct {
	Value utils.INTEGER `lb:0,ub:30`
}
