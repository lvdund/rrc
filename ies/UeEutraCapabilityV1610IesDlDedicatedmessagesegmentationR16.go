package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1610-IEs-dl-DedicatedMessageSegmentation-r16 ::= ENUMERATED
type UeEutraCapabilityV1610IesDlDedicatedmessagesegmentationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1610IesDlDedicatedmessagesegmentationR16EnumeratedNothing = iota
	UeEutraCapabilityV1610IesDlDedicatedmessagesegmentationR16EnumeratedSupported
)
