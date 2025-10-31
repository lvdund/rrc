package ies

import "rrc/utils"

// UE-NR-Capability-v1700-nonTerrestrialNetwork-r17 ::= ENUMERATED
type UeNrCapabilityV1700NonterrestrialnetworkR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700NonterrestrialnetworkR17EnumeratedNothing = iota
	UeNrCapabilityV1700NonterrestrialnetworkR17EnumeratedSupported
)
