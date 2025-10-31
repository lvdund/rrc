package ies

// UE-EUTRA-Capability-v1440-IEs ::= SEQUENCE
type UeEutraCapabilityV1440 struct {
	LwaParametersV1440   LwaParametersV1440
	MacParametersV1440   MacParametersV1440
	Noncriticalextension *UeEutraCapabilityV1450
}
