package ies

// SL-PreconfigCommRxPoolList-r13 ::= SEQUENCE OF SL-PreconfigCommPool-r12
// SIZE (1..maxSL-CommRxPoolPreconf-v1310)
type SlPreconfigcommrxpoollistR13 struct {
	Value []SlPreconfigcommpoolR12 `lb:1,ub:maxSLCommrxpoolpreconfV1310`
}
