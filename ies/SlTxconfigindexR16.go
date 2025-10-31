package ies

import "rrc/utils"

// SL-TxConfigIndex-r16 ::= utils.INTEGER (0..maxTxConfig-1-r16)
type SlTxconfigindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxTxConfig1R16`
}
