package ies

// SupportedBandCombinationAdd-v1530 ::= SEQUENCE OF BandCombinationParameters-v1530
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1530 struct {
	Value []BandcombinationparametersV1530 `lb:1,ub:maxBandCombR11`
}
