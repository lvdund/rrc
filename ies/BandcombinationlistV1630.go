package ies

// BandCombinationList-v1630 ::= SEQUENCE OF BandCombination-v1630
// SIZE (1..maxBandComb)
type BandcombinationlistV1630 struct {
	Value []BandcombinationV1630 `lb:1,ub:maxBandComb`
}
