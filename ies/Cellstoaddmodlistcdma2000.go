package ies

// CellsToAddModListCDMA2000 ::= SEQUENCE OF CellsToAddModCDMA2000
// SIZE (1..maxCellMeas)
type Cellstoaddmodlistcdma2000 struct {
	Value []Cellstoaddmodcdma2000 `lb:1,ub:maxCellMeas`
}
