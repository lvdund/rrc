package ies

// UE-EUTRA-Capability-v1350-IEs ::= SEQUENCE
type UeEutraCapabilityV1350 struct {
	UeCategorydlV1350    *UeEutraCapabilityV1350IesUeCategorydlV1350
	UeCategoryulV1350    *UeEutraCapabilityV1350IesUeCategoryulV1350
	CeParametersV1350    CeParametersV1350
	Noncriticalextension *UeEutraCapabilityV1360
}
