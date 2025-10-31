package ies

// CellsToAddModListUTRA-FDD ::= SEQUENCE OF CellsToAddModUTRA-FDD
// SIZE (1..maxCellMeas)
type CellstoaddmodlistutraFdd struct {
	Value []CellstoaddmodutraFdd `lb:1,ub:maxCellMeas`
}
