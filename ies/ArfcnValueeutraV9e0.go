package ies

import "rrc/utils"

// ARFCN-ValueEUTRA-v9e0 ::= utils.INTEGER (maxEARFCN-Plus1..maxEARFCN2)
type ArfcnValueeutraV9e0 struct {
	Value utils.INTEGER `lb:maxEARFCNPlus1,ub:maxEARFCN2`
}
