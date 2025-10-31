package ies

import "rrc/utils"

// SRS-RSRP-Range-r16 ::= utils.INTEGER (0..98)
type SrsRsrpRangeR16 struct {
	Value utils.INTEGER `lb:0,ub:98`
}
