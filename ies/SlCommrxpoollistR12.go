package ies

// SL-CommRxPoolList-r12 ::= SEQUENCE OF SL-CommResourcePool-r12
// SIZE (1..maxSL-RxPool-r12)
type SlCommrxpoollistR12 struct {
	Value []SlCommresourcepoolR12 `lb:1,ub:maxSLRxpoolR12`
}
