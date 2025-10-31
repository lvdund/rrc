package ies

// SupportedBandCombination-v1320 ::= SEQUENCE OF BandCombinationParameters-v1320
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1320 struct {
	Value []BandcombinationparametersV1320 `lb:1,ub:maxBandCombR10`
}
