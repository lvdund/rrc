package ies

// SL-PreconfigDiscRxPoolList-r13 ::= SEQUENCE OF SL-PreconfigDiscPool-r13
// SIZE (1..maxSL-DiscRxPoolPreconf-r13)
type SlPreconfigdiscrxpoollistR13 struct {
	Value []SlPreconfigdiscpoolR13 `lb:1,ub:maxSLDiscrxpoolpreconfR13`
}
