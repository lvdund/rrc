package ies

import "rrc/utils"

// UE-EUTRA-Capability-v940-IEs ::= SEQUENCE
type UeEutraCapabilityV940Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV1020Ies
}
