package ies

import "rrc/utils"

// SupportedBandListUTRA-FDD ::= SEQUENCE OF SupportedBandUTRA-FDD
// SIZE (1..maxBands)
type SupportedbandlistutraFdd struct {
	Value utils.Sequence[SupportedbandutraFdd]
}
