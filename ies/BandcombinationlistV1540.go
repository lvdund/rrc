package ies

// BandCombinationList-v1540 ::= SEQUENCE OF BandCombination-v1540
// SIZE (1..maxBandComb)
type BandcombinationlistV1540 struct {
	Value []BandcombinationV1540 `lb:1,ub:maxBandComb`
}
