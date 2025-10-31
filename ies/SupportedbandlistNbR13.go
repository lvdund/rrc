package ies

// SupportedBandList-NB-r13 ::= SEQUENCE OF SupportedBand-NB-r13
// SIZE (1..maxBands)
type SupportedbandlistNbR13 struct {
	Value []SupportedbandNbR13 `lb:1,ub:maxBands`
}
