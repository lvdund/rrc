package ies

// BandCombinationListSidelinkEUTRA-NR-v1710 ::= SEQUENCE OF BandCombinationParametersSidelinkEUTRA-NR-v1710
// SIZE (1..maxBandComb)
type BandcombinationlistsidelinkeutraNrV1710 struct {
	Value []BandcombinationparameterssidelinkeutraNrV1710 `lb:1,ub:maxBandComb`
}
