package ies

import "rrc/utils"

// UE-EUTRA-Capability-v14b0-IEs ::= SEQUENCE
type UeEutraCapabilityV14b0Ies struct {
	RfParametersV14b0    *RfParametersV14b0
	Noncriticalextension *UeEutraCapabilityV14x0Ies
}
