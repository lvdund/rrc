package ies

import "rrc/utils"

// CellsToAddModNR-r15 ::= SEQUENCE
type CellstoaddmodnrR15 struct {
	CellindexR15  utils.INTEGER `lb:0,ub:maxCellMeas`
	PhyscellidR15 PhyscellidnrR15
}
