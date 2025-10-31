package ies

import "rrc/utils"

// SCellIndex-r13 ::= utils.INTEGER (1..31)
type ScellindexR13 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
