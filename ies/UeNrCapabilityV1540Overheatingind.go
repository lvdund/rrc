package ies

import "rrc/utils"

// UE-NR-Capability-v1540-overheatingInd ::= ENUMERATED
type UeNrCapabilityV1540Overheatingind struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1540OverheatingindEnumeratedNothing = iota
	UeNrCapabilityV1540OverheatingindEnumeratedSupported
)
