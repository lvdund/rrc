package ies

import "rrc/utils"

// RSRQ-RangeEUTRA ::= utils.INTEGER (0..34)
type RsrqRangeeutra struct {
	Value utils.INTEGER `lb:0,ub:34`
}
