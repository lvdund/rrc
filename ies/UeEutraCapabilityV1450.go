package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1450-IEs ::= SEQUENCE
type UeEutraCapabilityV1450 struct {
	PhylayerparametersV1450 *PhylayerparametersV1450
	RfParametersV1450       *RfParametersV1450
	OtherparametersV1450    OtherparametersV1450
	UeCategorydlV1450       *utils.INTEGER `lb:0,ub:20`
	Noncriticalextension    *UeEutraCapabilityV1460
}
