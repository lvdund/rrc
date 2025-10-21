package ies

import "rrc/utils"

// GWUS-ResourceConfig-NB-r16 ::= SEQUENCE
type GwusResourceconfigNbR16 struct {
	ResourcepositionR16     utils.ENUMERATED
	NumgroupslistR16        *GwusNumgroupslistNbR16
	GroupsforservicelistR16 *GwusGroupsforservicelistNbR16
}
