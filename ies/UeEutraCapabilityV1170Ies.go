package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1170-IEs ::= SEQUENCE
type UeEutraCapabilityV1170Ies struct {
	PhylayerparametersV1170 *PhylayerparametersV1170
	UeCategoryV1170         *utils.INTEGER
	Noncriticalextension    *UeEutraCapabilityV1180Ies
}
