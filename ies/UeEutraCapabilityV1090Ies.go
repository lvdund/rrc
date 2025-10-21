package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1090-IEs ::= SEQUENCE
type UeEutraCapabilityV1090Ies struct {
	RfParametersV1090    *RfParametersV1090
	Noncriticalextension *UeEutraCapabilityV1130Ies
}
