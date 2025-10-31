package ies

import "rrc/utils"

// RAN-AreaCode-r15 ::= utils.INTEGER (0..255)
type RanAreacodeR15 struct {
	Value utils.INTEGER `lb:0,ub:255`
}
