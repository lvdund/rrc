package ies

// BandCombinationList ::= SEQUENCE OF BandCombination
// SIZE (1..maxBandComb)
type Bandcombinationlist struct {
	Value []Bandcombination `lb:1,ub:maxBandComb`
}
