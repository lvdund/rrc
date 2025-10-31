package ies

// SupportedBandListEUTRA ::= SEQUENCE OF SupportedBandEUTRA
// SIZE (1..maxBands)
type Supportedbandlisteutra struct {
	Value []Supportedbandeutra `lb:1,ub:maxBands`
}
