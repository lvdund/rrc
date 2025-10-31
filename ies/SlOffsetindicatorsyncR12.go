package ies

import "rrc/utils"

// SL-OffsetIndicatorSync-r12 ::= utils.INTEGER (0..39)
type SlOffsetindicatorsyncR12 struct {
	Value utils.INTEGER `lb:0,ub:39`
}
