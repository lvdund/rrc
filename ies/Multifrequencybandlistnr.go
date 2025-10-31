package ies

// MultiFrequencyBandListNR ::= SEQUENCE OF FreqBandIndicatorNR
// SIZE (1..maxNrofMultiBands)
type Multifrequencybandlistnr struct {
	Value []Freqbandindicatornr `lb:1,ub:maxNrofMultiBands`
}
