package ies

// BandCombinationParametersSidelinkNR-v1710 ::= SEQUENCE OF BandParametersSidelink-v1710
// SIZE (1..maxSimultaneousBands)
type BandcombinationparameterssidelinknrV1710 struct {
	Value []BandparameterssidelinkV1710 `lb:1,ub:maxSimultaneousBands`
}
