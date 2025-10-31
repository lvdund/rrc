package ies

// BandListEUTRA ::= SEQUENCE OF BandInfoEUTRA
// SIZE (1..maxBands)
type Bandlisteutra struct {
	Value []Bandinfoeutra `lb:1,ub:maxBands`
}
