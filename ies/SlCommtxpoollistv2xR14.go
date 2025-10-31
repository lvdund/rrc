package ies

// SL-CommTxPoolListV2X-r14 ::= SEQUENCE OF SL-CommResourcePoolV2X-r14
// SIZE (1..maxSL-V2X-TxPool-r14)
type SlCommtxpoollistv2xR14 struct {
	Value []SlCommresourcepoolv2xR14 `lb:1,ub:maxSLV2xTxpoolR14`
}
