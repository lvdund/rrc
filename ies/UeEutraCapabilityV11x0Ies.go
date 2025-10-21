package ies

import "rrc/utils"

// UE-EUTRA-Capability-v11x0-IEs ::= SEQUENCE
type UeEutraCapabilityV11x0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV12b0Ies
}
