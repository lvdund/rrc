package ies

import "rrc/utils"

// CellsToAddMod ::= SEQUENCE
type Cellstoaddmod struct {
	Cellindex            utils.INTEGER `lb:0,ub:maxCellMeas`
	Physcellid           Physcellid
	Cellindividualoffset QOffsetrange
}
