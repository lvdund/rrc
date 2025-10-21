package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1450-IEs ::= SEQUENCE
type UeEutraCapabilityV1450Ies struct {
	PhylayerparametersV1450 *PhylayerparametersV1450
	RfParametersV1450       *RfParametersV1450
	OtherparametersV1450    OtherparametersV1450
	UeCategorydlV1450       *utils.INTEGER
	Noncriticalextension    *UeEutraCapabilityV1460Ies
}
