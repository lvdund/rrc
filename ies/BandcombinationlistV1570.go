package ies

// BandCombinationList-v1570 ::= SEQUENCE OF BandCombination-v1570
// SIZE (1..maxBandComb)
type BandcombinationlistV1570 struct {
	Value []BandcombinationV1570 `lb:1,ub:maxBandComb`
}
