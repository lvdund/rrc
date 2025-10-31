package ies

import "rrc/utils"

// RNTI-Value ::= utils.INTEGER (0..65535)
type RntiValue struct {
	Value utils.INTEGER `lb:0,ub:65535`
}
