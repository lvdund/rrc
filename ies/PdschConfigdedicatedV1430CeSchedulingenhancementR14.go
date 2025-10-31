package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1430-ce-SchedulingEnhancement-r14 ::= ENUMERATED
type PdschConfigdedicatedV1430CeSchedulingenhancementR14 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1430CeSchedulingenhancementR14EnumeratedNothing = iota
	PdschConfigdedicatedV1430CeSchedulingenhancementR14EnumeratedRange1
	PdschConfigdedicatedV1430CeSchedulingenhancementR14EnumeratedRange2
)
