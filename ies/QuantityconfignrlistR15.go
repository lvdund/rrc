package ies

// QuantityConfigNRList-r15 ::= SEQUENCE OF QuantityConfigNR-r15
// SIZE (1..maxQuantSetsNR-r15)
type QuantityconfignrlistR15 struct {
	Value []QuantityconfignrR15 `lb:1,ub:maxQuantSetsNRR15`
}
