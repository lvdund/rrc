package ies

import "rrc/utils"

// UE-NR-Capability-v1610-dl-DedicatedMessageSegmentation-r16 ::= ENUMERATED
type UeNrCapabilityV1610DlDedicatedmessagesegmentationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610DlDedicatedmessagesegmentationR16EnumeratedNothing = iota
	UeNrCapabilityV1610DlDedicatedmessagesegmentationR16EnumeratedSupported
)
