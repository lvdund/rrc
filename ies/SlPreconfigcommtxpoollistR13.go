package ies

// SL-PreconfigCommTxPoolList-r13 ::= SEQUENCE OF SL-PreconfigCommPool-r12
// SIZE (1..maxSL-CommTxPoolPreconf-v1310)
type SlPreconfigcommtxpoollistR13 struct {
	Value []SlPreconfigcommpoolR12 `lb:1,ub:maxSLCommtxpoolpreconfV1310`
}
