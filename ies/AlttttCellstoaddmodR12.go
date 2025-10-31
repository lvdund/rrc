package ies

import "rrc/utils"

// AltTTT-CellsToAddMod-r12 ::= SEQUENCE
type AlttttCellstoaddmodR12 struct {
	CellindexR12       utils.INTEGER `lb:0,ub:maxCellMeas`
	PhyscellidrangeR12 Physcellidrange
}
