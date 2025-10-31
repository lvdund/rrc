package ies

// UE-NR-Capability-v1570 ::= SEQUENCE
type UeNrCapabilityV1570 struct {
	NrdcParametersV1570  *NrdcParametersV1570
	Noncriticalextension *UeNrCapabilityV1610
}
