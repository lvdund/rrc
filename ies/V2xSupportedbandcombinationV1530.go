package ies

// V2X-SupportedBandCombination-v1530 ::= SEQUENCE OF V2X-BandCombinationParameters-v1530
// SIZE (1..maxBandComb-r13)
type V2xSupportedbandcombinationV1530 struct {
	Value []V2xBandcombinationparametersV1530 `lb:1,ub:maxBandCombR13`
}
