package ies

// SL-CommTxPoolToAddModListExt-r13 ::= SEQUENCE OF SL-CommTxPoolToAddModExt-r13
// SIZE (1..maxSL-TxPool-v1310)
type SlCommtxpooltoaddmodlistextR13 struct {
	Value []SlCommtxpooltoaddmodextR13 `lb:1,ub:maxSLTxpoolV1310`
}
