package ies

// SupportedBandCombination-r10 ::= SEQUENCE OF BandCombinationParameters-r10
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationR10 struct {
	Value []BandcombinationparametersR10 `lb:1,ub:maxBandCombR10`
}
