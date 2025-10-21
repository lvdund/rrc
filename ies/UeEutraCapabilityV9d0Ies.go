package ies

import "rrc/utils"

// UE-EUTRA-Capability-v9d0-IEs ::= SEQUENCE
type UeEutraCapabilityV9d0Ies struct {
	PhylayerparametersV9d0 *PhylayerparametersV9d0
	Noncriticalextension   *UeEutraCapabilityV9e0Ies
}
