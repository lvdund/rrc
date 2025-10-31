package ies

import "rrc/utils"

// RSRP-RangeEUTRA ::= utils.INTEGER (0..97)
type RsrpRangeeutra struct {
	Value utils.INTEGER `lb:0,ub:97`
}
