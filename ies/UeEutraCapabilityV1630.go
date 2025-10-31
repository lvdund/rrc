package ies

// UE-EUTRA-Capability-v1630-IEs ::= SEQUENCE
type UeEutraCapabilityV1630 struct {
	RfParametersV1630              *RfParametersV1630
	SlParametersV1630              *SlParametersV1630
	EarlysecurityreactivationR16   *UeEutraCapabilityV1630IesEarlysecurityreactivationR16
	MacParametersV1630             MacParametersV1630
	MeasparametersV1630            *MeasparametersV1630
	FddAddUeEutraCapabilitiesV1630 UeEutraCapabilityaddxddModeV1630
	TddAddUeEutraCapabilitiesV1630 UeEutraCapabilityaddxddModeV1630
	Noncriticalextension           *UeEutraCapabilityV1650
}
