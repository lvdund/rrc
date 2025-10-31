package ies

// UE-EUTRA-Capability-v1370-IEs ::= SEQUENCE
type UeEutraCapabilityV1370 struct {
	CeParametersV1370              *CeParametersV1370
	FddAddUeEutraCapabilitiesV1370 *UeEutraCapabilityaddxddModeV1370
	TddAddUeEutraCapabilitiesV1370 *UeEutraCapabilityaddxddModeV1370
	Noncriticalextension           *UeEutraCapabilityV1380
}
