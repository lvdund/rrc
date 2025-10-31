package ies

// BandCombinationList-v1680 ::= SEQUENCE OF BandCombination-v1680
// SIZE (1..maxBandComb)
type BandcombinationlistV1680 struct {
	Value []BandcombinationV1680 `lb:1,ub:maxBandComb`
}
