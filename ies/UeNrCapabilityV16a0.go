package ies

// UE-NR-Capability-v16a0 ::= SEQUENCE
type UeNrCapabilityV16a0 struct {
	PhyParametersV16a0   *PhyParametersV16a0
	RfParametersV16a0    *RfParametersV16a0
	Noncriticalextension *UeNrCapabilityV16a0Noncriticalextension
}
