package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1180-IEs ::= SEQUENCE
type UeEutraCapabilityV1180Ies struct {
	RfParametersV1180              *RfParametersV1180
	MbmsParametersR11              *MbmsParametersR11
	FddAddUeEutraCapabilitiesV1180 *UeEutraCapabilityaddxddModeV1180
	TddAddUeEutraCapabilitiesV1180 *UeEutraCapabilityaddxddModeV1180
	Noncriticalextension           *UeEutraCapabilityV11a0Ies
}
