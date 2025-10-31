package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1530-IEs-reducedCP-Latency-r15 ::= ENUMERATED
type UeEutraCapabilityV1530IesReducedcpLatencyR15 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1530IesReducedcpLatencyR15EnumeratedNothing = iota
	UeEutraCapabilityV1530IesReducedcpLatencyR15EnumeratedSupported
)
