package ies

// SupportedBandCombination-v1250 ::= SEQUENCE OF BandCombinationParameters-v1250
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1250 struct {
	Value []BandcombinationparametersV1250 `lb:1,ub:maxBandCombR10`
}
