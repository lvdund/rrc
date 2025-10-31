package ies

import "rrc/utils"

// ARFCN-ValueEUTRA-r9 ::= utils.INTEGER (0..maxEARFCN2)
type ArfcnValueeutraR9 struct {
	Value utils.INTEGER `lb:0,ub:maxEARFCN2`
}
