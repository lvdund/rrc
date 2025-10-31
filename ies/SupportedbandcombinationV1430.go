package ies

// SupportedBandCombination-v1430 ::= SEQUENCE OF BandCombinationParameters-v1430
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1430 struct {
	Value []BandcombinationparametersV1430 `lb:1,ub:maxBandCombR10`
}
