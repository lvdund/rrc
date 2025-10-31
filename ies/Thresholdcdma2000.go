package ies

import "rrc/utils"

// ThresholdCDMA2000 ::= utils.INTEGER (0..63)
type Thresholdcdma2000 struct {
	Value utils.INTEGER `lb:0,ub:63`
}
