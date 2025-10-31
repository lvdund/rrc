package ies

// BandCombinationList-v1720 ::= SEQUENCE OF BandCombination-v1720
// SIZE (1..maxBandComb)
type BandcombinationlistV1720 struct {
	Value []BandcombinationV1720 `lb:1,ub:maxBandComb`
}
