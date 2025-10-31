package ies

import "rrc/utils"

// MBSFN-RSRQ-Range-r12 ::= utils.INTEGER (0..31)
type MbsfnRsrqRangeR12 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
