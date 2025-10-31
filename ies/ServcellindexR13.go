package ies

import "rrc/utils"

// ServCellIndex-r13 ::= utils.INTEGER (0..31)
type ServcellindexR13 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
