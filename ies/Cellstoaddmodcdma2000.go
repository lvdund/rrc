package ies

import "rrc/utils"

// CellsToAddModCDMA2000 ::= SEQUENCE
type Cellstoaddmodcdma2000 struct {
	Cellindex  utils.INTEGER `lb:0,ub:maxCellMeas`
	Physcellid Physcellidcdma2000
}
