package ies

// BandCombinationList-v1700 ::= SEQUENCE OF BandCombination-v1700
// SIZE (1..maxBandComb)
type BandcombinationlistV1700 struct {
	Value []BandcombinationV1700 `lb:1,ub:maxBandComb`
}
