package ies

import "rrc/utils"

// RSRQ-RangeEUTRA-r16 ::= utils.INTEGER (-30..46)
type RsrqRangeeutraR16 struct {
	Value utils.INTEGER `lb:0,ub:46`
}
