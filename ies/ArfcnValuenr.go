package ies

import "rrc/utils"

// ARFCN-ValueNR ::= utils.INTEGER (0..maxNARFCN)
type ArfcnValuenr struct {
	Value utils.INTEGER `lb:0,ub:maxNARFCN`
}
