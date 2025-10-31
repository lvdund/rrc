package ies

import "rrc/utils"

// TDD-UE-Capability-NB-v1610-subframeResourceResvUL-r16 ::= ENUMERATED
type TddUeCapabilityNbV1610SubframeresourceresvulR16 struct {
	Value utils.ENUMERATED
}

const (
	TddUeCapabilityNbV1610SubframeresourceresvulR16EnumeratedNothing = iota
	TddUeCapabilityNbV1610SubframeresourceresvulR16EnumeratedSupported
)
