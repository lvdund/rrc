package ies

// UE-NR-Capability-v1550 ::= SEQUENCE
type UeNrCapabilityV1550 struct {
	ReducedcpLatency     *UeNrCapabilityV1550ReducedcpLatency
	Noncriticalextension *UeNrCapabilityV1560
}
