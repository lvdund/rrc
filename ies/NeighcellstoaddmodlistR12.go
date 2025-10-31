package ies

// NeighCellsToAddModList-r12 ::= SEQUENCE OF NeighCellsInfo-r12
// SIZE (1..maxNeighCell-r12)
type NeighcellstoaddmodlistR12 struct {
	Value []NeighcellsinfoR12 `lb:1,ub:maxNeighCellR12`
}
