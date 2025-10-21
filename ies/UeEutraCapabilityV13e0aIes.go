package ies

import "rrc/utils"

// UE-EUTRA-Capability-v13e0a-IEs ::= SEQUENCE
type UeEutraCapabilityV13e0aIes struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV1470Ies
}
