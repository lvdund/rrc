package ies

import "rrc/utils"

// SL-DiscTxPoolToAddModList-r12 ::= SEQUENCE OF SL-DiscTxPoolToAddMod-r12
// SIZE (1..maxSL-TxPool-r12)
type SlDisctxpooltoaddmodlistR12 struct {
	Value utils.Sequence[SlDisctxpooltoaddmodR12]
}
