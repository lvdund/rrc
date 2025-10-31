package ies

import "rrc/utils"

// UE-NR-Capability-v1610-bh-RLF-Indication-r16 ::= ENUMERATED
type UeNrCapabilityV1610BhRlfIndicationR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610BhRlfIndicationR16EnumeratedNothing = iota
	UeNrCapabilityV1610BhRlfIndicationR16EnumeratedSupported
)
