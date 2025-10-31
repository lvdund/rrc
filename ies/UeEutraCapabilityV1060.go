package ies

// UE-EUTRA-Capability-v1060-IEs ::= SEQUENCE
type UeEutraCapabilityV1060 struct {
	FddAddUeEutraCapabilitiesV1060 *UeEutraCapabilityaddxddModeV1060
	TddAddUeEutraCapabilitiesV1060 *UeEutraCapabilityaddxddModeV1060
	RfParametersV1060              *RfParametersV1060
	Noncriticalextension           *UeEutraCapabilityV1090
}
