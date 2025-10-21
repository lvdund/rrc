package ies

import "rrc/utils"

// MeasSubframeCellList-r10 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellMeas)
type MeassubframecelllistR10 struct {
	Value utils.Sequence[Physcellidrange]
}
