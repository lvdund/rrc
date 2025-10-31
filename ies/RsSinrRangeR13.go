package ies

import "rrc/utils"

// RS-SINR-Range-r13 ::= utils.INTEGER (0..127)
type RsSinrRangeR13 struct {
	Value utils.INTEGER `lb:0,ub:127`
}
