package ies

import "rrc/utils"

// UE-EUTRA-Capability-v12x0-IEs ::= SEQUENCE
type UeEutraCapabilityV12x0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV1370Ies
}
