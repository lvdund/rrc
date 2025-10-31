package ies

import "rrc/utils"

// RSRQ-Range-r13 ::= utils.INTEGER (-30..46)
type RsrqRangeR13 struct {
	Value utils.INTEGER `lb:0,ub:46`
}
