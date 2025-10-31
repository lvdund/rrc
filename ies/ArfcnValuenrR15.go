package ies

import "rrc/utils"

// ARFCN-ValueNR-r15 ::= utils.INTEGER (0.. 3279165)
type ArfcnValuenrR15 struct {
	Value utils.INTEGER `lb:0,ub:3279165`
}
