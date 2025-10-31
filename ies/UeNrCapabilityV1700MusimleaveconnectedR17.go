package ies

import "rrc/utils"

// UE-NR-Capability-v1700-musimLeaveConnected-r17 ::= ENUMERATED
type UeNrCapabilityV1700MusimleaveconnectedR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700MusimleaveconnectedR17EnumeratedNothing = iota
	UeNrCapabilityV1700MusimleaveconnectedR17EnumeratedSupported
)
