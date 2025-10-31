package ies

import "rrc/utils"

// CellsToAddModNR-r16 ::= SEQUENCE
type CellstoaddmodnrR16 struct {
	CellindexR16            utils.INTEGER `lb:0,ub:maxCellMeas`
	PhyscellidR16           PhyscellidnrR15
	CellindividualoffsetR16 QOffsetrange
}
