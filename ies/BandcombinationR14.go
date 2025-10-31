package ies

// BandCombination-r14 ::= SEQUENCE OF BandIndication-r14
// SIZE (1..maxSimultaneousBands-r10)
type BandcombinationR14 struct {
	Value []BandindicationR14 `lb:1,ub:maxSimultaneousBandsR10`
}
