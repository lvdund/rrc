package ies

// CellsToAddModListNR-r16 ::= SEQUENCE OF CellsToAddModNR-r16
// SIZE (1..maxCellMeas)
type CellstoaddmodlistnrR16 struct {
	Value []CellstoaddmodnrR16 `lb:1,ub:maxCellMeas`
}
