package ies

import "rrc/utils"

// ARFCN-ValueEUTRA ::= utils.INTEGER (0..maxEARFCN)
type ArfcnValueeutra struct {
	Value utils.INTEGER `lb:0,ub:maxEARFCN`
}
