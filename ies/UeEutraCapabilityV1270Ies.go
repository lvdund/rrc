package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1270-IEs ::= SEQUENCE
type UeEutraCapabilityV1270Ies struct {
	RfParametersV1270    *RfParametersV1270
	Noncriticalextension *UeEutraCapabilityV1280Ies
}
