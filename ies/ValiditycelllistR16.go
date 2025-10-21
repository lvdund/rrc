package ies

import "rrc/utils"

// ValidityCellList-r16 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1.. maxCellMeasIdle-r15)
type ValiditycelllistR16 struct {
	Value []Physcellidrange
}
