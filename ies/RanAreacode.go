package ies

import "rrc/utils"

// RAN-AreaCode ::= utils.INTEGER (0..255)
type RanAreacode struct {
	Value utils.INTEGER `lb:0,ub:255`
}
