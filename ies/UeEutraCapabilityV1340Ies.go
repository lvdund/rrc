package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1340-IEs ::= SEQUENCE
type UeEutraCapabilityV1340Ies struct {
	UeCategoryulV1340    *utils.INTEGER
	Noncriticalextension *UeEutraCapabilityV1350Ies
}
