package ies

import "rrc/utils"

// UE-NR-Capability-v1700-ra-SDT-r17 ::= ENUMERATED
type UeNrCapabilityV1700RaSdtR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700RaSdtR17EnumeratedNothing = iota
	UeNrCapabilityV1700RaSdtR17EnumeratedSupported
)
