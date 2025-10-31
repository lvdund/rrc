package ies

// SL-PreconfigV2X-RxPoolList-r14 ::= SEQUENCE OF SL-V2X-PreconfigCommPool-r14
// SIZE (1..maxSL-V2X-RxPoolPreconf-r14)
type SlPreconfigv2xRxpoollistR14 struct {
	Value []SlV2xPreconfigcommpoolR14 `lb:1,ub:maxSLV2xRxpoolpreconfR14`
}
