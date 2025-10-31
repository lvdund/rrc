package ies

import "rrc/utils"

// SL-RSRP-Range-r16 ::= utils.INTEGER (0..13)
type SlRsrpRangeR16 struct {
	Value utils.INTEGER `lb:0,ub:13`
}
