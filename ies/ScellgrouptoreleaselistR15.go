package ies

// SCellGroupToReleaseList-r15 ::= SEQUENCE OF SCellGroupIndex-r15
// SIZE (1..maxSCellGroups-r15)
type ScellgrouptoreleaselistR15 struct {
	Value []ScellgroupindexR15 `lb:1,ub:maxSCellGroupsR15`
}
