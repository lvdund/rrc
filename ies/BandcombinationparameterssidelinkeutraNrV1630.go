package ies

// BandCombinationParametersSidelinkEUTRA-NR-v1630 ::= SEQUENCE OF BandParametersSidelinkEUTRA-NR-v1630
// SIZE (1..maxSimultaneousBands)
type BandcombinationparameterssidelinkeutraNrV1630 struct {
	Value []BandparameterssidelinkeutraNrV1630 `lb:1,ub:maxSimultaneousBands`
}
