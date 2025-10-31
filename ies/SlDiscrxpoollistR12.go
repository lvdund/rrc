package ies

// SL-DiscRxPoolList-r12 ::= SEQUENCE OF SL-DiscResourcePool-r12
// SIZE (1..maxSL-RxPool-r12)
type SlDiscrxpoollistR12 struct {
	Value []SlDiscresourcepoolR12 `lb:1,ub:maxSLRxpoolR12`
}
