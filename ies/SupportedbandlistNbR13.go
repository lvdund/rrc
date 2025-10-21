package ies

import "rrc/utils"

// SupportedBandList-NB-r13 ::= SEQUENCE OF SupportedBand-NB-r13
// SIZE (1..maxBands)
type SupportedbandlistNbR13 struct {
	Value utils.Sequence[SupportedbandNbR13]
}
