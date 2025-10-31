package ies

// BandCombinationList-v15g0 ::= SEQUENCE OF BandCombination-v15g0
// SIZE (1..maxBandComb)
type BandcombinationlistV15g0 struct {
	Value []BandcombinationV15g0 `lb:1,ub:maxBandComb`
}
