package ies

import "rrc/utils"

// SupportedBandListUTRA-TDD128 ::= SEQUENCE OF SupportedBandUTRA-TDD128
// SIZE (1..maxBands)
type SupportedbandlistutraTdd128 struct {
	Value utils.Sequence[SupportedbandutraTdd128]
}
