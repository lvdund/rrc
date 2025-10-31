package ies

// SupportedBandCombination-v1270 ::= SEQUENCE OF BandCombinationParameters-v1270
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1270 struct {
	Value []BandcombinationparametersV1270 `lb:1,ub:maxBandCombR10`
}
