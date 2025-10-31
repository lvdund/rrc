package ies

// SL-CommTxPoolList-r12 ::= SEQUENCE OF SL-CommResourcePool-r12
// SIZE (1..maxSL-TxPool-r12)
type SlCommtxpoollistR12 struct {
	Value []SlCommresourcepoolR12 `lb:1,ub:maxSLTxpoolR12`
}
