package ies

import "rrc/utils"

// UE-EUTRA-Capability-v14a0-IEs ::= SEQUENCE
type UeEutraCapabilityV14a0Ies struct {
	PhylayerparametersV14a0 PhylayerparametersV14a0
	Noncriticalextension    *UeEutraCapabilityV14b0Ies
}
