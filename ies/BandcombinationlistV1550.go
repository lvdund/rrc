package ies

// BandCombinationList-v1550 ::= SEQUENCE OF BandCombination-v1550
// SIZE (1..maxBandComb)
type BandcombinationlistV1550 struct {
	Value []BandcombinationV1550 `lb:1,ub:maxBandComb`
}
