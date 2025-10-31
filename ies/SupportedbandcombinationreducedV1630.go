package ies

// SupportedBandCombinationReduced-v1630 ::= SEQUENCE OF BandCombinationParameters-v1630
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1630 struct {
	Value []BandcombinationparametersV1630 `lb:1,ub:maxBandCombR13`
}
