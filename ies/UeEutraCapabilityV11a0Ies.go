package ies

import "rrc/utils"

// UE-EUTRA-Capability-v11a0-IEs ::= SEQUENCE
type UeEutraCapabilityV11a0Ies struct {
	UeCategoryV11a0      *utils.INTEGER
	MeasparametersV11a0  *MeasparametersV11a0
	Noncriticalextension *UeEutraCapabilityV1250Ies
}
