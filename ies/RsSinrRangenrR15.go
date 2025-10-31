package ies

import "rrc/utils"

// RS-SINR-RangeNR-r15 ::= utils.INTEGER (0..127)
type RsSinrRangenrR15 struct {
	Value utils.INTEGER `lb:0,ub:127`
}
