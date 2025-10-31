package ies

// BandCombinationList-v16a0 ::= SEQUENCE OF BandCombination-v16a0
// SIZE (1..maxBandComb)
type BandcombinationlistV16a0 struct {
	Value []BandcombinationV16a0 `lb:1,ub:maxBandComb`
}
