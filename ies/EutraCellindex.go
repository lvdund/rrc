package ies

import "rrc/utils"

// EUTRA-CellIndex ::= utils.INTEGER (1..maxCellMeasEUTRA)
type EutraCellindex struct {
	Value utils.INTEGER `lb:0,ub:maxCellMeasEUTRA`
}
