package ies

// BandCombinationList-v1560 ::= SEQUENCE OF BandCombination-v1560
// SIZE (1..maxBandComb)
type BandcombinationlistV1560 struct {
	Value []BandcombinationV1560 `lb:1,ub:maxBandComb`
}
