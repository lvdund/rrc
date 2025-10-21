package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1520-IEs ::= SEQUENCE
type UeEutraCapabilityV1520Ies struct {
	MeasparametersV1520  MeasparametersV1520
	Noncriticalextension *UeEutraCapabilityV1530Ies
}
