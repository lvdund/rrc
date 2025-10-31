package ies

import "rrc/utils"

// PUCCH-ResourceGroupId-r16 ::= utils.INTEGER (0..maxNrofPUCCH-ResourceGroups-1-r16)
type PucchResourcegroupidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUCCHResourcegroups1R16`
}
