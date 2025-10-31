package ies

// SupportedBandCombinationAdd-v1450 ::= SEQUENCE OF BandCombinationParameters-v1450
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1450 struct {
	Value []BandcombinationparametersV1450 `lb:1,ub:maxBandCombR11`
}
