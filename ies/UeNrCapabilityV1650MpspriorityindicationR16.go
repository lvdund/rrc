package ies

import "rrc/utils"

// UE-NR-Capability-v1650-mpsPriorityIndication-r16 ::= ENUMERATED
type UeNrCapabilityV1650MpspriorityindicationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1650MpspriorityindicationR16EnumeratedNothing = iota
	UeNrCapabilityV1650MpspriorityindicationR16EnumeratedSupported
)
