package ies

import "rrc/utils"

// NRSRQ-Range-NB-r14 ::= utils.INTEGER (-30..46)
type NrsrqRangeNbR14 struct {
	Value utils.INTEGER `lb:0,ub:46`
}
