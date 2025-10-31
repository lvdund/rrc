package ies

import "rrc/utils"

// SL-V2X-TxPoolIdentity-r14 ::= utils.INTEGER (1.. maxSL-V2X-TxPool-r14)
type SlV2xTxpoolidentityR14 struct {
	Value utils.INTEGER `lb:0,ub:maxSLV2xTxpoolR14`
}
