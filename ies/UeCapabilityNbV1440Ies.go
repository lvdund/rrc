package ies

import "rrc/utils"

// UE-Capability-NB-v1440-IEs ::= SEQUENCE
type UeCapabilityNbV1440Ies struct {
	PhylayerparametersV1440 *PhylayerparametersNbV1440
	Noncriticalextension    *UeCapabilityNbV14x0Ies
}
