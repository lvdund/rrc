package ies

import "rrc/utils"

// UE-NR-Capability-v1700-inactiveStatePO-Determination-r17 ::= ENUMERATED
type UeNrCapabilityV1700InactivestatepoDeterminationR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700InactivestatepoDeterminationR17EnumeratedNothing = iota
	UeNrCapabilityV1700InactivestatepoDeterminationR17EnumeratedSupported
)
