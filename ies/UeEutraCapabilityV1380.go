package ies

// UE-EUTRA-Capability-v1380-IEs ::= SEQUENCE
type UeEutraCapabilityV1380 struct {
	RfParametersV1380              *RfParametersV1380
	CeParametersV1380              CeParametersV1380
	FddAddUeEutraCapabilitiesV1380 UeEutraCapabilityaddxddModeV1380
	TddAddUeEutraCapabilitiesV1380 UeEutraCapabilityaddxddModeV1380
	Noncriticalextension           *UeEutraCapabilityV1390
}
