package ies

import "rrc/utils"

// RB-SetGroup-r17 ::= SEQUENCE
type RbSetgroupR17 struct {
	ResourceavailabilityR17 *[]utils.INTEGER `lb:1,ub:maxNrofResourceAvailabilityPerCombinationR16`
	RbSetsR17               *[]utils.INTEGER `lb:1,ub:maxNrofRBSetsR17`
}
