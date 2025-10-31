package ies

// UE-EUTRA-Capability-v1510-IEs ::= SEQUENCE
type UeEutraCapabilityV1510 struct {
	IratParametersnrR15            *IratParametersnrR15
	FeaturesetseutraR15            *FeaturesetseutraR15
	PdcpParametersnrR15            *PdcpParametersnrR15
	FddAddUeEutraCapabilitiesV1510 *UeEutraCapabilityaddxddModeV1510
	TddAddUeEutraCapabilitiesV1510 *UeEutraCapabilityaddxddModeV1510
	Noncriticalextension           *UeEutraCapabilityV1520
}
