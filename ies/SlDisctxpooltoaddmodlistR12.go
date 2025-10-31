package ies

// SL-DiscTxPoolToAddModList-r12 ::= SEQUENCE OF SL-DiscTxPoolToAddMod-r12
// SIZE (1..maxSL-TxPool-r12)
type SlDisctxpooltoaddmodlistR12 struct {
	Value []SlDisctxpooltoaddmodR12 `lb:1,ub:maxSLTxpoolR12`
}
