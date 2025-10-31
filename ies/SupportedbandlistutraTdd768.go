package ies

// SupportedBandListUTRA-TDD768 ::= SEQUENCE OF SupportedBandUTRA-TDD768
// SIZE (1..maxBands)
type SupportedbandlistutraTdd768 struct {
	Value []SupportedbandutraTdd768 `lb:1,ub:maxBands`
}
