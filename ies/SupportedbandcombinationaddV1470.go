package ies

// SupportedBandCombinationAdd-v1470 ::= SEQUENCE OF BandCombinationParameters-v1470
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1470 struct {
	Value []BandcombinationparametersV1470 `lb:1,ub:maxBandCombR11`
}
