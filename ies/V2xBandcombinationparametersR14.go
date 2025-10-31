package ies

// V2X-BandCombinationParameters-r14 ::= SEQUENCE OF V2X-BandParameters-r14
// SIZE (1.. maxSimultaneousBands-r10)
type V2xBandcombinationparametersR14 struct {
	Value []V2xBandparametersR14 `lb:1,ub:maxSimultaneousBandsR10`
}
