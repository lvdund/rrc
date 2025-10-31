package ies

// UE-EUTRA-Capability-v1130-IEs ::= SEQUENCE
type UeEutraCapabilityV1130 struct {
	PdcpParametersV1130             PdcpParametersV1130
	PhylayerparametersV1130         *PhylayerparametersV1130
	RfParametersV1130               RfParametersV1130
	MeasparametersV1130             MeasparametersV1130
	InterratParameterscdma2000V1130 IratParameterscdma2000V1130
	OtherparametersR11              OtherParametersR11
	FddAddUeEutraCapabilitiesV1130  *UeEutraCapabilityaddxddModeV1130
	TddAddUeEutraCapabilitiesV1130  *UeEutraCapabilityaddxddModeV1130
	Noncriticalextension            *UeEutraCapabilityV1170
}
