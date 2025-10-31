package ies

// SupportedBandCombinationAdd-v1270 ::= SEQUENCE OF BandCombinationParameters-v1270
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1270 struct {
	Value []BandcombinationparametersV1270 `lb:1,ub:maxBandCombR11`
}
