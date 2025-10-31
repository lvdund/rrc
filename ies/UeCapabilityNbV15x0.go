package ies

import "rrc/utils"

// UE-Capability-NB-v15x0-IEs ::= SEQUENCE
type UeCapabilityNbV15x0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeCapabilityNbV1610
}
