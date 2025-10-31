package ies

import "rrc/utils"

// UE-NR-Capability-v1610-inDeviceCoexInd-r16 ::= ENUMERATED
type UeNrCapabilityV1610IndevicecoexindR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610IndevicecoexindR16EnumeratedNothing = iota
	UeNrCapabilityV1610IndevicecoexindR16EnumeratedSupported
)
