package ies

// BandCombinationList-v1640 ::= SEQUENCE OF BandCombination-v1640
// SIZE (1..maxBandComb)
type BandcombinationlistV1640 struct {
	Value []BandcombinationV1640 `lb:1,ub:maxBandComb`
}
