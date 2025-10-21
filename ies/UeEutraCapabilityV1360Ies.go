package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1360-IEs ::= SEQUENCE
type UeEutraCapabilityV1360Ies struct {
	OtherParametersV1360 *OtherParametersV1360
	Noncriticalextension *UeEutraCapabilityV1430Ies
}
