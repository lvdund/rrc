package ies

import "rrc/utils"

// AvailabilityCombination-r16 ::= SEQUENCE
type AvailabilitycombinationR16 struct {
	AvailabilitycombinationidR16 AvailabilitycombinationidR16
	ResourceavailabilityR16      []utils.INTEGER `lb:1,ub:maxNrofResourceAvailabilityPerCombinationR16`
}
