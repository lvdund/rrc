package ies

// BandCombinationParametersSidelinkEUTRA-NR-r16 ::= SEQUENCE OF BandParametersSidelinkEUTRA-NR-r16
// SIZE (1..maxSimultaneousBands)
type BandcombinationparameterssidelinkeutraNrR16 struct {
	Value []BandparameterssidelinkeutraNrR16 `lb:1,ub:maxSimultaneousBands`
}
