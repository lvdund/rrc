package ies

// BandCombinationList-v1580 ::= SEQUENCE OF BandCombination-v1580
// SIZE (1..maxBandComb)
type BandcombinationlistV1580 struct {
	Value []BandcombinationV1580 `lb:1,ub:maxBandComb`
}
