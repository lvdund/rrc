package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1340-IEs ::= SEQUENCE
type UeEutraCapabilityV1340 struct {
	UeCategoryulV1340    *utils.INTEGER `lb:0,ub:15`
	Noncriticalextension *UeEutraCapabilityV1350
}
