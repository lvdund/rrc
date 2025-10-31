package ies

import "rrc/utils"

// UE-Capability-NB-v1610-IEs-earlySecurityReactivation-r16 ::= ENUMERATED
type UeCapabilityNbV1610IesEarlysecurityreactivationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeCapabilityNbV1610IesEarlysecurityreactivationR16EnumeratedNothing = iota
	UeCapabilityNbV1610IesEarlysecurityreactivationR16EnumeratedSupported
)
