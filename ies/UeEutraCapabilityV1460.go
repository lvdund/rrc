package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1460-IEs ::= SEQUENCE
type UeEutraCapabilityV1460 struct {
	UeCategorydlV1460    *utils.INTEGER `lb:0,ub:21`
	OtherparametersV1460 OtherParametersV1460
	Noncriticalextension *UeEutraCapabilityV1510
}
