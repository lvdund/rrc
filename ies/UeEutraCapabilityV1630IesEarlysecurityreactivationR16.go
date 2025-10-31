package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1630-IEs-earlySecurityReactivation-r16 ::= ENUMERATED
type UeEutraCapabilityV1630IesEarlysecurityreactivationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1630IesEarlysecurityreactivationR16EnumeratedNothing = iota
	UeEutraCapabilityV1630IesEarlysecurityreactivationR16EnumeratedSupported
)
