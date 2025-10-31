package ies

// SupportedBandListUTRA-TDD384 ::= SEQUENCE OF SupportedBandUTRA-TDD384
// SIZE (1..maxBands)
type SupportedbandlistutraTdd384 struct {
	Value []SupportedbandutraTdd384 `lb:1,ub:maxBands`
}
