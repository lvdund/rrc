package ies

// SCellGroupToAddModList-r15 ::= SEQUENCE OF SCellGroupToAddMod-r15
// SIZE (1..maxSCellGroups-r15)
type ScellgrouptoaddmodlistR15 struct {
	Value []ScellgrouptoaddmodR15 `lb:1,ub:maxSCellGroupsR15`
}
