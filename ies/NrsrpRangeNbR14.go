package ies

import "rrc/utils"

// NRSRP-Range-NB-r14 ::= utils.INTEGER (0..113)
type NrsrpRangeNbR14 struct {
	Value utils.INTEGER `lb:0,ub:113`
}
