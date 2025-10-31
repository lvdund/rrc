package ies

// BandCombinationParametersSidelinkNR-r16 ::= SEQUENCE OF BandParametersSidelink-r16
// SIZE (1..maxSimultaneousBands)
type BandcombinationparameterssidelinknrR16 struct {
	Value []BandparameterssidelinkR16 `lb:1,ub:maxSimultaneousBands`
}
