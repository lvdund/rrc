package ies

// SupportedBandCombinationAdd-v1320 ::= SEQUENCE OF BandCombinationParameters-v1320
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1320 struct {
	Value []BandcombinationparametersV1320 `lb:1,ub:maxBandCombR11`
}
