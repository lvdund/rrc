package ies

// SupportedBandCombinationReduced-v1430 ::= SEQUENCE OF BandCombinationParameters-v1430
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1430 struct {
	Value []BandcombinationparametersV1430 `lb:1,ub:maxBandCombR13`
}
