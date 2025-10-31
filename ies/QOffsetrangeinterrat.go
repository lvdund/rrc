package ies

import "rrc/utils"

// Q-OffsetRangeInterRAT ::= utils.INTEGER (-15..15)
type QOffsetrangeinterrat struct {
	Value utils.INTEGER `lb:0,ub:15`
}
