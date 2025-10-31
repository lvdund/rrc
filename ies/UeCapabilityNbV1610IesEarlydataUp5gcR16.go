package ies

import "rrc/utils"

// UE-Capability-NB-v1610-IEs-earlyData-UP-5GC-r16 ::= ENUMERATED
type UeCapabilityNbV1610IesEarlydataUp5gcR16 struct {
	Value utils.ENUMERATED
}

const (
	UeCapabilityNbV1610IesEarlydataUp5gcR16EnumeratedNothing = iota
	UeCapabilityNbV1610IesEarlydataUp5gcR16EnumeratedSupported
)
