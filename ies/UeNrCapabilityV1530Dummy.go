package ies

import "rrc/utils"

// UE-NR-Capability-v1530-dummy ::= ENUMERATED
type UeNrCapabilityV1530Dummy struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1530DummyEnumeratedNothing = iota
	UeNrCapabilityV1530DummyEnumeratedSupported
)
