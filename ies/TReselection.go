package ies

import "rrc/utils"

// T-Reselection ::= utils.INTEGER (0..7)
type TReselection struct {
	Value utils.INTEGER `lb:0,ub:7`
}
