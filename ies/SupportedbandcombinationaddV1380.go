package ies

// SupportedBandCombinationAdd-v1380 ::= SEQUENCE OF BandCombinationParameters-v1380
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV1380 struct {
	Value []BandcombinationparametersV1380 `lb:1,ub:maxBandCombR11`
}
