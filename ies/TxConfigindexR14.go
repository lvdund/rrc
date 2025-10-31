package ies

import "rrc/utils"

// Tx-ConfigIndex-r14 ::= utils.INTEGER (0..maxSL-V2X-TxConfig-1-r14)
type TxConfigindexR14 struct {
	Value utils.INTEGER `lb:0,ub:maxSLV2xTxconfig1R14`
}
