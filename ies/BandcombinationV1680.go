package ies

// BandCombination-v1680 ::= SEQUENCE
type BandcombinationV1680 struct {
	IntrabandconcurrentoperationpowerclassR16 *[]IntrabandpowerclassR16 `lb:1,ub:maxBandComb`
}
