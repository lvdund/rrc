package ies

// SupportedBandCombinationReduced-v1530 ::= SEQUENCE OF BandCombinationParameters-v1530
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1530 struct {
	Value []BandcombinationparametersV1530 `lb:1,ub:maxBandCombR13`
}
