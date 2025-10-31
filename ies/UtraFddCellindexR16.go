package ies

import "rrc/utils"

// UTRA-FDD-CellIndex-r16 ::= utils.INTEGER (1..maxCellMeasUTRA-FDD-r16)
type UtraFddCellindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxCellMeasUTRAFddR16`
}
