package ies

// SupportedBandCombinationAdd-v11d0 ::= SEQUENCE OF BandCombinationParameters-v10i0
// SIZE (1..maxBandComb-r11)
type SupportedbandcombinationaddV11d0 struct {
	Value []BandcombinationparametersV10i0 `lb:1,ub:maxBandCombR11`
}
