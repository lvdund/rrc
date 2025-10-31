package ies

import "rrc/utils"

// UAC-BarringInfoSetIndex ::= utils.INTEGER (1..maxBarringInfoSet)
type UacBarringinfosetindex struct {
	Value utils.INTEGER `lb:0,ub:maxBarringInfoSet`
}
