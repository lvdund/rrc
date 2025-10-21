package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1440-IEs ::= SEQUENCE
type UeEutraCapabilityV1440Ies struct {
	LwaParametersV1440   LwaParametersV1440
	MacParametersV1440   MacParametersV1440
	Noncriticalextension *UeEutraCapabilityV1450Ies
}
