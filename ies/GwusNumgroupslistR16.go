package ies

// GWUS-NumGroupsList-r16 ::= SEQUENCE OF GWUS-NumGroups-r16
// SIZE (1..maxGWUS-Resources-r16)
type GwusNumgroupslistR16 struct {
	Value []GwusNumgroupsR16 `lb:1,ub:maxGWUSResourcesR16`
}
