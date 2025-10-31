package ies

import "rrc/utils"

// UE-NR-Capability-v1530-inactiveState ::= ENUMERATED
type UeNrCapabilityV1530Inactivestate struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1530InactivestateEnumeratedNothing = iota
	UeNrCapabilityV1530InactivestateEnumeratedSupported
)
