package ies

// UTRA-FDD-CellIndexList-r16 ::= SEQUENCE OF UTRA-FDD-CellIndex-r16
// SIZE (1..maxCellMeasUTRA-FDD-r16)
type UtraFddCellindexlistR16 struct {
	Value []UtraFddCellindexR16 `lb:1,ub:maxCellMeasUTRAFddR16`
}
