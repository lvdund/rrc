package ies

// SupportedBandCombination-v1130 ::= SEQUENCE OF BandCombinationParameters-v1130
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1130 struct {
	Value []BandcombinationparametersV1130 `lb:1,ub:maxBandCombR10`
}
