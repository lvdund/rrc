package ies

// BandCombinationList-v1610 ::= SEQUENCE OF BandCombination-v1610
// SIZE (1..maxBandComb)
type BandcombinationlistV1610 struct {
	Value []BandcombinationV1610 `lb:1,ub:maxBandComb`
}
