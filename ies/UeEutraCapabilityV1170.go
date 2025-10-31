package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1170-IEs ::= SEQUENCE
type UeEutraCapabilityV1170 struct {
	PhylayerparametersV1170 *PhylayerparametersV1170
	UeCategoryV1170         *utils.INTEGER `lb:0,ub:10`
	Noncriticalextension    *UeEutraCapabilityV1180
}
