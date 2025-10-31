package ies

import "rrc/utils"

// SINR-RangeEUTRA ::= utils.INTEGER (0..127)
type SinrRangeeutra struct {
	Value utils.INTEGER `lb:0,ub:127`
}
