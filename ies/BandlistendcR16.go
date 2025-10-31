package ies

// BandListENDC-r16 ::= SEQUENCE OF FreqBandIndicatorNR-r15
// SIZE (1.. maxBandsENDC-r16)
type BandlistendcR16 struct {
	Value []FreqbandindicatornrR15 `lb:1,ub:maxBandsENDCR16`
}
