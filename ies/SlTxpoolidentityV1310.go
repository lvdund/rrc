package ies

import "rrc/utils"

// SL-TxPoolIdentity-v1310 ::= utils.INTEGER (maxSL-TxPool-r12Plus1-r13.. maxSL-TxPool-r13)
type SlTxpoolidentityV1310 struct {
	Value utils.INTEGER `lb:maxSLTxpoolR12plus1R13,ub:maxSLTxpoolR13`
}
