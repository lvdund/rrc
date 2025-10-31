package ies

// InterFreqBandList ::= SEQUENCE OF InterFreqBandInfo
// SIZE (1..maxBands)
type Interfreqbandlist struct {
	Value []Interfreqbandinfo `lb:1,ub:maxBands`
}
