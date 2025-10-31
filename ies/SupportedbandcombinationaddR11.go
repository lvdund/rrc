package ies

// SupportedBandCombinationAdd-r11 ::= SEQUENCE OF BandCombinationParameters-r11
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddR11 struct {
	Value []BandcombinationparametersR11 `lb:1,ub:maxBandCombR11`
}
