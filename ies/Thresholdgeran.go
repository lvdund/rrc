package ies

import "rrc/utils"

// ThresholdGERAN ::= utils.INTEGER (0..63)
type Thresholdgeran struct {
	Value utils.INTEGER `lb:0,ub:63`
}
