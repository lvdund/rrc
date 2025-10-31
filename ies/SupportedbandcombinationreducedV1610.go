package ies

// SupportedBandCombinationReduced-v1610 ::= SEQUENCE OF BandCombinationParameters-v1610
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1610 struct {
	Value []BandcombinationparametersV1610 `lb:1,ub:maxBandCombR13`
}
