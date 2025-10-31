package ies

// SupportedBandCombination-v10i0 ::= SEQUENCE OF BandCombinationParameters-v10i0
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV10i0 struct {
	Value []BandcombinationparametersV10i0 `lb:1,ub:maxBandCombR10`
}
