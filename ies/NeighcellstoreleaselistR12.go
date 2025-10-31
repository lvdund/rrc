package ies

// NeighCellsToReleaseList-r12 ::= SEQUENCE OF PhysCellId
// SIZE (1..maxNeighCell-r12)
type NeighcellstoreleaselistR12 struct {
	Value []Physcellid `lb:1,ub:maxNeighCellR12`
}
