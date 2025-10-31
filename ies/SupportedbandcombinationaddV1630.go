package ies

// SupportedBandCombinationAdd-v1630 ::= SEQUENCE OF BandCombinationParameters-v1630
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1630 struct {
	Value []BandcombinationparametersV1630 `lb:1,ub:maxBandCombR11`
}
