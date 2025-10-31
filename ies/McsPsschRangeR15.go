package ies

import "rrc/utils"

// MCS-PSSCH-Range-r15 ::= SEQUENCE
type McsPsschRangeR15 struct {
	MinmcsPsschR15 utils.INTEGER `lb:0,ub:31`
	MaxmcsPsschR15 utils.INTEGER `lb:0,ub:31`
}
