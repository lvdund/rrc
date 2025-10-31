package ies

// BandCombinationParameters-r10 ::= SEQUENCE OF BandParameters-r10
// SIZE (1..maxSimultaneousBands-r10)
type BandcombinationparametersR10 struct {
	Value []BandparametersR10 `lb:1,ub:maxSimultaneousBandsR10`
}
