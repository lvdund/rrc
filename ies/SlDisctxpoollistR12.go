package ies

import "rrc/utils"

// SL-DiscTxPoolList-r12 ::= SEQUENCE OF SL-DiscResourcePool-r12
// SIZE (1..maxSL-TxPool-r12)
type SlDisctxpoollistR12 struct {
	Value []SlDiscresourcepoolR12
}
