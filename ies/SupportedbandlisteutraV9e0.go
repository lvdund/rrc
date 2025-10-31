package ies

// SupportedBandListEUTRA-v9e0 ::= SEQUENCE OF SupportedBandEUTRA-v9e0
// SIZE (1..maxBands)
type SupportedbandlisteutraV9e0 struct {
	Value []SupportedbandeutraV9e0 `lb:1,ub:maxBands`
}
