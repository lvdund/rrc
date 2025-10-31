package ies

import "rrc/utils"

// CellGroupId ::= utils.INTEGER (0.. maxSecondaryCellGroups)
type Cellgroupid struct {
	Value utils.INTEGER `lb:0,ub:maxSecondaryCellGroups`
}
