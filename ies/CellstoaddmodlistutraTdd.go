package ies

import "rrc/utils"

// CellsToAddModListUTRA-TDD ::= SEQUENCE OF CellsToAddModUTRA-TDD
// SIZE (1..maxCellMeas)
type CellstoaddmodlistutraTdd struct {
	Value []CellstoaddmodutraTdd
}
