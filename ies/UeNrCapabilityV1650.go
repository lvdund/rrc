package ies

// UE-NR-Capability-v1650 ::= SEQUENCE
type UeNrCapabilityV1650 struct {
	MpspriorityindicationR16 *UeNrCapabilityV1650MpspriorityindicationR16
	HighspeedparametersV1650 *HighspeedparametersV1650
	Noncriticalextension     *UeNrCapabilityV1690
}
