package ies

import "rrc/utils"

// ARFCN-ValueCDMA2000 ::= utils.INTEGER (0..2047)
type ArfcnValuecdma2000 struct {
	Value utils.INTEGER `lb:0,ub:2047`
}
