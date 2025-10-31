package ies

import "rrc/utils"

// UE-NR-Capability-v1610-directSN-AdditionFirstRRC-IAB-r16 ::= ENUMERATED
type UeNrCapabilityV1610DirectsnAdditionfirstrrcIabR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610DirectsnAdditionfirstrrcIabR16EnumeratedNothing = iota
	UeNrCapabilityV1610DirectsnAdditionfirstrrcIabR16EnumeratedSupported
)
