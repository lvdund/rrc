package ies

import "rrc/utils"

// UE-EUTRA-Capability-v11a0-IEs ::= SEQUENCE
type UeEutraCapabilityV11a0 struct {
	UeCategoryV11a0      *utils.INTEGER `lb:0,ub:12`
	MeasparametersV11a0  *MeasparametersV11a0
	Noncriticalextension *UeEutraCapabilityV1250
}
