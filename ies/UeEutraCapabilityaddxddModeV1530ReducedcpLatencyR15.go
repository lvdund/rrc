package ies

import "rrc/utils"

// UE-EUTRA-CapabilityAddXDD-Mode-v1530-reducedCP-Latency-r15 ::= ENUMERATED
type UeEutraCapabilityaddxddModeV1530ReducedcpLatencyR15 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityaddxddModeV1530ReducedcpLatencyR15EnumeratedNothing = iota
	UeEutraCapabilityaddxddModeV1530ReducedcpLatencyR15EnumeratedSupported
)
