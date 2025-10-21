package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1320-IEs ::= SEQUENCE
type UeEutraCapabilityV1320Ies struct {
	CeParametersV1320              *CeParametersV1320
	PhylayerparametersV1320        *PhylayerparametersV1320
	RfParametersV1320              *RfParametersV1320
	FddAddUeEutraCapabilitiesV1320 *UeEutraCapabilityaddxddModeV1320
	TddAddUeEutraCapabilitiesV1320 *UeEutraCapabilityaddxddModeV1320
	Noncriticalextension           *UeEutraCapabilityV1330Ies
}
