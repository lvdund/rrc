package ies

// BandCombinationListSidelinkEUTRA-NR-r16 ::= SEQUENCE OF BandCombinationParametersSidelinkEUTRA-NR-r16
// SIZE (1..maxBandComb)
type BandcombinationlistsidelinkeutraNrR16 struct {
	Value []BandcombinationparameterssidelinkeutraNrR16 `lb:1,ub:maxBandComb`
}
