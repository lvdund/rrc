package ies

import "rrc/utils"

// ServCellIndex-r10 ::= utils.INTEGER (0..7)
type ServcellindexR10 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
