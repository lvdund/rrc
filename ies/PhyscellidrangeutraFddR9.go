package ies

import "rrc/utils"

// PhysCellIdRangeUTRA-FDD-r9 ::= SEQUENCE
type PhyscellidrangeutraFddR9 struct {
	StartR9 PhyscellidutraFdd
	RangeR9 *utils.INTEGER `lb:0,ub:512`
}
