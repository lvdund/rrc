package ies

// MBS-NeighbourCellList-r17 ::= SEQUENCE OF MBS-NeighbourCell-r17
// SIZE (0..maxNeighCellMBS-r17)
type MbsNeighbourcelllistR17 struct {
	Value []MbsNeighbourcellR17 `lb:0,ub:maxNeighCellMBSR17`
}
