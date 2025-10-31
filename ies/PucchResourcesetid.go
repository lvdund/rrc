package ies

import "rrc/utils"

// PUCCH-ResourceSetId ::= utils.INTEGER (0..maxNrofPUCCH-ResourceSets-1)
type PucchResourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUCCHResourcesets1`
}
