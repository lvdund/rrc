package ies

import "rrc/utils"

// RSRP-RangeSL4-r13 ::= utils.INTEGER (0..49)
type RsrpRangesl4R13 struct {
	Value utils.INTEGER `lb:0,ub:49`
}
