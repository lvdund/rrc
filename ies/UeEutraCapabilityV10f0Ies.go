package ies

import "rrc/utils"

// UE-EUTRA-Capability-v10f0-IEs ::= SEQUENCE
type UeEutraCapabilityV10f0Ies struct {
	RfParametersV10f0    *RfParametersV10f0
	Noncriticalextension *UeEutraCapabilityV10i0Ies
}
