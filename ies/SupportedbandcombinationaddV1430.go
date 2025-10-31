package ies

// SupportedBandCombinationAdd-v1430 ::= SEQUENCE OF BandCombinationParameters-v1430
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1430 struct {
	Value []BandcombinationparametersV1430 `lb:1,ub:maxBandCombR11`
}
