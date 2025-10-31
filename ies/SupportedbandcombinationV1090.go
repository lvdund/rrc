package ies

// SupportedBandCombination-v1090 ::= SEQUENCE OF BandCombinationParameters-v1090
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1090 struct {
	Value []BandcombinationparametersV1090 `lb:1,ub:maxBandCombR10`
}
