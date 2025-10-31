package ies

import "rrc/utils"

// CellsToAddModUTRA-TDD ::= SEQUENCE
type CellstoaddmodutraTdd struct {
	Cellindex  utils.INTEGER `lb:0,ub:maxCellMeas`
	Physcellid PhyscellidutraTdd
}
