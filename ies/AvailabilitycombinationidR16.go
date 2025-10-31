package ies

import "rrc/utils"

// AvailabilityCombinationId-r16 ::= utils.INTEGER (0..maxNrofAvailabilityCombinationsPerSet-1-r16)
type AvailabilitycombinationidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofAvailabilityCombinationsPerSet1R16`
}
