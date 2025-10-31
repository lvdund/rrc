package ies

import "rrc/utils"

// RSRP-RangeSL-r12 ::= utils.INTEGER (0..13)
type RsrpRangeslR12 struct {
	Value utils.INTEGER `lb:0,ub:13`
}
