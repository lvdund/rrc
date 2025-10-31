package ies

// BandCombinationList-v1730 ::= SEQUENCE OF BandCombination-v1730
// SIZE (1..maxBandComb)
type BandcombinationlistV1730 struct {
	Value []BandcombinationV1730 `lb:1,ub:maxBandComb`
}
