package ies

// BandCombinationListSidelinkEUTRA-NR-v1630 ::= SEQUENCE OF BandCombinationParametersSidelinkEUTRA-NR-v1630
// SIZE (1..maxBandComb)
type BandcombinationlistsidelinkeutraNrV1630 struct {
	Value []BandcombinationparameterssidelinkeutraNrV1630 `lb:1,ub:maxBandComb`
}
