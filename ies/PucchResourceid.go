package ies

import "rrc/utils"

// PUCCH-ResourceId ::= utils.INTEGER (0..maxNrofPUCCH-Resources-1)
type PucchResourceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUCCHResources1`
}
