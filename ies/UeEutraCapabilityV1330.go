package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1330-IEs ::= SEQUENCE
type UeEutraCapabilityV1330 struct {
	UeCategorydlV1330       *utils.INTEGER `lb:0,ub:19`
	PhylayerparametersV1330 *PhylayerparametersV1330
	UeCeNeedulgapsR13       *bool
	Noncriticalextension    *UeEutraCapabilityV1340
}
