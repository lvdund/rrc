package ies

import "rrc/utils"

// SL-BetaOffsets-r16 ::= utils.INTEGER (0..31)
type SlBetaoffsetsR16 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
