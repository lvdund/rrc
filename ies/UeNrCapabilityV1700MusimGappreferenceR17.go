package ies

import "rrc/utils"

// UE-NR-Capability-v1700-musim-GapPreference-r17 ::= ENUMERATED
type UeNrCapabilityV1700MusimGappreferenceR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700MusimGappreferenceR17EnumeratedNothing = iota
	UeNrCapabilityV1700MusimGappreferenceR17EnumeratedSupported
)
