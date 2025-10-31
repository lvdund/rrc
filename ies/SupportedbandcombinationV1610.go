package ies

// SupportedBandCombination-v1610 ::= SEQUENCE OF BandCombinationParameters-v1610
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1610 struct {
	Value []BandcombinationparametersV1610 `lb:1,ub:maxBandCombR10`
}
