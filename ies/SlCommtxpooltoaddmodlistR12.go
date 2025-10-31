package ies

// SL-CommTxPoolToAddModList-r12 ::= SEQUENCE OF SL-CommTxPoolToAddMod-r12
// SIZE (1..maxSL-TxPool-r12)
type SlCommtxpooltoaddmodlistR12 struct {
	Value []SlCommtxpooltoaddmodR12 `lb:1,ub:maxSLTxpoolR12`
}
