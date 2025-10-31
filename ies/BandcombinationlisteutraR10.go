package ies

// BandCombinationListEUTRA-r10 ::= SEQUENCE OF BandInfoEUTRA
// SIZE (1..maxBandComb-r10)
type BandcombinationlisteutraR10 struct {
	Value []Bandinfoeutra `lb:1,ub:maxBandCombR10`
}
