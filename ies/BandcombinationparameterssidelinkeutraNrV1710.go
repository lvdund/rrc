package ies

// BandCombinationParametersSidelinkEUTRA-NR-v1710 ::= SEQUENCE OF BandParametersSidelinkEUTRA-NR-v1710
// SIZE (1..maxSimultaneousBands)
type BandcombinationparameterssidelinkeutraNrV1710 struct {
	Value []BandparameterssidelinkeutraNrV1710 `lb:1,ub:maxSimultaneousBands`
}
