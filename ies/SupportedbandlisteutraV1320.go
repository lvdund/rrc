package ies

// SupportedBandListEUTRA-v1320 ::= SEQUENCE OF SupportedBandEUTRA-v1320
// SIZE (1..maxBands)
type SupportedbandlisteutraV1320 struct {
	Value []SupportedbandeutraV1320 `lb:1,ub:maxBands`
}
