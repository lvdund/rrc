package ies

import "rrc/utils"

// InterRAT-BandList ::= SEQUENCE OF InterRAT-BandInfo
// SIZE (1..maxBands)
type InterratBandlist struct {
	Value utils.Sequence[InterratBandinfo]
}
