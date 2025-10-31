package ies

// SL-PreconfigDiscTxPoolList-r13 ::= SEQUENCE OF SL-PreconfigDiscPool-r13
// SIZE (1..maxSL-DiscTxPoolPreconf-r13)
type SlPreconfigdisctxpoollistR13 struct {
	Value []SlPreconfigdiscpoolR13 `lb:1,ub:maxSLDisctxpoolpreconfR13`
}
