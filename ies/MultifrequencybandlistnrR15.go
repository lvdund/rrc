package ies

// MultiFrequencyBandListNR-r15 ::= SEQUENCE OF FreqBandIndicatorNR-r15
// SIZE (1.. maxMultiBandsNR-r15)
type MultifrequencybandlistnrR15 struct {
	Value []FreqbandindicatornrR15 `lb:1,ub:maxMultiBandsNRR15`
}
