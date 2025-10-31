package ies

import "rrc/utils"

// UE-NR-Capability-v1610-onDemandSIB-Connected-r16 ::= ENUMERATED
type UeNrCapabilityV1610OndemandsibConnectedR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610OndemandsibConnectedR16EnumeratedNothing = iota
	UeNrCapabilityV1610OndemandsibConnectedR16EnumeratedSupported
)
