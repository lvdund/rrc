package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1390-IEs ::= SEQUENCE
type UeEutraCapabilityV1390Ies struct {
	RfParametersV1390    *RfParametersV1390
	Noncriticalextension *UeEutraCapabilityV13e0aIes
}
