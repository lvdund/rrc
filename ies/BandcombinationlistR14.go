package ies

// BandCombinationList-r14 ::= SEQUENCE OF BandCombination-r14
// SIZE (1..maxBandComb-r13)
type BandcombinationlistR14 struct {
	Value []BandcombinationR14 `lb:1,ub:maxBandCombR13`
}
