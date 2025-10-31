package ies

import "rrc/utils"

// NumberOfCarriers ::= utils.INTEGER (1..16)
type Numberofcarriers struct {
	Value utils.INTEGER `lb:0,ub:16`
}
