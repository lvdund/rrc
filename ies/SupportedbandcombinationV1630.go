package ies

// SupportedBandCombination-v1630 ::= SEQUENCE OF BandCombinationParameters-v1630
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationV1630 struct {
	Value []BandcombinationparametersV1630 `lb:1,ub:maxBandCombR10`
}
