package ies

// SupportedBandCombinationReduced-v1390 ::= SEQUENCE OF BandCombinationParameters-v1390
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV1390 struct {
	Value []BandcombinationparametersV1390 `lb:1,ub:maxBandCombR13`
}
