package ies

import "rrc/utils"

// UAC-BarringInfoSetIndex-r15 ::= utils.INTEGER (1..maxBarringInfoSet-r15)
type UacBarringinfosetindexR15 struct {
	Value utils.INTEGER `lb:0,ub:maxBarringInfoSetR15`
}
