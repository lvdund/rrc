package ies

// CellsToAddModListNR-r15 ::= SEQUENCE OF CellsToAddModNR-r15
// SIZE (1..maxCellMeas)
type CellstoaddmodlistnrR15 struct {
	Value []CellstoaddmodnrR15 `lb:1,ub:maxCellMeas`
}
