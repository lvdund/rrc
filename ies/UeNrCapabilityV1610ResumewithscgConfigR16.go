package ies

import "rrc/utils"

// UE-NR-Capability-v1610-resumeWithSCG-Config-r16 ::= ENUMERATED
type UeNrCapabilityV1610ResumewithscgConfigR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610ResumewithscgConfigR16EnumeratedNothing = iota
	UeNrCapabilityV1610ResumewithscgConfigR16EnumeratedSupported
)
