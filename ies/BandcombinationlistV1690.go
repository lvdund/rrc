package ies

// BandCombinationList-v1690 ::= SEQUENCE OF BandCombination-v1690
// SIZE (1..maxBandComb)
type BandcombinationlistV1690 struct {
	Value []BandcombinationV1690 `lb:1,ub:maxBandComb`
}
