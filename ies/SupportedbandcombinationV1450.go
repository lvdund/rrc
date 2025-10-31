package ies

// SupportedBandCombination-v1450 ::= SEQUENCE OF BandCombinationParameters-v1450
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1450 struct {
	Value []BandcombinationparametersV1450 `lb:1,ub:maxBandCombR10`
}
