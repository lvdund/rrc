package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1260-IEs ::= SEQUENCE
type UeEutraCapabilityV1260 struct {
	UeCategorydlV1260    *utils.INTEGER `lb:0,ub:16`
	Noncriticalextension *UeEutraCapabilityV1270
}
