package ies

// MultiBandInfoList ::= SEQUENCE OF FreqBandIndicator
// SIZE (1..maxMultiBands)
type Multibandinfolist struct {
	Value []Freqbandindicator `lb:1,ub:maxMultiBands`
}
