package ies

// SupportedBandCombination-v1530 ::= SEQUENCE OF BandCombinationParameters-v1530
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1530 struct {
	Value []BandcombinationparametersV1530 `lb:1,ub:maxBandCombR10`
}
