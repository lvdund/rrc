package ies

import "rrc/utils"

// SL-OffsetIndicatorSync-r14 ::= utils.INTEGER (0..159)
type SlOffsetindicatorsyncR14 struct {
	Value utils.INTEGER `lb:0,ub:159`
}
