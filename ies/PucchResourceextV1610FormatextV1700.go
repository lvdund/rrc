package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-formatExt-v1700 ::= SEQUENCE
type PucchResourceextV1610FormatextV1700 struct {
	NrofprbsR17 utils.INTEGER `lb:0,ub:16`
}
