package ies

// BandCombinationParameters-v1090 ::= SEQUENCE OF BandParameters-v1090
// SIZE (1..maxSimultaneousBands-r10)
type BandcombinationparametersV1090 struct {
	Value []BandparametersV1090 `lb:1,ub:maxSimultaneousBandsR10`
}
