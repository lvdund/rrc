package ies

import "rrc/utils"

// CellsToAddModListCDMA2000 ::= SEQUENCE OF CellsToAddModCDMA2000
// SIZE (1..maxCellMeas)
type Cellstoaddmodlistcdma2000 struct {
	Value utils.Sequence[Cellstoaddmodcdma2000]
}
