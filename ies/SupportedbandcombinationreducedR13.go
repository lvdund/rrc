package ies

// SupportedBandCombinationReduced-r13 ::= SEQUENCE OF BandCombinationParameters-r13
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedR13 struct {
	Value []BandcombinationparametersR13 `lb:1,ub:maxBandCombR13`
}
