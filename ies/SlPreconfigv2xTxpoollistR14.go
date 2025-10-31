package ies

// SL-PreconfigV2X-TxPoolList-r14 ::= SEQUENCE OF SL-V2X-PreconfigCommPool-r14
// SIZE (1..maxSL-V2X-TxPoolPreconf-r14)
type SlPreconfigv2xTxpoollistR14 struct {
	Value []SlV2xPreconfigcommpoolR14 `lb:1,ub:maxSLV2xTxpoolpreconfR14`
}
