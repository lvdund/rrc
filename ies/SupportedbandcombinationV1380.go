package ies

// SupportedBandCombination-v1380 ::= SEQUENCE OF BandCombinationParameters-v1380
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1380 struct {
	Value []BandcombinationparametersV1380 `lb:1,ub:maxBandCombR10`
}
