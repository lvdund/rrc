package ies

import "rrc/utils"

// SupportedBandListNR-r15 ::= SEQUENCE OF SupportedBandNR-r15
// SIZE (1..maxBandsNR-r15)
type SupportedbandlistnrR15 struct {
	Value utils.Sequence[SupportedbandnrR15]
}
