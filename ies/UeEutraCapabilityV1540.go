package ies

// UE-EUTRA-Capability-v1540-IEs ::= SEQUENCE
type UeEutraCapabilityV1540 struct {
	PhylayerparametersV1540        *PhylayerparametersV1540
	OtherparametersV1540           OtherParametersV1540
	FddAddUeEutraCapabilitiesV1540 *UeEutraCapabilityaddxddModeV1540
	TddAddUeEutraCapabilitiesV1540 *UeEutraCapabilityaddxddModeV1540
	SlParametersV1540              *SlParametersV1540
	IratParametersnrV1540          *IratParametersnrV1540
	Noncriticalextension           *UeEutraCapabilityV1550
}
