package ies

// BandCombinationListSidelinkNR-r16 ::= SEQUENCE OF BandCombinationParametersSidelinkNR-r16
// SIZE (1..maxBandComb)
type BandcombinationlistsidelinknrR16 struct {
	Value []BandcombinationparameterssidelinknrR16 `lb:1,ub:maxBandComb`
}
