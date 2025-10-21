package ies

import "rrc/utils"

// UE-EUTRA-Capability-v9e0-IEs ::= SEQUENCE
type UeEutraCapabilityV9e0Ies struct {
	RfParametersV9e0     *RfParametersV9e0
	Noncriticalextension *UeEutraCapabilityV9h0Ies
}
