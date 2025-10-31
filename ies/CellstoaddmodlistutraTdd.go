package ies

// CellsToAddModListUTRA-TDD ::= SEQUENCE OF CellsToAddModUTRA-TDD
// SIZE (1..maxCellMeas)
type CellstoaddmodlistutraTdd struct {
	Value []CellstoaddmodutraTdd `lb:1,ub:maxCellMeas`
}
