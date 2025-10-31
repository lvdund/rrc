package ies

import "rrc/utils"

// AvailabilityIndicator-r16 ::= SEQUENCE
// Extensible
type AvailabilityindicatorR16 struct {
	AiRntiR16                     AiRntiR16
	DciPayloadsizeaiR16           utils.INTEGER                              `lb:0,ub:maxAIDciPayloadsizeR16`
	AvailablecombtoaddmodlistR16  *[]AvailabilitycombinationspercellR16      `lb:1,ub:maxNrofDUCellsR16`
	AvailablecombtoreleaselistR16 *[]AvailabilitycombinationspercellindexR16 `lb:1,ub:maxNrofDUCellsR16`
}
