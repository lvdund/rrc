package ies

// SupportedBandListEUTRA-v1250 ::= SEQUENCE OF SupportedBandEUTRA-v1250
// SIZE (1..maxBands)
type SupportedbandlisteutraV1250 struct {
	Value []SupportedbandeutraV1250 `lb:1,ub:maxBands`
}
