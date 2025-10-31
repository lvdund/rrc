package ies

import "rrc/utils"

// SL-Reliability-r15 ::= utils.INTEGER (1..8)
type SlReliabilityR15 struct {
	Value utils.INTEGER `lb:0,ub:8`
}
