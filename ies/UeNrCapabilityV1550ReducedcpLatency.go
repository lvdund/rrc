package ies

import "rrc/utils"

// UE-NR-Capability-v1550-reducedCP-Latency ::= ENUMERATED
type UeNrCapabilityV1550ReducedcpLatency struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1550ReducedcpLatencyEnumeratedNothing = iota
	UeNrCapabilityV1550ReducedcpLatencyEnumeratedSupported
)
