package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1350-IEs ::= SEQUENCE
type UeEutraCapabilityV1350Ies struct {
	UeCategorydlV1350    *utils.ENUMERATED
	UeCategoryulV1350    *utils.ENUMERATED
	CeParametersV1350    CeParametersV1350
	Noncriticalextension *UeEutraCapabilityV1360Ies
}
