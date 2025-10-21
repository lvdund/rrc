package ies

import "rrc/utils"

// SL-CommTxPoolList-r12 ::= SEQUENCE OF SL-CommResourcePool-r12
// SIZE (1..maxSL-TxPool-r12)
type SlCommtxpoollistR12 struct {
	Value []SlCommresourcepoolR12
}
