package ies

// SupportedBandCombinationReduced-v1320 ::= SEQUENCE OF BandCombinationParameters-v1320
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1320 struct {
	Value []BandcombinationparametersV1320 `lb:1,ub:maxBandCombR13`
}
