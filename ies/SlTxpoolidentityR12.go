package ies

import "rrc/utils"

// SL-TxPoolIdentity-r12 ::= utils.INTEGER (1.. maxSL-TxPool-r12)
type SlTxpoolidentityR12 struct {
	Value utils.INTEGER `lb:0,ub:maxSLTxpoolR12`
}
