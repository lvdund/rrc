package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1510-IEs ::= SEQUENCE
type UeEutraCapabilityV1510Ies struct {
	IratParametersnrR15            *IratParametersnrR15
	FeaturesetseutraR15            *FeaturesetseutraR15
	PdcpParametersnrR15            *PdcpParametersnrR15
	FddAddUeEutraCapabilitiesV1510 *UeEutraCapabilityaddxddModeV1510
	TddAddUeEutraCapabilitiesV1510 *UeEutraCapabilityaddxddModeV1510
	Noncriticalextension           *UeEutraCapabilityV1520Ies
}
