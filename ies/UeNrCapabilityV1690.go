package ies

// UE-NR-Capability-v1690 ::= SEQUENCE
type UeNrCapabilityV1690 struct {
	UlRrcSegmentationR16 *UeNrCapabilityV1690UlRrcSegmentationR16
	Noncriticalextension *UeNrCapabilityV1700
}
