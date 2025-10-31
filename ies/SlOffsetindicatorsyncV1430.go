package ies

import "rrc/utils"

// SL-OffsetIndicatorSync-v1430 ::= utils.INTEGER (40..159)
type SlOffsetindicatorsyncV1430 struct {
	Value utils.INTEGER `lb:0,ub:159`
}
