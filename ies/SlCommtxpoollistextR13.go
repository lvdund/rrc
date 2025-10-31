package ies

// SL-CommTxPoolListExt-r13 ::= SEQUENCE OF SL-CommResourcePool-r12
// SIZE (1..maxSL-TxPool-v1310)
type SlCommtxpoollistextR13 struct {
	Value []SlCommresourcepoolR12 `lb:1,ub:maxSLTxpoolV1310`
}
