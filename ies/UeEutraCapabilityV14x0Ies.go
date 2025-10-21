package ies

import "rrc/utils"

// UE-EUTRA-Capability-v14x0-IEs ::= SEQUENCE
type UeEutraCapabilityV14x0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV15x0Ies
}
