package ies

// BandCombinationListSidelinkNR-v1710 ::= SEQUENCE OF BandCombinationParametersSidelinkNR-v1710
// SIZE (1..maxBandComb)
type BandcombinationlistsidelinknrV1710 struct {
	Value []BandcombinationparameterssidelinknrV1710 `lb:1,ub:maxBandComb`
}
