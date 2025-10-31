package ies

// ChoCandidateCellList-r17 ::= SEQUENCE OF ChoCandidateCell-r17
// SIZE (1..maxNrofCondCells-r16)
type ChocandidatecelllistR17 struct {
	Value []ChocandidatecellR17 `lb:1,ub:maxNrofCondCellsR16`
}
