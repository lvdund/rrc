package ies

import "rrc/utils"

// UE-NR-Capability-v1690-ul-RRC-Segmentation-r16 ::= ENUMERATED
type UeNrCapabilityV1690UlRrcSegmentationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1690UlRrcSegmentationR16EnumeratedNothing = iota
	UeNrCapabilityV1690UlRrcSegmentationR16EnumeratedSupported
)
