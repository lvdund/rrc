package ies

// GWUS-ResourceConfig-r16 ::= SEQUENCE
type GwusResourceconfigR16 struct {
	ResourcemappingpatternR16 GwusResourceconfigR16ResourcemappingpatternR16
	NumgroupslistR16          *GwusNumgroupslistR16
	GroupsforservicelistR16   *GwusGroupsforservicelistR16
}
