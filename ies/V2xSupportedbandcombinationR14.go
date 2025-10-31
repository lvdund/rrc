package ies

// V2X-SupportedBandCombination-r14 ::= SEQUENCE OF V2X-BandCombinationParameters-r14
// SIZE (1..maxBandComb-r13)
type V2xSupportedbandcombinationR14 struct {
	Value []V2xBandcombinationparametersR14 `lb:1,ub:maxBandCombR13`
}
