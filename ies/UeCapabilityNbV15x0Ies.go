package ies

import "rrc/utils"

// UE-Capability-NB-v15x0-IEs ::= SEQUENCE
type UeCapabilityNbV15x0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeCapabilityNbV1610Ies
}
