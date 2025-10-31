package ies

import "rrc/utils"

// CellsToAddModUTRA-FDD ::= SEQUENCE
type CellstoaddmodutraFdd struct {
	Cellindex  utils.INTEGER `lb:0,ub:maxCellMeas`
	Physcellid PhyscellidutraFdd
}
