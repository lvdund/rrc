package ies

// CellIndexList ::= SEQUENCE OF CellIndex
// SIZE (1..maxCellMeas)
type Cellindexlist struct {
	Value []Cellindex `lb:1,ub:maxCellMeas`
}
