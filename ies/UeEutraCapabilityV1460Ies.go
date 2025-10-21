package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1460-IEs ::= SEQUENCE
type UeEutraCapabilityV1460Ies struct {
	UeCategorydlV1460    *utils.INTEGER
	OtherparametersV1460 OtherParametersV1460
	Noncriticalextension *UeEutraCapabilityV1510Ies
}
