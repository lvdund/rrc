package ies

import "rrc/utils"

// UE-NR-Capability-v1700-srb-SDT-r17 ::= ENUMERATED
type UeNrCapabilityV1700SrbSdtR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700SrbSdtR17EnumeratedNothing = iota
	UeNrCapabilityV1700SrbSdtR17EnumeratedSupported
)
