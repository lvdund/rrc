package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1280-IEs ::= SEQUENCE
type UeEutraCapabilityV1280Ies struct {
	PhylayerparametersV1280 *PhylayerparametersV1280
	Noncriticalextension    *UeEutraCapabilityV1310Ies
}
