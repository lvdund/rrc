package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1330-IEs ::= SEQUENCE
type UeEutraCapabilityV1330Ies struct {
	UeCategorydlV1330       *utils.INTEGER
	PhylayerparametersV1330 *PhylayerparametersV1330
	UeCeNeedulgapsR13       *utils.ENUMERATED
	Noncriticalextension    *UeEutraCapabilityV1340Ies
}
