package ies

import "rrc/utils"

// UE-EUTRA-Capability-v9c0-IEs ::= SEQUENCE
type UeEutraCapabilityV9c0Ies struct {
	InterratParametersutraV9c0 *IratParametersutraV9c0
	Noncriticalextension       *UeEutraCapabilityV9d0Ies
}
