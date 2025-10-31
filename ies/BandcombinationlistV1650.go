package ies

// BandCombinationList-v1650 ::= SEQUENCE OF BandCombination-v1650
// SIZE (1..maxBandComb)
type BandcombinationlistV1650 struct {
	Value []BandcombinationV1650 `lb:1,ub:maxBandComb`
}
