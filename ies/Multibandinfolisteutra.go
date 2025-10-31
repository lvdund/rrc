package ies

// MultiBandInfoListEUTRA ::= SEQUENCE OF FreqBandIndicatorEUTRA
// SIZE (1..maxMultiBands)
type Multibandinfolisteutra struct {
	Value []Freqbandindicatoreutra `lb:1,ub:maxMultiBands`
}
