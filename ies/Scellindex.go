package ies

import "rrc/utils"

// SCellIndex ::= utils.INTEGER (1..31)
type Scellindex struct {
	Value utils.INTEGER `lb:0,ub:31`
}
