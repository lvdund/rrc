package ies

import "rrc/utils"

// AvailabilityCombinationRB-Groups-r17 ::= SEQUENCE
type AvailabilitycombinationrbGroupsR17 struct {
	AvailabilitycombinationidR17 AvailabilitycombinationidR16
	RbSetgroupsR17               *[]RbSetgroupR17 `lb:1,ub:maxNrofRBSetgroupsR17`
	ResourceavailabilityR17      *[]utils.INTEGER `lb:1,ub:maxNrofResourceAvailabilityPerCombinationR16`
}
