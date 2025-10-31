package ies

import "rrc/utils"

// WhiteCellsToAddMod-r13 ::= SEQUENCE
type WhitecellstoaddmodR13 struct {
	CellindexR13       utils.INTEGER `lb:0,ub:maxCellMeas`
	PhyscellidrangeR13 Physcellidrange
}
