package ies

import "rrc/utils"

// CellIndex ::= utils.INTEGER (1..maxCellMeas)
type Cellindex struct {
	Value utils.INTEGER `lb:0,ub:maxCellMeas`
}
