package ies

import "rrc/utils"

// MeasIdToRemoveList ::= SEQUENCE OF MeasId
// SIZE (1..maxMeasId)
type Measidtoremovelist struct {
	Value utils.Sequence[Measid]
}
