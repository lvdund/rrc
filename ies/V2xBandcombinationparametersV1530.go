package ies

// V2X-BandCombinationParameters-v1530 ::= SEQUENCE OF V2X-BandParameters-v1530
// SIZE (1.. maxSimultaneousBands-r10)
type V2xBandcombinationparametersV1530 struct {
	Value []V2xBandparametersV1530 `lb:1,ub:maxSimultaneousBandsR10`
}
