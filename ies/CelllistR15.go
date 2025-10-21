package ies

import "rrc/utils"

// CellList-r15 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellMeasIdle-r15)
type CelllistR15 struct {
	Value utils.Sequence[Physcellidrange]
}
