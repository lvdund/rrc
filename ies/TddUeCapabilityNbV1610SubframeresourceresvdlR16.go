package ies

import "rrc/utils"

// TDD-UE-Capability-NB-v1610-subframeResourceResvDL-r16 ::= ENUMERATED
type TddUeCapabilityNbV1610SubframeresourceresvdlR16 struct {
	Value utils.ENUMERATED
}

const (
	TddUeCapabilityNbV1610SubframeresourceresvdlR16EnumeratedNothing = iota
	TddUeCapabilityNbV1610SubframeresourceresvdlR16EnumeratedSupported
)
