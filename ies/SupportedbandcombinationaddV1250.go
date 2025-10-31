package ies

// SupportedBandCombinationAdd-v1250 ::= SEQUENCE OF BandCombinationParameters-v1250
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1250 struct {
	Value []BandcombinationparametersV1250 `lb:1,ub:maxBandCombR11`
}
