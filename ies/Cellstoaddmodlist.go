package ies

import "rrc/utils"

// CellsToAddModList ::= SEQUENCE OF CellsToAddMod
// SIZE (1..maxCellMeas)
type Cellstoaddmodlist struct {
	Value utils.Sequence[Cellstoaddmod]
}
