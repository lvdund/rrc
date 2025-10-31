package ies

import "rrc/utils"

// RSRP-RangeSL3-r12 ::= utils.INTEGER (0..11)
type RsrpRangesl3R12 struct {
	Value utils.INTEGER `lb:0,ub:11`
}
