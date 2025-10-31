package ies

import "rrc/utils"

// RSRP-RangeNR-r15 ::= utils.INTEGER (0..127)
type RsrpRangenrR15 struct {
	Value utils.INTEGER `lb:0,ub:127`
}
