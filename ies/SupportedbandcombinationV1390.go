package ies

// SupportedBandCombination-v1390 ::= SEQUENCE OF BandCombinationParameters-v1390
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1390 struct {
	Value []BandcombinationparametersV1390 `lb:1,ub:maxBandCombR10`
}
