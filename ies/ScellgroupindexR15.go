package ies

import "rrc/utils"

// SCellGroupIndex-r15 ::= utils.INTEGER (1..maxSCellGroups-r15)
type ScellgroupindexR15 struct {
	Value utils.INTEGER `lb:0,ub:maxSCellGroupsR15`
}
