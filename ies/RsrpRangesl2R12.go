package ies

import "rrc/utils"

// RSRP-RangeSL2-r12 ::= utils.INTEGER (0..7)
type RsrpRangesl2R12 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
