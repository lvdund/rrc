package ies

import "rrc/utils"

// SL-TxPercentageConfig-r16 ::= SEQUENCE
type SlTxpercentageconfigR16 struct {
	SlPriorityR16     utils.INTEGER `lb:0,ub:8`
	SlTxpercentageR16 SlTxpercentageconfigR16SlTxpercentageR16
}
