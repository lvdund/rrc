package ies

import "rrc/utils"

// GWUS-ResourceConfig-r16 ::= SEQUENCE
type GwusResourceconfigR16 struct {
	ResourcemappingpatternR16 interface{}
	NumgroupslistR16          *GwusNumgroupslistR16
	GroupsforservicelistR16   *GwusGroupsforservicelistR16
}
