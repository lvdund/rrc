package ies

import "rrc/utils"

// SCellIndex-r10 ::= utils.INTEGER (1..7)
type ScellindexR10 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
