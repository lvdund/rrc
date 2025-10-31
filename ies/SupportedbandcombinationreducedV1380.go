package ies

// SupportedBandCombinationReduced-v1380 ::= SEQUENCE OF BandCombinationParameters-v1380
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1380 struct {
	Value []BandcombinationparametersV1380 `lb:1,ub:maxBandCombR13`
}
