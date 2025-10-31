package ies

// SL-CommRxPoolListV2X-r14 ::= SEQUENCE OF SL-CommResourcePoolV2X-r14
// SIZE (1..maxSL-V2X-RxPool-r14)
type SlCommrxpoollistv2xR14 struct {
	Value []SlCommresourcepoolv2xR14 `lb:1,ub:maxSLV2xRxpoolR14`
}
