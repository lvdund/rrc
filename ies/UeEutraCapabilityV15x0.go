package ies

import "rrc/utils"

// UE-EUTRA-Capability-v15x0-IEs ::= SEQUENCE
type UeEutraCapabilityV15x0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV16c0
}
