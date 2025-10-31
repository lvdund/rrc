package ies

import "rrc/utils"

// AvailabilityCombinationsPerCell-r16 ::= SEQUENCE
// Extensible
type AvailabilitycombinationspercellR16 struct {
	AvailabilitycombinationspercellindexR16 AvailabilitycombinationspercellindexR16
	IabDuCellidentityR16                    Cellidentity
	PositionindciAiR16                      *utils.INTEGER                        `lb:0,ub:maxAIDciPayloadsize1R16`
	AvailabilitycombinationsR16             []AvailabilitycombinationR16          `lb:1,ub:maxNrofAvailabilityCombinationsPerSetR16`
	AvailabilitycombinationsrbGroupsR17     *[]AvailabilitycombinationrbGroupsR17 `lb:1,ub:maxNrofAvailabilityCombinationsPerSetR16,ext`
	PositionindciAiRbgroupsV1720            *utils.INTEGER                        `lb:0,ub:maxAIDciPayloadsize1R16,ext`
}
