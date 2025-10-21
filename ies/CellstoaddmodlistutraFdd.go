package ies

import "rrc/utils"

// CellsToAddModListUTRA-FDD ::= SEQUENCE OF CellsToAddModUTRA-FDD
// SIZE (1..maxCellMeas)
type CellstoaddmodlistutraFdd struct {
	Value []CellstoaddmodutraFdd
}
