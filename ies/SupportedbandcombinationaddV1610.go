package ies

// SupportedBandCombinationAdd-v1610 ::= SEQUENCE OF BandCombinationParameters-v1610
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1610 struct {
	Value []BandcombinationparametersV1610 `lb:1,ub:maxBandCombR11`
}
