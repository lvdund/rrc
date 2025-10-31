package ies

// FreqBandList ::= SEQUENCE OF FreqBandInformation
// SIZE (1..maxBandsMRDC)
type Freqbandlist struct {
	Value []Freqbandinformation `lb:1,ub:maxBandsMRDC`
}
