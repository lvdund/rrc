package ies

import "rrc/utils"

// UE-NR-Capability-v1610-referenceTimeProvision-r16 ::= ENUMERATED
type UeNrCapabilityV1610ReferencetimeprovisionR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610ReferencetimeprovisionR16EnumeratedNothing = iota
	UeNrCapabilityV1610ReferencetimeprovisionR16EnumeratedSupported
)
