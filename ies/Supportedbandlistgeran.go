package ies

import "rrc/utils"

// SupportedBandListGERAN ::= SEQUENCE OF SupportedBandGERAN
// SIZE (1..maxBands)
type Supportedbandlistgeran struct {
	Value utils.Sequence[Supportedbandgeran]
}
