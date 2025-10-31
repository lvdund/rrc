package ies

import "rrc/utils"

// P0-PUCCH-Id ::= utils.INTEGER (1..8)
type P0PucchId struct {
	Value utils.INTEGER `lb:0,ub:8`
}
