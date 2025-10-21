package ies

import "rrc/utils"

// UE-EUTRA-Capability-v11d0-IEs ::= SEQUENCE
type UeEutraCapabilityV11d0Ies struct {
	RfParametersV11d0    *RfParametersV11d0
	OtherparametersV11d0 *OtherParametersV11d0
	Noncriticalextension *UeEutraCapabilityV11x0Ies
}
