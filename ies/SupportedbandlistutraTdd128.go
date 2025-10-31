package ies

// SupportedBandListUTRA-TDD128 ::= SEQUENCE OF SupportedBandUTRA-TDD128
// SIZE (1..maxBands)
type SupportedbandlistutraTdd128 struct {
	Value []SupportedbandutraTdd128 `lb:1,ub:maxBands`
}
