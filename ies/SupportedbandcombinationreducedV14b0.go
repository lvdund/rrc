package ies

// SupportedBandCombinationReduced-v14b0 ::= SEQUENCE OF BandCombinationParameters-v14b0
// SIZE (1..maxBandComb-r13)
type SupportedbandcombinationreducedV14b0 struct {
	Value []BandcombinationparametersV14b0 `lb:1,ub:maxBandCombR13`
}
