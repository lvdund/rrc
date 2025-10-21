package ies

import "rrc/utils"

// UE-EUTRA-Capability-v12b0-IEs ::= SEQUENCE
type UeEutraCapabilityV12b0Ies struct {
	RfParametersV12b0    *RfParametersV12b0
	Noncriticalextension *UeEutraCapabilityV12x0Ies
}
