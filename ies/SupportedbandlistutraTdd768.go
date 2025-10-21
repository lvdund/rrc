package ies

import "rrc/utils"

// SupportedBandListUTRA-TDD768 ::= SEQUENCE OF SupportedBandUTRA-TDD768
// SIZE (1..maxBands)
type SupportedbandlistutraTdd768 struct {
	Value utils.Sequence[SupportedbandutraTdd768]
}
