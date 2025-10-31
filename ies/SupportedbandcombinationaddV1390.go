package ies

// SupportedBandCombinationAdd-v1390 ::= SEQUENCE OF BandCombinationParameters-v1390
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1390 struct {
	Value []BandcombinationparametersV1390 `lb:1,ub:maxBandCombR11`
}
