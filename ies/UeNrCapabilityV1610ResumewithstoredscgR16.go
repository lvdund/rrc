package ies

import "rrc/utils"

// UE-NR-Capability-v1610-resumeWithStoredSCG-r16 ::= ENUMERATED
type UeNrCapabilityV1610ResumewithstoredscgR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610ResumewithstoredscgR16EnumeratedNothing = iota
	UeNrCapabilityV1610ResumewithstoredscgR16EnumeratedSupported
)
