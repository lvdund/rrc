package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-SchedulingEnhancement-r14 ::= ENUMERATED
type PhylayerparametersV1430CeSchedulingenhancementR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CeSchedulingenhancementR14EnumeratedNothing = iota
	PhylayerparametersV1430CeSchedulingenhancementR14EnumeratedSupported
)
