package ies

// BandCombinationList-v1590 ::= SEQUENCE OF BandCombination-v1590
// SIZE (1..maxBandComb)
type BandcombinationlistV1590 struct {
	Value []BandcombinationV1590 `lb:1,ub:maxBandComb`
}
