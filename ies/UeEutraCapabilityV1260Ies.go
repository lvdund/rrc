package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1260-IEs ::= SEQUENCE
type UeEutraCapabilityV1260Ies struct {
	UeCategorydlV1260    *utils.INTEGER
	Noncriticalextension *UeEutraCapabilityV1270Ies
}
