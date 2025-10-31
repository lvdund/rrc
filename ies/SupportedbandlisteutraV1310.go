package ies

// SupportedBandListEUTRA-v1310 ::= SEQUENCE OF SupportedBandEUTRA-v1310
// SIZE (1..maxBands)
type SupportedbandlisteutraV1310 struct {
	Value []SupportedbandeutraV1310 `lb:1,ub:maxBands`
}
