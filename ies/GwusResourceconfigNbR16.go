package ies

// GWUS-ResourceConfig-NB-r16 ::= SEQUENCE
type GwusResourceconfigNbR16 struct {
	ResourcepositionR16     GwusResourceconfigNbR16ResourcepositionR16
	NumgroupslistR16        *GwusNumgroupslistNbR16
	GroupsforservicelistR16 *GwusGroupsforservicelistNbR16
}
