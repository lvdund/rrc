package ies

// SupportedBandCombination-v14b0 ::= SEQUENCE OF BandCombinationParameters-v14b0
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV14b0 struct {
	Value []BandcombinationparametersV14b0 `lb:1,ub:maxBandCombR10`
}
