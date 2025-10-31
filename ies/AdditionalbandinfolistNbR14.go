package ies

// AdditionalBandInfoList-NB-r14 ::= SEQUENCE OF FreqBandIndicator-NB-r13
// SIZE (1..maxMultiBands)
type AdditionalbandinfolistNbR14 struct {
	Value []FreqbandindicatorNbR13 `lb:1,ub:maxMultiBands`
}
